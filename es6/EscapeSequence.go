package es6

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
