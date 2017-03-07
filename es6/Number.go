package es6

import "unicode"

// see See 11.8.3

const decimalDigits = "0123456789"

func hasNumericLiteral(l *Lexer) bool {
	defer l.reset()
	l.accept("-")
	return l.accept("123456789") || (l.accept(".") && l.accept("0123456789")) || (l.accept("0") && (l.accept("oOxXbB") || !hasIdentifierNameStartPrefix(l)))
}

// lexNumericLiteral inspired by Rob Pike's talk
func lexNumericLiteral(l *Lexer) stateFunc {

	// Next thing mustn't be alphanumeric.
	mustNotHaveNextAlpha := func(l *Lexer) stateFunc {
		if unicode.IsLetter(l.peek()) {
			l.next()
			return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
		}
		l.emit(NumericLiteral)
		return l.state
	}

	// Optional leading sign.
	l.accept("-")
	// Is it hex?
	if l.accept("0") {
		if l.accept("xX") {
			l.acceptRun("0123456789abcdefABCDEF")
			return mustNotHaveNextAlpha(l)
		} else if l.accept("oO") { // Is it octal?
			l.acceptRun(decimalDigits[:8])
			return mustNotHaveNextAlpha(l)
		} else if l.accept("bB") { // Is it bin?
			l.acceptRun("01")
			return mustNotHaveNextAlpha(l)
		}
	}

	if l.accept(decimalDigits[1:]) {
		l.acceptRun(decimalDigits)
	}

	if l.accept(".") {
		l.acceptRun(decimalDigits)
	}

	if l.accepted() && l.accept("eE") {
		l.accept("+-")
		l.acceptRun(decimalDigits)
	}

	return mustNotHaveNextAlpha(l)
}
