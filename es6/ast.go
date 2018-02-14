package es6

import (
	"fmt"

	"github.com/pkg/errors"
)

// ASTNode ...
type ASTNode interface {
	Positioner
}

// ParseFunc ...
type parseFunc func(l *Lexer) (ASTNode, error)

// TryParseFuncs
func tryParseFuncs(l *Lexer, pfs ...parseFunc) (ASTNode, error) {
	var (
		child ASTNode
		err   error
	)
	for _, pf := range pfs {
		child, err = pf(l)
		if err != nil {
			if _, ok := err.(IncorrectTokenError); !ok {
				break
			}
		}
	}
	return child, err
}

// Positioner ...
type Positioner interface {
	Position() (filename string, offset int, line int, column int)
}

// Parser ...
type Parser interface {
	Parse(l *Lexer) (ASTNode, error)
}

type node struct {
	FilePosition
}

// IncorrectTokenError is returned when an unexpected token is found
// tokens recieved by function returning this error have only been peeked
// so that the calling function may try another Parse* function
type IncorrectTokenError Token

func (tok IncorrectTokenError) Error() string {
	return fmt.Sprintf("IncorrectTokenError at %s", tok.FilePosition.String())
}

//
//  A.2 Expressions
//

// IdentifierReferenceNode [Yield] : [See 12.1]
//  Identifier
//  [~Yield] yield
// implements: Parser and ASTNode
type IdentifierReferenceNode struct {
	node
}

// ParseIdentifierReferenceNode ...
func ParseIdentifierReferenceNode(l *Lexer) (ASTNode, error) {
	panic("ParseIdentifierReferenceNode not implemented")
	// return nil, nil
}

// BindingIdentifierNode [Yield] : [See 12.1]
//  Identifier
//  [~Yield] yield
//
// implements: Parser and ASTNode
type BindingIdentifierNode struct {
	node
}

// ParseBindingIdentifierNode ...
func ParseBindingIdentifierNode(l *Lexer) (ASTNode, error) {
	panic("ParseBindingIdentifierNode not implemented")
	// return nil, nil
}

// IdentifierNode  : [See 12.1]
//  IdentifierName but not ReservedWord
// implements: Parser and ASTNode
type IdentifierNode struct {
	node
	Name string
}

// ParseIdentifierNode ...
func ParseIdentifierNode(l *Lexer) (ASTNode, error) {
	n := IdentifierNode{node: node{l.CurrentPosition()}}
	tt := l.Next(l.goal)
	if tt.Type == ReservedWordToken {
		return n, errors.Errorf("IdentifierNode must not be a ReservedWordToken found %q", tt.Value)
	}
	n.Name = tt.Value
	return n, nil
}

// ParseCoverParenthesizedExpressionAndArrowParameterListNode ...
func ParseCoverParenthesizedExpressionAndArrowParameterListNode(l *Lexer) (ASTNode, error) {
	panic("ParseCoverParenthesizedExpressionAndArrowParameterListNode not implemented")
	// return nil, nil
}

// ParenthesizedExpressionNode [Yield] : [See 12.2]
//  ( Expression[In, ?Yield] )
// implements: Parser and ASTNode
type ParenthesizedExpressionNode struct {
	node
	ExpressionNode
}

// ParseParenthesizedExpressionNode ...
func ParseParenthesizedExpressionNode(l *Lexer) (ASTNode, error) {
	n := ParenthesizedExpressionNode{node: node{l.CurrentPosition()}}

	tt := l.Next(l.goal)
	if tt.Value != "(" {
		return n, errors.New("expected '('")
	}

	expressionNode, err := ParseExpressionNode(l)
	n.ExpressionNode = expressionNode.(ExpressionNode)
	if err != nil {
		return n, err
	}

	tt = l.Next(l.goal)
	if tt.Value != ")" {
		return n, errors.New("expected ')'")
	}

	return n, nil
}

// ElementListNode [Yield] : [See 12.2.5]
//  Elisionopt AssignmentExpression[In, ?Yield]
//  Elisionopt SpreadElement[?Yield]
//  ElementList[?Yield] , Elisionopt AssignmentExpression[In, ?Yield]
//  ElementList[?Yield] , Elisionopt SpreadElement[?Yield]
// implements: Parser and ASTNode
type ElementListNode struct {
	node
}

// ParseElementListNode ...
func ParseElementListNode(l *Lexer) (ASTNode, error) {
	panic("ParseElementListNode not implemented")
	// return nil, nil
}

// ElisionNode  : [See 12.2.5]
//  ,
//  Elision ,
// implements: Parser and ASTNode
type ElisionNode struct {
	node
}

// ParseElisionNode ...
func ParseElisionNode(l *Lexer) (ASTNode, error) {
	panic("ParseElisionNode not implemented")
	// return nil, nil
}

// SpreadElementNode [Yield] : [See 12.2.5]
//  ... AssignmentExpression[In, ?Yield]
// implements: Parser and ASTNode
type SpreadElementNode struct {
	node
}

// ParseSpreadElementNode ...
func ParseSpreadElementNode(l *Lexer) (ASTNode, error) {
	panic("ParseSpreadElementNode not implemented")
	// return nil, nil
}

// PropertyDefinitionListNode [Yield] : [See 12.2.6]
//  PropertyDefinition[?Yield]
//  PropertyDefinitionList[?Yield] , PropertyDefinition[?Yield]
// implements: Parser and ASTNode
type PropertyDefinitionListNode struct {
	node
}

// ParsePropertyDefinitionListNode ...
func ParsePropertyDefinitionListNode(l *Lexer) (ASTNode, error) {
	panic("ParsePropertyDefinitionListNode not implemented")
	// return nil, nil
}

// PropertyDefinitionNode [Yield] : [See 12.2.6]
//  IdentifierReference[?Yield]
//  CoverInitializedName[?Yield]
//  PropertyName[?Yield] : AssignmentExpression[In, ?Yield]
//  MethodDefinition[?Yield]
// implements: Parser and ASTNode
type PropertyDefinitionNode struct {
	node
}

// ParsePropertyDefinitionNode ...
func ParsePropertyDefinitionNode(l *Lexer) (ASTNode, error) {
	panic("ParsePropertyDefinitionNode not implemented")
	// return nil, nil
}

// PropertyNameNode [Yield] : [See 12.2.6]
//  LiteralPropertyName
//  ComputedPropertyName[?Yield]
// implements: Parser and ASTNode
type PropertyNameNode struct {
	node
}

// ParsePropertyNameNode ...
func ParsePropertyNameNode(l *Lexer) (ASTNode, error) {
	panic("ParsePropertyNameNode not implemented")
	// return nil, nil
}

// LiteralPropertyNameNode  : [See 12.2.6]
//  IdentifierName
//  StringLiteral
//  NumericLiteral
// implements: Parser and ASTNode
type LiteralPropertyNameNode struct {
	node
}

// ParseLiteralPropertyNameNode ...
func ParseLiteralPropertyNameNode(l *Lexer) (ASTNode, error) {
	panic("ParseLiteralPropertyNameNode not implemented")
	// return nil, nil
}

// ComputedPropertyNameNode [Yield] : [See 12.2.6]
//  [ AssignmentExpression[In, ?Yield] ]
// implements: Parser and ASTNode
type ComputedPropertyNameNode struct {
	node
}

// ParseComputedPropertyNameNode ...
func ParseComputedPropertyNameNode(l *Lexer) (ASTNode, error) {
	panic("ParseComputedPropertyNameNode not implemented")
	// return nil, nil
}

// CoverInitializedNameNode [Yield] : [See 12.2.6]
//  IdentifierReference[?Yield] Initializer[In, ?Yield]
// implements: Parser and ASTNode
type CoverInitializedNameNode struct {
	node
}

// ParseCoverInitializedNameNode ...
func ParseCoverInitializedNameNode(l *Lexer) (ASTNode, error) {
	panic("ParseCoverInitializedNameNode not implemented")
	// return nil, nil
}

// InitializerNode [In, Yield] : [See 12.2.6]
//  = AssignmentExpression[?In, ?Yield]
// implements: Parser and ASTNode
type InitializerNode struct {
	node
}

// ParseInitializerNode ...
func ParseInitializerNode(l *Lexer) (ASTNode, error) {
	panic("ParseInitializerNode not implemented")
	// return nil, nil
}

// TemplateSpansNode [Yield] : [See 12.2.9]
//  TemplateTail
//  TemplateMiddleList[?Yield] TemplateTail
// implements: Parser and ASTNode
type TemplateSpansNode struct {
	node
}

// ParseTemplateSpansNode ...
func ParseTemplateSpansNode(l *Lexer) (ASTNode, error) {
	panic("ParseTemplateSpansNode not implemented")
	// return nil, nil
}

// TemplateMiddleListNode [Yield] : [See 12.2.9]
//  TemplateMiddle Expression[In, ?Yield]
//  TemplateMiddleList[?Yield] TemplateMiddle Expression[In, ?Yield]
// implements: Parser and ASTNode
type TemplateMiddleListNode struct {
	node
}

// ParseTemplateMiddleListNode ...
func ParseTemplateMiddleListNode(l *Lexer) (ASTNode, error) {
	panic("ParseTemplateMiddleListNode not implemented")
	// return nil, nil
}

// MemberExpressionNode [Yield] : [See 12.3]
//  PrimaryExpression[?Yield]
//  MemberExpression[?Yield] [ Expression[In, ?Yield] ]
//  MemberExpression[?Yield] . IdentifierName
//  MemberExpression[?Yield] TemplateLiteral[?Yield]
//  SuperProperty[?Yield]
//  MetaProperty
//  new MemberExpression[?Yield] Arguments[?Yield]
// implements: Parser and ASTNode
type MemberExpressionNode struct {
	node
}

// ParseMemberExpressionNode ...
func ParseMemberExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseMemberExpressionNode not implemented")
	// return nil, nil
}

// SuperPropertyNode [Yield] : [See 12.3]
//  super [ Expression[In, ?Yield] ]
//  super . IdentifierName
// implements: Parser and ASTNode
type SuperPropertyNode struct {
	node
}

// ParseSuperPropertyNode ...
func ParseSuperPropertyNode(l *Lexer) (ASTNode, error) {
	panic("ParseSuperPropertyNode not implemented")
	// return nil, nil
}

// MetaPropertyNode  : [See 12.3]
//  NewTarget
// implements: Parser and ASTNode
type MetaPropertyNode struct {
	node
}

// ParseMetaPropertyNode ...
func ParseMetaPropertyNode(l *Lexer) (ASTNode, error) {
	panic("ParseMetaPropertyNode not implemented")
	// return nil, nil
}

// NewTargetNode  : [See 12.3]
//  new . target
// implements: Parser and ASTNode
type NewTargetNode struct {
	node
}

// ParseNewTargetNode ...
func ParseNewTargetNode(l *Lexer) (ASTNode, error) {
	panic("ParseNewTargetNode not implemented")
	// return nil, nil
}

// NewExpressionNode [Yield] : [See 12.3]
//  MemberExpression[?Yield]
//  new NewExpression[?Yield]
// implements: Parser and ASTNode
type NewExpressionNode struct {
	node
}

// ParseNewExpressionNode ...
func ParseNewExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseNewExpressionNode not implemented")
	// return nil, nil
}

// CallExpressionNode [Yield] : [See 12.3]
//  MemberExpression[?Yield] Arguments[?Yield]
//  SuperCall[?Yield]
//  CallExpression[?Yield] Arguments[?Yield]
//  CallExpression[?Yield] [ Expression[In, ?Yield] ]
//  CallExpression[?Yield] . IdentifierName
//  CallExpression[?Yield] TemplateLiteral[?Yield]
// implements: Parser and ASTNode
type CallExpressionNode struct {
	node
}

// ParseCallExpressionNode ...
func ParseCallExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseCallExpressionNode not implemented")
	// return nil, nil
}

// SuperCallNode [Yield] : [See 12.3]
//  super Arguments[?Yield]
// implements: Parser and ASTNode
type SuperCallNode struct {
	node
}

// ParseSuperCallNode ...
func ParseSuperCallNode(l *Lexer) (ASTNode, error) {
	panic("ParseSuperCallNode not implemented")
	// return nil, nil
}

// ArgumentsNode [Yield] : [See 12.3]
//  ( )
//  ( ArgumentList[?Yield] )
// implements: Parser and ASTNode
type ArgumentsNode struct {
	node
}

// ParseArgumentsNode ...
func ParseArgumentsNode(l *Lexer) (ASTNode, error) {
	panic("ParseArgumentsNode not implemented")
	// return nil, nil
}

// ArgumentListNode [Yield] : [See 12.3]
//  AssignmentExpression[In, ?Yield]
//  ... AssignmentExpression[In, ?Yield]
//  ArgumentList[?Yield] , AssignmentExpression[In, ?Yield]
//  ArgumentList[?Yield] , ... AssignmentExpression[In, ?Yield]
// implements: Parser and ASTNode
type ArgumentListNode struct {
	node
}

// ParseArgumentListNode ...
func ParseArgumentListNode(l *Lexer) (ASTNode, error) {
	panic("ParseArgumentListNode not implemented")
	// return nil, nil
}

// LeftHandSideExpressionNode [Yield] : [See 12.3]
//  NewExpression[?Yield]
//  CallExpression[?Yield]
// implements: Parser and ASTNode
type LeftHandSideExpressionNode struct {
	node
}

// ParseLeftHandSideExpressionNode ...
func ParseLeftHandSideExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseLeftHandSideExpressionNode not implemented")
	// return nil, nil
}

// PostfixExpressionNode [Yield] : [See 12.4]
//  LeftHandSideExpression[?Yield]
//  LeftHandSideExpression[?Yield] [no LineTerminator here] ++
//  LeftHandSideExpression[?Yield] [no LineTerminator here] --
// implements: Parser and ASTNode
type PostfixExpressionNode struct {
	node
}

// ParsePostfixExpressionNode ...
func ParsePostfixExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParsePostfixExpressionNode not implemented")
	// return nil, nil
}

// UnaryExpressionNode [Yield] : [See 12.5]
//  PostfixExpression[?Yield]
//  delete UnaryExpression[?Yield]
//  void UnaryExpression[?Yield]
//  typeof UnaryExpression[?Yield]
//  ++ UnaryExpression[?Yield]
//  -- UnaryExpression[?Yield]
//  + UnaryExpression[?Yield]
//  - UnaryExpression[?Yield]
//  ~ UnaryExpression[?Yield]
//  ! UnaryExpression[?Yield]
// implements: Parser and ASTNode
type UnaryExpressionNode struct {
	node
}

// ParseUnaryExpressionNode ...
func ParseUnaryExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseUnaryExpressionNode not implemented")
	// return nil, nil
}

// MultiplicativeExpressionNode [Yield] : [See 12.6]
//  UnaryExpression[?Yield]
//  MultiplicativeExpression[?Yield] MultiplicativeOperator UnaryExpression[?Yield]
// implements: Parser and ASTNode
type MultiplicativeExpressionNode struct {
	node
}

// ParseMultiplicativeExpressionNode ...
func ParseMultiplicativeExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseMultiplicativeExpressionNode not implemented")
	// return nil, nil
}

// MultiplicativeOperatorNode  : one of [See 12.6]
//  * / %
// implements: Parser and ASTNode
type MultiplicativeOperatorNode struct {
	node
}

// ParseMultiplicativeOperatorNode ...
func ParseMultiplicativeOperatorNode(l *Lexer) (ASTNode, error) {
	panic("ParseMultiplicativeOperatorNode not implemented")
	// return nil, nil
}

// AdditiveExpressionNode [Yield] : [See 12.7]
//  MultiplicativeExpression[?Yield]
//  AdditiveExpression[?Yield] + MultiplicativeExpression[?Yield]
//  AdditiveExpression[?Yield] - MultiplicativeExpression[?Yield]
// implements: Parser and ASTNode
type AdditiveExpressionNode struct {
	node
}

// ParseAdditiveExpressionNode ...
func ParseAdditiveExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseAdditiveExpressionNode not implemented")
	// return nil, nil
}

// ShiftExpressionNode [Yield] : [See 12.8]
//  AdditiveExpression[?Yield]
//  ShiftExpression[?Yield] << AdditiveExpression[?Yield]
//  ShiftExpression[?Yield] >> AdditiveExpression[?Yield]
//  ShiftExpression[?Yield] >>> AdditiveExpression[?Yield]
// implements: Parser and ASTNode
type ShiftExpressionNode struct {
	node
}

// ParseShiftExpressionNode ...
func ParseShiftExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseShiftExpressionNode not implemented")
	// return nil, nil
}

// RelationalExpressionNode [In, Yield] : [See 12.9]
//  ShiftExpression[?Yield]
//  RelationalExpression[?In, ?Yield] < ShiftExpression[?Yield]
//  RelationalExpression[?In, ?Yield] > ShiftExpression[?Yield]
//  RelationalExpression[?In, ?Yield] <= ShiftExpression[? Yield]
//  RelationalExpression[?In, ?Yield] >= ShiftExpression[?Yield]
//  RelationalExpression[?In, ?Yield] instanceof ShiftExpression[?Yield]
//  [+In] RelationalExpression[In, ?Yield] in ShiftExpression[?Yield]
// implements: Parser and ASTNode
type RelationalExpressionNode struct {
	node
}

// ParseRelationalExpressionNode ...
func ParseRelationalExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseRelationalExpressionNode not implemented")
	// return nil, nil
}

// EqualityExpressionNode [In, Yield] : [See 12.10]
//  RelationalExpression[?In, ?Yield]
//  EqualityExpression[?In, ?Yield] == RelationalExpression[?In, ?Yield]
//  EqualityExpression[?In, ?Yield] != RelationalExpression[?In, ?Yield]
//  EqualityExpression[?In, ?Yield] === RelationalExpression[?In, ?Yield]
//  EqualityExpression[?In, ?Yield] !== RelationalExpression[?In, ?Yield]
// implements: Parser and ASTNode
type EqualityExpressionNode struct {
	node
}

// ParseEqualityExpressionNode ...
func ParseEqualityExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseEqualityExpressionNode not implemented")
	// return nil, nil
}

// BitwiseANDExpressionNode [In, Yield] : [See 12.11]
//  EqualityExpression[?In, ?Yield]
//  BitwiseANDExpression[?In, ?Yield] & EqualityExpression[?In, ?Yield]
// implements: Parser and ASTNode
type BitwiseANDExpressionNode struct {
	node
}

// ParseBitwiseANDExpressionNode ...
func ParseBitwiseANDExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseBitwiseANDExpressionNode not implemented")
	// return nil, nil
}

// BitwiseXORExpressionNode [In, Yield] : [See 12.11]
//  BitwiseANDExpression[?In, ?Yield]
//  BitwiseXORExpression[?In, ?Yield] ^ BitwiseANDExpression[?In, ?Yield]
// implements: Parser and ASTNode
type BitwiseXORExpressionNode struct {
	node
}

// ParseBitwiseXORExpressionNode ...
func ParseBitwiseXORExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseBitwiseXORExpressionNode not implemented")
	// return nil, nil
}

// BitwiseORExpressionNode [In, Yield] : [See 12.11]
//  BitwiseXORExpression[?In, ?Yield]
//  BitwiseORExpression[?In, ?Yield] | BitwiseXORExpression[?In, ?Yield]
// implements: Parser and ASTNode
type BitwiseORExpressionNode struct {
	node
}

// ParseBitwiseORExpressionNode ...
func ParseBitwiseORExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseBitwiseORExpressionNode not implemented")
	// return nil, nil
}

// LogicalANDExpressionNode [In, Yield] : [See 12.12]
//  BitwiseORExpression[?In, ?Yield]
//  LogicalANDExpression[?In, ?Yield] && BitwiseORExpression[?In, ?Yield]
// implements: Parser and ASTNode
type LogicalANDExpressionNode struct {
	node
}

// ParseLogicalANDExpressionNode ...
func ParseLogicalANDExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseLogicalANDExpressionNode not implemented")
	// return nil, nil
}

// LogicalORExpressionNode [In, Yield] : [See 12.12]
//  LogicalANDExpression[?In, ?Yield]
//  LogicalORExpression[?In, ?Yield] || LogicalANDExpression[?In, ?Yield]
// implements: Parser and ASTNode
type LogicalORExpressionNode struct {
	node
}

// ParseLogicalORExpressionNode ...
func ParseLogicalORExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseLogicalORExpressionNode not implemented")
	// return nil, nil
}

// ConditionalExpressionNode [In, Yield] : [See 12.13]
//  LogicalORExpression[?In, ?Yield]
//  LogicalORExpression[?In,?Yield] ? AssignmentExpression[In, ?Yield] : AssignmentExpression[?In, ?Yield]
// implements: Parser and ASTNode
type ConditionalExpressionNode struct {
	node
}

// ParseConditionalExpressionNode ...
func ParseConditionalExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseConditionalExpressionNode not implemented")
	// return nil, nil
}

// AssignmentExpressionNode [In, Yield] : [See 12.14]
//  ConditionalExpression[?In, ?Yield]
//  [+Yield] YieldExpression[?In]
//  ArrowFunction[?In, ?Yield]
//  LeftHandSideExpression[?Yield] = AssignmentExpression[?In, ?Yield]
//  LeftHandSideExpression[?Yield] AssignmentOperator AssignmentExpression[?In, ?Yield]
// implements: Parser and ASTNode
type AssignmentExpressionNode struct {
	node
}

// ParseAssignmentExpressionNode ...
func ParseAssignmentExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseAssignmentExpressionNode not implemented")
	// return nil, nil
}

// AssignmentOperatorNode  : one of [See 12.14]
//  *= /= %= += -= <<= >>= >>>= &= ^= |=
// implements: Parser and ASTNode
type AssignmentOperatorNode struct {
	node
	Operator string
}

// ParseAssignmentOperatorNode ...
func ParseAssignmentOperatorNode(l *Lexer) (ASTNode, error) {
	n := AssignmentOperatorNode{node: node{l.CurrentPosition()}}

	err := errors.New("Assignment operation expected one of: *= /= %= += -= <<= >>= >>>= &= ^= |=")
	tok := l.Next(l.goal)
	if tok.Type != PunctuatorToken {
		return n, err
	}

	for _, op := range []string{"*=", "/=", "%=", "+=", "-=", "<=", ">>=", ">>>=", "&=", "^=", "|"} {
		if tok.Value == op {
			n.Operator = op
			break
		}
	}

	if n.Operator == "" {
		return n, err
	}
	return n, nil
}

// ExpressionNode [In, Yield] : [See 12.15]
//  AssignmentExpression[?In, ?Yield]
//  Expression[?In, ?Yield] , AssignmentExpression[?In, ?Yield]
// implements: Parser and ASTNode
type ExpressionNode struct {
	node
	child ASTNode
}

// ThisExpressionNode is a helper for ExpressionNode
type ThisExpressionNode struct {
	node
}

// ParseExpressionNode ...
func ParseExpressionNode(l *Lexer) (ASTNode, error) {
	var (
		child ASTNode
		err   error
		node  = ExpressionNode{}
	)
	tok := l.Peek(InputElementDiv)
	if tok.Type == ReservedWordToken && tok.Value == "this" {
		this := ThisExpressionNode{}
		this.FilePosition = tok.FilePosition
		return this, nil
	}
	if child != nil {
		node.child = child
		return node, nil
	}
	child, err = tryParseFuncs(l,
		ParseAssignmentExpressionNode,
	)

	if child != nil {
		node.child = child
	}
	node.FilePosition = l.CurrentPosition()
	return node, err
}

//
// A.3 Statements
//

// StatementNode [Yield, Return] : [See clause 13]
//  BlockStatement[?Yield, ?Return]
//  VariableStatement[?Yield]
//  EmptyStatement
//  ExpressionStatement[?Yield]
//  IfStatement[?Yield, ?Return]
//  BreakableStatement[?Yield, ?Return]
//  ContinueStatement[?Yield]
//  BreakStatement[?Yield]
//  [+Return] ReturnStatement[?Yield]
//  WithStatement[?Yield, ?Return]
//  LabelledStatement[?Yield, ?Return]
//  ThrowStatement[?Yield]
//  TryStatement[?Yield, ?Return]
//  DebuggerStatement
// implements: Parser and ASTNode
type StatementNode struct {
	child ASTNode
	node
}

// ParseStatementNode ...
func ParseStatementNode(l *Lexer) (ASTNode, error) {
	node := StatementNode{}
	child, err := tryParseFuncs(l,
		ParseExpressionStatementNode,
		// ParseBlockStatementNode,
		// ParseVariableStatementNode,
		// ParseIfStatementNode,
		// ParseBreakableStatementNode,
		// ParseContinueStatementNode,
		// ParseBreakStatementNode,
		// ParseReturnStatementNode,
		// ParseWithStatementNode,
		// ParseLabelledStatementNode,
		// ParseThrowStatementNode,
		// ParseTryStatementNode,
		// ParseDebuggerStatementNode,
		// ParseEmptyStatementNode,
	)
	if child != nil {
		node.child = child
	}
	node.FilePosition = l.CurrentPosition()
	return node, err
}

// DeclarationNode [Yield] : [See clause 13]
//  HoistableDeclaration[?Yield]
//  ClassDeclaration[?Yield]
//  LexicalDeclaration[In, ?Yield]
// implements: Parser and ASTNode
type DeclarationNode struct {
	child ASTNode
	node
}

// ParseDeclarationNode ...
func ParseDeclarationNode(l *Lexer) (ASTNode, error) {
	node := DeclarationNode{}
	child, err := tryParseFuncs(l,
		ParseHoistableDeclarationNode,
		ParseClassDeclarationNode,
		ParseLexicalDeclarationNode,
	)
	if err != nil {
		return nil, err
	}
	if child != nil {
		node.child = child
	}
	node.FilePosition = l.CurrentPosition()
	return node, nil
}

// HoistableDeclarationNode [Yield, Default] : [See clause 13]
//  FunctionDeclaration[?Yield,?Default]
//  GeneratorDeclaration[?Yield, ?Default]
// implements: Parser and ASTNode
type HoistableDeclarationNode struct {
	child ASTNode
	node
}

// ParseHoistableDeclarationNode ...
func ParseHoistableDeclarationNode(l *Lexer) (ASTNode, error) {
	node := HoistableDeclarationNode{}
	child, err := tryParseFuncs(l,
		ParseFunctionDeclarationNode,
		ParseGeneratorDeclarationNode,
	)
	if err != nil {
		return nil, err
	}
	if child != nil {
		node.child = node
	}
	node.FilePosition = l.CurrentPosition()
	return node, err
}

// BreakableStatementNode [Yield, Return] : [See clause 13]
//  IterationStatement[?Yield, ?Return]
//  SwitchStatement[?Yield, ?Return]
// implements: Parser and ASTNode
type BreakableStatementNode struct {
	node
}

// ParseBreakableStatementNode ...
func ParseBreakableStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseBreakableStatementNode not implemented")
	// return nil, nil
}

// BlockStatementNode [Yield, Return] : [See 13.2]
//  Block[?Yield, ?Return]
// implements: Parser and ASTNode
type BlockStatementNode struct {
	node
}

// ParseBlockStatementNode ...
func ParseBlockStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseBlockStatementNode not implemented")
	// return nil, nil
}

// BlockNode [Yield, Return] : [See 13.2]
//  { StatementList[?Yield, ?Return]opt }
// implements: Parser and ASTNode
type BlockNode struct {
	node
}

// ParseBlockNode ...
func ParseBlockNode(l *Lexer) (ASTNode, error) {
	panic("ParseBlockNode not implemented")
	// return nil, nil
}

// StatementListNode [Yield, Return] : [See 13.2]
//  StatementListItem[?Yield, ?Return]
//  StatementList[?Yield, ?Return] StatementListItem[?Yield, ?Return]
// implements: Parser and ASTNode
type StatementListNode struct {
	children []ASTNode
	node
}

// ParseStatementListNode ...
func ParseStatementListNode(l *Lexer) (StatementListNode, error) {
	node := StatementListNode{}
	for {
		child, err := ParseStatementListItemNode(l)
		if err != nil {
			return node, err
		}
		node.children = append(node.children, child)
	}
}

// StatementListItemNode [Yield, Return] : [See 13.2]
//  Statement[?Yield, ?Return]
//  Declaration[?Yield]
// implements: Parser and ASTNode
type StatementListItemNode struct {
	child ASTNode
	node
}

// ParseStatementListItemNode ...
func ParseStatementListItemNode(l *Lexer) (StatementListItemNode, error) {
	node := StatementListItemNode{}
	child, err := tryParseFuncs(l, ParseStatementNode, ParseDeclarationNode)
	if child != nil {
		node.child = child
	}
	node.FilePosition = l.CurrentPosition()
	return node, err
}

// LexicalDeclarationNode [In, Yield] : [See 13.3.1]
//  LetOrConst BindingList[?In, ?Yield] ;
// implements: Parser and ASTNode
type LexicalDeclarationNode struct {
	node
}

// ParseLexicalDeclarationNode ...
func ParseLexicalDeclarationNode(l *Lexer) (ASTNode, error) {
	panic("ParseLexicalDeclarationNode not implemented")
	// return nil, nil
}

// LetOrConstNode  : [See 13.3.1]
//  let
//  const
// implements: Parser and ASTNode
type LetOrConstNode struct {
	node
	Value string
}

// ParseLetOrConstNode ...
func ParseLetOrConstNode(l *Lexer) (ASTNode, error) {
	n := LetOrConstNode{node: node{l.CurrentPosition()}}

	tok := l.Next(l.goal)
	if tok.Value == "const" || tok.Value == "let" {
		n.Value = tok.Value
		return n, nil
	}
	return n, errors.New("expected const or let")
}

// BindingListNode [In, Yield] : [See 13.3.1]
//  LexicalBinding[?In, ?Yield]
//  BindingList[?In, ?Yield] , LexicalBinding[?In, ?Yield]
// implements: Parser and ASTNode
type BindingListNode struct {
	node
}

// ParseBindingListNode ...
func ParseBindingListNode(l *Lexer) (ASTNode, error) {
	panic("ParseBindingListNode not implemented")
	// return nil, nil
}

// LexicalBindingNode [In, Yield] : [See 13.3.1]
//  BindingIdentifier[?Yield] Initializer[?In, ?Yield]opt
//  BindingPattern[?Yield] Initializer[?In, ?Yield]
// implements: Parser and ASTNode
type LexicalBindingNode struct {
	node
}

// ParseLexicalBindingNode ...
func ParseLexicalBindingNode(l *Lexer) (ASTNode, error) {
	panic("ParseLexicalBindingNode not implemented")
	// return nil, nil
}

// VariableStatementNode [Yield] : [See 13.3.2]
//  var VariableDeclarationList[In, ?Yield] ;
// implements: Parser and ASTNode
type VariableStatementNode struct {
	node
}

// ParseVariableStatementNode ...
func ParseVariableStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseVariableStatementNode not implemented")
	// return nil, nil
}

// VariableDeclarationListNode [In, Yield] : [See 13.3.2]
//  VariableDeclaration[?In, ?Yield]
//  VariableDeclarationList[?In, ?Yield] , VariableDeclaration[?In, ?Yield]
// implements: Parser and ASTNode
type VariableDeclarationListNode struct {
	node
}

// ParseVariableDeclarationListNode ...
func ParseVariableDeclarationListNode(l *Lexer) (ASTNode, error) {
	panic("ParseVariableDeclarationListNode not implemented")
	// return nil, nil
}

// VariableDeclarationNode [In, Yield] : [See 13.3.2]
//  BindingIdentifier[?Yield] Initializer[?In, ?Yield]opt
//  BindingPattern[?Yield] Initializer[?In, ?Yield]
// implements: Parser and ASTNode
type VariableDeclarationNode struct {
	node
}

// ParseVariableDeclarationNode ...
func ParseVariableDeclarationNode(l *Lexer) (ASTNode, error) {
	panic("ParseVariableDeclarationNode not implemented")
	// return nil, nil
}

// BindingPatternNode [Yield] : [See 13.3.3]
//  ObjectBindingPattern[?Yield]
//  ArrayBindingPattern[?Yield]
// implements: Parser and ASTNode
type BindingPatternNode struct {
	node
}

// ParseBindingPatternNode ...
func ParseBindingPatternNode(l *Lexer) (ASTNode, error) {
	panic("ParseBindingPatternNode not implemented")
	// return nil, nil
}

// ObjectBindingPatternNode [Yield] : [See 13.3.3]
//  { }
//  { BindingPropertyList[?Yield] }
//  { BindingPropertyList[?Yield] , }
// implements: Parser and ASTNode
type ObjectBindingPatternNode struct {
	node
}

// ParseObjectBindingPatternNode ...
func ParseObjectBindingPatternNode(l *Lexer) (ASTNode, error) {
	panic("ParseObjectBindingPatternNode not implemented")
	// return nil, nil
}

// ArrayBindingPatternNode [Yield] : [See 13.3.3]
//  [ Elisionopt BindingRestElement[?Yield]opt ]
//  [ BindingElementList[?Yield] ]
//  [ BindingElementList[?Yield] , Elisionopt BindingRestElement[?Yield]opt ]
// implements: Parser and ASTNode
type ArrayBindingPatternNode struct {
	node
}

// ParseArrayBindingPatternNode ...
func ParseArrayBindingPatternNode(l *Lexer) (ASTNode, error) {
	panic("ParseArrayBindingPatternNode not implemented")
	// return nil, nil
}

// BindingPropertyListNode [Yield] : [See 13.3.3]
//  BindingProperty[?Yield]
//  BindingPropertyList[?Yield] , BindingProperty[?Yield]
// implements: Parser and ASTNode
type BindingPropertyListNode struct {
	node
}

// ParseBindingPropertyListNode ...
func ParseBindingPropertyListNode(l *Lexer) (ASTNode, error) {
	panic("ParseBindingPropertyListNode not implemented")
	// return nil, nil
}

// BindingElementListNode [Yield] : [See 13.3.3]
//  BindingElisionElement[?Yield]
//  BindingElementList[?Yield] , BindingElisionElement[?Yield]
// implements: Parser and ASTNode
type BindingElementListNode struct {
	node
}

// ParseBindingElementListNode ...
func ParseBindingElementListNode(l *Lexer) (ASTNode, error) {
	panic("ParseBindingElementListNode not implemented")
	// return nil, nil
}

// BindingElisionElementNode [Yield] : [See 13.3.3]
//  Elisionopt BindingElement[?Yield]
// implements: Parser and ASTNode
type BindingElisionElementNode struct {
	node
}

// ParseBindingElisionElementNode ...
func ParseBindingElisionElementNode(l *Lexer) (ASTNode, error) {
	panic("ParseBindingElisionElementNode not implemented")
	// return nil, nil
}

// BindingPropertyNode [Yield] : [See 13.3.3]
//  SingleNameBinding[?Yield]
//  PropertyName[?Yield] : BindingElement[?Yield]
// implements: Parser and ASTNode
type BindingPropertyNode struct {
	node
}

// ParseBindingPropertyNode ...
func ParseBindingPropertyNode(l *Lexer) (ASTNode, error) {
	panic("ParseBindingPropertyNode not implemented")
	// return nil, nil
}

// BindingElementNode [Yield] : [See 13.3.3]
//  SingleNameBinding[?Yield]
//  BindingPattern[?Yield] Initializer[In, ?Yield]opt
// implements: Parser and ASTNode
type BindingElementNode struct {
	node
}

// ParseBindingElementNode ...
func ParseBindingElementNode(l *Lexer) (ASTNode, error) {
	panic("ParseBindingElementNode not implemented")
	// return nil, nil
}

// SingleNameBindingNode [Yield] : [See 13.3.3]
//  BindingIdentifier[?Yield] Initializer[In, ?Yield]opt
// implements: Parser and ASTNode
type SingleNameBindingNode struct {
	node
}

// ParseSingleNameBindingNode ...
func ParseSingleNameBindingNode(l *Lexer) (ASTNode, error) {
	panic("ParseSingleNameBindingNode not implemented")
	// return nil, nil
}

// BindingRestElementNode [Yield] : [See 13.3.3]
//  ... BindingIdentifier[?Yield]
// implements: Parser and ASTNode
type BindingRestElementNode struct {
	node
}

// ParseBindingRestElementNode ...
func ParseBindingRestElementNode(l *Lexer) (ASTNode, error) {
	panic("ParseBindingRestElementNode not implemented")
	// return nil, nil
}

// EmptyStatementNode  : [See 13.4]
//  ;
// implements: Parser and ASTNode
type EmptyStatementNode struct {
	node
}

// ParseEmptyStatementNode ...
func ParseEmptyStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseEmptyStatementNode not implemented")
	// return nil, nil
}

// ExpressionStatementNode [Yield] : [See 13.5]
//  [lookahead ∉ {{, function, class, let [}] Expression[In, ?Yield] ;
// implements: Parser and ASTNode
type ExpressionStatementNode struct {
	child ASTNode
	node
}

// ParseExpressionStatementNode ...
func ParseExpressionStatementNode(l *Lexer) (ASTNode, error) {
	var err error
	tok := l.Peek(InputElementDiv)
	switch tok.Type {
	case ReservedWordToken:
		switch tok.Value {
		case "function", "class":
			return nil, IncorrectTokenError(tok)
		case "let":
			tok2 := l.Peek(InputElementDiv)
			if tok2.Type == PunctuatorToken && tok2.Value == "[" {
				return nil, IncorrectTokenError(tok2)
			}
		default:
		}
	case PunctuatorToken:
		if tok.Type == PunctuatorToken && tok.Value == "{" {
			return nil, IncorrectTokenError(tok)
		}
	default:
	}
	node := ExpressionStatementNode{}

	node.child, err = ParseExpressionNode(l)
	if err != nil {
		return nil, err
	}
	tok3 := l.Peek(InputElementDiv)
	if tok3.Type != PunctuatorToken || tok3.Value != ";" {
		return nil, IncorrectTokenError(tok)
	}
	node.FilePosition = l.CurrentPosition()
	return node, err
}

// IfStatementNode [Yield, Return] : [See 13.6]
//  if ( Expression[In, ?Yield] ) Statement[?Yield, ?Return] else Statement[?Yield, ?Return]
//  if ( Expression[In, ?Yield] ) Statement[?Yield, ?Return]
// implements: Parser and ASTNode
type IfStatementNode struct {
	node
}

// ParseIfStatementNode ...
func ParseIfStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseIfStatementNode not implemented")
	// return nil, nil
}

// IterationStatementNode [Yield, Return] : [See 13.7]
//  do Statement[?Yield, ?Return] while ( Expression[In, ?Yield] ) ;
//  while ( Expression[In, ?Yield] ) Statement[?Yield, ?Return]
//  for ( [lookahead ∉ {let [}] Expression[?Yield]opt ; Expression[In, ?Yield]opt ; Expression[In, ?Yield]opt ) Statement[?Yield, ?Return]
//  for ( var VariableDeclarationList[?Yield] ; Expression[In, ?Yield]opt ; Expression[In, ?Yield]opt ) Statement[?Yield, ?Return]
//  for ( LexicalDeclaration[?Yield] Expression[In, ?Yield]opt ; Expression[In, ?Yield]opt ) Statement[?Yield, ?Return]
//  for ( [lookahead ∉ {let [}] LeftHandSideExpression[?Yield] in Expression[In, ?Yield] ) Statement[?Yield, ?Return]
//  for ( var ForBinding[?Yield] in Expression[In, ?Yield] ) Statement[?Yield, ?Return]
//  for ( ForDeclaration[?Yield] in Expression[In, ?Yield] ) Statement[?Yield, ?Return]
//  for ( [lookahead ≠ let ] LeftHandSideExpression[?Yield] of AssignmentExpression[In, ?Yield] ) Statement[?Yield, ?Return]
//  for ( var ForBinding[?Yield] of AssignmentExpression[In, ?Yield] ) Statement[?Yield, ?Return]
//  for ( ForDeclaration[?Yield] of AssignmentExpression[In, ?Yield] ) Statement[?Yield, ?Return]
// implements: Parser and ASTNode
type IterationStatementNode struct {
	node
}

// ParseIterationStatementNode ...
func ParseIterationStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseIterationStatementNode not implemented")
	// return nil, nil
}

// ForDeclarationNode [Yield] : [See 13.7]
//  LetOrConst ForBinding[?Yield]
// implements: Parser and ASTNode
type ForDeclarationNode struct {
	node
}

// ParseForDeclarationNode ...
func ParseForDeclarationNode(l *Lexer) (ASTNode, error) {
	panic("ParseForDeclarationNode not implemented")
	// return nil, nil
}

// ForBindingNode [Yield] : [See 13.7]
//  BindingIdentifier[?Yield]
//  BindingPattern[?Yield]
// implements: Parser and ASTNode
type ForBindingNode struct {
	node
}

// ParseForBindingNode ...
func ParseForBindingNode(l *Lexer) (ASTNode, error) {
	panic("ParseForBindingNode not implemented")
	// return nil, nil
}

// ContinueStatementNode [Yield] : [See 13.8]
//  continue ;
//  continue [no LineTerminator here] LabelIdentifier[?Yield] ;
// implements: Parser and ASTNode
type ContinueStatementNode struct {
	node
}

// ParseContinueStatementNode ...
func ParseContinueStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseContinueStatementNode not implemented")
	// return nil, nil
}

// BreakStatementNode [Yield] : [See 13.9]
//  break ;
//  break [no LineTerminator here] LabelIdentifier[?Yield] ;
// implements: Parser and ASTNode
type BreakStatementNode struct {
	node
}

// ParseBreakStatementNode ...
func ParseBreakStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseBreakStatementNode not implemented")
	// return nil, nil
}

// ReturnStatementNode [Yield] : [See 13.10]
//  return ;
//  return [no LineTerminator here] Expression[In, ?Yield] ;
// implements: Parser and ASTNode
type ReturnStatementNode struct {
	node
}

// ParseReturnStatementNode ...
func ParseReturnStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseReturnStatementNode not implemented")
	// return nil, nil
}

// WithStatementNode [Yield, Return] : [See 13.11]
//  with ( Expression[In, ?Yield] ) Statement[?Yield, ?Return]
// implements: Parser and ASTNode
type WithStatementNode struct {
	node
}

// ParseWithStatementNode ...
func ParseWithStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseWithStatementNode not implemented")
	// return nil, nil
}

// SwitchStatementNode [Yield, Return] : [See 13.12]
//  switch ( Expression[In, ?Yield] ) CaseBlock[?Yield, ?Return]
// implements: Parser and ASTNode
type SwitchStatementNode struct {
	node
}

// ParseSwitchStatementNode ...
func ParseSwitchStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseSwitchStatementNode not implemented")
	// return nil, nil
}

// CaseBlockNode [Yield, Return] : [See 13.12]
//  { CaseClauses[?Yield, ?Return]opt }
//  { CaseClauses[?Yield, ?Return]opt DefaultClause[?Yield, ?Return] CaseClauses[?Yield, ?Return]opt }
// implements: Parser and ASTNode
type CaseBlockNode struct {
	node
}

// ParseCaseBlockNode ...
func ParseCaseBlockNode(l *Lexer) (ASTNode, error) {
	panic("ParseCaseBlockNode not implemented")
	// return nil, nil
}

// CaseClausesNode [Yield, Return] : [See 13.12]
//  CaseClause[?Yield, ?Return]
//  CaseClauses[?Yield, ?Return] CaseClause[?Yield, ?Return]
// implements: Parser and ASTNode
type CaseClausesNode struct {
	node
}

// ParseCaseClausesNode ...
func ParseCaseClausesNode(l *Lexer) (ASTNode, error) {
	panic("ParseCaseClausesNode not implemented")
	// return nil, nil
}

// CaseClauseNode [Yield, Return] : [See 13.12]
//  case Expression[In, ?Yield] : StatementList[?Yield, ?Return]opt
// implements: Parser and ASTNode
type CaseClauseNode struct {
	node
}

// ParseCaseClauseNode ...
func ParseCaseClauseNode(l *Lexer) (ASTNode, error) {
	panic("ParseCaseClauseNode not implemented")
	// return nil, nil
}

// DefaultClauseNode [Yield, Return] : [See 13.12]
//  default : StatementList[?Yield, ?Return]opt
// implements: Parser and ASTNode
type DefaultClauseNode struct {
	node
}

// ParseDefaultClauseNode ...
func ParseDefaultClauseNode(l *Lexer) (ASTNode, error) {
	panic("ParseDefaultClauseNode not implemented")
	// return nil, nil
}

// LabelledStatementNode [Yield, Return] : [See 13.13]
//  LabelIdentifier[?Yield] : LabelledItem[?Yield, ?Return]
// implements: Parser and ASTNode
type LabelledStatementNode struct {
	node
}

// ParseLabelledStatementNode ...
func ParseLabelledStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseLabelledStatementNode not implemented")
	// return nil, nil
}

// LabelledItemNode [Yield, Return] : [See 13.13]
//  Statement[?Yield, ?Return]
//  FunctionDeclaration[?Yield]
// implements: Parser and ASTNode
type LabelledItemNode struct {
	node
}

// ParseLabelledItemNode ...
func ParseLabelledItemNode(l *Lexer) (ASTNode, error) {
	panic("ParseLabelledItemNode not implemented")
	// return nil, nil
}

// ThrowStatementNode [Yield] : [See 13.14]
//  throw [no LineTerminator here] Expression[In, ?Yield] ;
// implements: Parser and ASTNode
type ThrowStatementNode struct {
	node
}

// ParseThrowStatementNode ...
func ParseThrowStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseThrowStatementNode not implemented")
	// return nil, nil
}

// TryStatementNode [Yield, Return] : [See 13.15]
//  try Block[?Yield, ?Return] Catch[?Yield, ?Return]
//  try Block[?Yield, ?Return] Finally[?Yield, ?Return]
//  try Block[?Yield, ?Return] Catch[?Yield, ?Return] Finally[?Yield, ?Return]
// implements: Parser and ASTNode
type TryStatementNode struct {
	node
}

// ParseTryStatementNode ...
func ParseTryStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseTryStatementNode not implemented")
	// return nil, nil
}

// CatchNode [Yield, Return] : [See 13.15]
//  catch ( CatchParameter[?Yield] ) Block[?Yield, ?Return]
// implements: Parser and ASTNode
type CatchNode struct {
	node
}

// ParseCatchNode ...
func ParseCatchNode(l *Lexer) (ASTNode, error) {
	panic("ParseCatchNode not implemented")
	// return nil, nil
}

// FinallyNode [Yield, Return] : [See 13.15]
//  finally Block[?Yield, ?Return]
// implements: Parser and ASTNode
type FinallyNode struct {
	node
}

// ParseFinallyNode ...
func ParseFinallyNode(l *Lexer) (ASTNode, error) {
	panic("ParseFinallyNode not implemented")
	// return nil, nil
}

// CatchParameterNode [Yield] : [See 13.15]
//  BindingIdentifier[?Yield]
//  BindingPattern[?Yield]
// implements: Parser and ASTNode
type CatchParameterNode struct {
	node
}

// ParseCatchParameterNode ...
func ParseCatchParameterNode(l *Lexer) (ASTNode, error) {
	panic("ParseCatchParameterNode not implemented")
	// return nil, nil
}

// DebuggerStatementNode  : [See 13.16]
//  debugger ;
// implements: Parser and ASTNode
type DebuggerStatementNode struct {
	node
}

// ParseDebuggerStatementNode ...
func ParseDebuggerStatementNode(l *Lexer) (ASTNode, error) {
	panic("ParseDebuggerStatementNode not implemented")
	// return nil, nil
}

//
// A.4 Functions and Classes
//

// FunctionDeclarationNode [Yield, Default] :
//  function BindingIdentifier[?Yield] ( FormalParameters ) { FunctionBody }
//  [+Default] function ( FormalParameters ) { FunctionBody }
// implements: Parser and ASTNode
type FunctionDeclarationNode struct {
	BindingIdentifier BindingIdentifierNode
	FormalParameters  FormalParametersNode
	FunctionBody      FunctionBodyNode
	node
}

// ParseFunctionDeclarationNode ...
func ParseFunctionDeclarationNode(l *Lexer) (ASTNode, error) {
	var (
		node = FunctionDeclarationNode{}
		// err error
	)
	tokPeek0 := l.Peek(InputElementDiv)
	if tokPeek0.Type != ReservedWordToken {
		return nil, IncorrectTokenError(tokPeek0)
	}
	if tokPeek0.Value != "function" {
		return nil, IncorrectTokenError(tokPeek0)
	}
	l.Next(InputElementDiv) // accept input token

	tokPeek1 := l.Peek(InputElementDiv)
	node.FilePosition = l.CurrentPosition()
	fmt.Println(tokPeek1)
	fmt.Println(node)
	fmt.Println()
	panic("ParseFunctionDeclarationNode not implemented")
	// // if next token is not "(" attempt BindingIdentifier
	// if !(tokPeek1.Type == PunctuatorToken && tokPeek0.Value == "(") {
	// 	var identifierChild ASTNode
	// 	identifierChild, err = ParseBindingIdentifierNode(l)
	// 	if err != nil {
	// 		_, isIncorrectToken := err.(IncorrectTokenError)
	// 		if !isIncorrectToken {
	// 			return nil, err
	// 		}
	// 	}
	// 	var ok bool
	// 	if node.BindingIdentifier, ok = identifierChild.(BindingIdentifierNode); ok {
	// 		return node, IncorrectTokenError(tokPeek1)
	// 	}
	// }
	// return nil, nil
}

// StrictFormalParametersNode [Yield] : [See 14.1]
//  FormalParameters[?Yield]
// implements: Parser and ASTNode
type StrictFormalParametersNode struct {
	node
}

// ParseStrictFormalParametersNode ...
func ParseStrictFormalParametersNode(l *Lexer) (ASTNode, error) {
	panic("ParseStrictFormalParametersNode not implemented")
	// return nil, nil
}

// FormalParametersNode [Yield] : [See 14.1]
//  [empty]
//  FormalParameterList[?Yield]
// implements: Parser and ASTNode
type FormalParametersNode struct {
	node
}

// ParseFormalParametersNode ...
func ParseFormalParametersNode(l *Lexer) (ASTNode, error) {
	panic("ParseFormalParametersNode not implemented")
	// return nil, nil
}

// FormalParameterListNode [Yield] : [See 14.1]
//  FunctionRestParameter[?Yield]
//  FormalsList[?Yield]
//  FormalsList[?Yield] , FunctionRestParameter[?Yield]
// implements: Parser and ASTNode
type FormalParameterListNode struct {
	node
}

// ParseFormalParameterListNode ...
func ParseFormalParameterListNode(l *Lexer) (ASTNode, error) {
	panic("ParseFormalParameterListNode not implemented")
	// return nil, nil
}

// FormalsListNode [Yield] : [See 14.1]
//  FormalParameter[?Yield]
//  FormalsList[?Yield] , FormalParameter[?Yield]
// implements: Parser and ASTNode
type FormalsListNode struct {
	node
}

// ParseFormalsListNode ...
func ParseFormalsListNode(l *Lexer) (ASTNode, error) {
	panic("ParseFormalsListNode not implemented")
	// return nil, nil
}

// FunctionRestParameterNode [Yield] : [See 14.1]
//  BindingRestElement[?Yield]
// implements: Parser and ASTNode
type FunctionRestParameterNode struct {
	node
}

// ParseFunctionRestParameterNode ...
func ParseFunctionRestParameterNode(l *Lexer) (ASTNode, error) {
	panic("ParseFunctionRestParameterNode not implemented")
	// return nil, nil
}

// FormalParameterNode [Yield] : [See 14.1]
//  BindingElement[?Yield]
// implements: Parser and ASTNode
type FormalParameterNode struct {
	node
}

// ParseFormalParameterNode ...
func ParseFormalParameterNode(l *Lexer) (ASTNode, error) {
	panic("ParseFormalParameterNode not implemented")
	// return nil, nil
}

// FunctionBodyNode [Yield] : [See 14.1]
//  FunctionStatementList[?Yield]
// implements: Parser and ASTNode
type FunctionBodyNode struct {
	node
}

// ParseFunctionBodyNode ...
func ParseFunctionBodyNode(l *Lexer) (ASTNode, error) {
	panic("ParseFunctionBodyNode not implemented")
	// return nil, nil
}

// FunctionStatementListNode [Yield] : [See 14.1]
//  StatementList[?Yield, Return]opt
// implements: Parser and ASTNode
type FunctionStatementListNode struct {
	node
}

// ParseFunctionStatementListNode ...
func ParseFunctionStatementListNode(l *Lexer) (ASTNode, error) {
	panic("ParseFunctionStatementListNode not implemented")
	// return nil, nil
}

// ArrowFunctionNode [In, Yield] : [See 14.2]
//  ArrowParameters[?Yield] [no LineTerminator here] => ConciseBody[?In]
// implements: Parser and ASTNode
type ArrowFunctionNode struct {
	node
}

// ParseArrowFunctionNode ...
func ParseArrowFunctionNode(l *Lexer) (ASTNode, error) {
	panic("ParseArrowFunctionNode not implemented")
	// return nil, nil
}

// ArrowParametersNode [Yield] : [See 14.2]
//  BindingIdentifier[?Yield]
//  CoverParenthesizedExpressionAndArrowParameterList[?Yield]
// ArrowParameters[Yield] : CoverParenthesizedExpressionAndArrowParameterList[?Yield]
// is recognized the following grammar is used to refine the interpretation of CoverParenthesizedExpressionAndArrowParameterList :
// implements: Parser and ASTNode
type ArrowParametersNode struct {
	node
}

// ParseArrowParametersNode ...
func ParseArrowParametersNode(l *Lexer) (ASTNode, error) {
	panic("ParseArrowParametersNode not implemented")
	// return nil, nil
}

// ConciseBodyNode [In] : [See 14.2]
//  [lookahead ≠ { ] AssignmentExpression[?In]
//  { FunctionBody }
//  ArrowFormalParameters[Yield] :
//  ( StrictFormalParameters[?Yield] )
// implements: Parser and ASTNode
type ConciseBodyNode struct {
	node
}

// ParseConciseBodyNode ...
func ParseConciseBodyNode(l *Lexer) (ASTNode, error) {
	panic("ParseConciseBodyNode not implemented")
	// return nil, nil
}

// MethodDefinitionNode [Yield] : [See 14.3]
//  PropertyName[?Yield] ( StrictFormalParameters ) { FunctionBody }
//  GeneratorMethod[?Yield]
//  get PropertyName[?Yield] ( ) { FunctionBody }
//  set PropertyName[?Yield] ( PropertySetParameterList ) { FunctionBody }
// implements: Parser and ASTNode
type MethodDefinitionNode struct {
	node
}

// ParseMethodDefinitionNode ...
func ParseMethodDefinitionNode(l *Lexer) (ASTNode, error) {
	panic("ParseMethodDefinitionNode not implemented")
	// return nil, nil
}

// PropertySetParameterListNode  : [See 14.3]
// FormalParameter
// implements: Parser and ASTNode
type PropertySetParameterListNode struct {
	node
}

// ParsePropertySetParameterListNode ...
func ParsePropertySetParameterListNode(l *Lexer) (ASTNode, error) {
	panic("ParsePropertySetParameterListNode not implemented")
	// return nil, nil
}

// GeneratorMethodNode [Yield] : [See 14.4]
//  * PropertyName[?Yield] ( StrictFormalParameters[Yield] ) { GeneratorBody }
// implements: Parser and ASTNode
type GeneratorMethodNode struct {
	node
}

// ParseGeneratorMethodNode ...
func ParseGeneratorMethodNode(l *Lexer) (ASTNode, error) {
	panic("ParseGeneratorMethodNode not implemented")
	// return nil, nil
}

// GeneratorDeclarationNode [Yield, Default] : [See 14.4]
//  function * BindingIdentifier[?Yield] ( FormalParameters[Yield] ) { GeneratorBody }
//  [+Default] function * ( FormalParameters[Yield] ) { GeneratorBody }
// implements: Parser and ASTNode
type GeneratorDeclarationNode struct {
	node
}

// ParseGeneratorDeclarationNode ...
func ParseGeneratorDeclarationNode(l *Lexer) (ASTNode, error) {
	panic("ParseGeneratorDeclarationNode not implemented")
	// return nil, nil
}

// ParseGeneratorBodyNode ...
func ParseGeneratorBodyNode(l *Lexer) (ASTNode, error) {
	panic("ParseGeneratorBodyNode not implemented")
	// return nil, nil
}

// YieldExpressionNode [In] : [See 14.4]
//  yield
//  yield [no LineTerminator here] AssignmentExpression[?In, Yield]
//  yield [no LineTerminator here] * AssignmentExpression[?In, Yield]
// implements: Parser and ASTNode
type YieldExpressionNode struct {
	node
}

// ParseYieldExpressionNode ...
func ParseYieldExpressionNode(l *Lexer) (ASTNode, error) {
	panic("ParseYieldExpressionNode not implemented")
	// return nil, nil
}

// ClassDeclarationNode [Yield, Default] : [See 14.5]
//  class BindingIdentifier[?Yield] ClassTail[?Yield]
//  [+Default] class ClassTail[?Yield]
// implements: Parser and ASTNode
type ClassDeclarationNode struct {
	node
}

// ParseClassDeclarationNode ...
func ParseClassDeclarationNode(l *Lexer) (ASTNode, error) {
	panic("ParseClassDeclarationNode not implemented")
	// return nil, nil
}

// ClassTailNode [Yield] : [See 14.5]
//  ClassHeritage[?Yield]opt { ClassBody[?Yield]opt }
// implements: Parser and ASTNode
type ClassTailNode struct {
	node
}

// ParseClassTailNode ...
func ParseClassTailNode(l *Lexer) (ASTNode, error) {
	panic("ParseClassTailNode not implemented")
	// return nil, nil
}

// ClassHeritageNode [Yield] : [See 14.5]
//  extends LeftHandSideExpression[?Yield]
// implements: Parser and ASTNode
type ClassHeritageNode struct {
	node
}

// ParseClassHeritageNode ...
func ParseClassHeritageNode(l *Lexer) (ASTNode, error) {
	panic("ParseClassHeritageNode not implemented")
	// return nil, nil
}

// ClassBodyNode [Yield] : [See 14.5]
//  ClassElementList[?Yield]
// implements: Parser and ASTNode
type ClassBodyNode struct {
	node
}

// ParseClassBodyNode ...
func ParseClassBodyNode(l *Lexer) (ASTNode, error) {
	panic("ParseClassBodyNode not implemented")
	// return nil, nil
}

// ClassElementListNode [Yield] : [See 14.5]
//  ClassElement[?Yield]
//  ClassElementList[?Yield] ClassElement[?Yield]
// implements: Parser and ASTNode
type ClassElementListNode struct {
	node
}

// ParseClassElementListNode ...
func ParseClassElementListNode(l *Lexer) (ASTNode, error) {
	panic("ParseClassElementListNode not implemented")
	// return nil, nil
}

// ClassElementNode [Yield] : [See 14.5]
//  MethodDefinition[?Yield]
//  static MethodDefinition[?Yield]
//  ;
// implements: Parser and ASTNode
type ClassElementNode struct {
	node
}

// ParseClassElementNode ...
func ParseClassElementNode(l *Lexer) (ASTNode, error) {
	panic("ParseClassElementNode not implemented")
	// return nil, nil
}

//
// A.5 Scripts and Modules
//

// ScriptNode [See 15.1]
// Script : ScriptBody*
// implements: Parser and ASTNode
type ScriptNode struct {
	node
	child ASTNode
}

// ParseScriptNode ...
func ParseScriptNode(l *Lexer) (ASTNode, error) {
	c, err := ParseScriptBodyNode(l)
	return ScriptNode{child: c}, err
}

// ScriptBodyNode [See 15.1]
//  StatementList
// implements: Parser and ASTNode
type ScriptBodyNode struct {
	node
	child ASTNode
}

// ParseScriptBodyNode ...
func ParseScriptBodyNode(l *Lexer) (ASTNode, error) {
	c, err := ParseStatementListNode(l)
	node := ScriptBodyNode{
		child: c,
	}
	node.FilePosition = l.CurrentPosition()
	return node, err
}

// ModuleNode [See 15.2]
//  ModuleBodyopt
// implements: Parser and ASTNode
type ModuleNode struct {
	node
}

// ParseModuleNode ...
func ParseModuleNode(l *Lexer) (ASTNode, error) {
	panic("ParseModuleNode not implemented")
	// return nil, nil
}

// ModuleBodyNode [See 15.2]
//  ModuleItemList*
// implements: Parser and ASTNode
type ModuleBodyNode struct {
	node
}

// ParseModuleBodyNode ...
func ParseModuleBodyNode(l *Lexer) (ASTNode, error) {
	panic("ParseModuleBodyNode not implemented")
	// return nil, nil
}

// ModuleItemListNode [See 15.2]
//  ModuleItem
//  ModuleItemList ModuleItem
// implements: Parser and ASTNode
type ModuleItemListNode struct {
	node
}

// ParseModuleItemListNode ...
func ParseModuleItemListNode(l *Lexer) (ASTNode, error) {
	panic("ParseModuleItemListNode not implemented")
	// return nil, nil
}

// ModuleItemNode [See 15.2]
//  ImportDeclaration
//  ExportDeclaration
//  StatementListItem
// implements: Parser and ASTNode
type ModuleItemNode struct {
	node
}

// ParseModuleItemNode ...
func ParseModuleItemNode(l *Lexer) (ASTNode, error) {
	panic("ParseModuleItemNode not implemented")
	// return nil, nil
}

// ImportDeclarationNode [See 15.2.2]
//  import ImportClause FromClause ;
//  import ModuleSpecifier ;
// implements: Parser and ASTNode
type ImportDeclarationNode struct {
	node
}

// ParseImportDeclarationNode ...
func ParseImportDeclarationNode(l *Lexer) (ASTNode, error) {
	panic("ParseImportDeclarationNode not implemented")
	// return nil, nil
}

// ImportClauseNode [See 15.2.2]
//  ImportedDefaultBinding
//  NameSpaceImport
//  NamedImports
//  ImportedDefaultBinding , NameSpaceImport
//  ImportedDefaultBinding , NamedImports
// implements: Parser and ASTNode
type ImportClauseNode struct {
	node
}

// ParseImportClauseNode ...
func ParseImportClauseNode(l *Lexer) (ASTNode, error) {
	panic("ParseImportClauseNode not implemented")
	// return nil, nil
}

// ImportedDefaultBindingNode [See 15.2.2]
//  ImportedBinding
type ImportedDefaultBindingNode struct {
	node
}

// NameSpaceImportNode [See 15.2.2]
//  * as ImportedBinding
type NameSpaceImportNode struct {
	node
}

// NamedImportsNode [See 15.2]
//  { }
//  { ImportsList }
//  { ImportsList , }
type NamedImportsNode struct {
	node
}

// ImportsListNode [See 15.2.2]
//  ImportSpecifier
//  ImportsList , ImportSpecifier
// implements: Parser and ASTNode
type ImportsListNode struct {
	node
}

// ParseImportsListNode ...
func ParseImportsListNode(l *Lexer) (ASTNode, error) {
	panic("ParseImportsListNode not implemented")
	// return nil, nil
}

// ImportSpecifierNode [See 15.2.2]
//  ImportedBinding
//  IdentifierName as ImportedBinding
// implements: Parser and ASTNode
type ImportSpecifierNode struct {
	node
}

// ParseImportSpecifierNode ...
func ParseImportSpecifierNode(l *Lexer) (ASTNode, error) {
	panic("ParseImportSpecifierNode not implemented")
	// return nil, nil
}

// ModuleSpecifierNode [See 15.2.2]
//  StringLiteral
// implements: Parser and ASTNode
type ModuleSpecifierNode struct {
	node
}

// ParseModuleSpecifierNode ...
func ParseModuleSpecifierNode(l *Lexer) (ASTNode, error) {
	panic("ParseModuleSpecifierNode not implemented")
	// return nil, nil
}

// ImportedBindingNode [See 15.2.2]
//  BindingIdentifier
// implements: Parser and ASTNode
type ImportedBindingNode struct {
	node
}

// ParseImportedBindingNode ...
func ParseImportedBindingNode(l *Lexer) (ASTNode, error) {
	panic("ParseImportedBindingNode not implemented")
	// return nil, nil
}

// ExportDeclarationNode [See 15.2.3]
//  export * FromClause ;
//  export ExportClause FromClause ;
//  export ExportClause ;
//  export VariableStatement
//  export Declaration
//  export default HoistableDeclaration[Default]
//  export default ClassDeclaration[Default]
//  export default [lookahead ∉ {function, class}] AssignmentExpression[In] ;
// implements: Parser and ASTNode
type ExportDeclarationNode struct {
	node
}

// ParseExportDeclarationNode ...
func ParseExportDeclarationNode(l *Lexer) (ASTNode, error) {
	panic("ParseExportDeclarationNode not implemented")
	// return nil, nil
}

// ExportClauseNode [See 15.2.3]
//  { }
//  { ExportsList }
//  { ExportsList , }
// implements: Parser and ASTNode
type ExportClauseNode struct {
	node
}

// ParseExportClauseNode ...
func ParseExportClauseNode(l *Lexer) (ASTNode, error) {
	panic("ParseExportClauseNode not implemented")
	// return nil, nil
}

// ExportsListNode [See 15.2.3]
//  ExportSpecifier
//  ExportsList , ExportSpecifier
// implements: Parser and ASTNode
type ExportsListNode struct {
	node
	List []ExportSpecifierNode
}

// ParseExportsListNode ...
func ParseExportsListNode(l *Lexer) (ASTNode, error) {
	n := ExportsListNode{node: node{l.CurrentPosition()}}

	n.List = []ExportSpecifierNode{}

	for {
		exportSpecifier, err := ParseExportSpecifierNode(l)
		if err != nil {
			return n, err
		}
		n.List = append(n.List, exportSpecifier.(ExportSpecifierNode))

		if comma := l.Peek(l.goal); comma.Value != "," {
			return n, err
		}
		l.Next(l.goal)
	}
}

// ExportSpecifierNode [See 15.2.3]
//  IdentifierName
//  IdentifierName as IdentifierName
// implements: Parser and ASTNode
type ExportSpecifierNode struct {
	node
	IdentifierNode
	As IdentifierNode
}

// ParseExportSpecifierNode ...
func ParseExportSpecifierNode(l *Lexer) (ASTNode, error) {
	n := ExportSpecifierNode{node: node{l.CurrentPosition()}}

	identifierNode, err := ParseIdentifierNode(l)
	if err != nil {
		return n, err
	}
	n.IdentifierNode = identifierNode.(IdentifierNode)

	if as := l.Peek(l.goal); as.Value != "as" {
		return n, err
	}
	l.Next(l.goal)

	identifierNode, err = ParseIdentifierNode(l)
	n.As = identifierNode.(IdentifierNode)

	return n, err
}
