package es6

import (
  "strings"
  "unicode"
)

// lexMux multiplexes the various states based on
// input prefix checks. it follows the following rules from
// the specification when determinint what states are allowed
//
// InputElementDiv :: 									WhiteSpace | LineTerminator | Comment | CommonToken | DivPunctuator | RightBracePunctuator
// InputElementRegExp :: 								WhiteSpace | LineTerminator | Comment | CommonToken | 								RightBracePunctuator | 	RegularExpressionLiteral
// InputElementRegExpOrTemplateTail :: 	WhiteSpace | LineTerminator | Comment | CommonToken | 																				RegularExpressionLiteral | 	TemplateSubstitutionTail
// InputElementTemplateTail ::  				WhiteSpace | LineTerminator | Comment | CommonToken | DivPunctuator | 																										TemplateSubstitutionTail
//
// CommonToken :: IdentifierName | Punctuator | NumericLiteral | StringLiteral | Template
func lexInputElement(l *Lexer) stateFunc {
	switch {
	case hasWhiteSpacePrefix(l):
		return lexWhiteSpace
	case hasLineTerminatorPrefix(l):
		return lexLineTerminator
	case strings.HasPrefix(l.input[l.pos:], "//"): // SingleLineComment
		return lexSingleLineComment
	case strings.HasPrefix(l.input[l.pos:], "/*"): // MultiLineComment
		return lexMultiLineComment
	case hasReservedWord(l, l.input[l.pos:]):
		return lexReservedWord
	case hasNumericLiteral(l):
		return lexNumericLiteral
	case hasPunctuator(l):
		return lexPunctuator
	case hasIdentifierNameStartPrefix(l): // IdentifierName
		return lexIdentifierName
	case strings.HasPrefix(l.input[l.pos:], "\""): // StringLiteral
		return lexStringLiteralDouble
	case strings.HasPrefix(l.input[l.pos:], "'"): // StringLiteral
		return lexStringLiteralSingle
	case strings.HasPrefix(l.input[l.pos:], "`"): // TemplateLiteral
		return lexTemplateLiteral
	default:
		switch l.goal {
		case InputElementRegExp:
			if strings.HasPrefix(l.input[l.pos:], "}") { // RightBracePunctuator
				return lexRightBracePunctuator
			}
			if strings.HasPrefix(l.input[l.pos:], "/") {
				return lexRegex
			}
		case InputElementRegExpOrTemplateTail:
			if strings.HasPrefix(l.input[l.pos:], "}") { // TemplateSubstitutionTail
				return lexTemplateSubstitutionTail
			}
			if strings.HasPrefix(l.input[l.pos:], "/") {
				return lexRegex
			}
		case InputElementTemplateTail:
			if strings.HasPrefix(l.input[l.pos:], "}") { // TemplateSubstitutionTail
				return lexTemplateSubstitutionTail
			}
		case InputElementDiv:
			if hasDivPunctuator(l) {
				return lexDivPunctuator
			}
			if strings.HasPrefix(l.input[l.pos:], "}") { // RightBracePunctuator
				return lexRightBracePunctuator
			}
		}
	}
	if l.pos != len(l.input) {
		l.errorf("unexpected end of input")
		return nil
	}
	l.emit(EOF)
	return nil
}

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

//
// newline
//

var lineTerminators = "\u000A\u000D\u2028\u2029"

func hasLineTerminatorPrefix(l *Lexer) bool {
	defer l.reset()
	return l.accept(lineTerminators)
}

func lexLineTerminator(l *Lexer) stateFunc {
	l.accept(lineTerminators)
	l.emit(LineTerminator)
	return l.state
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
