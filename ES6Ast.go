package gobel

type ASTNode struct {
	Token
	Children []*ASTNode
}

func (ast *ASTNode) Fmt() string {
	return ""
}
