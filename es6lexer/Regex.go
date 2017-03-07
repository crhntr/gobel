package es6lexer

func lexRegex(l *Lexer) stateFunc {
	l.accept("/")

	l.acceptSourceCharacterRunExcept(lineTerminators + "/")
	if !l.accept("/") {
		l.errorf("regex did not close with '/' ")
		return nil
	}

	for {
		r := l.next()
		if !isIdentifierPart(r) {
			l.backup()
			break
		}
	}

	l.emit(RegEx)

	return nil
}
