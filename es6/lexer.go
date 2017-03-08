package es6

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

// Lex lexes a string into tokens
func Lex(name, input string, safe bool) *Lexer {
	l := &Lexer{
		name:   name,
		input:  input,
		state:  lexInputElement,
		tokens: []Token{},
		strict: true,
		goal:   InputElementDiv,
		line:   1,
		column: 0,
	}
	if safe {
		l.setStrict()
	} else {
		l.unsetStrict()
	}
	return l
}

// A Lexer represents the state of the lexing algorithm use func Lex to return
// an initalized Lexer
type Lexer struct {
	name           string // used for error reports
	state          stateFunc
	input          string  // the string being scanned
	start          int     // start position of this item
	pos            int     // current position of this input
	width          int     // width of last rune read
	tokens         []Token // chan Token // channel if scanned tokens
	reservedWords  []string
	strict         bool
	goal           LexerGoal
	SkipWhitespace bool

	line   int
	column int
}

// LexerGoal represents a lexing goal
// There are several situations where the identification of lexical input
// elements is sensitive to the syntactic grammar context that is consuming the
// input elements. This requires multiple goal symbols for the lexical grammar.
type LexerGoal int

const (
	// InputElementDiv is used as the lexical goal symbol when none of the following
	// goals are set
	InputElementDiv LexerGoal = iota
	// InputElementRegExp goal symbol is used in all syntactic grammar
	// contexts where a RegularExpressionLiteral is permitted but neither a
	// TemplateMiddle, nor a TemplateTail is permitted
	InputElementRegExp
	// InputElementRegExpOrTemplateTail goal is used in syntactic grammar
	// contexts where a RegularExpressionLiteral, a TemplateMiddle, or a
	// TemplateTail is permitted
	InputElementRegExpOrTemplateTail
	// InputElementTemplateTail goal is used in all syntactic grammar contexts
	// where a TemplateMiddle or a TemplateTail is permitted but a
	// RegularExpressionLiteral is not permitted
	InputElementTemplateTail
)

func (goal LexerGoal) String() string {
	switch goal {
	case InputElementDiv:
		return "InputElementDiv"
	case InputElementRegExp:
		return "InputElementRegExp"
	case InputElementRegExpOrTemplateTail:
		return "InputElementRegExpOrTemplateTail"
	case InputElementTemplateTail:
		return "InputElementTemplateTail"
	default:
		return "UnknownGoal"
	}
}

// Next returns the next token
func (l *Lexer) Next(goal LexerGoal) Token {
	l.goal = goal
	l.state = lexInputElement
	for len(l.tokens) < 1 {
		l.state = l.state(l)
	}
	tok := l.tokens[0]
	l.tokens = l.tokens[1:]
	return tok
}

// Peek returns the next token
func (l *Lexer) Peek(goal LexerGoal) Token {
	l.goal = goal
	l.state = lexInputElement
	for len(l.tokens) < 1 {
		l.state = l.state(l)
	}
	return l.tokens[0]
}

// CurrentPosition returns the Lexer's current position
func (l *Lexer) CurrentPosition() (offset, line, column int) {
	return l.pos, l.line, l.column
}

const eof rune = -1

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

// emit passes an item back to the client.
func (l *Lexer) emit(typ TokenType) {
	val := l.input[l.start:l.pos]
	// l.tokens <- Token{typ, val}
	if typ != WhiteSpaceToken || (!l.SkipWhitespace && typ == WhiteSpaceToken) {
		l.tokens = append(
			l.tokens,
			Token{
				Type:  typ,
				Value: val,
				FilePosition: FilePosition{
					fileName: l.name,
					offset:   l.pos,
					line:     l.line,
					column:   l.column,
				},
			},
		)
	}
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

var maxUtf8Rune = utf16.DecodeRune(0xDBFF, 0xDFFF)

// accept consumes the next rune if it is from
// a valid set
func (l *Lexer) acceptSourceCharacterRunExcept(invalidSet string) bool {
	n := 0
	for {
		r := l.next()
		if r < '\u0000' || r > maxUtf8Rune || strings.ContainsRune(invalidSet, r) {
			break
		}
		n++
	}
	l.backup()
	return n > 0
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

// acceptRunN consumes a run of runes from the
// valid set of exactly length n
// func (l *Lexer) acceptRunN(validSet string, n int) bool {
// 	acceptedWidth := 0
// 	for strings.IndexRune(validSet, l.next()) >= 0 {
// 		n--
// 		acceptedWidth += l.width
// 	}
// 	if n != 0 {
// 		l.pos -= acceptedWidth
// 		return false
// 	}
// 	return true
// }

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
	// l.tokens <- Token{
	// 	Error,
	// 	fmt.Sprintf(format, args...),
	// }
	l.tokens = append(l.tokens, Token{
		Type:  ErrorToken,
		Value: fmt.Sprintf(format, args...),
	})
	return nil
}
