package es6

import (
  "fmt"
  "io"
  "io/ioutil"
  "github.com/crhntr/gobel/es6lexer"
)

// Parser ...
type Parser func(es6lexer.Token) (ASTNode, error)

// DecodeES6 ...
func DecodeES6(r io.Reader) (*ASTNode, error) {
  node := &ASTNode{}
  b, err := ioutil.ReadAll(r)
  if err != nil {
    return node, err
  }

  l := es6lexer.Lex("", string(b), true)

  for {
    tok := l.Next(es6lexer.InputElementDiv)
    if tok.Type == es6lexer.EOF {
      break
    }

    fmt.Println(tok.String())
  }
  return node, err
}

// ParseES6 ...
func ParseES6(lexer *es6lexer.Lexer) (*ASTNode, error) {
  return nil, nil
}
