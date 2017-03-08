package es6

// Positioner ...
type Positioner interface {
  Position() (filename string, offset int, line int, column int)
}

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
