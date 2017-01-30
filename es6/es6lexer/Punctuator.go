package es6lexer

var punctuators = []string{
	"{", "(", ")", ";", "]", "[", ",",
	">>>=", "<<=", "!==", "===", ">>>", "...", ".", ">>=", ">=",
	"%=", "*=", "-=", "<=", "&=", "==", "!=", "|=",
	"^=", "+=", "<<", "||", "&&", "++", "--", "=>",
	">>", "-", "&", "|", "^", "!", "~", "%",
	"*", "?", ":", "=", "+", ">", "<"}

func hasPunctuator(l *Lexer) bool {
	defer l.reset()
	return l.acceptAnyString(punctuators)
}

func lexPunctuator(l *Lexer) stateFunc {
	l.acceptAnyString(punctuators)
	l.emit(Punctuator)
	return lexMux
}

func hasDivPunctuator(l *Lexer) bool {
	defer l.reset()
	return l.acceptAnyString([]string{"/=", "/"})
}

func lexDivPunctuator(l *Lexer) stateFunc {
	l.acceptAnyString([]string{"/=", "/"})
	l.emit(DivPunctuator)
	return lexMux
	// }
}

func lexRightBracePunctuator(l *Lexer) stateFunc {
	// if strings.HasPrefix(l.input[l.pos:], "}") {
	l.accept("}")
	l.emit(RightBracePunctuator)
	return lexMux
	// }
	// return l.error(fmt.Errorf("div punctuator not found")) // Paranoic (should never happen)
}
