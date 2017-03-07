package es6

import "testing"

func TestLex_MultiLineComment1(t *testing.T) {
	expected := []Token{
		Token{MultiLineComment, "Hello World!"},
	}
	js := "/*Hello World!*/"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_MultiLineComment2(t *testing.T) {
	expected := []Token{
		Token{Error, "no multi line comment terminator \"*/\""},
	}
	js := "/*Hello World!"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_MultiLineComment3(t *testing.T) {
	expected := []Token{
		Token{Error, "no multi line comment terminator \"*/\""},
	}
	js := "/* \""
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_SingleLineComment1(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!"},
	}
	js := "// Hello World!"
	expectedTokens(t, expected, Lex("", js, true))
}
func TestLex_SingleLineComment2(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!"},
	}
	js := "// Hello World!\n"
	expectedTokens(t, expected, Lex("", js, true))
}
func TestLex_Comments(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!"},
		Token{MultiLineComment, "This is a multi\nline comment "},
	}
	js := "// Hello World!\n/*This is a multi\nline comment */"
	expectedTokens(t, expected, Lex("", js, true))
}
