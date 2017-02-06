package es6lexer

var lineTerminators = "\u000A\u000D\u2028\u2029"

func hasLineTerminatorPrefix(l *Lexer) bool {
	defer l.reset()
	return l.accept(lineTerminators)
}

func lexLineTerminator(l *Lexer) stateFunc {
	l.accept(lineTerminators)
	l.emit(LineTerminator)
	return l.state
}
