package phplint

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shyim/go-phplint/internal/php/pkg/ast"
	"github.com/shyim/go-phplint/internal/php/pkg/visitor"
	"github.com/shyim/go-phplint/internal/php/pkg/visitor/traverser"
)

type compileValidator struct {
	visitor.Null

	filename    string
	source      []byte
	version     Version
	stack       []ast.Vertex
	diagnostics []Diagnostic
}

func validate(root ast.Vertex, source []byte, filename string, version Version) []Diagnostic {
	validator := &compileValidator{
		filename: filename,
		source:   source,
		version:  version,
	}
	traverser.NewTraverser(validator).Traverse(root)
	return validator.diagnostics
}

func (v *compileValidator) EnterNode(node ast.Vertex) bool {
	switch current := node.(type) {
	case *ast.Root:
		v.validateRoot(current)
	case *ast.StmtNamespace:
		v.validateDeclarationList(current.Stmts)
	case *ast.StmtClass:
		v.validateClass(current)
	case *ast.StmtInterface:
		v.validateClassStatements("interface", false, current.Stmts)
	case *ast.StmtTrait:
		v.validateClassStatements("trait", false, current.Stmts)
	case *ast.StmtEnum:
		v.requireVersion(current, PHP81, "enums")
		v.validateClassStatements("enum", false, current.Stmts)
	case *ast.StmtClassMethod:
		v.validateMethod(current)
	case *ast.StmtFunction:
		v.validateFunction(current)
	case *ast.ExprClosure:
		v.validateClosure(current)
	case *ast.ExprArrowFunction:
		v.validateArrowFunction(current)
	case *ast.ExprAssignCoalesce:
		v.requireVersion(current, PHP74, "null coalescing assignment")
	case *ast.ExprArrayItem:
		v.validateArrayItem(current)
	case *ast.StmtPropertyList:
		v.validatePropertyList(current)
	case *ast.Parameter:
		v.validateParameter(current)
	case *ast.Nullable:
		v.validateNullable(current)
	case *ast.Union:
		v.validateUnion(current)
	case *ast.Intersection:
		v.validateIntersection(current)
	case *ast.StmtBreak:
		v.validateLoopControl(current, "break", current.Expr)
	case *ast.StmtContinue:
		v.validateLoopControl(current, "continue", current.Expr)
	case *ast.StmtReturn:
		v.validateReturn(current)
	case *ast.ExprYield:
		v.validateYield(current)
	case *ast.ExprYieldFrom:
		v.validateYield(current)
	case *ast.ExprFunctionCall:
		v.validateTrailingCallComma(current, len(current.Args), len(current.SeparatorTkns))
		if current.EllipsisTkn != nil {
			v.requireVersion(current, PHP81, "first-class callable syntax")
		}
	case *ast.ExprMethodCall:
		v.validateTrailingCallComma(current, len(current.Args), len(current.SeparatorTkns))
		if current.EllipsisTkn != nil {
			v.requireVersion(current, PHP81, "first-class callable syntax")
		}
	case *ast.ExprStaticCall:
		v.validateTrailingCallComma(current, len(current.Args), len(current.SeparatorTkns))
		if current.EllipsisTkn != nil {
			v.requireVersion(current, PHP81, "first-class callable syntax")
		}
	case *ast.ExprNew:
		v.validateTrailingCallComma(current, len(current.Args), len(current.SeparatorTkns))
	case *ast.StmtClassConstList:
		v.validateClassConstants(current)
	case *ast.EnumCase:
		v.validateEnumCase(current)
	case *ast.StmtConstList:
		v.validateConstants(current.Consts, "constant", constantContextGlobalConstant)
	case *ast.StmtStaticVar:
		v.validateStaticVariable(current)
	}

	v.stack = append(v.stack, node)
	return true
}

func (v *compileValidator) LeaveNode(ast.Vertex) {
	if len(v.stack) > 0 {
		v.stack = v.stack[:len(v.stack)-1]
	}
}

func (v *compileValidator) requireVersion(node ast.Vertex, introduced Version, feature string) {
	if v.version >= introduced {
		return
	}
	v.report(node, fmt.Sprintf("%s require PHP %s", feature, introduced))
}

func (v *compileValidator) report(node ast.Vertex, message string) {
	if node == nil {
		return
	}
	start, end := positionFromInternal(node.GetPosition())
	v.diagnostics = append(v.diagnostics, Diagnostic{
		Filename: v.filename,
		Message:  message,
		Phase:    PhaseCompile,
		Start:    start,
		End:      end,
	})
}

func (v *compileValidator) validateRoot(root *ast.Root) {
	v.validateDeclarationList(root.Stmts)

	var namespaceMode string
	seenNamespace := false
	for _, statement := range root.Stmts {
		namespace, ok := statement.(*ast.StmtNamespace)
		if !ok {
			if seenNamespace && namespaceMode == "bracketed" && !allowedOutsideNamespace(statement) {
				v.report(statement, "code may not exist outside a bracketed namespace")
			}
			continue
		}

		mode := "unbracketed"
		if namespace.OpenCurlyBracketTkn != nil {
			mode = "bracketed"
		}
		if namespaceMode != "" && namespaceMode != mode {
			v.report(namespace, "cannot mix bracketed and unbracketed namespace declarations")
		}
		namespaceMode = mode
		seenNamespace = true
	}

	for _, statement := range root.Stmts {
		if _, ok := statement.(*ast.StmtNamespace); ok {
			break
		}
		if !allowedBeforeNamespace(statement) && containsNamespace(root.Stmts) {
			v.report(statement, "namespace declaration must be the first statement or follow declare")
			break
		}
	}
}

func (v *compileValidator) validateDeclarationList(statements []ast.Vertex) {
	functions := make(map[string]ast.Vertex)

	for _, statement := range statements {
		switch current := statement.(type) {
		case *ast.StmtFunction:
			v.trackDuplicate(functions, nodeName(current.Name), current, "function")
		}
	}
}

func (v *compileValidator) trackDuplicate(
	seen map[string]ast.Vertex,
	name string,
	node ast.Vertex,
	kind string,
) {
	key := strings.ToLower(name)
	if key == "" {
		return
	}
	if _, exists := seen[key]; exists {
		v.report(node, fmt.Sprintf("cannot redeclare %s %s", kind, name))
		return
	}
	seen[key] = node
}

func (v *compileValidator) validateClass(class *ast.StmtClass) {
	modifiers := v.validateModifiers(class, class.Modifiers)
	readonly := modifiers["readonly"]
	if readonly {
		v.requireVersion(class, PHP82, "readonly classes")
	}
	v.validateClassStatements("class", readonly, class.Stmts)
}

func (v *compileValidator) validateClassStatements(
	kind string,
	readonlyClass bool,
	statements []ast.Vertex,
) {
	methods := make(map[string]ast.Vertex)
	properties := make(map[string]ast.Vertex)
	constants := make(map[string]ast.Vertex)
	enumCases := make(map[string]ast.Vertex)

	for _, statement := range statements {
		switch current := statement.(type) {
		case *ast.StmtClassMethod:
			v.trackDuplicate(methods, nodeName(current.Name), current, "method")
		case *ast.StmtPropertyList:
			if kind == "interface" {
				v.report(current, "interfaces may not declare properties")
			}
			if kind == "enum" {
				v.report(current, "enums may not declare properties")
			}
			for _, property := range current.Props {
				v.trackDuplicateExact(properties, propertyName(property), property, "property")
			}
			if readonlyClass {
				modifiers := modifierSet(current.Modifiers)
				if modifiers["static"] {
					v.report(current, "readonly classes cannot declare static properties")
				}
				if current.Type == nil {
					v.report(current, "readonly properties must have a type")
				}
			}
		case *ast.StmtClassConstList:
			if kind == "trait" && v.version < PHP82 {
				v.report(current, "trait constants require PHP 8.2")
			}
			for _, constant := range current.Consts {
				v.trackDuplicateExact(constants, constantName(constant), constant, "class constant")
			}
		case *ast.EnumCase:
			v.trackDuplicateExact(enumCases, nodeName(current.Name), current, "enum case")
			v.trackDuplicateExact(constants, nodeName(current.Name), current, "class constant")
		}
	}
}

func (v *compileValidator) trackDuplicateExact(
	seen map[string]ast.Vertex,
	name string,
	node ast.Vertex,
	kind string,
) {
	if name == "" {
		return
	}
	if _, exists := seen[name]; exists {
		v.report(node, fmt.Sprintf("cannot redeclare %s %s", kind, name))
		return
	}
	seen[name] = node
}

func (v *compileValidator) validateMethod(method *ast.StmtClassMethod) {
	modifiers := v.validateModifiers(method, method.Modifiers)
	name := strings.ToLower(nodeName(method.Name))
	_, noBody := method.Stmt.(*ast.StmtNop)
	owner := v.enclosingClassKind()

	if modifiers["readonly"] {
		v.report(method, "methods cannot be readonly")
	}
	if modifiers["abstract"] && !noBody {
		v.report(method, "abstract methods cannot contain a body")
	}
	if !modifiers["abstract"] && noBody && owner != "interface" {
		v.report(method, "non-abstract methods must contain a body")
	}
	if owner == "interface" && !noBody {
		v.report(method, "interface methods cannot contain a body")
	}
	if name == "__construct" && method.ReturnType != nil {
		v.report(method.ReturnType, "constructors cannot declare a return type")
	}
	if name == "__destruct" {
		if method.ReturnType != nil {
			v.report(method.ReturnType, "destructors cannot declare a return type")
		}
		if len(method.Params) > 0 {
			v.report(method, "destructors cannot accept arguments")
		}
	}

	v.validateParameterNames(method.Params)
	v.validateType(method.ReturnType, typeContextReturn)
}

func (v *compileValidator) validateFunction(function *ast.StmtFunction) {
	v.validateParameterNames(function.Params)
	v.validateType(function.ReturnType, typeContextReturn)
}

func (v *compileValidator) validateClosure(closure *ast.ExprClosure) {
	v.validateParameterNames(closure.Params)
	v.validateType(closure.ReturnType, typeContextReturn)
}

func (v *compileValidator) validateArrowFunction(function *ast.ExprArrowFunction) {
	v.requireVersion(function, PHP74, "arrow functions")
	v.validateParameterNames(function.Params)
	v.validateType(function.ReturnType, typeContextReturn)
}

func (v *compileValidator) validateArrayItem(item *ast.ExprArrayItem) {
	if item.EllipsisTkn != nil {
		v.requireVersion(item, PHP74, "array unpacking")
	}
	if item.AmpersandTkn != nil && v.version < PHP73 && v.inDestructuringTarget(item) {
		v.report(item, "references in list assignments require PHP 7.3")
	}
}

func (v *compileValidator) validateTrailingCallComma(node ast.Vertex, arguments, separators int) {
	if v.version < PHP73 && arguments > 0 && separators >= arguments {
		v.report(node, "trailing commas in function calls require PHP 7.3")
	}
}

func (v *compileValidator) inDestructuringTarget(node ast.Vertex) bool {
	for index := len(v.stack) - 1; index >= 0; index-- {
		assignment, ok := v.stack[index].(*ast.ExprAssign)
		if !ok {
			continue
		}
		return positionContains(assignment.Var, node)
	}
	return false
}

func positionContains(container, child ast.Vertex) bool {
	if container == nil || child == nil ||
		container.GetPosition() == nil || child.GetPosition() == nil {
		return false
	}
	return child.GetPosition().StartPos >= container.GetPosition().StartPos &&
		child.GetPosition().EndPos <= container.GetPosition().EndPos
}

func (v *compileValidator) validateParameterNames(parameters []ast.Vertex) {
	seen := make(map[string]ast.Vertex)
	for _, parameterNode := range parameters {
		parameter, ok := parameterNode.(*ast.Parameter)
		if !ok {
			continue
		}
		name := variableName(parameter.Var)
		if _, exists := seen[name]; exists && name != "" {
			v.report(parameter, fmt.Sprintf("duplicate parameter $%s", name))
		}
		seen[name] = parameter
	}
}

func (v *compileValidator) validateParameter(parameter *ast.Parameter) {
	modifiers := v.validateModifiers(parameter, parameter.Modifiers)
	readonly := modifiers["readonly"] && v.version >= PHP81
	if modifiers["readonly"] && v.version < PHP81 && parameter.Type != nil {
		v.requireVersion(parameter, PHP81, "readonly promoted properties")
	}
	promoted := modifiers["public"] || modifiers["protected"] ||
		modifiers["private"] || readonly

	if readonly {
		if parameter.Type == nil {
			v.report(parameter, "readonly properties must have a type")
		}
	}
	if promoted {
		method := v.enclosingMethod()
		if method == nil || !strings.EqualFold(nodeName(method.Name), "__construct") {
			v.report(parameter, "promoted properties are only allowed in constructors")
		}
		if parameter.VariadicTkn != nil {
			v.report(parameter, "promoted properties cannot be variadic")
		}
	}

	v.validateType(parameter.Type, typeContextParameter)
	if parameter.DefaultValue != nil &&
		!isConstantExpression(parameter.DefaultValue, v.version, constantContextParameter) {
		v.report(parameter.DefaultValue, "default parameter value must be a constant expression")
	}
}

func (v *compileValidator) validatePropertyList(property *ast.StmtPropertyList) {
	modifiers := v.validateModifiers(property, property.Modifiers)
	readonly := modifiers["readonly"] && v.version >= PHP81
	if property.Type != nil {
		v.requireVersion(property.Type, PHP74, "typed properties")
	}
	if modifiers["readonly"] && v.version < PHP81 && property.Type != nil {
		v.requireVersion(property, PHP81, "readonly properties")
	}
	if modifiers["abstract"] {
		v.report(property, "properties cannot be abstract")
	}
	if modifiers["final"] {
		v.requireVersion(property, PHP84, "final properties")
	}
	if readonly {
		if property.Type == nil {
			v.report(property, "readonly properties must have a type")
		}
		if modifiers["static"] {
			v.report(property, "static properties cannot be readonly")
		}
	}

	v.validateType(property.Type, typeContextProperty)
	for _, propertyNode := range property.Props {
		current, ok := propertyNode.(*ast.StmtProperty)
		if !ok || current.Expr == nil {
			continue
		}
		if readonly || v.enclosingReadonlyClass() {
			v.report(current, "readonly properties cannot have default values")
		}
		if !isConstantExpression(current.Expr, v.version, constantContextProperty) {
			v.report(current.Expr, "property default value must be a constant expression")
		}
	}
}

func (v *compileValidator) validateClassConstants(constants *ast.StmtClassConstList) {
	modifiers := v.validateModifiers(constants, constants.Modifiers)
	if hasAttribute(constants.AttrGroups, "override") &&
		v.version >= PHP83 && v.version < PHP86 {
		v.report(constants, "the #[Override] attribute on class constants requires PHP 8.6")
	}
	if modifiers["abstract"] {
		v.report(constants, "class constants cannot be abstract")
	}
	if modifiers["static"] {
		v.report(constants, "class constants cannot be static")
	}
	if modifiers["readonly"] {
		v.report(constants, "class constants cannot be readonly")
	}
	if modifiers["final"] {
		v.requireVersion(constants, PHP81, "final class constants")
	}
	v.validateConstants(constants.Consts, "class constant", constantContextClassConstant)
}

func (v *compileValidator) validateEnumCase(enumCase *ast.EnumCase) {
	if hasAttribute(enumCase.AttrGroups, "override") &&
		v.version >= PHP83 && v.version < PHP86 {
		v.report(enumCase, "the #[Override] attribute on enum cases requires PHP 8.6")
	}
}

func (v *compileValidator) validateConstants(
	constants []ast.Vertex,
	kind string,
	context constantContext,
) {
	for _, constantNode := range constants {
		constant, ok := constantNode.(*ast.StmtConstant)
		if !ok || constant.Expr == nil {
			continue
		}
		if !isConstantExpression(constant.Expr, v.version, context) {
			v.report(constant.Expr, fmt.Sprintf("%s value must be a constant expression", kind))
		}
	}
}

func (v *compileValidator) validateStaticVariable(variable *ast.StmtStaticVar) {
	if variable.Expr == nil || v.version >= PHP83 {
		return
	}
	if !isConstantExpression(variable.Expr, v.version, constantContextStatic) {
		v.report(variable.Expr, "non-constant static variable initializers require PHP 8.3")
	}
}

func (v *compileValidator) validateModifiers(
	node ast.Vertex,
	modifierNodes []ast.Vertex,
) map[string]bool {
	seen := make(map[string]bool)
	accessCount := 0
	for _, modifier := range modifierNodes {
		name := strings.ToLower(nodeName(modifier))
		if name == "" {
			continue
		}
		if seen[name] {
			v.report(modifier, fmt.Sprintf("duplicate %s modifier", name))
		}
		seen[name] = true
		if name == "public" || name == "protected" || name == "private" {
			accessCount++
		}
	}

	if accessCount > 1 {
		v.report(node, "multiple access type modifiers are not allowed")
	}
	if seen["abstract"] && seen["final"] {
		v.report(node, "a declaration cannot be both abstract and final")
	}
	if seen["abstract"] && seen["private"] {
		v.report(node, "a private declaration cannot be abstract")
	}
	return seen
}

func (v *compileValidator) validateLoopControl(node ast.Vertex, keyword string, level ast.Vertex) {
	available := v.breakableLevels()
	requested := 1
	if level != nil {
		value, ok := integerLiteral(level)
		if !ok || value < 1 {
			v.report(node, fmt.Sprintf("%s accepts only positive integer levels", keyword))
			return
		}
		requested = value
	}

	if available == 0 {
		v.report(node, fmt.Sprintf("%s is not in a loop or switch context", keyword))
		return
	}
	if requested > available {
		v.report(
			node,
			fmt.Sprintf("cannot %s %d level(s); only %d available", keyword, requested, available),
		)
	}
}

func (v *compileValidator) breakableLevels() int {
	levels := 0
	for index := len(v.stack) - 1; index >= 0; index-- {
		switch v.stack[index].(type) {
		case *ast.StmtWhile, *ast.StmtFor, *ast.StmtForeach, *ast.StmtDo, *ast.StmtSwitch:
			levels++
		case *ast.StmtFunction, *ast.StmtClassMethod, *ast.ExprClosure, *ast.ExprArrowFunction:
			return levels
		}
	}
	return levels
}

func (v *compileValidator) validateReturn(statement *ast.StmtReturn) {
	returnType := v.enclosingReturnType()
	if statement.Expr != nil && strings.EqualFold(simpleTypeName(returnType), "void") {
		v.report(statement, "a void function must not return a value")
	}
	if statement.Expr == nil && strings.EqualFold(simpleTypeName(returnType), "never") {
		v.report(statement, "a never-returning function must not return")
	}
}

func (v *compileValidator) validateYield(node ast.Vertex) {
	returnType := v.enclosingReturnType()
	name := strings.ToLower(simpleTypeName(returnType))
	if name == "" {
		return
	}
	switch name {
	case "generator", "iterator", "traversable", "iterable", "object", "mixed":
		return
	default:
		v.report(node, fmt.Sprintf("generator return type %s is not compatible with Generator", name))
	}
}

func (v *compileValidator) enclosingReturnType() ast.Vertex {
	for index := len(v.stack) - 1; index >= 0; index-- {
		switch current := v.stack[index].(type) {
		case *ast.StmtFunction:
			return current.ReturnType
		case *ast.StmtClassMethod:
			return current.ReturnType
		case *ast.ExprClosure:
			return current.ReturnType
		case *ast.ExprArrowFunction:
			return current.ReturnType
		}
	}
	return nil
}

func (v *compileValidator) enclosingMethod() *ast.StmtClassMethod {
	for index := len(v.stack) - 1; index >= 0; index-- {
		switch current := v.stack[index].(type) {
		case *ast.StmtClassMethod:
			return current
		case *ast.StmtFunction, *ast.ExprClosure, *ast.ExprArrowFunction:
			return nil
		}
	}
	return nil
}

func (v *compileValidator) enclosingClassKind() string {
	for index := len(v.stack) - 1; index >= 0; index-- {
		switch v.stack[index].(type) {
		case *ast.StmtClass:
			return "class"
		case *ast.StmtInterface:
			return "interface"
		case *ast.StmtTrait:
			return "trait"
		case *ast.StmtEnum:
			return "enum"
		}
	}
	return ""
}

func (v *compileValidator) enclosingReadonlyClass() bool {
	for index := len(v.stack) - 1; index >= 0; index-- {
		class, ok := v.stack[index].(*ast.StmtClass)
		if ok {
			return modifierSet(class.Modifiers)["readonly"]
		}
	}
	return false
}

func modifierSet(modifiers []ast.Vertex) map[string]bool {
	result := make(map[string]bool)
	for _, modifier := range modifiers {
		result[strings.ToLower(nodeName(modifier))] = true
	}
	return result
}

func hasAttribute(groups []ast.Vertex, expected string) bool {
	for _, groupNode := range groups {
		group, ok := groupNode.(*ast.AttributeGroup)
		if !ok {
			continue
		}
		for _, attributeNode := range group.Attrs {
			attribute, ok := attributeNode.(*ast.Attribute)
			if !ok {
				continue
			}
			name := strings.ToLower(nodeName(attribute.Name))
			if name == `\`+expected {
				return true
			}
		}
	}
	return false
}

func nodeName(node ast.Vertex) string {
	switch current := node.(type) {
	case *ast.Identifier:
		return string(current.Value)
	case *ast.NamePart:
		return string(current.Value)
	case *ast.Name:
		return joinNameParts(current.Parts)
	case *ast.NameFullyQualified:
		return `\` + joinNameParts(current.Parts)
	case *ast.NameRelative:
		return `namespace\` + joinNameParts(current.Parts)
	default:
		return ""
	}
}

func joinNameParts(parts []ast.Vertex) string {
	names := make([]string, 0, len(parts))
	for _, part := range parts {
		if name := nodeName(part); name != "" {
			names = append(names, name)
		}
	}
	return strings.Join(names, `\`)
}

func variableName(node ast.Vertex) string {
	variable, ok := node.(*ast.ExprVariable)
	if !ok {
		return ""
	}
	return strings.TrimPrefix(nodeName(variable.Name), "$")
}

func propertyName(node ast.Vertex) string {
	property, ok := node.(*ast.StmtProperty)
	if !ok {
		return ""
	}
	return variableName(property.Var)
}

func constantName(node ast.Vertex) string {
	constant, ok := node.(*ast.StmtConstant)
	if !ok {
		return ""
	}
	return nodeName(constant.Name)
}

func integerLiteral(node ast.Vertex) (int, bool) {
	switch current := node.(type) {
	case *ast.ScalarLnumber:
		value, err := strconv.ParseInt(strings.ReplaceAll(string(current.Value), "_", ""), 0, 32)
		return int(value), err == nil
	case *ast.ExprBrackets:
		return integerLiteral(current.Expr)
	case *ast.ExprUnaryPlus:
		return integerLiteral(current.Expr)
	case *ast.ExprUnaryMinus:
		value, ok := integerLiteral(current.Expr)
		return -value, ok
	default:
		return 0, false
	}
}

func containsNamespace(statements []ast.Vertex) bool {
	for _, statement := range statements {
		if _, ok := statement.(*ast.StmtNamespace); ok {
			return true
		}
	}
	return false
}

func allowedBeforeNamespace(statement ast.Vertex) bool {
	switch current := statement.(type) {
	case *ast.StmtDeclare, *ast.StmtNop:
		return true
	case *ast.StmtInlineHtml:
		return strings.TrimSpace(string(current.Value)) == ""
	default:
		return false
	}
}

func allowedOutsideNamespace(statement ast.Vertex) bool {
	if allowedBeforeNamespace(statement) {
		return true
	}
	_, isNamespace := statement.(*ast.StmtNamespace)
	return isNamespace
}
