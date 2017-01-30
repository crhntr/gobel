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
func lexMux(l *Lexer) stateFunc {
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
		return lexReservedWord(l)
	case hasNumericLiteral(l):
		return lexNumericLiteral(l)
	case hasPunctuator(l):
		return lexPunctuator(l)
	case hasIdentifierNameStartPrefix(l): // IdentifierName
		return lexIdentifierName(l)
	case strings.HasPrefix(l.input[l.pos:], "`"): // TemplateLiteral
		return lexTemplateLiteral(l)
	case strings.HasPrefix(l.input[l.pos:], "\""): // StringLiteral
		return lexStringLiteralDouble(l)
	case strings.HasPrefix(l.input[l.pos:], "'"): // StringLiteral
		return lexStringLiteralSingle(l)
	case (l.goal == InputElementDiv || l.goal == InputElementTemplateTail) && hasDivPunctuator(l): // DivPunctuator
		return lexDivPunctuator(l)
	case (l.goal == InputElementDiv || l.goal == InputElementRegExp) && strings.HasPrefix(l.input[l.pos:], "}"): // RightBracePunctuator
		return lexRightBracePunctuator(l)
	}
	return nil
}

// func f(l *lexer) (stateFunc, bool) {
// 	return nil, false
// }

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
