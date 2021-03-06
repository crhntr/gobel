package es6

import (
	"fmt"

	"github.com/pkg/errors"
)

// ASTNode ...
type ASTNode interface {
	Positioner
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
func ParseIdentifierReferenceNode(l *Lexer) (IdentifierReferenceNode, error) {
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
func ParseBindingIdentifierNode(l *Lexer) (BindingIdentifierNode, error) {
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
func ParseIdentifierNode(l *Lexer) (IdentifierNode, error) {
	n := IdentifierNode{node: node{l.CurrentPosition()}}
	tt := l.Next(l.goal)
	if tt.Type == ReservedWordToken {
		return n, errors.Errorf("IdentifierNode must not be a ReservedWordToken found %q", tt.Value)
	}
	n.Name = tt.Value
	return n, nil
}

// ParseCoverParenthesizedExpressionAndArrowParameterListNode ...
func ParseCoverParenthesizedExpressionAndArrowParameterListNode(l *Lexer) (CoverParenthesizedExpressionAndArrowParameterListNode, error) {
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
func ParseParenthesizedExpressionNode(l *Lexer) (ParenthesizedExpressionNode, error) {
	n := ParenthesizedExpressionNode{node: node{l.CurrentPosition()}}

	tt := l.Next(l.goal)
	if tt.Value != "(" {
		return n, errors.New("expected '('")
	}
	var err error
	n.ExpressionNode, err = ParseExpressionNode(l)
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
func ParseElementListNode(l *Lexer) (ElementListNode, error) {
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
func ParseElisionNode(l *Lexer) (ElisionNode, error) {
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
func ParseSpreadElementNode(l *Lexer) (SpreadElementNode, error) {
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
func ParsePropertyDefinitionListNode(l *Lexer) (PropertyDefinitionListNode, error) {
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
func ParsePropertyDefinitionNode(l *Lexer) (PropertyDefinitionNode, error) {
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
func ParsePropertyNameNode(l *Lexer) (PropertyNameNode, error) {
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
func ParseLiteralPropertyNameNode(l *Lexer) (LiteralPropertyNameNode, error) {
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
func ParseComputedPropertyNameNode(l *Lexer) (ComputedPropertyNameNode, error) {
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
func ParseCoverInitializedNameNode(l *Lexer) (CoverInitializedNameNode, error) {
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
func ParseInitializerNode(l *Lexer) (InitializerNode, error) {
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
func ParseTemplateSpansNode(l *Lexer) (TemplateSpansNode, error) {
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
func ParseTemplateMiddleListNode(l *Lexer) (TemplateMiddleListNode, error) {
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
func ParseMemberExpressionNode(l *Lexer) (MemberExpressionNode, error) {
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
func ParseSuperPropertyNode(l *Lexer) (SuperPropertyNode, error) {
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
func ParseMetaPropertyNode(l *Lexer) (MetaPropertyNode, error) {
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
func ParseNewTargetNode(l *Lexer) (NewTargetNode, error) {
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
func ParseNewExpressionNode(l *Lexer) (NewExpressionNode, error) {
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
func ParseCallExpressionNode(l *Lexer) (CallExpressionNode, error) {
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
func ParseSuperCallNode(l *Lexer) (SuperCallNode, error) {
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
func ParseArgumentsNode(l *Lexer) (ArgumentsNode, error) {
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
func ParseArgumentListNode(l *Lexer) (ArgumentListNode, error) {
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
func ParseLeftHandSideExpressionNode(l *Lexer) (LeftHandSideExpressionNode, error) {
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
func ParsePostfixExpressionNode(l *Lexer) (PostfixExpressionNode, error) {
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
func ParseUnaryExpressionNode(l *Lexer) (UnaryExpressionNode, error) {
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
func ParseMultiplicativeExpressionNode(l *Lexer) (MultiplicativeExpressionNode, error) {
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
func ParseMultiplicativeOperatorNode(l *Lexer) (MultiplicativeOperatorNode, error) {
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
func ParseAdditiveExpressionNode(l *Lexer) (AdditiveExpressionNode, error) {
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
func ParseShiftExpressionNode(l *Lexer) (ShiftExpressionNode, error) {
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
func ParseRelationalExpressionNode(l *Lexer) (RelationalExpressionNode, error) {
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
func ParseEqualityExpressionNode(l *Lexer) (EqualityExpressionNode, error) {
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
func ParseBitwiseANDExpressionNode(l *Lexer) (BitwiseANDExpressionNode, error) {
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
func ParseBitwiseXORExpressionNode(l *Lexer) (BitwiseXORExpressionNode, error) {
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
func ParseBitwiseORExpressionNode(l *Lexer) (BitwiseORExpressionNode, error) {
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
func ParseLogicalANDExpressionNode(l *Lexer) (LogicalANDExpressionNode, error) {
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
func ParseLogicalORExpressionNode(l *Lexer) (LogicalORExpressionNode, error) {
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
func ParseConditionalExpressionNode(l *Lexer) (ConditionalExpressionNode, error) {
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
func ParseAssignmentExpressionNode(l *Lexer) (AssignmentExpressionNode, error) {
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
func ParseAssignmentOperatorNode(l *Lexer) (AssignmentOperatorNode, error) {
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
	isThis bool
	child  ASTNode
}

// ParseExpressionNode ...
func ParseExpressionNode(l *Lexer) (node ExpressionNode, err error) {
	if tok := l.Peek(l.goal); tok.Type == ReservedWordToken && tok.Value == "this" {
		node.isThis = true
		return node, nil
	}
	if child, err := ParseAssignmentExpressionNode(l); err == nil {
		node.child = child
	}
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
func ParseStatementNode(l *Lexer) (node StatementNode, err error) {
	defer func() { node.FilePosition = l.CurrentPosition() }()
	if node.child, err = ParseExpressionStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseBlockStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseVariableStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseIfStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseBreakableStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseContinueStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseBreakStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseReturnStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseWithStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseLabelledStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseThrowStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseTryStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseDebuggerStatementNode(l); err == nil {
		return
	} else if node.child, err = ParseEmptyStatementNode(l); err == nil {
		return
	} else {
		return node, err
	}
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
func ParseDeclarationNode(l *Lexer) (node DeclarationNode, err error) {
	defer func() { node.FilePosition = l.CurrentPosition() }()
	if node.child, err = ParseHoistableDeclarationNode(l); err == nil {
		return
	} else if node.child, err = ParseClassDeclarationNode(l); err == nil {
		return
	} else if node.child, err = ParseLexicalDeclarationNode(l); err == nil {
		return
	}
	return
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
func ParseHoistableDeclarationNode(l *Lexer) (node HoistableDeclarationNode, err error) {
	defer func() { node.FilePosition = l.CurrentPosition() }()
	if node.child, err = ParseFunctionDeclarationNode(l); err == nil {
		return
	}
	node.child, err = ParseGeneratorDeclarationNode(l)
	return
}

// BreakableStatementNode [Yield, Return] : [See clause 13]
//  IterationStatement[?Yield, ?Return]
//  SwitchStatement[?Yield, ?Return]
// implements: Parser and ASTNode
type BreakableStatementNode struct {
	node
}

// ParseBreakableStatementNode ...
func ParseBreakableStatementNode(l *Lexer) (BreakableStatementNode, error) {
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
func ParseBlockStatementNode(l *Lexer) (BlockStatementNode, error) {
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
func ParseBlockNode(l *Lexer) (BlockNode, error) {
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
func ParseStatementListItemNode(l *Lexer) (node StatementListItemNode, err error) {
	defer func() { node.FilePosition = l.CurrentPosition() }()
	if node.child, err = ParseStatementNode(l); err == nil {
		return
	}
	node.child, err = ParseDeclarationNode(l)
	return
}

// LexicalDeclarationNode [In, Yield] : [See 13.3.1]
//  LetOrConst BindingList[?In, ?Yield] ;
// implements: Parser and ASTNode
type LexicalDeclarationNode struct {
	node
}

// ParseLexicalDeclarationNode ...
func ParseLexicalDeclarationNode(l *Lexer) (LexicalDeclarationNode, error) {
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
func ParseLetOrConstNode(l *Lexer) (LetOrConstNode, error) {
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
func ParseBindingListNode(l *Lexer) (BindingListNode, error) {
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
func ParseLexicalBindingNode(l *Lexer) (LexicalBindingNode, error) {
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
func ParseVariableStatementNode(l *Lexer) (VariableStatementNode, error) {
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
func ParseVariableDeclarationListNode(l *Lexer) (VariableDeclarationListNode, error) {
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
func ParseVariableDeclarationNode(l *Lexer) (VariableDeclarationNode, error) {
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
func ParseBindingPatternNode(l *Lexer) (BindingPatternNode, error) {
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
func ParseObjectBindingPatternNode(l *Lexer) (ObjectBindingPatternNode, error) {
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
func ParseArrayBindingPatternNode(l *Lexer) (ArrayBindingPatternNode, error) {
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
func ParseBindingPropertyListNode(l *Lexer) (BindingPropertyListNode, error) {
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
func ParseBindingElementListNode(l *Lexer) (BindingElementListNode, error) {
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
func ParseBindingElisionElementNode(l *Lexer) (BindingElisionElementNode, error) {
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
func ParseBindingPropertyNode(l *Lexer) (BindingPropertyNode, error) {
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
func ParseBindingElementNode(l *Lexer) (BindingElementNode, error) {
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
func ParseSingleNameBindingNode(l *Lexer) (SingleNameBindingNode, error) {
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
func ParseBindingRestElementNode(l *Lexer) (BindingRestElementNode, error) {
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
func ParseEmptyStatementNode(l *Lexer) (EmptyStatementNode, error) {
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
func ParseExpressionStatementNode(l *Lexer) (node ExpressionStatementNode, err error) {
	tok := l.Peek(InputElementDiv)
	switch tok.Type {
	case ReservedWordToken:
		switch tok.Value {
		case "function", "class":
			return node, IncorrectTokenError(tok)
		case "let":
			tok2 := l.Peek(InputElementDiv)
			if tok2.Type == PunctuatorToken && tok2.Value == "[" {
				return node, IncorrectTokenError(tok2)
			}
		default:
		}
	case PunctuatorToken:
		if tok.Type == PunctuatorToken && tok.Value == "{" {
			return node, IncorrectTokenError(tok)
		}
	default:
	}
	node.child, err = ParseExpressionNode(l)
	if err != nil {
		return node, err
	}
	tok3 := l.Peek(InputElementDiv)
	if tok3.Type != PunctuatorToken || tok3.Value != ";" {
		return node, IncorrectTokenError(tok)
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
func ParseIfStatementNode(l *Lexer) (IfStatementNode, error) {
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
func ParseIterationStatementNode(l *Lexer) (IterationStatementNode, error) {
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
func ParseForDeclarationNode(l *Lexer) (ForDeclarationNode, error) {
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
func ParseForBindingNode(l *Lexer) (ForBindingNode, error) {
	panic("ParseForBindingNode not implemented")
	// return nil, nil
}

// ContinueStatementNode [Yield] : [See 13.8]
//  continue ;
//  continue [no LineTerminator here] LabelIdentifier[?Yield] ;
// implements: Parser and ASTNode
type ContinueStatementNode struct {
	node
	LabelIdentifier string
}

// ParseContinueStatementNode ...
func ParseContinueStatementNode(l *Lexer) (ContinueStatementNode, error) {
	node := ContinueStatementNode{}
	if token := l.Next(l.goal); token.Type != ReservedWordToken || token.Value != "continue" {
		return node, errors.Errorf("expected keyword %q", "continue")
	}
	idToken := l.Peek(l.goal)
	if idToken.Type == IdentifierNameToken {
		l.Next(l.goal)
		node.LabelIdentifier = idToken.Value
	}
	return node, nil
}

// BreakStatementNode [Yield] : [See 13.9]
//  break ;
//  break [no LineTerminator here] LabelIdentifier[?Yield] ;
// implements: Parser and ASTNode
type BreakStatementNode struct {
	node
}

// ParseBreakStatementNode ...
func ParseBreakStatementNode(l *Lexer) (BreakStatementNode, error) {
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
func ParseReturnStatementNode(l *Lexer) (ReturnStatementNode, error) {
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
func ParseWithStatementNode(l *Lexer) (WithStatementNode, error) {
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
func ParseSwitchStatementNode(l *Lexer) (SwitchStatementNode, error) {
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
func ParseCaseBlockNode(l *Lexer) (CaseBlockNode, error) {
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
func ParseCaseClausesNode(l *Lexer) (CaseClausesNode, error) {
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
func ParseCaseClauseNode(l *Lexer) (CaseClauseNode, error) {
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
func ParseDefaultClauseNode(l *Lexer) (DefaultClauseNode, error) {
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
func ParseLabelledStatementNode(l *Lexer) (LabelledStatementNode, error) {
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
func ParseLabelledItemNode(l *Lexer) (LabelledItemNode, error) {
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
func ParseThrowStatementNode(l *Lexer) (ThrowStatementNode, error) {
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
func ParseTryStatementNode(l *Lexer) (TryStatementNode, error) {
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
func ParseCatchNode(l *Lexer) (CatchNode, error) {
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
func ParseFinallyNode(l *Lexer) (FinallyNode, error) {
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
func ParseCatchParameterNode(l *Lexer) (CatchParameterNode, error) {
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
func ParseDebuggerStatementNode(l *Lexer) (DebuggerStatementNode, error) {
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
func ParseFunctionDeclarationNode(l *Lexer) (FunctionDeclarationNode, error) {
	var (
		node = FunctionDeclarationNode{}
		// err error
	)
	tokPeek0 := l.Peek(InputElementDiv)
	if tokPeek0.Type != ReservedWordToken {
		return node, IncorrectTokenError(tokPeek0)
	}
	if tokPeek0.Value != "function" {
		return node, IncorrectTokenError(tokPeek0)
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
func ParseStrictFormalParametersNode(l *Lexer) (StrictFormalParametersNode, error) {
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
func ParseFormalParametersNode(l *Lexer) (FormalParametersNode, error) {
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
func ParseFormalParameterListNode(l *Lexer) (FormalParameterListNode, error) {
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
func ParseFormalsListNode(l *Lexer) (FormalsListNode, error) {
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
func ParseFunctionRestParameterNode(l *Lexer) (FunctionRestParameterNode, error) {
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
func ParseFormalParameterNode(l *Lexer) (FormalParameterNode, error) {
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
func ParseFunctionBodyNode(l *Lexer) (FunctionBodyNode, error) {
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
func ParseFunctionStatementListNode(l *Lexer) (FunctionStatementListNode, error) {
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
func ParseArrowFunctionNode(l *Lexer) (ArrowFunctionNode, error) {
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
func ParseArrowParametersNode(l *Lexer) (ArrowParametersNode, error) {
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
func ParseConciseBodyNode(l *Lexer) (ConciseBodyNode, error) {
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
func ParseMethodDefinitionNode(l *Lexer) (MethodDefinitionNode, error) {
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
func ParsePropertySetParameterListNode(l *Lexer) (PropertySetParameterListNode, error) {
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
func ParseGeneratorMethodNode(l *Lexer) (GeneratorMethodNode, error) {
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
func ParseGeneratorDeclarationNode(l *Lexer) (GeneratorDeclarationNode, error) {
	panic("ParseGeneratorDeclarationNode not implemented")
	// return nil, nil
}

// ParseGeneratorBodyNode ...
func ParseGeneratorBodyNode(l *Lexer) (GeneratorBodyNode, error) {
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
func ParseYieldExpressionNode(l *Lexer) (YieldExpressionNode, error) {
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
func ParseClassDeclarationNode(l *Lexer) (ClassDeclarationNode, error) {
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
func ParseClassTailNode(l *Lexer) (ClassTailNode, error) {
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
func ParseClassHeritageNode(l *Lexer) (ClassHeritageNode, error) {
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
func ParseClassBodyNode(l *Lexer) (ClassBodyNode, error) {
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
func ParseClassElementListNode(l *Lexer) (ClassElementListNode, error) {
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
func ParseClassElementNode(l *Lexer) (ClassElementNode, error) {
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
func ParseScriptNode(l *Lexer) (ScriptNode, error) {
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
func ParseScriptBodyNode(l *Lexer) (ScriptBodyNode, error) {
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
func ParseModuleNode(l *Lexer) (ModuleNode, error) {
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
func ParseModuleBodyNode(l *Lexer) (ModuleBodyNode, error) {
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
func ParseModuleItemListNode(l *Lexer) (ModuleItemListNode, error) {
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
func ParseModuleItemNode(l *Lexer) (ModuleItemNode, error) {
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
func ParseImportDeclarationNode(l *Lexer) (ImportDeclarationNode, error) {
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
func ParseImportClauseNode(l *Lexer) (ImportClauseNode, error) {
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
func ParseImportsListNode(l *Lexer) (ImportsListNode, error) {
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
func ParseImportSpecifierNode(l *Lexer) (ImportSpecifierNode, error) {
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
func ParseModuleSpecifierNode(l *Lexer) (ModuleSpecifierNode, error) {
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
func ParseImportedBindingNode(l *Lexer) (ImportedBindingNode, error) {
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
func ParseExportDeclarationNode(l *Lexer) (ExportDeclarationNode, error) {
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
func ParseExportClauseNode(l *Lexer) (ExportClauseNode, error) {
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
func ParseExportsListNode(l *Lexer) (ExportsListNode, error) {
	n := ExportsListNode{node: node{l.CurrentPosition()}}

	n.List = []ExportSpecifierNode{}

	for {
		exportSpecifier, err := ParseExportSpecifierNode(l)
		if err != nil {
			return n, err
		}
		n.List = append(n.List, exportSpecifier)

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
func ParseExportSpecifierNode(l *Lexer) (ExportSpecifierNode, error) {
	n := ExportSpecifierNode{node: node{l.CurrentPosition()}}

	identifierNode, err := ParseIdentifierNode(l)
	if err != nil {
		return n, err
	}
	n.IdentifierNode = identifierNode

	if as := l.Peek(l.goal); as.Value != "as" {
		return n, nil
	}
	l.Next(l.goal)

	identifierNode, err = ParseIdentifierNode(l)
	n.As = identifierNode

	return n, err
}
