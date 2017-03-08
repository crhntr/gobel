package es6

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
  filename string
  line, column, offset int
}

func (n *node) Position() (filename string, offset int, line int, column int) {
  return n.filename, n.offset, n.line, n.column
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

// Parse implements Parser
func (n IdentifierReferenceNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BindingIdentifierNode [Yield] : [See 12.1]
//  Identifier
//  [~Yield] yield
//
// implements: Parser and ASTNode
type BindingIdentifierNode struct {
  node
}

// Parse implements Parser
func (n BindingIdentifierNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// LabelIdentifierNode [Yield] : [See 12.1]
//  Identifier
//  [~Yield] yield
// implements: Parser and ASTNode
type LabelIdentifierNode struct {
  node
}

// Parse implements Parser
func (n LabelIdentifierNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// IdentifierNode  : [See 12.1]
//  IdentifierName but not ReservedWord
// implements: Parser and ASTNode
type IdentifierNode struct {
  node
}

// Parse implements Parser
func (n IdentifierNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

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

// Parse implements Parser
func (n PrimaryExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n CoverParenthesizedExpressionAndArrowParameterListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ParenthesizedExpressionNode [Yield] : [See 12.2]
//  ( Expression[In, ?Yield] )
// implements: Parser and ASTNode
type ParenthesizedExpressionNode struct {
  node
}

// Parse implements Parser
func (n ParenthesizedExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n LiteralNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ArrayLiteralNode [Yield] : [See 12.2.5]
//  [ Elisionopt ]
//  [ ElementList[?Yield] ]
//  [ ElementList[?Yield] , Elisionopt ]
// implements: Parser and ASTNode
type ArrayLiteralNode struct {
  node
}

// Parse implements Parser
func (n ArrayLiteralNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n ElementListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ElisionNode  : [See 12.2.5]
//  ,
//  Elision ,
// implements: Parser and ASTNode
type ElisionNode struct {
  node
}

// Parse implements Parser
func (n ElisionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// SpreadElementNode [Yield] : [See 12.2.5]
//  ... AssignmentExpression[In, ?Yield]
// implements: Parser and ASTNode
type SpreadElementNode struct {
  node
}

// Parse implements Parser
func (n SpreadElementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ObjectLiteralNode [Yield] : [See 12.2.6]
//  { }
//  { PropertyDefinitionList[?Yield] }
//  { PropertyDefinitionList[?Yield] , }
// implements: Parser and ASTNode
type ObjectLiteralNode struct {
  node
}

// Parse implements Parser
func (n ObjectLiteralNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// PropertyDefinitionListNode [Yield] : [See 12.2.6]
//  PropertyDefinition[?Yield]
//  PropertyDefinitionList[?Yield] , PropertyDefinition[?Yield]
// implements: Parser and ASTNode
type PropertyDefinitionListNode struct {
  node
}

// Parse implements Parser
func (n PropertyDefinitionListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n PropertyDefinitionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// PropertyNameNode [Yield] : [See 12.2.6]
//  LiteralPropertyName
//  ComputedPropertyName[?Yield]
// implements: Parser and ASTNode
type PropertyNameNode struct {
  node
}

// Parse implements Parser
func (n PropertyNameNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// LiteralPropertyNameNode  : [See 12.2.6]
//  IdentifierName
//  StringLiteral
//  NumericLiteral
// implements: Parser and ASTNode
type LiteralPropertyNameNode struct {
  node
}

// Parse implements Parser
func (n LiteralPropertyNameNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ComputedPropertyNameNode [Yield] : [See 12.2.6]
//  [ AssignmentExpression[In, ?Yield] ]
// implements: Parser and ASTNode
type ComputedPropertyNameNode struct {
  node
}

// Parse implements Parser
func (n ComputedPropertyNameNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// CoverInitializedNameNode [Yield] : [See 12.2.6]
//  IdentifierReference[?Yield] Initializer[In, ?Yield]
// implements: Parser and ASTNode
type CoverInitializedNameNode struct {
  node
}

// Parse implements Parser
func (n CoverInitializedNameNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// InitializerNode [In, Yield] : [See 12.2.6]
//  = AssignmentExpression[?In, ?Yield]
// implements: Parser and ASTNode
type InitializerNode struct {
  node
}

// Parse implements Parser
func (n InitializerNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// TemplateLiteralNode [Yield] : [See 12.2.9]
//  NoSubstitutionTemplate
//  TemplateHead Expression[In, ?Yield] TemplateSpans[?Yield]
// implements: Parser and ASTNode
type TemplateLiteralNode struct {
  node
}

// Parse implements Parser
func (n TemplateLiteralNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// TemplateSpansNode [Yield] : [See 12.2.9]
//  TemplateTail
//  TemplateMiddleList[?Yield] TemplateTail
// implements: Parser and ASTNode
type TemplateSpansNode struct {
  node
}

// Parse implements Parser
func (n TemplateSpansNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// TemplateMiddleListNode [Yield] : [See 12.2.9]
//  TemplateMiddle Expression[In, ?Yield]
//  TemplateMiddleList[?Yield] TemplateMiddle Expression[In, ?Yield]
// implements: Parser and ASTNode
type TemplateMiddleListNode struct {
  node
}

// Parse implements Parser
func (n TemplateMiddleListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n MemberExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// SuperPropertyNode [Yield] : [See 12.3]
//  super [ Expression[In, ?Yield] ]
//  super . IdentifierName
// implements: Parser and ASTNode
type SuperPropertyNode struct {
  node
}

// Parse implements Parser
func (n SuperPropertyNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// MetaPropertyNode  : [See 12.3]
//  NewTarget
// implements: Parser and ASTNode
type MetaPropertyNode struct {
  node
}

// Parse implements Parser
func (n MetaPropertyNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// NewTargetNode  : [See 12.3]
//  new . target
// implements: Parser and ASTNode
type NewTargetNode struct {
  node
}

// Parse implements Parser
func (n NewTargetNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// NewExpressionNode [Yield] : [See 12.3]
//  MemberExpression[?Yield]
//  new NewExpression[?Yield]
// implements: Parser and ASTNode
type NewExpressionNode struct {
  node
}

// Parse implements Parser
func (n NewExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n CallExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// SuperCallNode [Yield] : [See 12.3]
//  super Arguments[?Yield]
// implements: Parser and ASTNode
type SuperCallNode struct {
  node
}

// Parse implements Parser
func (n SuperCallNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ArgumentsNode [Yield] : [See 12.3]
//  ( )
//  ( ArgumentList[?Yield] )
// implements: Parser and ASTNode
type ArgumentsNode struct {
  node
}

// Parse implements Parser
func (n ArgumentsNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n ArgumentListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// LeftHandSideExpressionNode [Yield] : [See 12.3]
//  NewExpression[?Yield]
//  CallExpression[?Yield]
// implements: Parser and ASTNode
type LeftHandSideExpressionNode struct {
  node
}

// Parse implements Parser
func (n LeftHandSideExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// PostfixExpressionNode [Yield] : [See 12.4]
//  LeftHandSideExpression[?Yield]
//  LeftHandSideExpression[?Yield] [no LineTerminator here] ++
//  LeftHandSideExpression[?Yield] [no LineTerminator here] --
// implements: Parser and ASTNode
type PostfixExpressionNode struct {
  node
}

// Parse implements Parser
func (n PostfixExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n UnaryExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// MultiplicativeExpressionNode [Yield] : [See 12.6]
//  UnaryExpression[?Yield]
//  MultiplicativeExpression[?Yield] MultiplicativeOperator UnaryExpression[?Yield]
// implements: Parser and ASTNode
type MultiplicativeExpressionNode struct {
  node
}

// Parse implements Parser
func (n MultiplicativeExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// MultiplicativeOperatorNode  : one of [See 12.6]
//  * / %
// implements: Parser and ASTNode
type MultiplicativeOperatorNode struct {
  node
}

// Parse implements Parser
func (n MultiplicativeOperatorNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// AdditiveExpressionNode [Yield] : [See 12.7]
//  MultiplicativeExpression[?Yield]
//  AdditiveExpression[?Yield] + MultiplicativeExpression[?Yield]
//  AdditiveExpression[?Yield] - MultiplicativeExpression[?Yield]
// implements: Parser and ASTNode
type AdditiveExpressionNode struct {
  node
}

// Parse implements Parser
func (n AdditiveExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n ShiftExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n RelationalExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n EqualityExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BitwiseANDExpressionNode [In, Yield] : [See 12.11]
//  EqualityExpression[?In, ?Yield]
//  BitwiseANDExpression[?In, ?Yield] & EqualityExpression[?In, ?Yield]
// implements: Parser and ASTNode
type BitwiseANDExpressionNode struct {
  node
}

// Parse implements Parser
func (n BitwiseANDExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BitwiseXORExpressionNode [In, Yield] : [See 12.11]
//  BitwiseANDExpression[?In, ?Yield]
//  BitwiseXORExpression[?In, ?Yield] ^ BitwiseANDExpression[?In, ?Yield]
// implements: Parser and ASTNode
type BitwiseXORExpressionNode struct {
  node
}

// Parse implements Parser
func (n BitwiseXORExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BitwiseORExpressionNode [In, Yield] : [See 12.11]
//  BitwiseXORExpression[?In, ?Yield]
//  BitwiseORExpression[?In, ?Yield] | BitwiseXORExpression[?In, ?Yield]
// implements: Parser and ASTNode
type BitwiseORExpressionNode struct {
  node
}

// Parse implements Parser
func (n BitwiseORExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// LogicalANDExpressionNode [In, Yield] : [See 12.12]
//  BitwiseORExpression[?In, ?Yield]
//  LogicalANDExpression[?In, ?Yield] && BitwiseORExpression[?In, ?Yield]
// implements: Parser and ASTNode
type LogicalANDExpressionNode struct {
  node
}

// Parse implements Parser
func (n LogicalANDExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// LogicalORExpressionNode [In, Yield] : [See 12.12]
//  LogicalANDExpression[?In, ?Yield]
//  LogicalORExpression[?In, ?Yield] || LogicalANDExpression[?In, ?Yield]
// implements: Parser and ASTNode
type LogicalORExpressionNode struct {
  node
}

// Parse implements Parser
func (n LogicalORExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ConditionalExpressionNode [In, Yield] : [See 12.13]
//  LogicalORExpression[?In, ?Yield]
//  LogicalORExpression[?In,?Yield] ? AssignmentExpression[In, ?Yield] : AssignmentExpression[?In, ?Yield]
// implements: Parser and ASTNode
type ConditionalExpressionNode struct {
  node
}

// Parse implements Parser
func (n ConditionalExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n AssignmentExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// AssignmentOperatorNode  : one of [See 12.14]
//  *= /= %= += -= <<= >>= >>>= &= ^= |=
// implements: Parser and ASTNode
type AssignmentOperatorNode struct {
  node
}

// Parse implements Parser
func (n AssignmentOperatorNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ExpressionNode [In, Yield] : [See 12.15]
//  AssignmentExpression[?In, ?Yield]
//  Expression[?In, ?Yield] , AssignmentExpression[?In, ?Yield]
// implements: Parser and ASTNode
type ExpressionNode struct {
  node
}

// Parse implements Parser
func (n ExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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
  node
}

// Parse implements Parser
func (n StatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// DeclarationNode [Yield] : [See clause 13]
//  HoistableDeclaration[?Yield]
//  ClassDeclaration[?Yield]
//  LexicalDeclaration[In, ?Yield]
// implements: Parser and ASTNode
type DeclarationNode struct {
  node
}

// Parse implements Parser
func (n DeclarationNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// HoistableDeclarationNode [Yield, Default] : [See clause 13]
//  FunctionDeclaration[?Yield,?Default]
//  GeneratorDeclaration[?Yield, ?Default]
// implements: Parser and ASTNode
type HoistableDeclarationNode struct {
  node
}

// Parse implements Parser
func (n HoistableDeclarationNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BreakableStatementNode [Yield, Return] : [See clause 13]
//  IterationStatement[?Yield, ?Return]
//  SwitchStatement[?Yield, ?Return]
// implements: Parser and ASTNode
type BreakableStatementNode struct {
  node
}

// Parse implements Parser
func (n BreakableStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BlockStatementNode [Yield, Return] : [See 13.2]
//  Block[?Yield, ?Return]
// implements: Parser and ASTNode
type BlockStatementNode struct {
  node
}

// Parse implements Parser
func (n BlockStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BlockNode [Yield, Return] : [See 13.2]
//  { StatementList[?Yield, ?Return]opt }
// implements: Parser and ASTNode
type BlockNode struct {
  node
}

// Parse implements Parser
func (n BlockNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// StatementListNode [Yield, Return] : [See 13.2]
//  StatementListItem[?Yield, ?Return]
//  StatementList[?Yield, ?Return] StatementListItem[?Yield, ?Return]
// implements: Parser and ASTNode
type StatementListNode struct {
  node
}

// Parse implements Parser
func (n StatementListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// StatementListItemNode [Yield, Return] : [See 13.2]
//  Statement[?Yield, ?Return]
//  Declaration[?Yield]
// implements: Parser and ASTNode
type StatementListItemNode struct {
  node
}

// Parse implements Parser
func (n StatementListItemNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// LexicalDeclarationNode [In, Yield] : [See 13.3.1]
//  LetOrConst BindingList[?In, ?Yield] ;
// implements: Parser and ASTNode
type LexicalDeclarationNode struct {
  node
}

// Parse implements Parser
func (n LexicalDeclarationNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// LetOrConstNode  : [See 13.3.1]
//  let
//  const
// implements: Parser and ASTNode
type LetOrConstNode struct {
  node
}

// Parse implements Parser
func (n LetOrConstNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BindingListNode [In, Yield] : [See 13.3.1]
//  LexicalBinding[?In, ?Yield]
//  BindingList[?In, ?Yield] , LexicalBinding[?In, ?Yield]
// implements: Parser and ASTNode
type BindingListNode struct {
  node
}

// Parse implements Parser
func (n BindingListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// LexicalBindingNode [In, Yield] : [See 13.3.1]
//  BindingIdentifier[?Yield] Initializer[?In, ?Yield]opt
//  BindingPattern[?Yield] Initializer[?In, ?Yield]
// implements: Parser and ASTNode
type LexicalBindingNode struct {
  node
}

// Parse implements Parser
func (n LexicalBindingNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// VariableStatementNode [Yield] : [See 13.3.2]
//  var VariableDeclarationList[In, ?Yield] ;
// implements: Parser and ASTNode
type VariableStatementNode struct {
  node
}

// Parse implements Parser
func (n VariableStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// VariableDeclarationListNode [In, Yield] : [See 13.3.2]
//  VariableDeclaration[?In, ?Yield]
//  VariableDeclarationList[?In, ?Yield] , VariableDeclaration[?In, ?Yield]
// implements: Parser and ASTNode
type VariableDeclarationListNode struct {
  node
}

// Parse implements Parser
func (n VariableDeclarationListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// VariableDeclarationNode [In, Yield] : [See 13.3.2]
//  BindingIdentifier[?Yield] Initializer[?In, ?Yield]opt
//  BindingPattern[?Yield] Initializer[?In, ?Yield]
// implements: Parser and ASTNode
type VariableDeclarationNode struct {
  node
}

// Parse implements Parser
func (n VariableDeclarationNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BindingPatternNode [Yield] : [See 13.3.3]
//  ObjectBindingPattern[?Yield]
//  ArrayBindingPattern[?Yield]
// implements: Parser and ASTNode
type BindingPatternNode struct {
  node
}

// Parse implements Parser
func (n BindingPatternNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ObjectBindingPatternNode [Yield] : [See 13.3.3]
//  { }
//  { BindingPropertyList[?Yield] }
//  { BindingPropertyList[?Yield] , }
// implements: Parser and ASTNode
type ObjectBindingPatternNode struct {
  node
}

// Parse implements Parser
func (n ObjectBindingPatternNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ArrayBindingPatternNode [Yield] : [See 13.3.3]
//  [ Elisionopt BindingRestElement[?Yield]opt ]
//  [ BindingElementList[?Yield] ]
//  [ BindingElementList[?Yield] , Elisionopt BindingRestElement[?Yield]opt ]
// implements: Parser and ASTNode
type ArrayBindingPatternNode struct {
  node
}

// Parse implements Parser
func (n ArrayBindingPatternNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BindingPropertyListNode [Yield] : [See 13.3.3]
//  BindingProperty[?Yield]
//  BindingPropertyList[?Yield] , BindingProperty[?Yield]
// implements: Parser and ASTNode
type BindingPropertyListNode struct {
  node
}

// Parse implements Parser
func (n BindingPropertyListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BindingElementListNode [Yield] : [See 13.3.3]
//  BindingElisionElement[?Yield]
//  BindingElementList[?Yield] , BindingElisionElement[?Yield]
// implements: Parser and ASTNode
type BindingElementListNode struct {
  node
}

// Parse implements Parser
func (n BindingElementListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BindingElisionElementNode [Yield] : [See 13.3.3]
//  Elisionopt BindingElement[?Yield]
// implements: Parser and ASTNode
type BindingElisionElementNode struct {
  node
}

// Parse implements Parser
func (n BindingElisionElementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BindingPropertyNode [Yield] : [See 13.3.3]
//  SingleNameBinding[?Yield]
//  PropertyName[?Yield] : BindingElement[?Yield]
// implements: Parser and ASTNode
type BindingPropertyNode struct {
  node
}

// Parse implements Parser
func (n BindingPropertyNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BindingElementNode [Yield] : [See 13.3.3]
//  SingleNameBinding[?Yield]
//  BindingPattern[?Yield] Initializer[In, ?Yield]opt
// implements: Parser and ASTNode
type BindingElementNode struct {
  node
}

// Parse implements Parser
func (n BindingElementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// SingleNameBindingNode [Yield] : [See 13.3.3]
//  BindingIdentifier[?Yield] Initializer[In, ?Yield]opt
// implements: Parser and ASTNode
type SingleNameBindingNode struct {
  node
}

// Parse implements Parser
func (n SingleNameBindingNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BindingRestElementNode [Yield] : [See 13.3.3]
//  ... BindingIdentifier[?Yield]
// implements: Parser and ASTNode
type BindingRestElementNode struct {
  node
}

// Parse implements Parser
func (n BindingRestElementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// EmptyStatementNode  : [See 13.4]
//  ;
// implements: Parser and ASTNode
type EmptyStatementNode struct {
  node
}

// Parse implements Parser
func (n EmptyStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ExpressionStatementNode [Yield] : [See 13.5]
//  [lookahead ∉ {{, function, class, let [}] Expression[In, ?Yield] ;
// implements: Parser and ASTNode
type ExpressionStatementNode struct {
  node
}

// Parse implements Parser
func (n ExpressionStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// IfStatementNode [Yield, Return] : [See 13.6]
//  if ( Expression[In, ?Yield] ) Statement[?Yield, ?Return] else Statement[?Yield, ?Return]
//  if ( Expression[In, ?Yield] ) Statement[?Yield, ?Return]
// implements: Parser and ASTNode
type IfStatementNode struct {
  node
}

// Parse implements Parser
func (n IfStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n IterationStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ForDeclarationNode [Yield] : [See 13.7]
//  LetOrConst ForBinding[?Yield]
// implements: Parser and ASTNode
type ForDeclarationNode struct {
  node
}

// Parse implements Parser
func (n ForDeclarationNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ForBindingNode [Yield] : [See 13.7]
//  BindingIdentifier[?Yield]
//  BindingPattern[?Yield]
// implements: Parser and ASTNode
type ForBindingNode struct {
  node
}

// Parse implements Parser
func (n ForBindingNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ContinueStatementNode [Yield] : [See 13.8]
//  continue ;
//  continue [no LineTerminator here] LabelIdentifier[?Yield] ;
// implements: Parser and ASTNode
type ContinueStatementNode struct {
  node
}

// Parse implements Parser
func (n ContinueStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// BreakStatementNode [Yield] : [See 13.9]
//  break ;
//  break [no LineTerminator here] LabelIdentifier[?Yield] ;
// implements: Parser and ASTNode
type BreakStatementNode struct {
  node
}

// Parse implements Parser
func (n BreakStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ReturnStatementNode [Yield] : [See 13.10]
//  return ;
//  return [no LineTerminator here] Expression[In, ?Yield] ;
// implements: Parser and ASTNode
type ReturnStatementNode struct {
  node
}

// Parse implements Parser
func (n ReturnStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// WithStatementNode [Yield, Return] : [See 13.11]
//  with ( Expression[In, ?Yield] ) Statement[?Yield, ?Return]
// implements: Parser and ASTNode
type WithStatementNode struct {
  node
}

// Parse implements Parser
func (n WithStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// SwitchStatementNode [Yield, Return] : [See 13.12]
//  switch ( Expression[In, ?Yield] ) CaseBlock[?Yield, ?Return]
// implements: Parser and ASTNode
type SwitchStatementNode struct {
  node
}

// Parse implements Parser
func (n SwitchStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// CaseBlockNode [Yield, Return] : [See 13.12]
//  { CaseClauses[?Yield, ?Return]opt }
//  { CaseClauses[?Yield, ?Return]opt DefaultClause[?Yield, ?Return] CaseClauses[?Yield, ?Return]opt }
// implements: Parser and ASTNode
type CaseBlockNode struct {
  node
}

// Parse implements Parser
func (n CaseBlockNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// CaseClausesNode [Yield, Return] : [See 13.12]
//  CaseClause[?Yield, ?Return]
//  CaseClauses[?Yield, ?Return] CaseClause[?Yield, ?Return]
// implements: Parser and ASTNode
type CaseClausesNode struct {
  node
}

// Parse implements Parser
func (n CaseClausesNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// CaseClauseNode [Yield, Return] : [See 13.12]
//  case Expression[In, ?Yield] : StatementList[?Yield, ?Return]opt
// implements: Parser and ASTNode
type CaseClauseNode struct {
  node
}

// Parse implements Parser
func (n CaseClauseNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}
// DefaultClauseNode [Yield, Return] : [See 13.12]
//  default : StatementList[?Yield, ?Return]opt
// implements: Parser and ASTNode
type DefaultClauseNode struct {
  node
}

// Parse implements Parser
func (n DefaultClauseNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// LabelledStatementNode [Yield, Return] : [See 13.13]
//  LabelIdentifier[?Yield] : LabelledItem[?Yield, ?Return]
// implements: Parser and ASTNode
type LabelledStatementNode struct {
  node
}

// Parse implements Parser
func (n LabelledStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// LabelledItemNode [Yield, Return] : [See 13.13]
//  Statement[?Yield, ?Return]
//  FunctionDeclaration[?Yield]
// implements: Parser and ASTNode
type LabelledItemNode struct {
  node
}

// Parse implements Parser
func (n LabelledItemNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ThrowStatementNode [Yield] : [See 13.14]
//  throw [no LineTerminator here] Expression[In, ?Yield] ;
// implements: Parser and ASTNode
type ThrowStatementNode struct {
  node
}

// Parse implements Parser
func (n ThrowStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// TryStatementNode [Yield, Return] : [See 13.15]
//  try Block[?Yield, ?Return] Catch[?Yield, ?Return]
//  try Block[?Yield, ?Return] Finally[?Yield, ?Return]
//  try Block[?Yield, ?Return] Catch[?Yield, ?Return] Finally[?Yield, ?Return]
// implements: Parser and ASTNode
type TryStatementNode struct {
  node
}

// Parse implements Parser
func (n TryStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// CatchNode [Yield, Return] : [See 13.15]
//  catch ( CatchParameter[?Yield] ) Block[?Yield, ?Return]
// implements: Parser and ASTNode
type CatchNode struct {
  node
}

// Parse implements Parser
func (n CatchNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// FinallyNode [Yield, Return] : [See 13.15]
//  finally Block[?Yield, ?Return]
// implements: Parser and ASTNode
type FinallyNode struct {
  node
}

// Parse implements Parser
func (n FinallyNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// CatchParameterNode [Yield] : [See 13.15]
//  BindingIdentifier[?Yield]
//  BindingPattern[?Yield]
// implements: Parser and ASTNode
type CatchParameterNode struct {
  node
}

// Parse implements Parser
func (n CatchParameterNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// DebuggerStatementNode  : [See 13.16]
//  debugger ;
// implements: Parser and ASTNode
type DebuggerStatementNode struct {
  node
}

// Parse implements Parser
func (n DebuggerStatementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

//
// A.4 Functions and Classes
//

// FunctionDeclarationNode [Yield, Default] :
//  function BindingIdentifier[?Yield] ( FormalParameters ) { FunctionBody }
//  [+Default] function ( FormalParameters ) { FunctionBody }
// implements: Parser and ASTNode
type FunctionDeclarationNode struct {
  node
}

// Parse implements Parser
func (n FunctionDeclarationNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// FunctionExpressionNode  : [See 14.1]
//  function BindingIdentifieropt ( FormalParameters ) { FunctionBody }
// implements: Parser and ASTNode
type FunctionExpressionNode struct {
  node
}

// Parse implements Parser
func (n FunctionExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}
// StrictFormalParametersNode [Yield] : [See 14.1]
//  FormalParameters[?Yield]
// implements: Parser and ASTNode
type StrictFormalParametersNode struct {
  node
}

// Parse implements Parser
func (n StrictFormalParametersNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}
// FormalParametersNode [Yield] : [See 14.1]
//  [empty]
//  FormalParameterList[?Yield]
// implements: Parser and ASTNode
type FormalParametersNode struct {
  node
}

// Parse implements Parser
func (n FormalParametersNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// FormalParameterListNode [Yield] : [See 14.1]
//  FunctionRestParameter[?Yield]
//  FormalsList[?Yield]
//  FormalsList[?Yield] , FunctionRestParameter[?Yield]
// implements: Parser and ASTNode
type FormalParameterListNode struct {
  node
}

// Parse implements Parser
func (n FormalParameterListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}
// FormalsListNode [Yield] : [See 14.1]
//  FormalParameter[?Yield]
//  FormalsList[?Yield] , FormalParameter[?Yield]
// implements: Parser and ASTNode
type FormalsListNode struct {
  node
}

// Parse implements Parser
func (n FormalsListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// FunctionRestParameterNode [Yield] : [See 14.1]
//  BindingRestElement[?Yield]
// implements: Parser and ASTNode
type FunctionRestParameterNode struct {
  node
}

// Parse implements Parser
func (n FunctionRestParameterNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// FormalParameterNode [Yield] : [See 14.1]
//  BindingElement[?Yield]
// implements: Parser and ASTNode
type FormalParameterNode struct {
  node
}

// Parse implements Parser
func (n FormalParameterNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}
// FunctionBodyNode [Yield] : [See 14.1]
//  FunctionStatementList[?Yield]
// implements: Parser and ASTNode
type FunctionBodyNode struct {
  node
}

// Parse implements Parser
func (n FunctionBodyNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}
// FunctionStatementListNode [Yield] : [See 14.1]
//  StatementList[?Yield, Return]opt
// implements: Parser and ASTNode
type FunctionStatementListNode struct {
  node
}

// Parse implements Parser
func (n FunctionStatementListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}
// ArrowFunctionNode [In, Yield] : [See 14.2]
//  ArrowParameters[?Yield] [no LineTerminator here] => ConciseBody[?In]
// implements: Parser and ASTNode
type ArrowFunctionNode struct {
  node
}

// Parse implements Parser
func (n ArrowFunctionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n ArrowParametersNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n ConciseBodyNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n MethodDefinitionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// PropertySetParameterListNode  : [See 14.3]
// FormalParameter
// implements: Parser and ASTNode
type PropertySetParameterListNode struct {
  node
}

// Parse implements Parser
func (n PropertySetParameterListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// GeneratorMethodNode [Yield] : [See 14.4]
//  * PropertyName[?Yield] ( StrictFormalParameters[Yield] ) { GeneratorBody }
// implements: Parser and ASTNode
type GeneratorMethodNode struct {
  node
}

// Parse implements Parser
func (n GeneratorMethodNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// GeneratorDeclarationNode [Yield, Default] : [See 14.4]
//  function * BindingIdentifier[?Yield] ( FormalParameters[Yield] ) { GeneratorBody }
//  [+Default] function * ( FormalParameters[Yield] ) { GeneratorBody }
// implements: Parser and ASTNode
type GeneratorDeclarationNode struct {
  node
}

// Parse implements Parser
func (n GeneratorDeclarationNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// GeneratorExpressionNode  : [See 14.4]
//  function * BindingIdentifier[Yield]opt ( FormalParameters[Yield] ) { GeneratorBody }
// implements: Parser and ASTNode
type GeneratorExpressionNode struct {
  node
}

// Parse implements Parser
func (n GeneratorExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}
// GeneratorBodyNode  : [See 14.4]
//  FunctionBody[Yield]
// implements: Parser and ASTNode
type GeneratorBodyNode struct {
  node
}

// Parse implements Parser
func (n GeneratorBodyNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// YieldExpressionNode [In] : [See 14.4]
//  yield
//  yield [no LineTerminator here] AssignmentExpression[?In, Yield]
//  yield [no LineTerminator here] * AssignmentExpression[?In, Yield]
// implements: Parser and ASTNode
type YieldExpressionNode struct {
  node
}

// Parse implements Parser
func (n YieldExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ClassDeclarationNode [Yield, Default] : [See 14.5]
//  class BindingIdentifier[?Yield] ClassTail[?Yield]
//  [+Default] class ClassTail[?Yield]
// implements: Parser and ASTNode
type ClassDeclarationNode struct {
  node
}

// Parse implements Parser
func (n ClassDeclarationNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ClassExpressionNode [Yield] : [See 14.5]
//  class BindingIdentifier[?Yield]opt ClassTail[?Yield]
// implements: Parser and ASTNode
type ClassExpressionNode struct {
  node
}

// Parse implements Parser
func (n ClassExpressionNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ClassTailNode [Yield] : [See 14.5]
//  ClassHeritage[?Yield]opt { ClassBody[?Yield]opt }
// implements: Parser and ASTNode
type ClassTailNode struct {
  node
}

// Parse implements Parser
func (n ClassTailNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ClassHeritageNode [Yield] : [See 14.5]
//  extends LeftHandSideExpression[?Yield]
// implements: Parser and ASTNode
type ClassHeritageNode struct {
  node
}

// Parse implements Parser
func (n ClassHeritageNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ClassBodyNode [Yield] : [See 14.5]
//  ClassElementList[?Yield]
// implements: Parser and ASTNode
type ClassBodyNode struct {
  node
}

// Parse implements Parser
func (n ClassBodyNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ClassElementListNode [Yield] : [See 14.5]
//  ClassElement[?Yield]
//  ClassElementList[?Yield] ClassElement[?Yield]
// implements: Parser and ASTNode
type ClassElementListNode struct {
  node
}

// Parse implements Parser
func (n ClassElementListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}
// ClassElementNode [Yield] : [See 14.5]
//  MethodDefinition[?Yield]
//  static MethodDefinition[?Yield]
//  ;
// implements: Parser and ASTNode
type ClassElementNode struct {
  node
}

// Parse implements Parser
func (n ClassElementNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}
//
// A.5 Scripts and Modules
//

// ScriptNode [See 15.1]
// Script : ScriptBody*
// implements: Parser and ASTNode
type ScriptNode struct {
  node
}

// Parse implements Parser
func (n ScriptNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ScriptBodyNode [See 15.1]
// StatementList
// implements: Parser and ASTNode
type ScriptBodyNode struct {
  node
}

// Parse implements Parser
func (n ScriptBodyNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ModuleNode [See 15.2]
//  ModuleBodyopt
// implements: Parser and ASTNode
type ModuleNode struct {
  node
}

// Parse implements Parser
func (n ModuleNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ModuleBodyNode [See 15.2]
//  ModuleItemList*
// implements: Parser and ASTNode
type ModuleBodyNode struct {
  node
}

// Parse implements Parser
func (n ModuleBodyNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ModuleItemListNode [See 15.2]
//  ModuleItem
//  ModuleItemList ModuleItem
// implements: Parser and ASTNode
type ModuleItemListNode struct {
  node
}

// Parse implements Parser
func (n ModuleItemListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ModuleItemNode [See 15.2]
//  ImportDeclaration
//  ExportDeclaration
//  StatementListItem
// implements: Parser and ASTNode
type ModuleItemNode struct {
  node
}

// Parse implements Parser
func (n ModuleItemNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ImportDeclarationNode [See 15.2.2]
//  import ImportClause FromClause ;
//  import ModuleSpecifier ;
// implements: Parser and ASTNode
type ImportDeclarationNode struct {
  node
}

// Parse implements Parser
func (n ImportDeclarationNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n ImportClauseNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ImportedDefaultBindingNode [See 15.2.2]
//  ImportedBinding
type ImportedDefaultBindingNode struct{
  node
}

// NameSpaceImportNode [See 15.2.2]
//  * as ImportedBinding
type NameSpaceImportNode struct{
  node
}

// NamedImportsNode [See 15.2]
//  { }
//  { ImportsList }
//  { ImportsList , }
type NamedImportsNode struct{
  node
}

// ImportsListNode [See 15.2.2]
//  ImportSpecifier
//  ImportsList , ImportSpecifier
// implements: Parser and ASTNode
type ImportsListNode struct {
  node
}

// Parse implements Parser
func (n ImportsListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ImportSpecifierNode [See 15.2.2]
//  ImportedBinding
//  IdentifierName as ImportedBinding
// implements: Parser and ASTNode
type ImportSpecifierNode struct {
  node
}

// Parse implements Parser
func (n ImportSpecifierNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ModuleSpecifierNode [See 15.2.2]
//  StringLiteral
// implements: Parser and ASTNode
type ModuleSpecifierNode struct {
  node
}

// Parse implements Parser
func (n ModuleSpecifierNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ImportedBindingNode [See 15.2.2]
//  BindingIdentifier
// implements: Parser and ASTNode
type ImportedBindingNode struct {
  node
}

// Parse implements Parser
func (n ImportedBindingNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
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

// Parse implements Parser
func (n ExportDeclarationNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ExportClauseNode [See 15.2.3]
//  { }
//  { ExportsList }
//  { ExportsList , }
// implements: Parser and ASTNode
type ExportClauseNode struct {
  node
}

// Parse implements Parser
func (n ExportClauseNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ExportsListNode [See 15.2.3]
//  ExportSpecifier
//  ExportsList , ExportSpecifier
// implements: Parser and ASTNode
type ExportsListNode struct {
  node
}

// Parse implements Parser
func (n ExportsListNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}

// ExportSpecifierNode [See 15.2.3]
//  IdentifierName
//  IdentifierName as IdentifierName
// implements: Parser and ASTNode
type ExportSpecifierNode struct {
  node
}

// Parse implements Parser
func (n ExportSpecifierNode) Parse(l *Lexer) (ASTNode, error) {
  return nil, nil
}
