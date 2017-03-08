package es6

import (
  "strings"
  "unicode"
)
// see 11.4

//
// comment
//

func lexMultiLineComment(l *Lexer) stateFunc {
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
			return l.state
		}
		if r = l.next(); r == eof {
			break
		}
	}
	return l.errorf("no multi line comment terminator \"*/\"")
}

func lexSingleLineComment(l *Lexer) stateFunc {
	l.acceptString("//")
	l.ignore()
	for {
		if strings.HasPrefix(l.input[l.pos:], "\n") || l.next() == eof {
			l.emit(SingleLineComment)
			l.accept("\n")
			l.ignore()
			return l.state
		}
	}
}

//
// identifier
//

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

// // EscapeSequence :: CharacterEscapeSequence || 0 [lookahead ∉ DecimalDigit] || HexEscapeSequence || UnicodeEscapeSequence
// func lexEscapeSequence(l *Lexer) {
// 	// CharacterEscapeSequence
// 	if l.accept("\\") {
// 		// SingleEscapeCharacter :: ' " \ b f n r t v
// 		if l.accept("'\"\\bfnrtv") {
// 			return
// 		}
//
// 		// 0 [lookahead ∉ DecimalDigit]
// 		if l.accept("0") && strings.ContainsRune(decimalDigits, l.peek()) {
// 			return
// 		}
//
// 		// HexEscapeSequence :: x HexDigit HexDigit
// 		if l.accept("x") && l.acceptRunN(decimalDigits+"abcdefABCDEF", 2) {
// 			return
// 		}
//
// 		// UnicodeEscapeSequence :: u Hex4Digits
// 		if l.accept("u") &&
// 			(l.acceptRunN(decimalDigits+"abcdefABCDEF", 4) || l.accept("{") && l.acceptRunN(decimalDigits+"abcdefABCDEF", 4) && l.accept("}")) {
// 			return
// 		}
// 	}
// }
