package es6

//
// lexer
//

func hasWhiteSpacePrefix(l *lexer) bool {
	defer l.reset()
	return l.accept("\u0009\u000B\u000C\u0020\u00A0\uFEFF\uFEFF")
}

func lexWhiteSpace(l *lexer) stateFunc {
	l.acceptRun("\u0009\u000B\u000C\u0020\u00A0\uFEFF\uFEFF")
	l.emit(WhiteSpace)
	return lexMux
}
