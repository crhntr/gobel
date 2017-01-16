package es6

import "testing"

func TestLex_MultiLineComment1(t *testing.T) {
	expected := []Token{
		Token{MultiLineComment, "Hello World!"},
	}
	js := "/*Hello World!*/"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_MultiLineComment2(t *testing.T) {
	expected := []Token{
		Token{Error, "no multi line comment terminator \"*/\""},
	}
	js := "/*Hello World!"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_MultiLineComment3(t *testing.T) {
	expected := []Token{
		Token{Error, "no multi line comment terminator \"*/\""},
	}
	js := "/* \""
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_SingleLineComment1(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!"},
	}
	js := "// Hello World!"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}
func TestLex_SingleLineComment2(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!"},
	}
	js := "// Hello World!\n"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}
func TestLex_Comments(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!"},
		Token{MultiLineComment, "This is a multi\nline comment "},
	}
	js := "// Hello World!\n/*This is a multi\nline comment */"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}
