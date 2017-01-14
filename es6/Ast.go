package es6

// ASTNode represents a node in the AST
type ASTNode interface {
	SourceLocator
}

// SourceLocator allows for accessing the source and the start and end
// position of a ASTNode
type SourceLocator interface {
	Source() []byte
	Start() Position
	End() Position
}

// Position objects consists of:
// a line number (1-indexed) and
// a column number (0-indexed)
type Position struct {
	Row     int // >= 1
	Collumn int // >= 0
}
