package es6lexer

// // EscapeSequence :: CharacterEscapeSequence || 0 [lookahead ∉ DecimalDigit] || HexEscapeSequence || UnicodeEscapeSequence
// func lexEscapeSequence(l *lexer) {
// 	// CharacterEscapeSequence
//
// 	// SingleEscapeCharacter :: ' " \ b f n r t v
// 	if l.accept("'\"\\bfnrtv") {
// 		return
// 	}
// 	l.reset()
//
// 	// 0 [lookahead ∉ DecimalDigit]
// 	if l.accept("0") && !l.accept(decimalDigits) {
// 		return
// 	}
//
// 	// HexEscapeSequence :: x HexDigit HexDigit
// 	if l.accept("x") && l.acceptRunN(decimalDigits[:8], 2) {
// 		return
// 	}
//
// 	// UnicodeEscapeSequence :: u Hex4Digits
// 	if l.accept("u") && l.acceptRunN(decimalDigits[:8], 4) {
// 		return
// 	}
// 	l.reset()
//
// 	// UnicodeEscapeSequence :: u{ HexDigits }
// 	if l.accept("u") && l.accept("{") && l.acceptRunN(decimalDigits[:8], 4) && l.accept("}") {
// 		return
// 	}
// 	l.reset()
// }
