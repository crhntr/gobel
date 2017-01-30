package es6lexer

func lexRegex(l *Lexer) stateFunc {
	l.accept("/")

	i := 0
	var r rune
	for r = l.next(); ; r = l.next() {
		if l.accept("/") {
			break
		}
		if isLineTerminator(r) {
			l.errorf("regex can't have new lines")
			break
		}
		if !isIdentifierPart(r) {
			l.errorf("invalid character [%c] in regex", r)
			break
		}
		i++
		if r == eof {
			l.errorf("regex did not close with /")
			break
		}

		if i > 100 {
			panic("should not reach 100")
		}
	}

	for hasIdentifierNameContinuePrefix(l) {
		if r = l.next(); r == eof {
			l.errorf("did not reach end of string literal reached eof")
			break
		}
	}
	l.emit(RegEx)

	return nil
}
