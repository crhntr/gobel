package es6lexer

func lexRegex(l *Lexer) stateFunc {
	l.accept("/")
	var r rune
	for r = l.next(); ; r = l.next() {
		if r == eof {
			l.errorf("regex did not close with '/' ")
			return nil
		}
		if isLineTerminator(r) {
			l.errorf("regex can't have new lines")
			return nil
		}
		if l.accept("/") {
			break
		}
		if !isIdentifierPart(r) {
			l.errorf("invalid character '%c' in regex", r)
			return nil
		}
	}

	for {
		r = l.next()
		if r == eof || !isIdentifierPart(r) {
			break
		}
	}
	l.emit(RegEx)

	return nil
}
