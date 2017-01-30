package es6lexer

import "testing"

func TestLex_RegEx00(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWord, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{IdentifierName, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{Punctuator, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{RegEx, "/abc/i"}, InputElementRegExp},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = /abc/i", true))
}

func TestLex_RegEx01(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWord, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{IdentifierName, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{Punctuator, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{Error, "regex did not close with '/' "}, InputElementRegExp},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = /abc", true))
}

func TestLex_RegEx02(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWord, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{IdentifierName, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{Punctuator, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{Error, "regex can't have new lines"}, InputElementRegExp},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = /abc\n/i", true))
}

func TestLex_RegEx03(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWord, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{IdentifierName, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{Punctuator, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{Error, "invalid character '\u0007' in regex"}, InputElementRegExp},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = /abc\u0007xyz/i", true))
}
