package es6lexer

import "strings"

var lineTerminators = "\u000A\u000D\u2028\u2029"

func isLineTerminator(r rune) bool {
	return strings.ContainsRune(lineTerminators, r)
}
func hasLineTerminatorPrefix(l *Lexer) bool {
	defer l.reset()
	return l.accept(lineTerminators)
}

func lexLineTerminator(l *Lexer) stateFunc {
	l.accept(lineTerminators)
	l.emit(LineTerminator)
	return l.state
}
