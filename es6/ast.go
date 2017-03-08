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
