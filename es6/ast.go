package es6

// Positioner ...
type Positioner interface {
  Position() (filename string, offset int, line int, column int)
}

// ASTNode ...
type ASTNode interface {
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
type IdentifierReferenceNode struct {
  node
}

// BindingIdentifierNode [Yield] : [See 12.1]
//  Identifier
//  [~Yield] yield
//
type BindingIdentifierNode struct {
  node
}

// LabelIdentifierNode [Yield] : [See 12.1]
//  Identifier
//  [~Yield] yield
type LabelIdentifierNode struct {
  node
}

// IdentifierNode  : [See 12.1]
//  IdentifierName but not ReservedWord
type IdentifierNode struct {
  node
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
type PrimaryExpressionNode struct {
  node
}

// CoverParenthesizedExpressionAndArrowParameterListNode [Yield] : [See 12.2]
// ( Expression[In, ?Yield] )
// ( )
// ( ... BindingIdentifier[?Yield] )
// ( Expression[In, ?Yield] , ... BindingIdentifier[?Yield] )
//  When processing the production
type CoverParenthesizedExpressionAndArrowParameterListNode struct {
  node
}

// ParenthesizedExpressionNode [Yield] : [See 12.2]
//  ( Expression[In, ?Yield] )
type ParenthesizedExpressionNode struct {
  node
}

// LiteralNode  : [See 12.2.4]
//  NullLiteral
//  BooleanLiteral
//  NumericLiteral
//  StringLiteral
type LiteralNode struct {
  node
}

// ArrayLiteralNode [Yield] : [See 12.2.5]
//  [ Elisionopt ]
//  [ ElementList[?Yield] ]
//  [ ElementList[?Yield] , Elisionopt ]
type ArrayLiteralNode struct {
  node
}

// ElementListNode [Yield] : [See 12.2.5]
//  Elisionopt AssignmentExpression[In, ?Yield]
//  Elisionopt SpreadElement[?Yield]
//  ElementList[?Yield] , Elisionopt AssignmentExpression[In, ?Yield]
//  ElementList[?Yield] , Elisionopt SpreadElement[?Yield]
type ElementListNode struct {
  node
}

// ElisionNode  : [See 12.2.5]
//  ,
//  Elision ,
type ElisionNode struct {
  node
}

// SpreadElementNode [Yield] : [See 12.2.5]
//  ... AssignmentExpression[In, ?Yield]
type SpreadElementNode struct {
  node
}

// ObjectLiteralNode [Yield] : [See 12.2.6]
//  { }
//  { PropertyDefinitionList[?Yield] }
//  { PropertyDefinitionList[?Yield] , }
type ObjectLiteralNode struct {
  node
}

// PropertyDefinitionListNode [Yield] : [See 12.2.6]
//  PropertyDefinition[?Yield]
//  PropertyDefinitionList[?Yield] , PropertyDefinition[?Yield]
type PropertyDefinitionListNode struct {
  node
}

// PropertyDefinitionNode [Yield] : [See 12.2.6]
//  IdentifierReference[?Yield]
//  CoverInitializedName[?Yield]
//  PropertyName[?Yield] : AssignmentExpression[In, ?Yield]
//  MethodDefinition[?Yield]
type PropertyDefinitionNode struct {
  node
}

// PropertyNameNode [Yield] : [See 12.2.6]
//  LiteralPropertyName
//  ComputedPropertyName[?Yield]
type PropertyNameNode struct {
  node
}

// LiteralPropertyNameNode  : [See 12.2.6]
//  IdentifierName
//  StringLiteral
//  NumericLiteral
type LiteralPropertyNameNode struct {
  node
}

// ComputedPropertyNameNode [Yield] : [See 12.2.6]
//  [ AssignmentExpression[In, ?Yield] ]
type ComputedPropertyNameNode struct {
  node
}

// CoverInitializedNameNode [Yield] : [See 12.2.6]
//  IdentifierReference[?Yield] Initializer[In, ?Yield]
type CoverInitializedNameNode struct {
  node
}

// InitializerNode [In, Yield] : [See 12.2.6]
//  = AssignmentExpression[?In, ?Yield]
type InitializerNode struct {
  node
}

// TemplateLiteralNode [Yield] : [See 12.2.9]
//  NoSubstitutionTemplate
//  TemplateHead Expression[In, ?Yield] TemplateSpans[?Yield]
type TemplateLiteralNode struct {
  node
}

// TemplateSpansNode [Yield] : [See 12.2.9]
//  TemplateTail
//  TemplateMiddleList[?Yield] TemplateTail
type TemplateSpansNode struct {
  node
}

// TemplateMiddleListNode [Yield] : [See 12.2.9]
//  TemplateMiddle Expression[In, ?Yield]
//  TemplateMiddleList[?Yield] TemplateMiddle Expression[In, ?Yield]
type TemplateMiddleListNode struct {
  node
}

// MemberExpressionNode [Yield] : [See 12.3]
//  PrimaryExpression[?Yield]
//  MemberExpression[?Yield] [ Expression[In, ?Yield] ]
//  MemberExpression[?Yield] . IdentifierName
//  MemberExpression[?Yield] TemplateLiteral[?Yield]
//  SuperProperty[?Yield]
//  MetaProperty
//  new MemberExpression[?Yield] Arguments[?Yield]
type MemberExpressionNode struct {
  node
}

// SuperPropertyNode [Yield] : [See 12.3]
//  super [ Expression[In, ?Yield] ]
//  super . IdentifierName
type SuperPropertyNode struct {
  node
}

// MetaPropertyNode  : [See 12.3]
//  NewTarget
type MetaPropertyNode struct {
  node
}

// NewTargetNode  : [See 12.3]
//  new . target
type NewTargetNode struct {
  node
}

// NewExpressionNode [Yield] : [See 12.3]
//  MemberExpression[?Yield]
//  new NewExpression[?Yield]
type NewExpressionNode struct {
  node
}

// CallExpressionNode [Yield] : [See 12.3]
//  MemberExpression[?Yield] Arguments[?Yield]
//  SuperCall[?Yield]
//  CallExpression[?Yield] Arguments[?Yield]
//  CallExpression[?Yield] [ Expression[In, ?Yield] ]
//  CallExpression[?Yield] . IdentifierName
//  CallExpression[?Yield] TemplateLiteral[?Yield]
type CallExpressionNode struct {
  node
}

// SuperCallNode [Yield] : [See 12.3]
//  super Arguments[?Yield]
type SuperCallNode struct {
  node
}

// ArgumentsNode [Yield] : [See 12.3]
//  ( )
//  ( ArgumentList[?Yield] )
type ArgumentsNode struct {
  node
}

// ArgumentListNode [Yield] : [See 12.3]
//  AssignmentExpression[In, ?Yield]
//  ... AssignmentExpression[In, ?Yield]
//  ArgumentList[?Yield] , AssignmentExpression[In, ?Yield]
//  ArgumentList[?Yield] , ... AssignmentExpression[In, ?Yield]
type ArgumentListNode struct {
  node
}

// LeftHandSideExpressionNode [Yield] : [See 12.3]
//  NewExpression[?Yield]
//  CallExpression[?Yield]
type LeftHandSideExpressionNode struct {
  node
}

// PostfixExpressionNode [Yield] : [See 12.4]
//  LeftHandSideExpression[?Yield]
//  LeftHandSideExpression[?Yield] [no LineTerminator here] ++
//  LeftHandSideExpression[?Yield] [no LineTerminator here] --
type PostfixExpressionNode struct {
  node
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
type UnaryExpressionNode struct {
  node
}

// MultiplicativeExpressionNode [Yield] : [See 12.6]
//  UnaryExpression[?Yield]
//  MultiplicativeExpression[?Yield] MultiplicativeOperator UnaryExpression[?Yield]
type MultiplicativeExpressionNode struct {
  node
}

// MultiplicativeOperatorNode  : one of [See 12.6]
//  * / %
type MultiplicativeOperatorNode struct {
  node
}

// AdditiveExpressionNode [Yield] : [See 12.7]
//  MultiplicativeExpression[?Yield]
//  AdditiveExpression[?Yield] + MultiplicativeExpression[?Yield]
//  AdditiveExpression[?Yield] - MultiplicativeExpression[?Yield]
type AdditiveExpressionNode struct {
  node
}

// ShiftExpressionNode [Yield] : [See 12.8]
//  AdditiveExpression[?Yield]
//  ShiftExpression[?Yield] << AdditiveExpression[?Yield]
//  ShiftExpression[?Yield] >> AdditiveExpression[?Yield]
//  ShiftExpression[?Yield] >>> AdditiveExpression[?Yield]
type ShiftExpressionNode struct {
  node
}

// RelationalExpressionNode [In, Yield] : [See 12.9]
//  ShiftExpression[?Yield]
//  RelationalExpression[?In, ?Yield] < ShiftExpression[?Yield]
//  RelationalExpression[?In, ?Yield] > ShiftExpression[?Yield]
//  RelationalExpression[?In, ?Yield] <= ShiftExpression[? Yield]
//  RelationalExpression[?In, ?Yield] >= ShiftExpression[?Yield]
//  RelationalExpression[?In, ?Yield] instanceof ShiftExpression[?Yield]
//  [+In] RelationalExpression[In, ?Yield] in ShiftExpression[?Yield]
type RelationalExpressionNode struct {
  node
}

// EqualityExpressionNode [In, Yield] : [See 12.10]
//  RelationalExpression[?In, ?Yield]
//  EqualityExpression[?In, ?Yield] == RelationalExpression[?In, ?Yield]
//  EqualityExpression[?In, ?Yield] != RelationalExpression[?In, ?Yield]
//  EqualityExpression[?In, ?Yield] === RelationalExpression[?In, ?Yield]
//  EqualityExpression[?In, ?Yield] !== RelationalExpression[?In, ?Yield]
type EqualityExpressionNode struct {
  node
}

// BitwiseANDExpressionNode [In, Yield] : [See 12.11]
//  EqualityExpression[?In, ?Yield]
//  BitwiseANDExpression[?In, ?Yield] & EqualityExpression[?In, ?Yield]
type BitwiseANDExpressionNode struct {
  node
}

// BitwiseXORExpressionNode [In, Yield] : [See 12.11]
//  BitwiseANDExpression[?In, ?Yield]
//  BitwiseXORExpression[?In, ?Yield] ^ BitwiseANDExpression[?In, ?Yield]
type BitwiseXORExpressionNode struct {
  node
}

// BitwiseORExpressionNode [In, Yield] : [See 12.11]
//  BitwiseXORExpression[?In, ?Yield]
//  BitwiseORExpression[?In, ?Yield] | BitwiseXORExpression[?In, ?Yield]
type BitwiseORExpressionNode struct {
  node
}

// LogicalANDExpressionNode [In, Yield] : [See 12.12]
//  BitwiseORExpression[?In, ?Yield]
//  LogicalANDExpression[?In, ?Yield] && BitwiseORExpression[?In, ?Yield]
type LogicalANDExpressionNode struct {
  node
}

// LogicalORExpressionNode [In, Yield] : [See 12.12]
//  LogicalANDExpression[?In, ?Yield]
//  LogicalORExpression[?In, ?Yield] || LogicalANDExpression[?In, ?Yield]
type LogicalORExpressionNode struct {
  node
}

// ConditionalExpressionNode [In, Yield] : [See 12.13]
//  LogicalORExpression[?In, ?Yield]
//  LogicalORExpression[?In,?Yield] ? AssignmentExpression[In, ?Yield] : AssignmentExpression[?In, ?Yield]
type ConditionalExpressionNode struct {
  node
}

// AssignmentExpressionNode [In, Yield] : [See 12.14]
//  ConditionalExpression[?In, ?Yield]
//  [+Yield] YieldExpression[?In]
//  ArrowFunction[?In, ?Yield]
//  LeftHandSideExpression[?Yield] = AssignmentExpression[?In, ?Yield]
//  LeftHandSideExpression[?Yield] AssignmentOperator AssignmentExpression[?In, ?Yield]
type AssignmentExpressionNode struct {
  node
}

// AssignmentOperatorNode  : one of [See 12.14]
//  *= /= %= += -= <<= >>= >>>= &= ^= |=
type AssignmentOperatorNode struct {
  node
}

// ExpressionNode [In, Yield] : [See 12.15]
//  AssignmentExpression[?In, ?Yield]
//  Expression[?In, ?Yield] , AssignmentExpression[?In, ?Yield]
type ExpressionNode struct {
  node
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
type StatementNode struct {
  node
}

// DeclarationNode [Yield] : [See clause 13]
//  HoistableDeclaration[?Yield]
//  ClassDeclaration[?Yield]
//  LexicalDeclaration[In, ?Yield]
type DeclarationNode struct {
  node
}

// HoistableDeclarationNode [Yield, Default] : [See clause 13]
//  FunctionDeclaration[?Yield,?Default]
//  GeneratorDeclaration[?Yield, ?Default]
type HoistableDeclarationNode struct {
  node
}

// BreakableStatementNode [Yield, Return] : [See clause 13]
//  IterationStatement[?Yield, ?Return]
//  SwitchStatement[?Yield, ?Return]
type BreakableStatementNode struct {
  node
}

// BlockStatementNode [Yield, Return] : [See 13.2]
//  Block[?Yield, ?Return]
type BlockStatementNode struct {
  node
}

// BlockNode [Yield, Return] : [See 13.2]
//  { StatementList[?Yield, ?Return]opt }
type BlockNode struct {
  node
}

// StatementListNode [Yield, Return] : [See 13.2]
//  StatementListItem[?Yield, ?Return]
//  StatementList[?Yield, ?Return] StatementListItem[?Yield, ?Return]
type StatementListNode struct {
  node
}

// StatementListItemNode [Yield, Return] : [See 13.2]
//  Statement[?Yield, ?Return]
//  Declaration[?Yield]
type StatementListItemNode struct {
  node
}

// LexicalDeclarationNode [In, Yield] : [See 13.3.1]
//  LetOrConst BindingList[?In, ?Yield] ;
type LexicalDeclarationNode struct {
  node
}

// LetOrConstNode  : [See 13.3.1]
//  let
//  const
type LetOrConstNode struct {
  node
}

// BindingListNode [In, Yield] : [See 13.3.1]
//  LexicalBinding[?In, ?Yield]
//  BindingList[?In, ?Yield] , LexicalBinding[?In, ?Yield]
type BindingListNode struct {
  node
}

// LexicalBindingNode [In, Yield] : [See 13.3.1]
//  BindingIdentifier[?Yield] Initializer[?In, ?Yield]opt
//  BindingPattern[?Yield] Initializer[?In, ?Yield]
type LexicalBindingNode struct {
  node
}

// VariableStatementNode [Yield] : [See 13.3.2]
//  var VariableDeclarationList[In, ?Yield] ;
type VariableStatementNode struct {
  node
}

// VariableDeclarationListNode [In, Yield] : [See 13.3.2]
//  VariableDeclaration[?In, ?Yield]
//  VariableDeclarationList[?In, ?Yield] , VariableDeclaration[?In, ?Yield]
type VariableDeclarationListNode struct {
  node
}

// VariableDeclarationNode [In, Yield] : [See 13.3.2]
//  BindingIdentifier[?Yield] Initializer[?In, ?Yield]opt
//  BindingPattern[?Yield] Initializer[?In, ?Yield]
type VariableDeclarationNode struct {
  node
}

// BindingPatternNode [Yield] : [See 13.3.3]
//  ObjectBindingPattern[?Yield]
//  ArrayBindingPattern[?Yield]
type BindingPatternNode struct {
  node
}

// ObjectBindingPatternNode [Yield] : [See 13.3.3]
//  { }
//  { BindingPropertyList[?Yield] }
//  { BindingPropertyList[?Yield] , }
type ObjectBindingPatternNode struct {
  node
}

// ArrayBindingPatternNode [Yield] : [See 13.3.3]
//  [ Elisionopt BindingRestElement[?Yield]opt ]
//  [ BindingElementList[?Yield] ]
//  [ BindingElementList[?Yield] , Elisionopt BindingRestElement[?Yield]opt ]
type ArrayBindingPatternNode struct {
  node
}

// BindingPropertyListNode [Yield] : [See 13.3.3]
//  BindingProperty[?Yield]
//  BindingPropertyList[?Yield] , BindingProperty[?Yield]
type BindingPropertyListNode struct {
  node
}

// BindingElementListNode [Yield] : [See 13.3.3]
//  BindingElisionElement[?Yield]
//  BindingElementList[?Yield] , BindingElisionElement[?Yield]
type BindingElementListNode struct {
  node
}

// BindingElisionElementNode [Yield] : [See 13.3.3]
//  Elisionopt BindingElement[?Yield]
type BindingElisionElementNode struct {
  node
}

// BindingPropertyNode [Yield] : [See 13.3.3]
//  SingleNameBinding[?Yield]
//  PropertyName[?Yield] : BindingElement[?Yield]
type BindingPropertyNode struct {
  node
}

// BindingElementNode [Yield] : [See 13.3.3]
//  SingleNameBinding[?Yield]
//  BindingPattern[?Yield] Initializer[In, ?Yield]opt
type BindingElementNode struct {
  node
}

// SingleNameBindingNode [Yield] : [See 13.3.3]
//  BindingIdentifier[?Yield] Initializer[In, ?Yield]opt
type SingleNameBindingNode struct {
  node
}

// BindingRestElementNode [Yield] : [See 13.3.3]
//  ... BindingIdentifier[?Yield]
type BindingRestElementNode struct {
  node
}

// EmptyStatementNode  : [See 13.4]
//  ;
type EmptyStatementNode struct {
  node
}

// ExpressionStatementNode [Yield] : [See 13.5]
//  [lookahead ∉ {{, function, class, let [}] Expression[In, ?Yield] ;
type ExpressionStatementNode struct {
  node
}

// IfStatementNode [Yield, Return] : [See 13.6]
//  if ( Expression[In, ?Yield] ) Statement[?Yield, ?Return] else Statement[?Yield, ?Return]
//  if ( Expression[In, ?Yield] ) Statement[?Yield, ?Return]
type IfStatementNode struct {
  node
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
type IterationStatementNode struct {
  node
}

// ForDeclarationNode [Yield] : [See 13.7]
//  LetOrConst ForBinding[?Yield]
type ForDeclarationNode struct {
  node
}

// ForBindingNode [Yield] : [See 13.7]
//  BindingIdentifier[?Yield]
//  BindingPattern[?Yield]
type ForBindingNode struct {
  node
}

// ContinueStatementNode [Yield] : [See 13.8]
//  continue ;
//  continue [no LineTerminator here] LabelIdentifier[?Yield] ;
type ContinueStatementNode struct {
  node
}

// BreakStatementNode [Yield] : [See 13.9]
//  break ;
//  break [no LineTerminator here] LabelIdentifier[?Yield] ;
type BreakStatementNode struct {
  node
}

// ReturnStatementNode [Yield] : [See 13.10]
//  return ;
//  return [no LineTerminator here] Expression[In, ?Yield] ;
type ReturnStatementNode struct {
  node
}

// WithStatementNode [Yield, Return] : [See 13.11]
//  with ( Expression[In, ?Yield] ) Statement[?Yield, ?Return]
type WithStatementNode struct {
  node
}

// SwitchStatementNode [Yield, Return] : [See 13.12]
//  switch ( Expression[In, ?Yield] ) CaseBlock[?Yield, ?Return]
type SwitchStatementNode struct {
  node
}

// CaseBlockNode [Yield, Return] : [See 13.12]
//  { CaseClauses[?Yield, ?Return]opt }
//  { CaseClauses[?Yield, ?Return]opt DefaultClause[?Yield, ?Return] CaseClauses[?Yield, ?Return]opt }
type CaseBlockNode struct {
  node
}

// CaseClausesNode [Yield, Return] : [See 13.12]
//  CaseClause[?Yield, ?Return]
//  CaseClauses[?Yield, ?Return] CaseClause[?Yield, ?Return]
type CaseClausesNode struct {
  node
}

// CaseClauseNode [Yield, Return] : [See 13.12]
//  case Expression[In, ?Yield] : StatementList[?Yield, ?Return]opt
type CaseClauseNode struct {
  node
}
// DefaultClauseNode [Yield, Return] : [See 13.12]
//  default : StatementList[?Yield, ?Return]opt
type DefaultClauseNode struct {
  node
}

// LabelledStatementNode [Yield, Return] : [See 13.13]
//  LabelIdentifier[?Yield] : LabelledItem[?Yield, ?Return]
type LabelledStatementNode struct {
  node
}

// LabelledItemNode [Yield, Return] : [See 13.13]
//  Statement[?Yield, ?Return]
//  FunctionDeclaration[?Yield]
type LabelledItemNode struct {
  node
}

// ThrowStatementNode [Yield] : [See 13.14]
//  throw [no LineTerminator here] Expression[In, ?Yield] ;
type ThrowStatementNode struct {
  node
}

// TryStatementNode [Yield, Return] : [See 13.15]
//  try Block[?Yield, ?Return] Catch[?Yield, ?Return]
//  try Block[?Yield, ?Return] Finally[?Yield, ?Return]
//  try Block[?Yield, ?Return] Catch[?Yield, ?Return] Finally[?Yield, ?Return]
type TryStatementNode struct {
  node
}

// CatchNode [Yield, Return] : [See 13.15]
//  catch ( CatchParameter[?Yield] ) Block[?Yield, ?Return]
type CatchNode struct {
  node
}

// FinallyNode [Yield, Return] : [See 13.15]
//  finally Block[?Yield, ?Return]
type FinallyNode struct {
  node
}

// CatchParameterNode [Yield] : [See 13.15]
//  BindingIdentifier[?Yield]
//  BindingPattern[?Yield]
type CatchParameterNode struct {
  node
}

// DebuggerStatementNode  : [See 13.16]
//  debugger ;
type DebuggerStatementNode struct {
  node
}

//
// A.4 Functions and Classes
//

// FunctionDeclarationNode [Yield, Default] :
//  function BindingIdentifier[?Yield] ( FormalParameters ) { FunctionBody }
//  [+Default] function ( FormalParameters ) { FunctionBody }
type FunctionDeclarationNode struct {
  node
}

// FunctionExpressionNode  : [See 14.1]
//  function BindingIdentifieropt ( FormalParameters ) { FunctionBody }
type FunctionExpressionNode struct {
  node
}
// StrictFormalParametersNode [Yield] : [See 14.1]
//  FormalParameters[?Yield]
type StrictFormalParametersNode struct {
  node
}
// FormalParametersNode [Yield] : [See 14.1]
//  [empty]
//  FormalParameterList[?Yield]
type FormalParametersNode struct {
  node
}

// FormalParameterListNode [Yield] : [See 14.1]
//  FunctionRestParameter[?Yield]
//  FormalsList[?Yield]
//  FormalsList[?Yield] , FunctionRestParameter[?Yield]
type FormalParameterListNode struct {
  node
}
// FormalsListNode [Yield] : [See 14.1]
//  FormalParameter[?Yield]
//  FormalsList[?Yield] , FormalParameter[?Yield]
type FormalsListNode struct {
  node
}

// FunctionRestParameterNode [Yield] : [See 14.1]
//  BindingRestElement[?Yield]
type FunctionRestParameterNode struct {
  node
}

// FormalParameterNode [Yield] : [See 14.1]
//  BindingElement[?Yield]
type FormalParameterNode struct {
  node
}
// FunctionBodyNode [Yield] : [See 14.1]
//  FunctionStatementList[?Yield]
type FunctionBodyNode struct {
  node
}
// FunctionStatementListNode [Yield] : [See 14.1]
//  StatementList[?Yield, Return]opt
type FunctionStatementListNode struct {
  node
}
// ArrowFunctionNode [In, Yield] : [See 14.2]
//  ArrowParameters[?Yield] [no LineTerminator here] => ConciseBody[?In]
type ArrowFunctionNode struct {
  node
}

// ArrowParametersNode [Yield] : [See 14.2]
//  BindingIdentifier[?Yield]
//  CoverParenthesizedExpressionAndArrowParameterList[?Yield]
// ArrowParameters[Yield] : CoverParenthesizedExpressionAndArrowParameterList[?Yield]
// is recognized the following grammar is used to refine the interpretation of CoverParenthesizedExpressionAndArrowParameterList :
type ArrowParametersNode struct {
  node
}

// ConciseBodyNode [In] : [See 14.2]
//  [lookahead ≠ { ] AssignmentExpression[?In]
//  { FunctionBody }
//  ArrowFormalParameters[Yield] :
//  ( StrictFormalParameters[?Yield] )
type ConciseBodyNode struct {
  node
}

// MethodDefinitionNode [Yield] : [See 14.3]
//  PropertyName[?Yield] ( StrictFormalParameters ) { FunctionBody }
//  GeneratorMethod[?Yield]
//  get PropertyName[?Yield] ( ) { FunctionBody }
//  set PropertyName[?Yield] ( PropertySetParameterList ) { FunctionBody }
type MethodDefinitionNode struct {
  node
}

// PropertySetParameterListNode  : [See 14.3]
// FormalParameter
type PropertySetParameterListNode struct {
  node
}

// GeneratorMethodNode [Yield] : [See 14.4]
//  * PropertyName[?Yield] ( StrictFormalParameters[Yield] ) { GeneratorBody }
type GeneratorMethodNode struct {
  node
}

// GeneratorDeclarationNode [Yield, Default] : [See 14.4]
//  function * BindingIdentifier[?Yield] ( FormalParameters[Yield] ) { GeneratorBody }
//  [+Default] function * ( FormalParameters[Yield] ) { GeneratorBody }
type GeneratorDeclarationNode struct {
  node
}

// GeneratorExpressionNode  : [See 14.4]
//  function * BindingIdentifier[Yield]opt ( FormalParameters[Yield] ) { GeneratorBody }
type GeneratorExpressionNode struct {
  node
}
// GeneratorBodyNode  : [See 14.4]
//  FunctionBody[Yield]
type GeneratorBodyNode struct {
  node
}

// YieldExpressionNode [In] : [See 14.4]
//  yield
//  yield [no LineTerminator here] AssignmentExpression[?In, Yield]
//  yield [no LineTerminator here] * AssignmentExpression[?In, Yield]
type YieldExpressionNode struct {
  node
}

// ClassDeclarationNode [Yield, Default] : [See 14.5]
//  class BindingIdentifier[?Yield] ClassTail[?Yield]
//  [+Default] class ClassTail[?Yield]
type ClassDeclarationNode struct {
  node
}

// ClassExpressionNode [Yield] : [See 14.5]
//  class BindingIdentifier[?Yield]opt ClassTail[?Yield]
type ClassExpressionNode struct {
  node
}

// ClassTailNode [Yield] : [See 14.5]
//  ClassHeritage[?Yield]opt { ClassBody[?Yield]opt }
type ClassTailNode struct {
  node
}

// ClassHeritageNode [Yield] : [See 14.5]
//  extends LeftHandSideExpression[?Yield]
type ClassHeritageNode struct {
  node
}

// ClassBodyNode [Yield] : [See 14.5]
//  ClassElementList[?Yield]
type ClassBodyNode struct {
  node
}

// ClassElementListNode [Yield] : [See 14.5]
//  ClassElement[?Yield]
//  ClassElementList[?Yield] ClassElement[?Yield]
type ClassElementListNode struct {
  node
}
// ClassElementNode [Yield] : [See 14.5]
//  MethodDefinition[?Yield]
//  static MethodDefinition[?Yield]
//  ;
type ClassElementNode struct {
  node
}
//
// A.5 Scripts and Modules
//

// ScriptNode [See 15.1]
// Script : ScriptBody*
type ScriptNode struct {
  node
}

// ScriptBodyNode [See 15.1]
// StatementList
type ScriptBodyNode struct {
  node
}

// ModuleNode [See 15.2]
//  ModuleBodyopt
type ModuleNode struct {
  node
}

// ModuleBodyNode [See 15.2]
//  ModuleItemList*
type ModuleBodyNode struct {
  node
}

// ModuleItemListNode [See 15.2]
//  ModuleItem
//  ModuleItemList ModuleItem
type ModuleItemListNode struct {
  node
}

// ModuleItemNode [See 15.2]
//  ImportDeclaration
//  ExportDeclaration
//  StatementListItem
type ModuleItemNode struct {
  node
}

// ImportDeclarationNode [See 15.2.2]
//  import ImportClause FromClause ;
//  import ModuleSpecifier ;
type ImportDeclarationNode struct {
  node
}

// ImportClauseNode [See 15.2.2]
//  ImportedDefaultBinding
//  NameSpaceImport
//  NamedImports
//  ImportedDefaultBinding , NameSpaceImport
//  ImportedDefaultBinding , NamedImports
type ImportClauseNode struct {
  node
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
type ImportsListNode struct {
  node
}

// ImportSpecifierNode [See 15.2.2]
//  ImportedBinding
//  IdentifierName as ImportedBinding
type ImportSpecifierNode struct {
  node
}

// ModuleSpecifierNode [See 15.2.2]
//  StringLiteral
type ModuleSpecifierNode struct {
  node
}

// ImportedBindingNode [See 15.2.2]
//  BindingIdentifier
type ImportedBindingNode struct {
  node
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
type ExportDeclarationNode struct {
  node
}

// ExportClauseNode [See 15.2.3]
//  { }
//  { ExportsList }
//  { ExportsList , }
type ExportClauseNode struct {
  node
}

// ExportsListNode [See 15.2.3]
//  ExportSpecifier
//  ExportsList , ExportSpecifier
type ExportsListNode struct {
  node
}

// ExportSpecifierNode [See 15.2.3]
//  IdentifierName
//  IdentifierName as IdentifierName
type ExportSpecifierNode struct {
  node
}
