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

//
// NumbericLiteral
//

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

//
// Punctuator
//

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
	return l.state
}

func hasDivPunctuator(l *Lexer) bool {
	defer l.reset()
	return l.acceptAnyString([]string{"/=", "/"})
}

func lexDivPunctuator(l *Lexer) stateFunc {
	l.acceptAnyString([]string{"/=", "/"})
	l.emit(DivPunctuator)
	return l.state
	// }
}

func lexRightBracePunctuator(l *Lexer) stateFunc {
	l.accept("}")
	l.emit(RightBracePunctuator)
	return l.state
}

//
// regex
//

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

//
// ReservedWord
//

var currentReservedWords = []string{
	"break", "do", "instanceof", "in",
	"typeof", "case", "else", "var",
	"catch", "export", "new", "void",
	"class", "extends", "return", "while",
	"const", "finally", "super", "with",
	"continue", "for", "switch", "yield",
	"debugger", "function", "this", "default",
	"if", "throw", "delete", "import", "try"}
var futureReservedWords = []string{"enum", "await"}
var futureResdervedWordsStrict = []string{"implements", "package", "protected", "interface", "private", "public"}
var literals = []string{"null", "true", "false"}

func hasReservedWord(l *Lexer, str string) bool {
	for _, word := range l.reservedWords {
		if strings.HasPrefix(str, word) {
			return true
		}
	}
	return false
}

func lexReservedWord(l *Lexer) stateFunc {
	l.acceptAnyString(l.reservedWords)
	l.emit(ReservedWord)
	return l.state
}

//
// helper code
//

// copied from: https://gobyexample.com/sorting-by-functions
type keywordSorter []string

func (s keywordSorter) Len() int {
	return len(s)
}
func (s keywordSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s keywordSorter) Less(i, j int) bool {
	return len(s[i]) > len(s[j]) // sorts reverse order
}

//
// template literals
//

func lexTemplateLiteral(l *Lexer) stateFunc {
	l.accept("`")
	var r rune
	for {
		if strings.HasPrefix(l.input[l.pos:], "${") {
			l.acceptString("${")
			l.emit(TemplateHead)
			return l.state
		}
		if strings.HasPrefix(l.input[l.pos:], "`") {
			l.accept("`")
			l.emit(NoSubstitutionTemplate)
			return l.state
		}
		if r = l.next(); r == eof {
			l.errorf("did not reach end of template literal reached eof")
			return nil
		}
	}
}

func lexTemplateSubstitutionTail(l *Lexer) stateFunc {
	l.accept("}")
	var r rune
	for {
		if strings.HasPrefix(l.input[l.pos:], "${") {
			l.acceptString("${")
			l.emit(TemplateMiddle)
			return l.state
		}
		if strings.HasPrefix(l.input[l.pos:], "`") {
			l.accept("`")
			l.emit(TemplateTail)
			return l.state
		}
		if r = l.next(); r == eof {
			l.errorf("did not reach TemplateMiddle or TemplateTail but reached eof")
			return nil
		}
	}
}

//
// WhiteSpace
//

// see 11.2

func hasWhiteSpacePrefix(l *Lexer) bool {
	defer l.reset()
	return l.accept("\u0009\u000B\u000C\u0020\u00A0\uFEFF\uFEFF")
}

func lexWhiteSpace(l *Lexer) stateFunc {
	l.acceptRun("\u0009\u000B\u000C\u0020\u00A0\uFEFF\uFEFF")
	l.emit(WhiteSpace)
	return l.state
}


//
// StringLiteral
//

// See 11.8.4

// lexStringLiteralDouble consumes a string literal surounded by
// a double quotation marks
func lexStringLiteralDouble(l *Lexer) stateFunc {
	l.accept("\"")
	var r rune
	for {
		if strings.HasPrefix(l.input[l.pos:], "\"") {
			break
		}
		if r = l.next(); r == eof {
			l.errorf("did not reach end of string literal reached eof")
			break
		}
	}
	l.accept("\"")
	l.emit(StringLiteral)
	return l.state
}

// lexStringLiteralSingle consumes a string literal surounded by
// a single quotation marks
func lexStringLiteralSingle(l *Lexer) stateFunc {
	l.accept("'")
	var r rune
	for {
		if l.accept("'") {
			break
		}
		if r = l.next(); r == eof {
			l.errorf("did not reach end of string literal reached eof")
			break
		}
	}
	l.accept("'")
	l.emit(StringLiteral)
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
