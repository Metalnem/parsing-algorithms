// Package scan implements lexical scanning for simple expressions.
package scan

import (
	"unicode"
	"unicode/utf8"
)

// Type identifies the token type.
type Type int

const (
	// EOF indicates there are no tokens left.
	EOF Type = iota

	// LeftParen is '('.
	LeftParen

	// RightParen is ')'.
	RightParen

	// Operator is '+', '-', '*', '/' or '^'.
	Operator

	// Number is integer or floating-point number.
	Number

	// Error represents any unknown character.
	Error
)

const eof = -1

// Token represents a single lexical unit.
type Token struct {
	Type  Type
	Value string
}

// Scanner keeps the state of the scanner.
type Scanner struct {
	tokens chan Token
	input  string
	state  stateFn
	start  int
	pos    int
	width  int
}

type stateFn func(*Scanner) stateFn

// NewScanner initializes a new scanner for the input string.
func NewScanner(input string) *Scanner {
	return &Scanner{
		tokens: make(chan Token, 1),
		input:  input,
		state:  lexAny,
	}
}

// Next returns the next token.
func (s *Scanner) Next() Token {
	for s.state != nil {
		select {
		case token := <-s.tokens:
			return token
		default:
			s.state = s.state(s)
		}
	}

	if s.tokens != nil {
		close(s.tokens)
		s.tokens = nil
	}

	return Token{EOF, "EOF"}
}

func (s *Scanner) next() rune {
	if s.pos == len(s.input) {
		s.width = 0
		return eof
	}

	r, w := utf8.DecodeRuneInString(s.input[s.pos:])
	s.width = w
	s.pos += w

	return r
}

func (s *Scanner) peek() rune {
	r := s.next()
	s.backup()

	return r
}

func (s *Scanner) backup() {
	s.pos -= s.width
}

func (s *Scanner) emit(t Type) {
	val := s.input[s.start:s.pos]
	s.tokens <- Token{Type: t, Value: val}

	s.start = s.pos
	s.width = 0
}

func lexAny(s *Scanner) stateFn {
	switch r := s.next(); {
	case r == eof:
		return nil
	case unicode.IsSpace(r):
		return lexSpace
	case r == '(':
		s.emit(LeftParen)
		return lexAny
	case r == ')':
		s.emit(RightParen)
		return lexAny
	case r == '+' || r == '-' || r == '*' || r == '/' || r == '^':
		s.emit(Operator)
		return lexAny
	case r == '.' || unicode.IsDigit(r):
		s.backup()
		return lexNumber
	default:
		s.emit(Error)
		return lexAny
	}
}

func lexSpace(s *Scanner) stateFn {
	for unicode.IsSpace(s.peek()) {
		s.next()
	}

	s.start = s.pos
	return lexAny
}

func lexNumber(s *Scanner) stateFn {
	for unicode.IsDigit(s.peek()) {
		s.next()
	}

	if s.peek() == '.' {
		s.next()

		for unicode.IsDigit(s.peek()) {
			s.next()
		}
	}

	s.emit(Number)
	return lexAny
}
