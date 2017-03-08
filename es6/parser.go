package es6

import (
	"io"
	"io/ioutil"
)

// Parser ...
type Parser func(Token) (ASTNode, error)

// DecodeES6Script ...
func DecodeES6Script(r io.Reader) (ASTNode, error) {
	var err error
	var node ASTNode
	b, _ := ioutil.ReadAll(r)

	l := Lex("", string(b), true)

	for {
		tok, _ := l.Next(InputElementDiv)
		if tok.Type == EOFToken {
			break
		}

	}
	return node, err
}

// ParseES6 ...
func ParseES6(lexer *Lexer) (*ASTNode, error) {
	return nil, nil
}
