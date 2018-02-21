package es6

// PrimaryExpressionNode [Yield] : [See 12.2]
//  this
//  IdentifierReference[?Yield]
//  Literal
//  ArrayLiteral[?Yield]
//  ObjectLiteral[?Yield]
//  FunctionExpression
//  ClassExpression[?Yield]
//  GeneratorExpression
//  RegularExpressionLiteral
//  TemplateLiteral[?Yield]
//  CoverParenthesizedExpressionAndArrowParameterList[?Yield]
// the interpretation of CoverParenthesizedExpressionAndArrowParameterList
// is refined using the following grammar:
// implements: Parser and ASTNode
type PrimaryExpressionNode struct {
	node
}

// ParsePrimaryExpressionNode ...
func ParsePrimaryExpressionNode(l *Lexer) (PrimaryExpressionNode, error) {
	panic("ParsePrimaryExpressionNode not implemented")
	// return nil, nil
}

// LabelIdentifierNode [Yield] : [See 12.1]
//  Identifier
//  [~Yield] yield
// implements: Parser and ASTNode
type LabelIdentifierNode struct {
	node
}

// ParseLabelIdentifierNode ...
func ParseLabelIdentifierNode(l *Lexer) (LabelIdentifierNode, error) {
	panic("ParseLabelIdentifierNode not implemented")
	// return nil, nil
}

// LiteralNode  : [See 12.2.4]
//  NullLiteral
//  BooleanLiteral
//  NumericLiteral
//  StringLiteral
// implements: Parser and ASTNode
type LiteralNode struct {
	node
}

// ParseLiteralNode ...
func ParseLiteralNode(l *Lexer) (LiteralNode, error) {
	panic("ParseLiteralNode not implemented")
	// return nil, nil
}

// ArrayLiteralNode [Yield] : [See 12.2.5]
//  [ Elisionopt ]
//  [ ElementList[?Yield] ]
//  [ ElementList[?Yield] , Elisionopt ]
// implements: Parser and ASTNode
type ArrayLiteralNode struct {
	node
}

// ParseArrayLiteralNode ...
func ParseArrayLiteralNode(l *Lexer) (ArrayLiteralNode, error) {
	panic("ParseArrayLiteralNode not implemented")
	// return nil, nil
}

// ObjectLiteralNode [Yield] : [See 12.2.6]
//  { }
//  { PropertyDefinitionList[?Yield] }
//  { PropertyDefinitionList[?Yield] , }
// implements: Parser and ASTNode
type ObjectLiteralNode struct {
	node
}

// ParseObjectLiteralNode ...
func ParseObjectLiteralNode(l *Lexer) (ObjectLiteralNode, error) {
	panic("ParseObjectLiteralNode not implemented")
	// return nil, nil
}

// FunctionExpressionNode  : [See 14.1]
//  function BindingIdentifieropt ( FormalParameters ) { FunctionBody }
// implements: Parser and ASTNode
type FunctionExpressionNode struct {
	node
}

// ParseFunctionExpressionNode ...
func ParseFunctionExpressionNode(l *Lexer) (FunctionExpressionNode, error) {
	panic("ParseFunctionExpressionNode not implemented")
	// return nil, nil
}

// ClassExpressionNode [Yield] : [See 14.5]
//  class BindingIdentifier[?Yield]opt ClassTail[?Yield]
// implements: Parser and ASTNode
type ClassExpressionNode struct {
	node
}

// ParseClassExpressionNode ...
func ParseClassExpressionNode(l *Lexer) (ClassExpressionNode, error) {
	panic("ParseClassExpressionNode not implemented")
	// return nil, nil
}

// CoverParenthesizedExpressionAndArrowParameterListNode [Yield] : [See 12.2]
// ( Expression[In, ?Yield] )
// ( )
// ( ... BindingIdentifier[?Yield] )
// ( Expression[In, ?Yield] , ... BindingIdentifier[?Yield] )
//  When processing the production
// implements: Parser and ASTNode
type CoverParenthesizedExpressionAndArrowParameterListNode struct {
	node
}

// GeneratorExpressionNode  : [See 14.4]
//  function * BindingIdentifier[Yield]opt ( FormalParameters[Yield] ) { GeneratorBody }
// implements: Parser and ASTNode
type GeneratorExpressionNode struct {
	node
}

// ParseGeneratorExpressionNode ...
func ParseGeneratorExpressionNode(l *Lexer) (GeneratorExpressionNode, error) {
	panic("ParseGeneratorExpressionNode not implemented")
	// return nil, nil
}

// GeneratorBodyNode  : [See 14.4]
//  FunctionBody[Yield]
// implements: Parser and ASTNode
type GeneratorBodyNode struct {
	node
}

// TemplateLiteralNode [Yield] : [See 12.2.9]
//  NoSubstitutionTemplate
//  TemplateHead Expression[In, ?Yield] TemplateSpans[?Yield]
// implements: Parser and ASTNode
type TemplateLiteralNode struct {
	node
}

// ParseTemplateLiteralNode ...
func ParseTemplateLiteralNode(l *Lexer) (TemplateLiteralNode, error) {
	panic("ParseTemplateLiteralNode not implemented")
	// return nil, nil
}
