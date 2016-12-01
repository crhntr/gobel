package ecgo

import (
	"fmt"
	"testing"
)

var errNotNull = fmt.Errorf("not null")

func Test_peek(t *testing.T) {
	l, _ := lex("", "101", false)

	n1 := l.next()
	p0 := l.peek()
	n0 := l.next()

	if p0 != n0 && p0 == '0' && n1 != '1' {
		t.Error("peek is broken")
	}
}

func Test_error(t *testing.T) {
	c := make(chan Token)
	l := &lexer{
		name:   "",
		input:  "3zzz",
		tokens: c,
	}

	go func() {
		l.next()
		l.next()
		l.next()
		l.next()
		l.error(errNotNull)
		close(c)
	}()

	for tok := range c {
		if tok.Err == nil || tok.Err.Error() != errNotNull.Error() {
			t.Error("l.error(err error) should set the l.Err")
		}
		if tok.Type != Error {
			t.Error("l.error(err error) should send Token with tokenType Error")
		}
		if tok.Value != l.input {
			t.Error("l.error(err error) should set Value to offending string")
		}
	}
}

func TestLex_MultiLineComment1(t *testing.T) {
	expected := []Token{
		Token{MultiLineComment, "Hello World!", nil},
	}
	js := "/*Hello World!*/"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_MultiLineComment2(t *testing.T) {
	expected := []Token{
		Token{Error, "Hello World!", errNotNull},
	}
	js := "/*Hello World!"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_MultiLineComment3(t *testing.T) {
	expected := []Token{
		Token{Error, " \"", errNotNull},
	}
	js := "/* \""
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_SingleLineComment1(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!", nil},
	}
	js := "// Hello World!"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}
func TestLex_SingleLineComment2(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!", nil},
	}
	js := "// Hello World!\n"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}
func TestLex_Comments(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!", nil},
		Token{MultiLineComment, "This is a multi\nline comment ", nil},
	}
	js := "// Hello World!\n/*This is a multi\nline comment */"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Whitespace1(t *testing.T) {
	js := " \t"
	expected := []Token{
		Token{WhiteSpace, js, nil},
	}
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Whitespace2(t *testing.T) {
	js := " "
	expected := []Token{
		Token{WhiteSpace, js, nil},
	}

	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Whitespace_AND_SingleLineComment(t *testing.T) {
	expected := []Token{
		Token{WhiteSpace, " \t", nil},
		Token{SingleLineComment, " Hello World!", nil},
	}
	js := " \t// Hello World!"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_LineTerminator1(t *testing.T) {
	js := "\n"
	expected := []Token{
		Token{LineTerminator, js, nil},
	}

	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_LineTerminator2(t *testing.T) {
	js := "\n\n"
	expected := []Token{
		Token{LineTerminator, "\n", nil},
		Token{LineTerminator, "\n", nil},
	}

	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Terminator_And_Whitespace(t *testing.T) {
	js := "\n\t\n"
	expected := []Token{
		Token{LineTerminator, "\n", nil},
		Token{WhiteSpace, "\t", nil},
		Token{LineTerminator, "\n", nil},
	}

	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_ReservedWord1(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	lJs := lexer{}
	lJs.setStrict()
	for _, word := range lJs.reservedWords {
		js += word + ws
	}
	l, tokens := lex("", js, true)
	for _, word := range l.reservedWords {
		expected = append(expected, Token{ReservedWord, word, nil})
		expected = append(expected, Token{WhiteSpace, ws, nil})
	}

	expectedTokens(t, expected, tokens)
}

func TestLex_ReservedWord2(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	lJs := lexer{}
	lJs.unsetStrict()
	for _, word := range lJs.reservedWords {
		js += word + ws
	}

	l, tokens := lex("", js, false)
	for _, word := range l.reservedWords {
		expected = append(expected, Token{ReservedWord, word, nil})
		expected = append(expected, Token{WhiteSpace, ws, nil})
	}

	expectedTokens(t, expected, tokens)
}

func TestLex_Punctuator1(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	for _, punct := range punctuators {
		expected = append(expected, Token{Punctuator, punct, nil})
		expected = append(expected, Token{WhiteSpace, ws, nil})
		js += punct + ws
	}

	_, tokens := lex("", js, false)
	expectedTokens(t, expected, tokens)
}

func TestLex_DivPunctuator1(t *testing.T) {
	expected := []Token{
		Token{DivPunctuator, "/", nil},
		Token{WhiteSpace, " ", nil},
		Token{DivPunctuator, "/=", nil},
	}
	js := "/ /="
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_RightBracePunctuator1(t *testing.T) {
	expected := []Token{
		Token{Punctuator, "{", nil},
		Token{RightBracePunctuator, "}", nil},
	}
	js := "{}"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Identifier1(t *testing.T) {
	expected := []Token{
		Token{IdentifierName, "$", nil},
		Token{Punctuator, "=", nil},
		Token{IdentifierName, "_", nil},
		Token{Punctuator, "=", nil},
		Token{IdentifierName, "foo", nil},
	}
	js := "$=_=foo"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Identifier2(t *testing.T) {
	expected := []Token{
		Token{IdentifierName, "X", nil},
		Token{Punctuator, "&", nil},
		Token{IdentifierName, "ooooooooooooo___", nil},
		Token{LineTerminator, "\n", nil},
	}
	js := "X&ooooooooooooo___\n"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Identifier2(t *testing.T) {
	expected := []Token{
		Token{IdentifierName, "X", nil},
		Token{Punctuator, "&", nil},
		Token{IdentifierName, "ooooooooooooo___", nil},
		Token{LineTerminator, "\n", nil},
	}
	js := "X&ooooooooooooo___\n"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex1(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function", nil},
		Token{WhiteSpace, " ", nil},
		Token{Punctuator, "(", nil},
		Token{Punctuator, ")", nil},
		Token{Punctuator, "{", nil},
		Token{RightBracePunctuator, "}", nil},
		Token{Punctuator, "(", nil},
		Token{Punctuator, ")", nil},
	}

	js := "function (){}()"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex2(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function", nil},
		Token{WhiteSpace, " ", nil},
		Token{IdentifierName, "foo", nil},
		Token{WhiteSpace, " ", nil},
		Token{Punctuator, "(", nil},
		Token{Punctuator, ")", nil},
		Token{Punctuator, "{", nil},
		Token{RightBracePunctuator, "}", nil},
		Token{Punctuator, "(", nil},
		Token{Punctuator, ")", nil},
	}

	js := "function foo (){}()"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex3(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function", nil},
		Token{WhiteSpace, " ", nil},
		Token{IdentifierName, "add", nil},
		Token{WhiteSpace, " ", nil},
		Token{Punctuator, "(", nil},
		Token{IdentifierName, "a", nil},
		Token{Punctuator, ",", nil},
		Token{WhiteSpace, " ", nil},
		Token{IdentifierName, "b", nil},
		Token{Punctuator, ")", nil},
		Token{Punctuator, "{", nil},
		Token{LineTerminator, "\n", nil},
		Token{WhiteSpace, "\t", nil},
		Token{ReservedWord, "return", nil},
		Token{WhiteSpace, " ", nil},
		Token{IdentifierName, "a", nil},
		Token{Punctuator, "+", nil},
		Token{IdentifierName, "b", nil},
		Token{LineTerminator, "\n", nil},
		Token{RightBracePunctuator, "}", nil},
	}

	js := "function add (a, b){\n\treturn a+b\n}"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestToken_String(t *testing.T) {
	t1 := Token{tokenType(-1), "", nil}
	if t1.String() == "" {
		t.Error("t1.String() string returns empty string")
	}
	t2 := Token{LineTerminator, "", nil}
	if t2.String() == "" {
		t.Error("t2.String() string returns empty string")
	}
	t3 := Token{MultiLineComment, " MultiLineComment", nil}
	if t3.String() == "" {
		t.Error("t3.String() string returns empty string")
	}
}

func expectedTokens(t *testing.T, expectedTokens []Token, tokens chan Token) {
	i := 0
	for tok := range tokens {
		t.Logf("%d: %s %s\n", i, expectedTokens[i], tok)
		if i+1 > len(expectedTokens) {
			t.Errorf("expected fewer tokens (expected: %d, got %d)", len(expectedTokens), i+1)
		}
		if !expectedTokens[i].equals(tok) {
			t.Errorf("expected and recived tokens do not match [%d](%s != %s)", i, tok, expectedTokens[i])
		}
		i++
	}
	if i+1 < len(expectedTokens) {
		t.Errorf("expected more tokens (expected: %d, got %d)", len(expectedTokens), i+1)
	}
}
