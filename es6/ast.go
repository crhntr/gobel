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
