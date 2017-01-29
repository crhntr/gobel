package es6lexer

func lexTemplateLiteral(l *Lexer) stateFunc {
	l.accept("`")
	return nil
}
