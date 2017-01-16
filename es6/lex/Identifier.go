package es6

import "unicode"

func hasIdentifierNameStartPrefix(l *lexer) bool {
	r := l.peek()
	return r == '$' || r == '_' ||
		(unicode.IsOneOf([]*unicode.RangeTable{unicode.L, unicode.Nd, unicode.Other_ID_Start}, r) &&
			!(unicode.IsOneOf([]*unicode.RangeTable{unicode.Pattern_Syntax, unicode.Pattern_White_Space}, r)))
}

func hasIdentifierNameContinuePrefix(l *lexer) bool {
	r := l.peek()
	return hasIdentifierNameStartPrefix(l) || unicode.IsOneOf([]*unicode.RangeTable{unicode.L, unicode.Mn, unicode.Mc, unicode.Nd, unicode.Pc, unicode.Other_ID_Start}, r)
}

func lexIdentifierName(l *lexer) stateFunc {
	l.next()
	for {
		if !hasIdentifierNameContinuePrefix(l) {
			l.emit(IdentifierName)
			return lexMux
		}
		l.next()
	}
}
