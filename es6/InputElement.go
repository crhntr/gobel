package es6lexer

import "strings"

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
