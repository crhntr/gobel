package es6lexer

import "strings"

func lexTemplateLiteral(l *Lexer) stateFunc {
	l.accept("`")
	var r rune
	for {
		if strings.HasPrefix(l.input[l.pos:], "${") {
			l.acceptString("${")
			l.emit(TemplateHead)
			l.goal = InputElementRegExpOrTemplateTail
			return l.state
		}
		if strings.HasPrefix(l.input[l.pos:], "`") {
			l.accept("`")
			l.emit(NoSubstitutionTemplate)
			return l.state
		}
		if r = l.next(); r == eof {
			l.errorf("did not reach end of template literal reached eof")
			return nil
		}
	}
}

func lexTemplateSubstitutionTail(l *Lexer) stateFunc {
	l.accept("}")
	var r rune
	for {
		if strings.HasPrefix(l.input[l.pos:], "${") {
			l.acceptString("${")
			l.emit(TemplateMiddle)
			return l.state
		}
		if strings.HasPrefix(l.input[l.pos:], "`") {
			l.accept("`")
			l.emit(TemplateTail)
			return l.state
		}
		if r = l.next(); r == eof {
			l.errorf("did not reach TemplateMiddle or TemplateTail but reached eof")
			return nil
		}
	}
}
