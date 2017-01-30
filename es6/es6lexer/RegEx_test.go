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
		TokenTest{Token{Error, "regex can't have new lines"}, InputElementRegExp},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = /ab\nc/i", true))
}
