package es6

//
// lexer
//

func hasLineTerminatorPrefix(l *lexer) bool {
	defer l.reset()
	return l.accept("\u000A\u000D\u2028\u2029")
}

func lexLineTerminator(l *lexer) stateFunc {
	l.next()
	l.emit(LineTerminator)
	return lexMux
}

//
// parser
//
