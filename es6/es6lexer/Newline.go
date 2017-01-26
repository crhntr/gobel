package es6lexer

func hasLineTerminatorPrefix(l *Lexer) bool {
	defer l.reset()
	return l.accept("\u000A\u000D\u2028\u2029")
}

func lexLineTerminator(l *Lexer) stateFunc {
	l.accept("\u000A\u000D\u2028\u2029")
	l.emit(LineTerminator)
	return lexMux
}
