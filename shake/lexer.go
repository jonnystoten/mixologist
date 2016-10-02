package shake

import (
	"bufio"
	"bytes"
	"io"
)

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isLetter(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isAlphaNum(r rune) bool {
	return isLetter(r) || isDigit(r)
}

type Token int

const (
	ILLEGAL Token = iota
	EOF
	EOL
	WS

	STRING
	NUMBER
)

const eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(reader)}
}

func (s *Scanner) Scan() (tok Token, lit string) {
	r := s.read()

	if isWhitespace(r) {
		s.unread()
		return s.scanWhitespace()
	}
	if isAlphaNum(r) {
		s.unread()
		return s.scanAlphaNum()
	}

	switch r {
	case eof:
		return EOF, ""
	case '\n':
		return EOL, string(r)
	}

	return ILLEGAL, string(r)
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	buf := bytes.Buffer{}
	buf.WriteRune(s.read())

	for {
		if r := s.read(); r == eof {
			break
		} else if !isWhitespace(r) {
			s.unread()
			break
		} else {
			buf.WriteRune(r)
		}
	}

	return WS, " "
}

func (s *Scanner) scanAlphaNum() (tok Token, lit string) {
	buf := bytes.Buffer{}
	allDigits := true

	for {
		if r := s.read(); r == eof {
			break
		} else if !isAlphaNum(r) {
			s.unread()
			break
		} else {
			if !isDigit(r) {
				allDigits = false
			}
			buf.WriteRune(r)
		}
	}

	if allDigits {
		return NUMBER, buf.String()
	}

	return STRING, buf.String()
}

func (s *Scanner) read() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return r
}

func (s *Scanner) unread() {
	s.r.UnreadRune()
}
