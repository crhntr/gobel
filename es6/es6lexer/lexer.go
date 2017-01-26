package es6lexer

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func Lex(name, input string, safe bool) (*Lexer, chan Token) {
	l := &Lexer{
		name:   name,
		input:  input,
		tokens: make(chan Token),
		strict: true,
	}
	l.flags.div = true
	if safe {
		l.setStrict()
	} else {
		l.unsetStrict()
	}

	go l.run()
	return l, l.tokens
}

type Lexer struct {
	name          string     // used for error reports
	input         string     // the string being scanned
	start         int        // start position of this item
	pos           int        // current position of this input
	width         int        // width of last rune read
	tokens        chan Token // channel if scanned tokens
	reservedWords []string
	strict        bool
	flags         struct {
		div          bool
		regExp       bool
		templateTail bool
	}
}

type tokenType int

// Token is a unit generated by the lexer whitch includes a type
// or value
type Token struct {
	Type  tokenType
	Value string
}

func (tok Token) String() string {
	var (
		ok bool

		val       string
		tokTypStr string
	)

	if tokTypStr, ok = tokenTypeStrings[tok.Type]; !ok {
		tokTypStr = "UNKOWN_TOKEN_TYPE"
	}
	if len(tok.Value) > 0 {
		val = " \"" + tok.Value + "\""
	}
	return fmt.Sprintf("<%s%s>", tokTypStr, val)
}

func (tok Token) equals(otherTok Token) bool {
	/*  ||
	!((tok.Err == nil) || (otherTok.Err == nil)) ||
	tok.Err.Error() == otherTok.Err.Error()
	*/
	return tok.Type == otherTok.Type &&
		tok.Value == otherTok.Value
}

const eof rune = -1

/*
// not handled TokenTypes
// TODO missing LineTerminatorSequence
*/
const (
	Error tokenType = iota
	EOF
	// Comment ::
	MultiLineComment
	SingleLineComment
	// WhiteSpace ::
	WhiteSpace
	// LineTerminator ::
	LineTerminator
	// ComomonToken ::
	IdentifierName
	ReservedWord
	//   Punctuator
	Punctuator
	RightBracePunctuator
	DivPunctuator

	NumericLiteral
	StringLiteral
	Template
)

var tokenTypeStrings = map[tokenType]string{
	Error:                "Error",
	EOF:                  "EOF",
	MultiLineComment:     "MultiLineComment",
	SingleLineComment:    "SingleLineComment",
	WhiteSpace:           "WhiteSpace",
	LineTerminator:       "LineTerminator",
	IdentifierName:       "IdentifierName",
	ReservedWord:         "ReservedWord",
	Punctuator:           "Punctuator",
	RightBracePunctuator: "RightBracePunctuator",
	DivPunctuator:        "DivPunctuator",
	NumericLiteral:       "NumericLiteral",
	StringLiteral:        "StringLiteral",
	Template:             "Template",
}

// func (l lexer) String() string {
// 	return fmt.Sprintf("name: \"%s\", start: %d, pos: %d, width: %d, input: \n------\n%s\n-----,",
// 		l.name, l.start, l.pos, l.width, l.input)
// }

func (l *Lexer) setStrict() {
	l.strict = true
	l.reservedWords = []string{}
	l.reservedWords = append(currentReservedWords, futureReservedWords...)
	l.reservedWords = append(currentReservedWords, literals...)
	l.reservedWords = append(l.reservedWords, futureResdervedWordsStrict...)
	sort.Sort(keywordSorter(l.reservedWords))
}

func (l *Lexer) unsetStrict() {
	l.strict = false
	l.reservedWords = []string{}
	l.reservedWords = append(currentReservedWords, futureReservedWords...)
	l.reservedWords = append(currentReservedWords, literals...)
	sort.Sort(keywordSorter(l.reservedWords))
}

type stateFunc func(*Lexer) stateFunc

// run lexes the input by executing state functions
// until the state is nil
func (l *Lexer) run() {
	for state := lexMux; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}

// emit passes an item back to the client.
func (l *Lexer) emit(tok tokenType) {
	l.tokens <- Token{tok, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *Lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r

}

// func (l *lexer) nextN(n int) string {
// 	str := ""
// 	for i := 0; i < n; i++ {
// 		str += string(l.next())
// 	}
// 	return str
// }

// peek returns but does not consume the next
// next rune in the input
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// accept consumes the next rune if it is from
// a valid set
func (l *Lexer) accept(validSet string) bool {
	if strings.IndexRune(validSet, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// // acceptN consumes the next runes n times
// func (l *lexer) acceptRunN(validSet string, n int) bool {
// 	i := 1
// 	for ; i <= n; i++ {
// 		l.accept(validSet)
// 	}
// 	return i == n
// }

// acceptRun consumes a run of runes from the
// valid set
func (l *Lexer) acceptRun(validSet string) bool {
	n := 0
	for strings.IndexRune(validSet, l.next()) >= 0 {
		n++
	}
	l.backup()
	return n > 0
}

// acceptString consumes a string
func (l *Lexer) acceptString(str string) bool {
	if strings.HasPrefix(l.input[l.pos:], str) {
		for _, r := range str {
			l.accept(string(r))
		}
		return true
	}
	return false
}

// acceptAnyString consumes any one of a slice of strings
func (l *Lexer) acceptAnyString(valids []string) bool {
	for _, valid := range valids {
		if l.acceptString(valid) {
			return true
		}
	}
	return false
}

// accepted
func (l *Lexer) accepted() bool {
	return l.pos > l.start
}

// ignore steps over the pending input before
// this point.
func (l *Lexer) ignore() {
	l.start = l.pos
}

func (l *Lexer) ignoreN(n int) {
	l.start += n
}

// reset
func (l *Lexer) reset() {
	l.pos = l.start
}

// backup steps back once per rune
// Can be called once per call of next
func (l *Lexer) backup() {
	l.pos -= l.width
}

// error returns an error token and terminates the scan
// by passing back a nil pointer that will be the next
// state, terminating l.run.
func (l *Lexer) errorf(format string, args ...interface{}) stateFunc {
	l.tokens <- Token{
		Error,
		fmt.Sprintf(format, args...),
	}
	return nil
}

// lexMux multiplexes the various states based on
// input prefix checks. it follows the following rules from
// the specification when determinint what states are allowed
//
// InputElementDiv :: WhiteSpace | LineTerminator | Comment | CommonToken | DivPunctuator | RightBracePunctuator
// InputElementRegExp :: WhiteSpace | LineTerminator | Comment | CommonToken | RightBracePunctuator | RegularExpressionLiteral
// InputElementRegExpOrTemplateTail :: WhiteSpace | LineTerminator | Comment | CommonToken | RegularExpressionLiteral | TemplateSubstitutionTail
// InputElementTemplateTail :: | WhiteSpace | LineTerminator | Comment | CommonToken | DivPunctuator | TemplateSubstitutionTail
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
	case l.flags.div && hasDivPunctuator(l): // DivPunctuator
		return lexDivPunctuator(l)
	case !l.flags.templateTail && strings.HasPrefix(l.input[l.pos:], "}"): // RightBracePunctuator
		return lexRightBracePunctuator(l)
	}
	return nil
}

// func f(l *lexer) (stateFunc, bool) {
// 	return nil, false
// }

var punctuators = []string{
	"{", "(", ")",
	">>>=", "<<=", "!==", "===", ">>>", "...", ">>=", ">=",
	"%=", "*=", "-=", "<=", "&=", "==", "!=", "|=",
	"^=", "+=", "<<", "||", "&&", "++", "--", "=>",
	">>", "-", "&", "|", "^", "!", "~", "%",
	"*", "?", ":", "=", "+", ">", "<", ",",
	";", ".", "]", "["}

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

func lexTemplateLiteral(l *Lexer) stateFunc {
	l.accept("`")
	return nil
}