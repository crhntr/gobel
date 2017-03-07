package es6lexer

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var errNotNull = fmt.Errorf("not null")

func Test_peek(t *testing.T) {
	l := Lex("", "101", false)

	n1 := l.next()
	p0 := l.peek()
	n0 := l.next()

	if p0 != n0 && p0 == '0' && n1 != '1' {
		t.Error("peek is broken")
	}
}

func TestLex_Whitespace_AND_SingleLineComment(t *testing.T) {
	expected := []Token{
		Token{WhiteSpace, " \t"},
		Token{SingleLineComment, " Hello World!"},
	}
	js := " \t// Hello World!"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex_Terminator_And_Whitespace(t *testing.T) {
	js := "\n\t\n"
	expected := []Token{
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "\t"},
		Token{LineTerminator, "\n"},
	}

	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex_ReservedWord1(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	lJs := Lexer{}
	lJs.setStrict()
	for _, word := range lJs.reservedWords {
		js += word + ws
	}
	l := Lex("", js, true)
	for _, word := range l.reservedWords {
		expected = append(expected, Token{ReservedWord, word})
		expected = append(expected, Token{WhiteSpace, ws})
	}

	expectedTokens(t, expected, l)
}

func TestLex_ReservedWord2(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	lJs := Lexer{}
	lJs.unsetStrict()
	for _, word := range lJs.reservedWords {
		js += word + ws
	}

	l := Lex("", js, false)
	for _, word := range l.reservedWords {
		expected = append(expected, Token{ReservedWord, word})
		expected = append(expected, Token{WhiteSpace, ws})
	}

	expectedTokens(t, expected, l)
}

// func TestLex_EscapeSequence0(t *testing.T) {
// 	expected := []Token{
// 		Token{IdentifierName, "X"},
// 		Token{Punctuator, "&"},
// 		Token{IdentifierName, "ooooooooooooo___"},
// 		Token{LineTerminator, "\n"},
// 	}
// 	js := `"\n"`
// 	l := lex("", js, true)
// 	expectedTokens(t, expected, l)
// }

func TestLex1(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "("},
		Token{Punctuator, ")"},
		Token{Punctuator, "{"},
		Token{RightBracePunctuator, "}"},
		Token{Punctuator, "("},
		Token{Punctuator, ")"},
	}

	js := "function (){}()"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex2(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "("},
		Token{Punctuator, ")"},
		Token{Punctuator, "{"},
		Token{RightBracePunctuator, "}"},
		Token{Punctuator, "("},
		Token{Punctuator, ")"},
	}

	js := "function foo (){}()"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex3(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "add"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "("},
		Token{IdentifierName, "a"},
		Token{Punctuator, ","},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "b"},
		Token{Punctuator, ")"},
		Token{Punctuator, "{"},
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "\t"},
		Token{ReservedWord, "return"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "a"},
		Token{Punctuator, "+"},
		Token{IdentifierName, "b"},
		Token{LineTerminator, "\n"},
		Token{RightBracePunctuator, "}"},
	}

	js := "function add (a, b){\n\treturn a+b\n}"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLexJS(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "fibonacci"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{Punctuator, ")"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "{"},
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "  "},
		Token{ReservedWord, "if"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{WhiteSpace, " "},
		Token{Punctuator, ">="},
		Token{WhiteSpace, " "},
		Token{NumericLiteral, "2"},
		Token{Punctuator, ")"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "{"},
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "    "},
		Token{ReservedWord, "return"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{NumericLiteral, "-1"},
		Token{Punctuator, ")"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "+"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{NumericLiteral, "-2"},
		Token{Punctuator, ")"},
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "  "},
		Token{RightBracePunctuator, "}"},
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "  "},
		Token{ReservedWord, "return"},
		Token{WhiteSpace, " "},
		Token{NumericLiteral, "1"},
		Token{LineTerminator, "\n"},
		Token{RightBracePunctuator, "}"},
		Token{LineTerminator, "\n"},
		Token{LineTerminator, "\n"},
		Token{IdentifierName, "console"},
		Token{Punctuator, "."},
		Token{IdentifierName, "log"},
		Token{Punctuator, "("},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{NumericLiteral, "7"},
		Token{Punctuator, ")"},
		Token{Punctuator, ")"},
		Token{LineTerminator, "\n"},
	}

	testData, err := ioutil.ReadFile("testdata/index01.js")
	if err != nil {
		t.Fatal(err)
	}

	js := string(testData)
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLexJS2(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{Punctuator, ")"},
		Token{Punctuator, "{"},
		Token{ReservedWord, "if"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{Punctuator, ">="},
		Token{NumericLiteral, "2"},
		Token{Punctuator, ")"},
		Token{Punctuator, "{"},
		Token{ReservedWord, "return"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{NumericLiteral, "-1"},
		Token{Punctuator, ")"},
		Token{Punctuator, "+"},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{NumericLiteral, "-2"},
		Token{Punctuator, ")"},
		Token{RightBracePunctuator, "}"},
		Token{ReservedWord, "return"},
		Token{WhiteSpace, " "},
		Token{NumericLiteral, "1"},
		Token{RightBracePunctuator, "}"},
		Token{Punctuator, ";"},
		Token{IdentifierName, "console"},
		Token{Punctuator, "."},
		Token{IdentifierName, "log"},
		Token{Punctuator, "("},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{NumericLiteral, "7"},
		Token{Punctuator, ")"},
		Token{Punctuator, ")"},
		Token{LineTerminator, "\n"},
	}

	testData, err := ioutil.ReadFile("testdata/index02.js")
	if err != nil {
		t.Fatal(err)
	}

	js := string(testData)
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLexJS3(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWord, "function"}, InputElementRegExp},
		TokenTest{Token{WhiteSpace, " "}, InputElementRegExp},
		TokenTest{Token{IdentifierName, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{Punctuator, "("}, InputElementRegExp},
		TokenTest{Token{IdentifierName, "sequence"}, InputElementRegExp},
		TokenTest{Token{Punctuator, ")"}, InputElementRegExp},
		TokenTest{Token{WhiteSpace, " "}, InputElementRegExp},
		TokenTest{Token{Punctuator, "{"}, InputElementRegExp},
		TokenTest{Token{LineTerminator, "\n"}, InputElementRegExp},
		TokenTest{Token{WhiteSpace, "  "}, InputElementRegExp},
		TokenTest{Token{ReservedWord, "return"}, InputElementRegExp},
		TokenTest{Token{WhiteSpace, " "}, InputElementRegExp},
		TokenTest{Token{IdentifierName, "valid"}, InputElementRegExp},
		TokenTest{Token{Punctuator, "."}, InputElementRegExp},
		TokenTest{Token{IdentifierName, "match"}, InputElementRegExp},
		TokenTest{Token{Punctuator, "("}, InputElementRegExp},
		TokenTest{Token{RegEx, "/([CGAT]{3}){1,}/g"}, InputElementRegExp},
		TokenTest{Token{Punctuator, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminator, "\n"}, InputElementRegExp},
		TokenTest{Token{RightBracePunctuator, "}"}, InputElementRegExp},
		TokenTest{Token{LineTerminator, "\n"}, InputElementRegExp},
		TokenTest{Token{IdentifierName, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{Punctuator, "("}, InputElementRegExp},
		TokenTest{Token{StringLiteral, "\"ATATTGGTGTTCATGTGCGCGGGGCCGACGAGCTACTGGCAGAACCACGAGGACAAGAGGTGA\""}, InputElementRegExp},
		TokenTest{Token{Punctuator, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminator, "\n"}, InputElementRegExp},
		TokenTest{Token{IdentifierName, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{Punctuator, "("}, InputElementRegExp},
		TokenTest{Token{StringLiteral, "\"FAIL\""}, InputElementRegExp},
		TokenTest{Token{Punctuator, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminator, "\n"}, InputElementRegExp},
		TokenTest{Token{IdentifierName, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{Punctuator, "("}, InputElementRegExp},
		TokenTest{Token{StringLiteral, "\"Alanine\""}, InputElementRegExp},
		TokenTest{Token{Punctuator, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminator, "\n"}, InputElementRegExp},
		TokenTest{Token{EOF, ""}, InputElementDiv},
	}
	testData, err := ioutil.ReadFile("testdata/TestLexJS3.js")
	if err != nil {
		t.Fatal(err)
	}

	js := string(testData)
	l := Lex("", js, true)
	expectedTokensTable(t, expected, l)
}

func TestToken_String(t *testing.T) {
	t1 := Token{Type(-1), ""}
	if t1.String() == "" {
		t.Error("t1.String() string returns empty string")
	}
	t2 := Token{LineTerminator, ""}
	if t2.String() == "" {
		t.Error("t2.String() string returns empty string")
	}
	t3 := Token{MultiLineComment, " some string value "}
	if t3.String() == "" {
		t.Error("t3.String() string returns empty string")
	}
}

func expectedTokens(t *testing.T, expectedTokens []Token, l *Lexer) {
	i := 0
	for _, expected := range expectedTokens {
		tok := l.Next(InputElementDiv)
		t.Logf("%d: %s %s\n", i, expectedTokens[i], tok)
		if !expected.Equals(tok) {
			t.Errorf("expected and recived tokens do not match [%d](%s != %s)", i, tok, expectedTokens[i])
		}
		i++
	}
	if i < len(expectedTokens) {
		t.Errorf("expected more tokens (expected: %d, got %d)", len(expectedTokens), i)
	}
}

type TokenTest struct {
	Token Token
	Goal  LexerGoal
}

func expectedTokensTable(t *testing.T, expectedTokenRows []TokenTest, l *Lexer) {
	i := 0
	for _, expected := range expectedTokenRows {
		tok := l.Next(expected.Goal)
		t.Logf("%d: %s %s\n", i, expected.Token, tok)
		if !expected.Token.Equals(tok) {
			t.Errorf("expected and recived tokens do not match [%d](%s != %s)", i, expected.Token, tok)
		}
		i++
	}
	if i < len(expectedTokenRows) {
		t.Errorf("expected more tokens (expected: %d, got %d)", len(expectedTokenRows), i)
	}
}
