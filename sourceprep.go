package phplint

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/shyim/go-phplint/internal/token"
)

type sourceLayout struct {
	classMember   []bool
	classKind     []string
	functionParam []bool
}

func prepareSource(source []byte, tokens []*token.Token, version Version, filename string) []Diagnostic {
	if len(tokens) == 0 {
		return nil
	}

	var diagnostics []Diagnostic
	layout := analyzeLayout(tokens)

	prepareNumericLiteralSeparators(tokens, version, filename, &diagnostics)
	prepareTypedClassConstants(source, tokens, layout, version, filename, &diagnostics)
	prepareReadonlyAnonymousClasses(source, tokens, version, filename, &diagnostics)
	prepareDynamicClassConstants(source, tokens, version, filename, &diagnostics)
	preparePropertyHooks(source, tokens, layout, version, filename, &diagnostics)
	prepareUnparenthesizedNewDereference(source, tokens, version, filename, &diagnostics)
	preparePipes(source, tokens, version, filename, &diagnostics)
	prepareVoidCasts(source, tokens, version, filename, &diagnostics)
	prepareCloneWith(source, tokens, version, filename, &diagnostics)
	prepareConstantAttributes(source, tokens, layout, version, filename, &diagnostics)
	prepareFinalPromotedProperties(source, tokens, layout, version, filename, &diagnostics)

	return diagnostics
}

func prepareNumericLiteralSeparators(
	tokens []*token.Token,
	version Version,
	filename string,
	diagnostics *[]Diagnostic,
) {
	for _, current := range tokens {
		if current.ID != token.T_LNUMBER && current.ID != token.T_DNUMBER {
			continue
		}
		if !bytes.ContainsRune(current.Value, '_') {
			continue
		}
		addFeatureDiagnostic(
			diagnostics,
			filename,
			current,
			current,
			version,
			PHP74,
			"numeric literal separators",
		)
	}
}

func analyzeLayout(tokens []*token.Token) sourceLayout {
	layout := sourceLayout{
		classMember:   make([]bool, len(tokens)),
		classKind:     make([]string, len(tokens)),
		functionParam: make([]bool, len(tokens)),
	}

	braceDepth := 0
	type classFrame struct {
		depth int
		kind  string
	}
	var classes []classFrame
	pendingClassKind := ""

	parenDepth := 0
	var functionParamDepths []int
	pendingFunction := false

	for index, current := range tokens {
		if len(classes) > 0 && braceDepth == classes[len(classes)-1].depth {
			layout.classMember[index] = true
			layout.classKind[index] = classes[len(classes)-1].kind
		}
		if len(functionParamDepths) > 0 && parenDepth >= functionParamDepths[len(functionParamDepths)-1] {
			layout.functionParam[index] = true
		}

		switch current.ID {
		case token.T_CLASS:
			if index == 0 || tokens[index-1].ID != token.T_PAAMAYIM_NEKUDOTAYIM {
				pendingClassKind = "class"
			}
		case token.T_TRAIT:
			pendingClassKind = "trait"
		case token.T_INTERFACE:
			pendingClassKind = "interface"
		case token.T_ENUM:
			pendingClassKind = "enum"
		case token.T_FUNCTION, token.T_FN:
			pendingFunction = true
		case token.ID('{'), token.T_PROPERTY_HOOKS:
			braceDepth++
			if pendingClassKind != "" {
				classes = append(classes, classFrame{depth: braceDepth, kind: pendingClassKind})
				pendingClassKind = ""
			}
		case token.ID('}'):
			if len(classes) > 0 && classes[len(classes)-1].depth == braceDepth {
				classes = classes[:len(classes)-1]
			}
			if braceDepth > 0 {
				braceDepth--
			}
		case token.ID('('):
			parenDepth++
			if pendingFunction {
				functionParamDepths = append(functionParamDepths, parenDepth)
				pendingFunction = false
			}
		case token.ID(')'):
			if len(functionParamDepths) > 0 &&
				functionParamDepths[len(functionParamDepths)-1] == parenDepth {
				functionParamDepths = functionParamDepths[:len(functionParamDepths)-1]
			}
			if parenDepth > 0 {
				parenDepth--
			}
		}
	}

	return layout
}

func prepareTypedClassConstants(
	source []byte,
	tokens []*token.Token,
	layout sourceLayout,
	version Version,
	filename string,
	diagnostics *[]Diagnostic,
) {
	for index, current := range tokens {
		if current.ID != token.T_CONST || !layout.classMember[index] {
			continue
		}

		equal := findBefore(tokens, index+1, token.ID('='), token.ID(';'), token.ID('{'), token.ID('}'))
		if equal < index+2 {
			continue
		}

		name := equal - 1
		if !isIdentifierToken(tokens[name]) {
			continue
		}

		typeStart := index + 1
		if typeStart == name {
			continue
		}

		if version >= PHP83 {
			validateTypedClassConstantDeclaration(
				tokens,
				typeStart,
				name,
				equal,
				filename,
				diagnostics,
			)
		}
		blankRange(source, tokenStart(tokens[typeStart]), tokenStart(tokens[name]))
		addFeatureDiagnostic(
			diagnostics,
			filename,
			tokens[typeStart],
			tokens[name-1],
			version,
			PHP83,
			"typed class constants",
		)
	}
}

func validateTypedClassConstantDeclaration(
	tokens []*token.Token,
	typeStart, name, firstEqual int,
	filename string,
	diagnostics *[]Diagnostic,
) {
	typeNames := make(map[string]bool)
	nullable := false
	for index := typeStart; index < name; index++ {
		value := strings.ToLower(string(tokens[index].Value))
		switch value {
		case "?":
			nullable = true
		case "void", "never", "callable":
			addSourceDiagnostic(
				diagnostics,
				filename,
				tokens[index],
				tokens[index],
				fmt.Sprintf("class constants cannot have type %s", value),
			)
		case "", "|", "&", "(", ")", "\\":
		default:
			typeNames[value] = true
		}
	}
	if nullable {
		typeNames["null"] = true
	}

	equal := firstEqual
	for equal >= 0 && equal < len(tokens) {
		end, separator := constantValueEnd(tokens, equal+1)
		if end > equal+1 {
			kind, known := constantLiteralKind(tokens, equal+1, end)
			if known && !constantLiteralMatchesType(kind, typeNames) {
				addSourceDiagnostic(
					diagnostics,
					filename,
					tokens[equal+1],
					tokens[end-1],
					fmt.Sprintf("value of type %s is not compatible with the class constant type", kind),
				)
			}
		}
		if separator < 0 || tokens[separator].ID == token.ID(';') {
			return
		}

		equal = findTopLevelToken(tokens, separator+1, token.ID('='), token.ID(';'))
	}
}

func constantValueEnd(tokens []*token.Token, start int) (int, int) {
	paren, square, curly := 0, 0, 0
	for index := start; index < len(tokens); index++ {
		switch tokens[index].ID {
		case token.ID('('):
			paren++
		case token.ID(')'):
			if paren > 0 {
				paren--
			}
		case token.ID('['):
			square++
		case token.ID(']'):
			if square > 0 {
				square--
			}
		case token.ID('{'):
			curly++
		case token.ID('}'):
			if curly == 0 {
				return index, index
			}
			curly--
		case token.ID(','), token.ID(';'):
			if paren == 0 && square == 0 && curly == 0 {
				return index, index
			}
		}
	}
	return len(tokens), -1
}

func findTopLevelToken(tokens []*token.Token, start int, wanted, stop token.ID) int {
	paren, square, curly := 0, 0, 0
	for index := start; index < len(tokens); index++ {
		switch tokens[index].ID {
		case token.ID('('):
			paren++
		case token.ID(')'):
			if paren > 0 {
				paren--
			}
		case token.ID('['):
			square++
		case token.ID(']'):
			if square > 0 {
				square--
			}
		case token.ID('{'):
			curly++
		case token.ID('}'):
			if curly > 0 {
				curly--
			}
		default:
			if paren == 0 && square == 0 && curly == 0 {
				if tokens[index].ID == wanted {
					return index
				}
				if tokens[index].ID == stop {
					return -1
				}
			}
		}
	}
	return -1
}

func constantLiteralKind(tokens []*token.Token, start, end int) (string, bool) {
	for start < end && (tokens[start].ID == token.ID('+') || tokens[start].ID == token.ID('-')) {
		start++
	}
	if start >= end {
		return "", false
	}

	switch tokens[start].ID {
	case token.T_LNUMBER:
		return "int", true
	case token.T_DNUMBER:
		return "float", true
	case token.T_CONSTANT_ENCAPSED_STRING:
		return "string", true
	case token.T_ARRAY, token.ID('['):
		return "array", true
	}

	switch strings.ToLower(string(tokens[start].Value)) {
	case "true":
		return "true", true
	case "false":
		return "false", true
	case "null":
		return "null", true
	default:
		return "", false
	}
}

func constantLiteralMatchesType(kind string, typeNames map[string]bool) bool {
	if typeNames["mixed"] {
		return true
	}
	switch kind {
	case "int":
		return typeNames["int"] || typeNames["float"]
	case "float":
		return typeNames["float"]
	case "string":
		return typeNames["string"]
	case "true":
		return typeNames["true"] || typeNames["bool"]
	case "false":
		return typeNames["false"] || typeNames["bool"]
	case "null":
		return typeNames["null"]
	case "array":
		return typeNames["array"] || typeNames["iterable"]
	default:
		return true
	}
}

func prepareReadonlyAnonymousClasses(
	source []byte,
	tokens []*token.Token,
	version Version,
	filename string,
	diagnostics *[]Diagnostic,
) {
	for index := 0; index+2 < len(tokens); index++ {
		if tokens[index].ID != token.T_NEW ||
			!isReadonlyToken(tokens[index+1]) ||
			tokens[index+2].ID != token.T_CLASS {
			continue
		}

		blankToken(source, tokens[index+1])
		addFeatureDiagnostic(
			diagnostics,
			filename,
			tokens[index+1],
			tokens[index+1],
			version,
			PHP83,
			"readonly anonymous classes",
		)
	}
}

func prepareDynamicClassConstants(
	source []byte,
	tokens []*token.Token,
	version Version,
	filename string,
	diagnostics *[]Diagnostic,
) {
	for index := 0; index+2 < len(tokens); index++ {
		if tokens[index].ID != token.T_PAAMAYIM_NEKUDOTAYIM ||
			tokens[index+1].ID != token.ID('{') {
			continue
		}

		closeIndex := matchingToken(tokens, index+1, token.ID('{'), token.ID('}'))
		if closeIndex < 0 {
			continue
		}

		replaceRange(source, tokenStart(tokens[index+1]), tokenEnd(tokens[closeIndex]), []byte("$x"))
		addFeatureDiagnostic(
			diagnostics,
			filename,
			tokens[index+1],
			tokens[closeIndex],
			version,
			PHP83,
			"dynamic class constant fetch",
		)
		index = closeIndex
	}
}

func preparePropertyHooks(
	source []byte,
	tokens []*token.Token,
	layout sourceLayout,
	version Version,
	filename string,
	diagnostics *[]Diagnostic,
) {
	for index := 0; index+1 < len(tokens); index++ {
		if tokens[index].ID != token.T_PROPERTY_HOOKS || !layout.classMember[index] {
			continue
		}
		hookNameIndex := propertyHookNameIndex(tokens, index+1, len(tokens))
		if hookNameIndex < 0 || !looksLikePropertyDeclaration(tokens, index) {
			continue
		}

		closeIndex := matchingHookBlock(tokens, index)
		if closeIndex < 0 {
			continue
		}

		if version >= PHP84 {
			validatePropertyHookBlock(
				tokens,
				layout,
				index,
				closeIndex,
				filename,
				diagnostics,
			)
		}
		replaceRange(source, tokenStart(tokens[index]), tokenEnd(tokens[closeIndex]), []byte(";"))
		addFeatureDiagnostic(
			diagnostics,
			filename,
			tokens[index],
			tokens[closeIndex],
			version,
			PHP84,
			"property hooks",
		)
		index = closeIndex
	}
}

func validatePropertyHookBlock(
	tokens []*token.Token,
	layout sourceLayout,
	openIndex, closeIndex int,
	filename string,
	diagnostics *[]Diagnostic,
) {
	declarationStart := classMemberStart(tokens, openIndex)
	hasAbstract := false
	for index := declarationStart; index < openIndex; index++ {
		switch tokens[index].ID {
		case token.T_STATIC:
			addSourceDiagnostic(
				diagnostics,
				filename,
				tokens[index],
				tokens[index],
				"hooked properties cannot be static",
			)
		case token.T_READONLY:
			addSourceDiagnostic(
				diagnostics,
				filename,
				tokens[index],
				tokens[index],
				"hooked properties cannot be readonly",
			)
		case token.T_ABSTRACT:
			hasAbstract = true
		case token.ID(','):
			addSourceDiagnostic(
				diagnostics,
				filename,
				tokens[index],
				tokens[index],
				"hooked properties cannot be declared in a property list",
			)
		}
	}

	isInterface := layout.classKind[openIndex] == "interface"
	sawAbstractHook := false
	defer func() {
		if hasAbstract && !sawAbstractHook {
			addSourceDiagnostic(
				diagnostics,
				filename,
				tokens[openIndex],
				tokens[openIndex],
				"abstract property must specify at least one abstract hook",
			)
		}
	}()

	seen := make(map[string]bool)
	for index := openIndex + 1; index < closeIndex; {
		for index < closeIndex && tokens[index].ID == token.T_ATTRIBUTE {
			attributeEnd := matchingAttributeEnd(tokens, index)
			if attributeEnd < 0 || attributeEnd >= closeIndex {
				return
			}
			index = attributeEnd + 1
		}
		hookFinal := false
		if index < closeIndex && tokens[index].ID == token.T_FINAL {
			hookFinal = true
			index++
		}
		if index < closeIndex && isAmpersand(tokens[index].ID) {
			index++
		}
		if index >= closeIndex {
			return
		}

		nameToken := tokens[index]
		name := strings.ToLower(string(nameToken.Value))
		if name != "get" && name != "set" {
			addSourceDiagnostic(
				diagnostics,
				filename,
				nameToken,
				nameToken,
				fmt.Sprintf("unknown property hook %q; expected get or set", name),
			)
			index = skipHook(tokens, index+1, closeIndex)
			continue
		}

		if seen[name] {
			addSourceDiagnostic(
				diagnostics,
				filename,
				nameToken,
				nameToken,
				fmt.Sprintf("cannot redeclare property hook %q", name),
			)
		}
		seen[name] = true
		index++

		if index < closeIndex && tokens[index].ID == token.ID('(') {
			parameterEnd := matchingToken(tokens, index, token.ID('('), token.ID(')'))
			if parameterEnd < 0 || parameterEnd >= closeIndex {
				return
			}
			if name == "get" {
				addSourceDiagnostic(
					diagnostics,
					filename,
					tokens[index],
					tokens[parameterEnd],
					"get hooks must not have a parameter list",
				)
			} else if !validSetHookParameter(tokens, index+1, parameterEnd) {
				addSourceDiagnostic(
					diagnostics,
					filename,
					tokens[index],
					tokens[parameterEnd],
					"set hooks must declare exactly one non-reference, non-variadic parameter without a default",
				)
			}
			index = parameterEnd + 1
		}
		if index >= closeIndex {
			return
		}

		switch tokens[index].ID {
		case token.T_DOUBLE_ARROW:
			if isInterface {
				addSourceDiagnostic(
					diagnostics,
					filename,
					tokens[index],
					tokens[index],
					"abstract property hook cannot have a body",
				)
			}
			semicolon := findHookTerminator(tokens, index+1, closeIndex)
			if semicolon < 0 {
				addSourceDiagnostic(
					diagnostics,
					filename,
					tokens[index],
					tokens[index],
					"expression property hooks must end with a semicolon",
				)
				return
			}
			index = semicolon + 1
		case token.ID('{'):
			if isInterface {
				addSourceDiagnostic(
					diagnostics,
					filename,
					tokens[index],
					tokens[index],
					"abstract property hook cannot have a body",
				)
			}
			bodyEnd := matchingToken(tokens, index, token.ID('{'), token.ID('}'))
			if bodyEnd < 0 || bodyEnd > closeIndex {
				return
			}
			index = bodyEnd + 1
		case token.ID(';'):
			if hookFinal && (hasAbstract || isInterface) {
				addSourceDiagnostic(
					diagnostics,
					filename,
					nameToken,
					tokens[index],
					"property hook cannot be both abstract and final",
				)
			} else {
				sawAbstractHook = true
				if !hasAbstract && !isInterface {
					addSourceDiagnostic(
						diagnostics,
						filename,
						nameToken,
						tokens[index],
						"non-abstract property hooks must have a body",
					)
				}
			}
			index++
		default:
			addSourceDiagnostic(
				diagnostics,
				filename,
				tokens[index],
				tokens[index],
				"property hook must have a block, expression body, or semicolon",
			)
			index = skipHook(tokens, index+1, closeIndex)
		}
	}
}

func propertyHookNameIndex(tokens []*token.Token, start, end int) int {
	index := start
	for index < end && tokens[index].ID == token.T_ATTRIBUTE {
		attributeEnd := matchingAttributeEnd(tokens, index)
		if attributeEnd < 0 {
			return -1
		}
		index = attributeEnd + 1
	}
	if index < end && tokens[index].ID == token.T_FINAL {
		index++
	}
	if index < end && isAmpersand(tokens[index].ID) {
		index++
	}
	if index < end && (tokenIsWord(tokens[index], "get") || tokenIsWord(tokens[index], "set")) {
		return index
	}
	return -1
}

func looksLikePropertyDeclaration(tokens []*token.Token, openIndex int) bool {
	start := classMemberStart(tokens, openIndex)
	hasVariable := false
	for index := start; index < openIndex; index++ {
		switch tokens[index].ID {
		case token.T_FUNCTION, token.T_CONST:
			return false
		case token.T_VARIABLE:
			hasVariable = true
		}
	}
	return hasVariable
}

func validSetHookParameter(tokens []*token.Token, start, end int) bool {
	if start >= end {
		return false
	}
	variables := 0
	for index := start; index < end; index++ {
		switch tokens[index].ID {
		case token.T_VARIABLE:
			variables++
		case token.T_ELLIPSIS, token.ID('='):
			return false
		}
		if isAmpersand(tokens[index].ID) && index+1 < end && tokens[index+1].ID == token.T_VARIABLE {
			return false
		}
	}
	return variables == 1
}

// isMethodNameUse reports whether the keyword token at index is used as a
// method name rather than as an operator. Keywords such as "new" and "clone"
// are only semi-reserved: after "function" or "::" they name a method. After
// "->" and "?->" the lexer already emits them as plain identifiers.
func isMethodNameUse(tokens []*token.Token, index int) bool {
	return index > 0 && (tokens[index-1].ID == token.T_FUNCTION ||
		tokens[index-1].ID == token.T_PAAMAYIM_NEKUDOTAYIM)
}

func isAmpersand(id token.ID) bool {
	return id == token.ID('&') ||
		id == token.T_AMPERSAND_FOLLOWED_BY_VAR_OR_VARARG ||
		id == token.T_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG
}

func classMemberStart(tokens []*token.Token, index int) int {
	for cursor := index - 1; cursor >= 0; cursor-- {
		switch tokens[cursor].ID {
		case token.ID(';'), token.ID('{'), token.ID('}'):
			return cursor + 1
		}
	}
	return 0
}

func skipHook(tokens []*token.Token, start, end int) int {
	for index := start; index < end; index++ {
		switch tokens[index].ID {
		case token.ID(';'):
			return index + 1
		case token.ID('{'):
			closeIndex := matchingToken(tokens, index, token.ID('{'), token.ID('}'))
			if closeIndex >= 0 {
				return closeIndex + 1
			}
		}
	}
	return end
}

func findHookTerminator(tokens []*token.Token, start, end int) int {
	paren, square, curly := 0, 0, 0
	for index := start; index < end; index++ {
		switch tokens[index].ID {
		case token.ID('('):
			paren++
		case token.ID(')'):
			paren--
		case token.ID('['):
			square++
		case token.ID(']'):
			square--
		case token.ID('{'), token.T_CURLY_OPEN, token.T_DOLLAR_OPEN_CURLY_BRACES:
			curly++
		case token.ID('}'):
			curly--
		case token.ID(';'):
			if paren == 0 && square == 0 && curly == 0 {
				return index
			}
		}
	}
	return -1
}

func addSourceDiagnostic(
	diagnostics *[]Diagnostic,
	filename string,
	startToken, endToken *token.Token,
	message string,
) {
	start, _ := positionFromInternal(startToken.Position)
	_, end := positionFromInternal(endToken.Position)
	*diagnostics = append(*diagnostics, Diagnostic{
		Filename: filename,
		Message:  message,
		Phase:    PhaseCompile,
		Start:    start,
		End:      end,
	})
}

func preparePipes(
	source []byte,
	tokens []*token.Token,
	version Version,
	filename string,
	diagnostics *[]Diagnostic,
) {
	for index := 0; index+1 < len(tokens); index++ {
		if tokens[index].ID == token.T_PIPE {
			addFeatureDiagnostic(
				diagnostics,
				filename,
				tokens[index],
				tokens[index],
				version,
				PHP85,
				"the pipe operator",
			)
			continue
		}
		if tokens[index].ID != token.ID('|') || tokens[index+1].ID != token.ID('>') {
			continue
		}
		if tokenEnd(tokens[index]) != tokenStart(tokens[index+1]) {
			continue
		}

		addFeatureDiagnostic(
			diagnostics,
			filename,
			tokens[index],
			tokens[index+1],
			version,
			PHP85,
			"the pipe operator",
		)
		index++
	}
}

func prepareUnparenthesizedNewDereference(
	source []byte,
	tokens []*token.Token,
	version Version,
	filename string,
	diagnostics *[]Diagnostic,
) {
	for index := 0; index+3 < len(tokens); index++ {
		if tokens[index].ID != token.T_NEW || isMethodNameUse(tokens, index) {
			continue
		}

		openIndex := -1
		for cursor := index + 1; cursor < len(tokens); cursor++ {
			switch tokens[cursor].ID {
			case token.ID('('):
				openIndex = cursor
			case token.ID(';'), token.ID('{'), token.ID('}'):
				openIndex = -1
			}
			if openIndex >= 0 || tokens[cursor].ID == token.ID(';') {
				break
			}
		}
		if openIndex < 0 {
			continue
		}

		closeIndex := matchingToken(tokens, openIndex, token.ID('('), token.ID(')'))
		if closeIndex < 0 || closeIndex+1 >= len(tokens) {
			continue
		}

		next := tokens[closeIndex+1].ID
		if next != token.T_OBJECT_OPERATOR &&
			next != token.T_NULLSAFE_OBJECT_OPERATOR &&
			next != token.T_PAAMAYIM_NEKUDOTAYIM &&
			next != token.ID('[') {
			continue
		}

		replaceRange(source, tokenStart(tokens[index]), tokenEnd(tokens[closeIndex]), []byte("$x"))
		addFeatureDiagnostic(
			diagnostics,
			filename,
			tokens[index],
			tokens[closeIndex],
			version,
			PHP84,
			"unparenthesized new-expression dereferencing",
		)
		index = closeIndex
	}
}

func prepareVoidCasts(
	source []byte,
	tokens []*token.Token,
	version Version,
	filename string,
	diagnostics *[]Diagnostic,
) {
	for index := 0; index < len(tokens); index++ {
		if tokens[index].ID == token.T_VOID_CAST {
			addFeatureDiagnostic(
				diagnostics,
				filename,
				tokens[index],
				tokens[index],
				version,
				PHP85,
				"void casts",
			)
			continue
		}
		if index+3 >= len(tokens) ||
			tokens[index].ID != token.ID('(') ||
			!tokenIsWord(tokens[index+1], "void") ||
			tokens[index+2].ID != token.ID(')') ||
			!canStartExpression(tokens[index+3]) {
			continue
		}

		addFeatureDiagnostic(
			diagnostics,
			filename,
			tokens[index],
			tokens[index+2],
			version,
			PHP85,
			"void casts",
		)
		index += 2
	}
}

func prepareCloneWith(
	source []byte,
	tokens []*token.Token,
	version Version,
	filename string,
	diagnostics *[]Diagnostic,
) {
	for index := 0; index+3 < len(tokens); index++ {
		if tokens[index].ID != token.T_CLONE || tokens[index+1].ID != token.ID('(') {
			continue
		}

		if isMethodNameUse(tokens, index) {
			continue
		}

		closeIndex := matchingToken(tokens, index+1, token.ID('('), token.ID(')'))
		if closeIndex < 0 || !hasTopLevelComma(tokens, index+2, closeIndex) {
			continue
		}

		addFeatureDiagnostic(
			diagnostics,
			filename,
			tokens[index],
			tokens[closeIndex],
			version,
			PHP85,
			"clone-with",
		)
		index = closeIndex
	}
}

func prepareConstantAttributes(
	source []byte,
	tokens []*token.Token,
	layout sourceLayout,
	version Version,
	filename string,
	diagnostics *[]Diagnostic,
) {
	// Before PHP 8, "#[" starts a line comment rather than an attribute.
	if version < PHP80 {
		return
	}

	for index := 0; index < len(tokens); index++ {
		if tokens[index].ID != token.T_ATTRIBUTE {
			continue
		}

		closeIndex := matchingAttributeEnd(tokens, index)
		if closeIndex < 0 || closeIndex+1 >= len(tokens) ||
			tokens[closeIndex+1].ID != token.T_CONST ||
			layout.classMember[closeIndex+1] {
			continue
		}

		blankRange(source, tokenStart(tokens[index]), tokenEnd(tokens[closeIndex]))
		addFeatureDiagnostic(
			diagnostics,
			filename,
			tokens[index],
			tokens[closeIndex],
			version,
			PHP85,
			"attributes on constants",
		)
		index = closeIndex
	}
}

func prepareFinalPromotedProperties(
	source []byte,
	tokens []*token.Token,
	layout sourceLayout,
	version Version,
	filename string,
	diagnostics *[]Diagnostic,
) {
	for index, current := range tokens {
		if current.ID != token.T_FINAL || !layout.functionParam[index] {
			continue
		}

		end := findBefore(tokens, index+1, token.T_VARIABLE, token.ID(','), token.ID(')'))
		if end < 0 || tokens[end].ID != token.T_VARIABLE {
			continue
		}

		hasVisibility := false
		for cursor := index + 1; cursor < end; cursor++ {
			if isVisibility(tokens[cursor].ID) {
				hasVisibility = true
				break
			}
		}
		if !hasVisibility {
			continue
		}

		blankToken(source, current)
		addFeatureDiagnostic(
			diagnostics,
			filename,
			current,
			current,
			version,
			PHP85,
			"final promoted properties",
		)
	}
}

func addFeatureDiagnostic(
	diagnostics *[]Diagnostic,
	filename string,
	startToken, endToken *token.Token,
	actual, introduced Version,
	feature string,
) {
	if actual >= introduced {
		return
	}

	start, _ := positionFromInternal(startToken.Position)
	_, end := positionFromInternal(endToken.Position)
	*diagnostics = append(*diagnostics, Diagnostic{
		Filename: filename,
		Message:  feature + " require PHP " + introduced.String(),
		Phase:    PhaseCompile,
		Start:    start,
		End:      end,
	})
}

func findBefore(tokens []*token.Token, start int, wanted token.ID, stops ...token.ID) int {
	for index := start; index < len(tokens); index++ {
		for _, stop := range stops {
			if tokens[index].ID == stop {
				return -1
			}
		}
		if tokens[index].ID == wanted {
			return index
		}
	}
	return -1
}

func matchingHookBlock(tokens []*token.Token, start int) int {
	depth := 0
	for index := start; index < len(tokens); index++ {
		switch tokens[index].ID {
		case token.ID('{'), token.T_PROPERTY_HOOKS, token.T_CURLY_OPEN, token.T_DOLLAR_OPEN_CURLY_BRACES:
			depth++
		case token.ID('}'):
			depth--
			if depth == 0 {
				return index
			}
		}
	}
	return -1
}

func matchingToken(tokens []*token.Token, start int, open, close token.ID) int {
	depth := 0
	for index := start; index < len(tokens); index++ {
		id := tokens[index].ID
		if id == open || (open == token.ID('{') && isInterpolatedCurly(id)) {
			depth++
			continue
		}
		if id == close {
			depth--
			if depth == 0 {
				return index
			}
		}
	}
	return -1
}

func isInterpolatedCurly(id token.ID) bool {
	return id == token.T_CURLY_OPEN || id == token.T_DOLLAR_OPEN_CURLY_BRACES
}

func matchingAttributeEnd(tokens []*token.Token, start int) int {
	depth := 1
	for index := start + 1; index < len(tokens); index++ {
		switch tokens[index].ID {
		case token.ID('['):
			depth++
		case token.ID(']'):
			depth--
			if depth == 0 {
				return index
			}
		}
	}
	return -1
}

func hasTopLevelComma(tokens []*token.Token, start, end int) bool {
	paren, square, curly := 0, 0, 0
	for index := start; index < end; index++ {
		switch tokens[index].ID {
		case token.ID('('):
			paren++
		case token.ID(')'):
			paren--
		case token.ID('['):
			square++
		case token.ID(']'):
			square--
		case token.ID('{'):
			curly++
		case token.ID('}'):
			curly--
		case token.ID(','):
			if paren == 0 && square == 0 && curly == 0 {
				return true
			}
		}
	}
	return false
}

func blankToken(source []byte, current *token.Token) {}

func blankRange(source []byte, start, end int) {}

func replaceRange(source []byte, start, end int, replacement []byte) {}

func tokenStart(current *token.Token) int {
	if current == nil || current.Position == nil {
		return 0
	}
	return current.Position.StartPos
}

func tokenEnd(current *token.Token) int {
	if current == nil || current.Position == nil {
		return 0
	}
	return current.Position.EndPos
}

func tokenIsWord(current *token.Token, word string) bool {
	return current != nil && strings.EqualFold(string(current.Value), word)
}

func isIdentifierToken(current *token.Token) bool {
	if current == nil {
		return false
	}
	switch current.ID {
	case token.T_STRING,
		token.T_NAME_FULLY_QUALIFIED,
		token.T_NAME_QUALIFIED,
		token.T_NAME_RELATIVE:
		return true
	default:
		return current.ID >= token.T_INCLUDE && current.ID <= token.T_ENUM &&
			len(current.Value) > 0
	}
}

func isVisibility(id token.ID) bool {
	return id == token.T_PUBLIC || id == token.T_PROTECTED || id == token.T_PRIVATE ||
		id == token.T_PUBLIC_SET || id == token.T_PROTECTED_SET || id == token.T_PRIVATE_SET
}

// isReadonlyToken reports whether the token is the `readonly` keyword. The
// unified lexer emits it as a plain identifier on PHP 7 profiles, where it
// is not yet reserved.
func isReadonlyToken(current *token.Token) bool {
	if current == nil {
		return false
	}
	return current.ID == token.T_READONLY ||
		(current.ID == token.T_STRING && tokenIsWord(current, "readonly"))
}

func canStartExpression(current *token.Token) bool {
	if current == nil {
		return false
	}
	switch current.ID {
	case token.T_VARIABLE,
		token.T_LNUMBER,
		token.T_DNUMBER,
		token.T_CONSTANT_ENCAPSED_STRING,
		token.T_STRING,
		token.T_NAME_FULLY_QUALIFIED,
		token.T_NAME_QUALIFIED,
		token.T_NAME_RELATIVE,
		token.T_NEW,
		token.T_CLONE,
		token.T_FUNCTION,
		token.T_FN,
		token.ID('('),
		token.ID('['):
		return true
	default:
		return false
	}
}
