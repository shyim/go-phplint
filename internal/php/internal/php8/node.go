package php8

import (
	"github.com/shyim/go-phplint/internal/php/pkg/ast"
	"github.com/shyim/go-phplint/internal/php/pkg/position"
	"github.com/shyim/go-phplint/internal/php/pkg/token"
)

type ParserBrackets struct {
	Position        *position.Position
	OpenBracketTkn  *token.Token
	Child           ast.Vertex
	CloseBracketTkn *token.Token
}

func (n *ParserBrackets) Accept(v ast.Visitor) {
	// do nothing
}

func (n *ParserBrackets) GetPosition() *position.Position {
	return n.Position
}

func (n *ParserBrackets) GetType() ast.Type {
    return ast.TypeNone
}

type ParserSeparatedList struct {
	Position      *position.Position
	Items         []ast.Vertex
	SeparatorTkns []*token.Token
}

func (n *ParserSeparatedList) Accept(v ast.Visitor) {
	// do nothing
}

func (n *ParserSeparatedList) GetPosition() *position.Position {
	return n.Position
}

func (n *ParserSeparatedList) GetType() ast.Type {
    return ast.TypeNone
}

// TraitAdaptationList node
type TraitAdaptationList struct {
	Position             *position.Position
	OpenCurlyBracketTkn  *token.Token
	Adaptations          []ast.Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *TraitAdaptationList) Accept(v ast.Visitor) {
	// do nothing
}

func (n *TraitAdaptationList) GetPosition() *position.Position {
	return n.Position
}

func (n *TraitAdaptationList) GetType() ast.Type {
    return ast.TypeNone
}

// ArgumentList node
type ArgumentList struct {
	Position            *position.Position
	OpenParenthesisTkn  *token.Token
	Arguments           []ast.Vertex
	SeparatorTkns       []*token.Token
	EllipsisTkn         *token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ArgumentList) Accept(v ast.Visitor) {
	// do nothing
}

func (n *ArgumentList) GetPosition() *position.Position {
	return n.Position
}

func (n *ArgumentList) GetType() ast.Type {
    return ast.TypeNone
}

type EnumCaseExpr struct {
	Position  *position.Position
	AssignTkn *token.Token
	Expr      ast.Vertex
}

func (n *EnumCaseExpr) Accept(v ast.Visitor) {
	// do nothing
}

func (n *EnumCaseExpr) GetPosition() *position.Position {
	return n.Position
}

func (n *EnumCaseExpr) GetType() ast.Type {
    return ast.TypeNone
}

type ReturnType struct {
	Position *position.Position
	ColonTkn *token.Token
	Type     ast.Vertex
}

func (n *ReturnType) Accept(v ast.Visitor) {
	// do nothing
}

func (n *ReturnType) GetPosition() *position.Position {
	return n.Position
}

func (n *ReturnType) GetType() ast.Type {
    return ast.TypeNone
}

// TraitMethodRef node
type TraitMethodRef struct {
	Position       *position.Position
	Trait          ast.Vertex
	DoubleColonTkn *token.Token
	Method         ast.Vertex
}

func (n *TraitMethodRef) Accept(v ast.Visitor) {
	// do nothing
}

func (n *TraitMethodRef) GetPosition() *position.Position {
	return n.Position
}

func (n *TraitMethodRef) GetType() ast.Type {
    return ast.TypeNone
}
