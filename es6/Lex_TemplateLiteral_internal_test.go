package es6

import (
	"testing"
)

func TestLex_TemplateLiteral01(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{NoSubstitutionTemplateToken, "`foo`"},
	}
	js := "var foo = `foo`"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_TemplateLiteral02(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{NoSubstitutionTemplateToken, "``"},
	}
	js := "var foo = ``"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_TemplateLiteral03(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{ErrorToken, "did not reach end of template literal reached eof"},
	}
	js := "var foo = `"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_TemplateLiteral04(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{TemplateHeadToken, "`Hello ${"},
		Token{IdentifierNameToken, "friend"},
	}
	js := "var foo = `Hello ${friend"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_TemplateLiteral05(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWordToken, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{PunctuatorToken, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "friend"}, InputElementDiv},
		TokenTest{Token{TemplateTailToken, "}!`"}, InputElementRegExpOrTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}!`", true))
}

func TestLex_TemplateLiteral06(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWordToken, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{PunctuatorToken, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "friend"}, InputElementDiv},
		TokenTest{Token{ErrorToken, "did not reach TemplateMiddle or TemplateTail but reached eof"}, InputElementRegExpOrTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}! ", true))
}

func TestLex_TemplateLiteral07(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWordToken, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{PunctuatorToken, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "friend"}, InputElementDiv},
		TokenTest{Token{TemplateMiddleToken, "}! ${"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{TemplateHeadToken, "` ${"}, InputElementDiv},
		TokenTest{Token{NumericLiteralToken, "4"}, InputElementDiv},
		TokenTest{Token{TemplateTailToken, "}`"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{PunctuatorToken, "+"}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`${"}, InputElementDiv},
		TokenTest{Token{NumericLiteralToken, "2"}, InputElementDiv},
		TokenTest{Token{TemplateTailToken, "} `"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{TemplateTailToken, "}`"}, InputElementRegExpOrTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}! ${` ${4}`+`${2} `}`", true))
}

func TestLex_TemplateLiteral08(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWordToken, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{PunctuatorToken, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "friend"}, InputElementDiv},
		TokenTest{Token{TemplateMiddleToken, "}! ${"}, InputElementTemplateTail},
		TokenTest{Token{TemplateHeadToken, "` ${"}, InputElementDiv},
		TokenTest{Token{NumericLiteralToken, "4"}, InputElementDiv},
		TokenTest{Token{TemplateTailToken, "}`"}, InputElementTemplateTail},
		TokenTest{Token{PunctuatorToken, "+"}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`${"}, InputElementDiv},
		TokenTest{Token{NumericLiteralToken, "2"}, InputElementDiv},
		TokenTest{Token{TemplateTailToken, "} `"}, InputElementTemplateTail},
		TokenTest{Token{TemplateTailToken, "}`"}, InputElementTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}! ${` ${4}`+`${2} `}`", true))
}
