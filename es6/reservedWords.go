package es6

import "strings"

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
