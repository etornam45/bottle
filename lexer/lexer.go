package lexer

import (
	"etornam.ben/polygon/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// l.readPosition = 0
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) Next() token.Token {

	var nTok token.Token
	l.eatWhitespace()

	switch l.ch {
	case '=':
		if l.lookAhead() == '=' {
			ch := l.ch
			l.readChar()
			nTok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			nTok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		nTok = newToken(token.PLUS, l.ch)
	case '-':
		nTok = newToken(token.MINUS, l.ch)
	case '!':
		if l.lookAhead() == '=' {
			ch := l.ch
			l.readChar()
			nTok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			nTok = newToken(token.BANG, l.ch)
		}
	case '*':
		nTok = newToken(token.ASTERISK, l.ch)
	case '/':
		nTok = newToken(token.SLASH, l.ch)
	case '%':
		nTok = newToken(token.MOD, l.ch)
	case '<':
		if l.lookAhead() == '=' {
			ch := l.ch
			l.readChar()
			nTok = token.Token{Type: token.LT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			nTok  = newToken(token.LT, l.ch)
		}
	case '>':
		if l.lookAhead() == '=' {
			ch := l.ch
			l.readChar()
			nTok = token.Token{Type: token.GT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			nTok = newToken(token.GT, l.ch)
		}
	case ',':
		nTok = newToken(token.COMMA, l.ch)
	case ';':
		nTok = newToken(token.SEMICOLON, l.ch)
	case ':':
		nTok = newToken(token.COLON, l.ch)
	case '(':
		nTok = newToken(token.LPAREN, l.ch)
	case ')':
		nTok = newToken(token.RPAREN, l.ch)
	case '{':
		nTok = newToken(token.LBRACE, l.ch)
	case '}':
		nTok = newToken(token.RBRACE, l.ch)
  case '[':
    nTok = newToken(token.LBRACKET, l.ch)
  case ']':
    nTok = newToken(token.RBRACKET, l.ch)
  case '"':
    nTok.Literal = l.readStringLiteral()
    nTok.Type = token.STR
  case '\'':
    nTok.Literal = l.readStringLiteral()
    nTok.Type = token.STR
	case 0:
		nTok.Literal = ""
		nTok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			nTok.Literal = l.readIdentifier()
			nTok.Type = token.LookupIdent(nTok.Literal)
			return nTok
		} else if isDigit(l.ch) {
			nTok.Type = token.INT
			nTok.Literal = l.readNumber()
			return nTok
		} else {
			nTok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return nTok
}

func (l *Lexer) lookAhead() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readStringLiteral() string {
  start := l.ch
  l.readChar()
  position := l.position
  for (l.ch != start) {
    l.readChar()
  }
  str := l.input[position:l.position]
  return str
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
