package es6

import (
  "fmt"
  "io"
  "io/ioutil"
)

// Parser ...
type Parser func(Token) (ASTNode, error)

// DecodeES6 ...
func DecodeES6(r io.Reader) (*ASTNode, error) {
  node := &ASTNode{}
  b, err := ioutil.ReadAll(r)
  if err != nil {
    return node, err
  }

  l := Lex("", string(b), true)

  for {
    tok := l.Next(InputElementDiv)
    if tok.Type == EOF {
      break
    }

    fmt.Println(tok.String())
  }
  return node, err
}

// ParseES6 ...
func ParseES6(lexer *Lexer) (*ASTNode, error) {
  return nil, nil
}
