package ast

import (
	"github.com/shyim/go-phplint/internal/position"
	"github.com/shyim/go-phplint/internal/token"
)

//go:generate go run node_funcs_gen.go

// Root node
type Root struct {
	Position *position.Position
	Stmts    []Vertex
	EndTkn   *token.Token
}

// Nullable node is ?Expr
type Nullable struct {
	Position    *position.Position
	QuestionTkn *token.Token
	Expr        Vertex
}

// Union node is Expr|Expr1|...
type Union struct {
	Position      *position.Position
	Types         []Vertex
	SeparatorTkns []*token.Token
}

// Intersection node is Expr&Expr1&...
type Intersection struct {
	Position      *position.Position
	Types         []Vertex
	SeparatorTkns []*token.Token
}

// Parameter node
type Parameter struct {
	Position     *position.Position
	AttrGroups   []Vertex
	Modifiers    []Vertex
	Type         Vertex
	AmpersandTkn *token.Token
	VariadicTkn  *token.Token
	Var          Vertex
	EqualTkn     *token.Token
	DefaultValue Vertex
}

// Identifier node
type Identifier struct {
	Position      *position.Position
	IdentifierTkn *token.Token
	Value         []byte
}

// Argument node
type Argument struct {
	Position     *position.Position
	Name         Vertex
	ColonTkn     *token.Token
	VariadicTkn  *token.Token
	AmpersandTkn *token.Token
	Expr         Vertex
}

// Attribute node
type Attribute struct {
	Position            *position.Position
	Name                Vertex
	OpenParenthesisTkn  *token.Token
	Args                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

// AttributeGroup node
type AttributeGroup struct {
	Position          *position.Position
	OpenAttributeTkn  *token.Token
	Attrs             []Vertex
	SeparatorTkns     []*token.Token
	CloseAttributeTkn *token.Token
}

// ScalarDnumber node
type ScalarDnumber struct {
	Position  *position.Position
	NumberTkn *token.Token
	Value     []byte
}

// ScalarEncapsed node
type ScalarEncapsed struct {
	Position      *position.Position
	OpenQuoteTkn  *token.Token
	Parts         []Vertex
	CloseQuoteTkn *token.Token
}

// ScalarEncapsedStringPart node
type ScalarEncapsedStringPart struct {
	Position       *position.Position
	EncapsedStrTkn *token.Token
	Value          []byte
}

// ScalarEncapsedStringVar node
type ScalarEncapsedStringVar struct {
	Position                  *position.Position
	DollarOpenCurlyBracketTkn *token.Token
	Name                      Vertex
	OpenSquareBracketTkn      *token.Token
	Dim                       Vertex
	CloseSquareBracketTkn     *token.Token
	CloseCurlyBracketTkn      *token.Token
}

// ScalarEncapsedStringVar node
type ScalarEncapsedStringBrackets struct {
	Position             *position.Position
	OpenCurlyBracketTkn  *token.Token
	Var                  Vertex
	CloseCurlyBracketTkn *token.Token
}

// ScalarHeredoc node
type ScalarHeredoc struct {
	Position        *position.Position
	OpenHeredocTkn  *token.Token
	Parts           []Vertex
	CloseHeredocTkn *token.Token
}

// ScalarLnumber node
type ScalarLnumber struct {
	Position  *position.Position
	NumberTkn *token.Token
	Value     []byte
}

// ScalarMagicConstant node
type ScalarMagicConstant struct {
	Position      *position.Position
	MagicConstTkn *token.Token
	Value         []byte
}

// ScalarString node
type ScalarString struct {
	Position  *position.Position
	MinusTkn  *token.Token
	StringTkn *token.Token
	Value     []byte
}

// StmtBreak node
type StmtBreak struct {
	Position     *position.Position
	BreakTkn     *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

// StmtCase node
type StmtCase struct {
	Position         *position.Position
	CaseTkn          *token.Token
	Cond             Vertex
	CaseSeparatorTkn *token.Token
	Stmts            []Vertex
}

// StmtCatch node
type StmtCatch struct {
	Position             *position.Position
	CatchTkn             *token.Token
	OpenParenthesisTkn   *token.Token
	Types                []Vertex
	SeparatorTkns        []*token.Token
	Var                  Vertex
	CloseParenthesisTkn  *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

// StmtEnum node
type StmtEnum struct {
	Position                *position.Position
	AttrGroups              []Vertex
	EnumTkn                 *token.Token
	Name                    Vertex
	ColonTkn                *token.Token
	Type                    Vertex
	ImplementsTkn           *token.Token
	Implements              []Vertex
	ImplementsSeparatorTkns []*token.Token
	OpenCurlyBracketTkn     *token.Token
	Stmts                   []Vertex
	CloseCurlyBracketTkn    *token.Token
}

// EnumCase node
type EnumCase struct {
	Position     *position.Position
	AttrGroups   []Vertex
	CaseTkn      *token.Token
	Name         Vertex
	EqualTkn     *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

// StmtClass node
type StmtClass struct {
	Position                *position.Position
	AttrGroups              []Vertex
	Modifiers               []Vertex
	ClassTkn                *token.Token
	Name                    Vertex
	OpenParenthesisTkn      *token.Token
	Args                    []Vertex
	SeparatorTkns           []*token.Token
	CloseParenthesisTkn     *token.Token
	ExtendsTkn              *token.Token
	Extends                 Vertex
	ImplementsTkn           *token.Token
	Implements              []Vertex
	ImplementsSeparatorTkns []*token.Token
	OpenCurlyBracketTkn     *token.Token
	Stmts                   []Vertex
	CloseCurlyBracketTkn    *token.Token
}

// StmtClassConstList node
type StmtClassConstList struct {
	Position      *position.Position
	AttrGroups    []Vertex
	Modifiers     []Vertex
	ConstTkn      *token.Token
	Consts        []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// StmtClassMethod node
type StmtClassMethod struct {
	Position            *position.Position
	AttrGroups          []Vertex
	Modifiers           []Vertex
	FunctionTkn         *token.Token
	AmpersandTkn        *token.Token
	Name                Vertex
	OpenParenthesisTkn  *token.Token
	Params              []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	ReturnType          Vertex
	Stmt                Vertex
}

// StmtConstList node
type StmtConstList struct {
	Position      *position.Position
	ConstTkn      *token.Token
	Consts        []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// StmtConstant node
type StmtConstant struct {
	Position *position.Position
	Name     Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// StmtContinue node
type StmtContinue struct {
	Position     *position.Position
	ContinueTkn  *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

// StmtDeclare node
type StmtDeclare struct {
	Position            *position.Position
	DeclareTkn          *token.Token
	OpenParenthesisTkn  *token.Token
	Consts              []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndDeclareTkn       *token.Token
	SemiColonTkn        *token.Token
}

// StmtDefault node
type StmtDefault struct {
	Position         *position.Position
	DefaultTkn       *token.Token
	CaseSeparatorTkn *token.Token
	Stmts            []Vertex
}

// StmtDo node
type StmtDo struct {
	Position            *position.Position
	DoTkn               *token.Token
	Stmt                Vertex
	WhileTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

// StmtEcho node
type StmtEcho struct {
	Position      *position.Position
	EchoTkn       *token.Token
	Exprs         []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// StmtElse node
type StmtElse struct {
	Position *position.Position
	ElseTkn  *token.Token
	ColonTkn *token.Token
	Stmt     Vertex
}

// StmtElseIf node
type StmtElseIf struct {
	Position            *position.Position
	ElseIfTkn           *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
}

// StmtExpression node
type StmtExpression struct {
	Position     *position.Position
	Expr         Vertex
	SemiColonTkn *token.Token
}

// StmtFinally node
type StmtFinally struct {
	Position             *position.Position
	FinallyTkn           *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

// StmtFor node
type StmtFor struct {
	Position            *position.Position
	ForTkn              *token.Token
	OpenParenthesisTkn  *token.Token
	Init                []Vertex
	InitSeparatorTkns   []*token.Token
	InitSemiColonTkn    *token.Token
	Cond                []Vertex
	CondSeparatorTkns   []*token.Token
	CondSemiColonTkn    *token.Token
	Loop                []Vertex
	LoopSeparatorTkns   []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndForTkn           *token.Token
	SemiColonTkn        *token.Token
}

// StmtForeach node
type StmtForeach struct {
	Position            *position.Position
	ForeachTkn          *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	AsTkn               *token.Token
	Key                 Vertex
	DoubleArrowTkn      *token.Token
	AmpersandTkn        *token.Token
	Var                 Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndForeachTkn       *token.Token
	SemiColonTkn        *token.Token
}

// StmtFunction node
type StmtFunction struct {
	Position             *position.Position
	AttrGroups           []Vertex
	FunctionTkn          *token.Token
	AmpersandTkn         *token.Token
	Name                 Vertex
	OpenParenthesisTkn   *token.Token
	Params               []Vertex
	SeparatorTkns        []*token.Token
	CloseParenthesisTkn  *token.Token
	ColonTkn             *token.Token
	ReturnType           Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

// StmtGlobal node
type StmtGlobal struct {
	Position      *position.Position
	GlobalTkn     *token.Token
	Vars          []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// StmtGoto node
type StmtGoto struct {
	Position     *position.Position
	GotoTkn      *token.Token
	Label        Vertex
	SemiColonTkn *token.Token
}

// StmtHaltCompiler node
type StmtHaltCompiler struct {
	Position            *position.Position
	HaltCompilerTkn     *token.Token
	OpenParenthesisTkn  *token.Token
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

// StmtIf node
type StmtIf struct {
	Position            *position.Position
	IfTkn               *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	ElseIf              []Vertex
	Else                Vertex
	EndIfTkn            *token.Token
	SemiColonTkn        *token.Token
}

// StmtInlineHtml node
type StmtInlineHtml struct {
	Position      *position.Position
	InlineHtmlTkn *token.Token
	Value         []byte
}

// StmtInterface node
type StmtInterface struct {
	Position             *position.Position
	AttrGroups           []Vertex
	InterfaceTkn         *token.Token
	Name                 Vertex
	ExtendsTkn           *token.Token
	Extends              []Vertex
	ExtendsSeparatorTkns []*token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

// StmtLabel node
type StmtLabel struct {
	Position *position.Position
	Name     Vertex
	ColonTkn *token.Token
}

// StmtNamespace node
type StmtNamespace struct {
	Position             *position.Position
	NsTkn                *token.Token
	Name                 Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
	SemiColonTkn         *token.Token
}

// StmtNop node
type StmtNop struct {
	Position     *position.Position
	SemiColonTkn *token.Token
}

// StmtProperty node
type StmtProperty struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// StmtPropertyList node
type StmtPropertyList struct {
	Position      *position.Position
	AttrGroups    []Vertex
	Modifiers     []Vertex
	Type          Vertex
	Props         []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// StmtReturn node
type StmtReturn struct {
	Position     *position.Position
	ReturnTkn    *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

// StmtStatic node
type StmtStatic struct {
	Position      *position.Position
	StaticTkn     *token.Token
	Vars          []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// StmtStaticVar node
type StmtStaticVar struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// StmtStmtList node
type StmtStmtList struct {
	Position             *position.Position
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

// StmtSwitch node
type StmtSwitch struct {
	Position             *position.Position
	SwitchTkn            *token.Token
	OpenParenthesisTkn   *token.Token
	Cond                 Vertex
	CloseParenthesisTkn  *token.Token
	ColonTkn             *token.Token
	OpenCurlyBracketTkn  *token.Token
	CaseSeparatorTkn     *token.Token
	Cases                []Vertex
	CloseCurlyBracketTkn *token.Token
	EndSwitchTkn         *token.Token
	SemiColonTkn         *token.Token
}

// StmtThrow node
type StmtThrow struct {
	Position     *position.Position
	ThrowTkn     *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

// StmtTrait node
type StmtTrait struct {
	Position             *position.Position
	AttrGroups           []Vertex
	TraitTkn             *token.Token
	Name                 Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

// StmtTraitUse node
type StmtTraitUse struct {
	Position             *position.Position
	UseTkn               *token.Token
	Traits               []Vertex
	SeparatorTkns        []*token.Token
	OpenCurlyBracketTkn  *token.Token
	Adaptations          []Vertex
	CloseCurlyBracketTkn *token.Token
	SemiColonTkn         *token.Token
}

// StmtTraitUseAlias node
type StmtTraitUseAlias struct {
	Position       *position.Position
	Trait          Vertex
	DoubleColonTkn *token.Token
	Method         Vertex
	AsTkn          *token.Token
	Modifier       Vertex
	Alias          Vertex
	SemiColonTkn   *token.Token
}

// StmtTraitUsePrecedence node
type StmtTraitUsePrecedence struct {
	Position       *position.Position
	Trait          Vertex
	DoubleColonTkn *token.Token
	Method         Vertex
	InsteadofTkn   *token.Token
	Insteadof      []Vertex
	SeparatorTkns  []*token.Token
	SemiColonTkn   *token.Token
}

// StmtTry node
type StmtTry struct {
	Position             *position.Position
	TryTkn               *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
	Catches              []Vertex
	Finally              Vertex
}

// StmtUnset node
type StmtUnset struct {
	Position            *position.Position
	UnsetTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Vars                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

// StmtUseList node
type StmtUseList struct {
	Position      *position.Position
	UseTkn        *token.Token
	Type          Vertex
	Uses          []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// StmtGroupUseList node
type StmtGroupUseList struct {
	Position              *position.Position
	UseTkn                *token.Token
	Type                  Vertex
	LeadingNsSeparatorTkn *token.Token
	Prefix                Vertex
	NsSeparatorTkn        *token.Token
	OpenCurlyBracketTkn   *token.Token
	Uses                  []Vertex
	SeparatorTkns         []*token.Token
	CloseCurlyBracketTkn  *token.Token
	SemiColonTkn          *token.Token
}

// StmtUse node
type StmtUse struct {
	Position       *position.Position
	Type           Vertex
	NsSeparatorTkn *token.Token
	Use            Vertex
	AsTkn          *token.Token
	Alias          Vertex
}

// StmtWhile node
type StmtWhile struct {
	Position            *position.Position
	WhileTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndWhileTkn         *token.Token
	SemiColonTkn        *token.Token
}

// ExprArray node
type ExprArray struct {
	Position        *position.Position
	ArrayTkn        *token.Token
	OpenBracketTkn  *token.Token
	Items           []Vertex
	SeparatorTkns   []*token.Token
	CloseBracketTkn *token.Token
}

// ExprArrayDimFetch node
type ExprArrayDimFetch struct {
	Position        *position.Position
	Var             Vertex
	OpenBracketTkn  *token.Token
	Dim             Vertex
	CloseBracketTkn *token.Token
}

// ExprArrayItem node
type ExprArrayItem struct {
	Position       *position.Position
	EllipsisTkn    *token.Token
	Key            Vertex
	DoubleArrowTkn *token.Token
	AmpersandTkn   *token.Token
	Val            Vertex
}

// ExprArrowFunction node
type ExprArrowFunction struct {
	Position            *position.Position
	AttrGroups          []Vertex
	StaticTkn           *token.Token
	FnTkn               *token.Token
	AmpersandTkn        *token.Token
	OpenParenthesisTkn  *token.Token
	Params              []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	ReturnType          Vertex
	DoubleArrowTkn      *token.Token
	Expr                Vertex
}

// ExprBitwiseNot node
type ExprBitwiseNot struct {
	Position *position.Position
	TildaTkn *token.Token
	Expr     Vertex
}

// ExprBooleanNot node
type ExprBooleanNot struct {
	Position       *position.Position
	ExclamationTkn *token.Token
	Expr           Vertex
}

type ExprBrackets struct {
	Position            *position.Position
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

// ExprClassConstFetch node
type ExprClassConstFetch struct {
	Position       *position.Position
	Class          Vertex
	DoubleColonTkn *token.Token
	Const          Vertex
}

// ExprClone node
type ExprClone struct {
	Position *position.Position
	CloneTkn *token.Token
	Expr     Vertex
}

// ExprClosure node
type ExprClosure struct {
	Position               *position.Position
	AttrGroups             []Vertex
	StaticTkn              *token.Token
	FunctionTkn            *token.Token
	AmpersandTkn           *token.Token
	OpenParenthesisTkn     *token.Token
	Params                 []Vertex
	SeparatorTkns          []*token.Token
	CloseParenthesisTkn    *token.Token
	UseTkn                 *token.Token
	UseOpenParenthesisTkn  *token.Token
	Uses                   []Vertex
	UseSeparatorTkns       []*token.Token
	UseCloseParenthesisTkn *token.Token
	ColonTkn               *token.Token
	ReturnType             Vertex
	OpenCurlyBracketTkn    *token.Token
	Stmts                  []Vertex
	CloseCurlyBracketTkn   *token.Token
}

// ExprClosureUse node
type ExprClosureUse struct {
	Position     *position.Position
	AmpersandTkn *token.Token
	Var          Vertex
}

// ExprConstFetch node
type ExprConstFetch struct {
	Position *position.Position
	Const    Vertex
}

// ExprEmpty node
type ExprEmpty struct {
	Position            *position.Position
	EmptyTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

// ExprErrorSuppress node
type ExprErrorSuppress struct {
	Position *position.Position
	AtTkn    *token.Token
	Expr     Vertex
}

// ExprEval node
type ExprEval struct {
	Position            *position.Position
	EvalTkn             *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

// ExprExit node
type ExprExit struct {
	Position            *position.Position
	ExitTkn             *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

// ExprFunctionCall node
type ExprFunctionCall struct {
	Position            *position.Position
	Function            Vertex
	OpenParenthesisTkn  *token.Token
	Args                []Vertex
	SeparatorTkns       []*token.Token
	EllipsisTkn         *token.Token
	CloseParenthesisTkn *token.Token
}

// ExprInclude node
type ExprInclude struct {
	Position   *position.Position
	IncludeTkn *token.Token
	Expr       Vertex
}

// ExprIncludeOnce node
type ExprIncludeOnce struct {
	Position       *position.Position
	IncludeOnceTkn *token.Token
	Expr           Vertex
}

// ExprInstanceOf node
type ExprInstanceOf struct {
	Position      *position.Position
	Expr          Vertex
	InstanceOfTkn *token.Token
	Class         Vertex
}

// ExprIsset node
type ExprIsset struct {
	Position            *position.Position
	IssetTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Vars                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

// ExprList node
type ExprList struct {
	Position        *position.Position
	ListTkn         *token.Token
	OpenBracketTkn  *token.Token
	Items           []Vertex
	SeparatorTkns   []*token.Token
	CloseBracketTkn *token.Token
}

// ExprMethodCall node
type ExprMethodCall struct {
	Position             *position.Position
	Var                  Vertex
	ObjectOperatorTkn    *token.Token
	OpenCurlyBracketTkn  *token.Token
	Method               Vertex
	CloseCurlyBracketTkn *token.Token
	OpenParenthesisTkn   *token.Token
	Args                 []Vertex
	SeparatorTkns        []*token.Token
	EllipsisTkn          *token.Token
	CloseParenthesisTkn  *token.Token
}

// ExprNullsafeMethodCall node is $a?->methodName()
type ExprNullsafeMethodCall struct {
	Position             *position.Position
	Var                  Vertex
	ObjectOperatorTkn    *token.Token
	OpenCurlyBracketTkn  *token.Token
	Method               Vertex
	CloseCurlyBracketTkn *token.Token
	OpenParenthesisTkn   *token.Token
	Args                 []Vertex
	SeparatorTkns        []*token.Token
	EllipsisTkn          *token.Token
	CloseParenthesisTkn  *token.Token
}

// ExprNew node
type ExprNew struct {
	Position            *position.Position
	NewTkn              *token.Token
	Class               Vertex
	OpenParenthesisTkn  *token.Token
	Args                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

// ExprPostDec node
type ExprPostDec struct {
	Position *position.Position
	Var      Vertex
	DecTkn   *token.Token
}

// ExprPostInc node
type ExprPostInc struct {
	Position *position.Position
	Var      Vertex
	IncTkn   *token.Token
}

// ExprPreDec node
type ExprPreDec struct {
	Position *position.Position
	DecTkn   *token.Token
	Var      Vertex
}

// ExprPreInc node
type ExprPreInc struct {
	Position *position.Position
	IncTkn   *token.Token
	Var      Vertex
}

// ExprPrint node
type ExprPrint struct {
	Position *position.Position
	PrintTkn *token.Token
	Expr     Vertex
}

// ExprPropertyFetch node
type ExprPropertyFetch struct {
	Position             *position.Position
	Var                  Vertex
	ObjectOperatorTkn    *token.Token
	OpenCurlyBracketTkn  *token.Token
	Prop                 Vertex
	CloseCurlyBracketTkn *token.Token
}

// ExprNullsafePropertyFetch node
type ExprNullsafePropertyFetch struct {
	Position             *position.Position
	Var                  Vertex
	ObjectOperatorTkn    *token.Token
	OpenCurlyBracketTkn  *token.Token
	Prop                 Vertex
	CloseCurlyBracketTkn *token.Token
}

// ExprRequire node
type ExprRequire struct {
	Position   *position.Position
	RequireTkn *token.Token
	Expr       Vertex
}

// ExprRequireOnce node
type ExprRequireOnce struct {
	Position       *position.Position
	RequireOnceTkn *token.Token
	Expr           Vertex
}

// ExprShellExec node
type ExprShellExec struct {
	Position         *position.Position
	OpenBacktickTkn  *token.Token
	Parts            []Vertex
	CloseBacktickTkn *token.Token
}

// ExprStaticCall node
type ExprStaticCall struct {
	Position             *position.Position
	Class                Vertex
	DoubleColonTkn       *token.Token
	OpenCurlyBracketTkn  *token.Token
	Call                 Vertex
	CloseCurlyBracketTkn *token.Token
	OpenParenthesisTkn   *token.Token
	Args                 []Vertex
	SeparatorTkns        []*token.Token
	EllipsisTkn          *token.Token
	CloseParenthesisTkn  *token.Token
}

// ExprStaticPropertyFetch node
type ExprStaticPropertyFetch struct {
	Position       *position.Position
	Class          Vertex
	DoubleColonTkn *token.Token
	Prop           Vertex
}

// ExprTernary node
type ExprTernary struct {
	Position    *position.Position
	Cond        Vertex
	QuestionTkn *token.Token
	IfTrue      Vertex
	ColonTkn    *token.Token
	IfFalse     Vertex
}

// ExprUnaryMinus node
type ExprUnaryMinus struct {
	Position *position.Position
	MinusTkn *token.Token
	Expr     Vertex
}

// ExprUnaryPlus node
type ExprUnaryPlus struct {
	Position *position.Position
	PlusTkn  *token.Token
	Expr     Vertex
}

// ExprVariable node
type ExprVariable struct {
	Position             *position.Position
	DollarTkn            *token.Token
	OpenCurlyBracketTkn  *token.Token
	Name                 Vertex
	CloseCurlyBracketTkn *token.Token
}

// ExprYield node
type ExprYield struct {
	Position       *position.Position
	YieldTkn       *token.Token
	Key            Vertex
	DoubleArrowTkn *token.Token
	Val            Vertex
}

// ExprYieldFrom node
type ExprYieldFrom struct {
	Position     *position.Position
	YieldFromTkn *token.Token
	Expr         Vertex
}

// ExprCastArray node
type ExprCastArray struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

// ExprCastBool node
type ExprCastBool struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

// ExprCastDouble node
type ExprCastDouble struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

// ExprCastInt node
type ExprCastInt struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

// ExprCastObject node
type ExprCastObject struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

// ExprCastString node
type ExprCastString struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

// ExprCastUnset node
type ExprCastUnset struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

// ExprAssign node
type ExprAssign struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignReference node
type ExprAssignReference struct {
	Position     *position.Position
	Var          Vertex
	EqualTkn     *token.Token
	AmpersandTkn *token.Token
	Expr         Vertex
}

// ExprAssignBitwiseAnd node
type ExprAssignBitwiseAnd struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignBitwiseOr node
type ExprAssignBitwiseOr struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignBitwiseXor node
type ExprAssignBitwiseXor struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignCoalesce node
type ExprAssignCoalesce struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignConcat node
type ExprAssignConcat struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignDiv node
type ExprAssignDiv struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignMinus node
type ExprAssignMinus struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignMod node
type ExprAssignMod struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignMul node
type ExprAssignMul struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignPlus node
type ExprAssignPlus struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignPow node
type ExprAssignPow struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignShiftLeft node
type ExprAssignShiftLeft struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprAssignShiftRight node
type ExprAssignShiftRight struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

// ExprBinaryBitwiseAnd node
type ExprBinaryBitwiseAnd struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryBitwiseOr node
type ExprBinaryBitwiseOr struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryBitwiseXor node
type ExprBinaryBitwiseXor struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryBooleanAnd node
type ExprBinaryBooleanAnd struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryBooleanOr node
type ExprBinaryBooleanOr struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryCoalesce node
type ExprBinaryCoalesce struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryConcat node
type ExprBinaryConcat struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryDiv node
type ExprBinaryDiv struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryEqual node
type ExprBinaryEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryGreater node
type ExprBinaryGreater struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryGreaterOrEqual node
type ExprBinaryGreaterOrEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryIdentical node
type ExprBinaryIdentical struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryLogicalAnd node
type ExprBinaryLogicalAnd struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryLogicalOr node
type ExprBinaryLogicalOr struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryLogicalXor node
type ExprBinaryLogicalXor struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryMinus node
type ExprBinaryMinus struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryMod node
type ExprBinaryMod struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryMul node
type ExprBinaryMul struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryNotEqual node
type ExprBinaryNotEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryNotIdentical node
type ExprBinaryNotIdentical struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryPlus node
type ExprBinaryPlus struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryPow node
type ExprBinaryPow struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryShiftLeft node
type ExprBinaryShiftLeft struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinaryShiftRight node
type ExprBinaryShiftRight struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinarySmaller node
type ExprBinarySmaller struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinarySmallerOrEqual node
type ExprBinarySmallerOrEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprBinarySpaceship node
type ExprBinarySpaceship struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

// ExprMatch node is match(expr) { list<MatchArm> }
type ExprMatch struct {
	Position             *position.Position
	MatchTkn             *token.Token
	OpenParenthesisTkn   *token.Token
	Expr                 Vertex
	CloseParenthesisTkn  *token.Token
	OpenCurlyBracketTkn  *token.Token
	Arms                 []Vertex
	SeparatorTkns        []*token.Token
	CloseCurlyBracketTkn *token.Token
}

// ExprThrow node is 'throw Expr'
type ExprThrow struct {
	Position     *position.Position
	ThrowTkn     *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

// MatchArm node is [expr, expr1, ...]|default => return_expr
type MatchArm struct {
	Position        *position.Position
	DefaultTkn      *token.Token
	DefaultCommaTkn *token.Token
	Exprs           []Vertex
	SeparatorTkns   []*token.Token
	DoubleArrowTkn  *token.Token
	ReturnExpr      Vertex
}

type Name struct {
	Position      *position.Position
	Parts         []Vertex
	SeparatorTkns []*token.Token
}

type NameFullyQualified struct {
	Position       *position.Position
	NsSeparatorTkn *token.Token
	Parts          []Vertex
	SeparatorTkns  []*token.Token
}

type NameRelative struct {
	Position       *position.Position
	NsTkn          *token.Token
	NsSeparatorTkn *token.Token
	Parts          []Vertex
	SeparatorTkns  []*token.Token
}

type NamePart struct {
	Position  *position.Position
	StringTkn *token.Token
	Value     []byte
}
