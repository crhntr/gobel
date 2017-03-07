package es6lexer

import "unicode"

func isIdentifierStart(r rune) bool {
	return r == '$' || r == '_' ||
		(unicode.IsOneOf([]*unicode.RangeTable{unicode.L, unicode.Nd, unicode.Other_ID_Start}, r) &&
			!(unicode.IsOneOf([]*unicode.RangeTable{unicode.Pattern_Syntax, unicode.Pattern_White_Space}, r)))
}

func isIdentifierPart(r rune) bool {
	return r == '$' || r == '_' || isIdentifierStart(r) || unicode.IsOneOf([]*unicode.RangeTable{unicode.L, unicode.Mn, unicode.Mc, unicode.Nd, unicode.Pc, unicode.Other_ID_Start}, r)
}
func hasIdentifierNameStartPrefix(l *Lexer) bool {
	return isIdentifierStart(l.peek())
}

func hasIdentifierNameContinuePrefix(l *Lexer) bool {
	return isIdentifierPart(l.peek())
}

func lexIdentifierName(l *Lexer) stateFunc {
	l.next()
	for {
		if !hasIdentifierNameContinuePrefix(l) {
			l.emit(IdentifierName)
			return l.state
		}
		l.next()
	}
}
