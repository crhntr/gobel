package es6

import "strings"

//
// lexer
//

// lexStringLiteralDouble consumes a string literal surounded by
// a double quotation marks
func lexStringLiteralDouble(l *lexer) stateFunc {
	l.accept("\"")
	var r rune
	for {
		if strings.HasPrefix(l.input[l.pos:], "\"") {
			break
		}
		if r = l.next(); r == eof {
			l.errorf("did not reach end of string literal reached eof")
			break
		}
	}
	l.accept("\"")
	l.emit(StringLiteral)
	return lexMux
}

// lexStringLiteralSingle consumes a string literal surounded by
// a single quotation marks
func lexStringLiteralSingle(l *lexer) stateFunc {
	l.accept("'")
	var r rune
	for {
		if l.accept("'") {
			break
		}
		if r = l.next(); r == eof {
			l.errorf("did not reach end of string literal reached eof")
			break
		}
	}
	l.accept("'")
	l.emit(StringLiteral)
	return lexMux
}
