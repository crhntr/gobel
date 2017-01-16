package es6

import "strings"

func lexMultiLineComment(l *lexer) stateFunc {
	l.acceptString("/*")
	l.ignoreN(len("/*"))
	var r rune
	for {
		if l.acceptString("*/") {
			l.pos -= len("*/")
			if l.pos >= l.start {
				l.emit(MultiLineComment)
			}
			l.ignoreN(len("*/"))
			return lexMux
		}
		if r = l.next(); r == eof {
			break
		}
	}
	return l.errorf("no multi line comment terminator \"*/\"")
}

func lexSingleLineComment(l *lexer) stateFunc {
	l.acceptString("//")
	l.ignore()
	for {
		if strings.HasPrefix(l.input[l.pos:], "\n") || l.next() == eof {
			l.emit(SingleLineComment)
			l.accept("\n")
			l.ignore()
			return lexMux
		}
	}
}
