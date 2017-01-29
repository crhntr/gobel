package es6lexer

// see 11.2

func hasWhiteSpacePrefix(l *Lexer) bool {
	defer l.reset()
	return l.accept("\u0009\u000B\u000C\u0020\u00A0\uFEFF\uFEFF")
}

func lexWhiteSpace(l *Lexer) stateFunc {
	l.acceptRun("\u0009\u000B\u000C\u0020\u00A0\uFEFF\uFEFF")
	l.emit(WhiteSpace)
	return lexMux
}
