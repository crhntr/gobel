package es6

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
		Token{WhiteSpaceToken, " \t"},
		Token{SingleLineCommentToken, " Hello World!"},
	}
	js := " \t// Hello World!"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex_Terminator_And_Whitespace(t *testing.T) {
	js := "\n\t\n"
	expected := []Token{
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "\t"},
		Token{LineTerminatorToken, "\n"},
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
		expected = append(expected, Token{ReservedWordToken, word})
		expected = append(expected, Token{WhiteSpaceToken, ws})
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
		expected = append(expected, Token{ReservedWordToken, word})
		expected = append(expected, Token{WhiteSpaceToken, ws})
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
		Token{ReservedWordToken, "function"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "("},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "{"},
		Token{RightBracePunctuatorToken, "}"},
		Token{PunctuatorToken, "("},
		Token{PunctuatorToken, ")"},
	}

	js := "function (){}()"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex2(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "function"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "("},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "{"},
		Token{RightBracePunctuatorToken, "}"},
		Token{PunctuatorToken, "("},
		Token{PunctuatorToken, ")"},
	}

	js := "function foo (){}()"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex3(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "function"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "add"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "a"},
		Token{PunctuatorToken, ","},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "b"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "{"},
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "\t"},
		Token{ReservedWordToken, "return"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "a"},
		Token{PunctuatorToken, "+"},
		Token{IdentifierNameToken, "b"},
		Token{LineTerminatorToken, "\n"},
		Token{RightBracePunctuatorToken, "}"},
	}

	js := "function add (a, b){\n\treturn a+b\n}"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLexJS(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "function"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "fibonacci"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{PunctuatorToken, ")"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "{"},
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "  "},
		Token{ReservedWordToken, "if"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, ">="},
		Token{WhiteSpaceToken, " "},
		Token{NumericLiteralToken, "2"},
		Token{PunctuatorToken, ")"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "{"},
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "    "},
		Token{ReservedWordToken, "return"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{NumericLiteralToken, "-1"},
		Token{PunctuatorToken, ")"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "+"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{NumericLiteralToken, "-2"},
		Token{PunctuatorToken, ")"},
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "  "},
		Token{RightBracePunctuatorToken, "}"},
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "  "},
		Token{ReservedWordToken, "return"},
		Token{WhiteSpaceToken, " "},
		Token{NumericLiteralToken, "1"},
		Token{LineTerminatorToken, "\n"},
		Token{RightBracePunctuatorToken, "}"},
		Token{LineTerminatorToken, "\n"},
		Token{LineTerminatorToken, "\n"},
		Token{IdentifierNameToken, "console"},
		Token{PunctuatorToken, "."},
		Token{IdentifierNameToken, "log"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{NumericLiteralToken, "7"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, ")"},
		Token{LineTerminatorToken, "\n"},
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
		Token{ReservedWordToken, "function"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "{"},
		Token{ReservedWordToken, "if"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{PunctuatorToken, ">="},
		Token{NumericLiteralToken, "2"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "{"},
		Token{ReservedWordToken, "return"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{NumericLiteralToken, "-1"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "+"},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{NumericLiteralToken, "-2"},
		Token{PunctuatorToken, ")"},
		Token{RightBracePunctuatorToken, "}"},
		Token{ReservedWordToken, "return"},
		Token{WhiteSpaceToken, " "},
		Token{NumericLiteralToken, "1"},
		Token{RightBracePunctuatorToken, "}"},
		Token{PunctuatorToken, ";"},
		Token{IdentifierNameToken, "console"},
		Token{PunctuatorToken, "."},
		Token{IdentifierNameToken, "log"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{NumericLiteralToken, "7"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, ")"},
		Token{LineTerminatorToken, "\n"},
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
		TokenTest{Token{ReservedWordToken, "function"}, InputElementRegExp},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "("}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "sequence"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, ")"}, InputElementRegExp},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "{"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{WhiteSpaceToken, "  "}, InputElementRegExp},
		TokenTest{Token{ReservedWordToken, "return"}, InputElementRegExp},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "valid"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "."}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "match"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "("}, InputElementRegExp},
		TokenTest{Token{RegExToken, "/([CGAT]{3}){1,}/g"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{RightBracePunctuatorToken, "}"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "("}, InputElementRegExp},
		TokenTest{Token{StringLiteralToken, "\"ATATTGGTGTTCATGTGCGCGGGGCCGACGAGCTACTGGCAGAACCACGAGGACAAGAGGTGA\""}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "("}, InputElementRegExp},
		TokenTest{Token{StringLiteralToken, "\"FAIL\""}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "("}, InputElementRegExp},
		TokenTest{Token{StringLiteralToken, "\"Alanine\""}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{EOFToken, ""}, InputElementDiv},
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
	t1 := Token{TokenType(-1), ""}
	if t1.String() == "" {
		t.Error("t1.String() string returns empty string")
	}
	t2 := Token{LineTerminatorToken, ""}
	if t2.String() == "" {
		t.Error("t2.String() string returns empty string")
	}
	t3 := Token{MultiLineCommentToken, " some string value "}
	if t3.String() == "" {
		t.Error("t3.String() string returns empty string")
	}
}

func expectedTokens(t *testing.T, expectedTokens []Token, l *Lexer) {
	i := 0
	for _, expected := range expectedTokens {
		tok, _ := l.Next(InputElementDiv)
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
		tok , _:= l.Next(expected.Goal)
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
