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
	return (r >= 'A' && r <= 'Z') || r == '∆' || r == '∏' || r == '∑'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isAlphaNum(r rune) bool {
	return isLetter(r) || isDigit(r)
}

func isCharCode(r rune) bool {
	return isAlphaNum(r) || r == ' ' || r == '.' || r == '"' || r == '$' || r == '<' || r == '>' || r == '@' || r == ';' || r == '\''
}

type Token int

const (
	ILLEGAL Token = iota
	EOF
	EOL
	WS

	STRING
	NUMBER

	PLUS
	MINUS
	ASTERISK
	DIVIDE
	SHIFTDIVIDE
	FIELDSIGN

	COMMA
	LPAREN
	RPAREN
	LITERALQUOTE

	STRINGLITERAL
	CHARCODE
)

const eof = rune(0)

type Scanner struct {
	r        *bufio.Reader
	line     int
	col      int
	lastCols map[int]int
}

func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(reader), line: 1, col: 0, lastCols: make(map[int]int)}
}

type Lexeme struct {
	Tok  Token
	Lit  string
	Line int
	Col  int
}

func (s *Scanner) Scan() Lexeme {
	tok, lit := s.scanToken()
	return Lexeme{tok, lit, s.line, s.col}
}

func (s *Scanner) scanToken() (tok Token, lit string) {
	r := s.read()

	if r == '#' {
		// ignore everything up to the end of the line
		for {
			if next := s.read(); next == '\n' {
				return EOL, string(r)
			}
		}
	}

	if r == '*' && s.col == 1 {
		// ignore the whole line
		for {
			if next := s.read(); next == '\n' {
				break
			}
		}

		return s.scanToken()
	}

	if isWhitespace(r) {
		s.unread()
		return s.scanWhitespace()
	}
	if isAlphaNum(r) {
		s.unread()
		return s.scanAlphaNum()
	}
	if r == '"' {
		return s.scanStringLiteral()
	}

	if r == '/' {
		next := s.read()
		if next == '/' {
			return SHIFTDIVIDE, "//"
		}
		s.unread()
		return DIVIDE, string(r)
	}

	switch r {
	case eof:
		return EOF, ""
	case '\n':
		return EOL, string(r)
	case '+':
		return PLUS, string(r)
	case '-':
		return MINUS, string(r)
	case '*':
		return ASTERISK, string(r)
	case ':':
		return FIELDSIGN, string(r)
	case ',':
		return COMMA, string(r)
	case '(':
		return LPAREN, string(r)
	case ')':
		return RPAREN, string(r)
	case '=':
		return LITERALQUOTE, string(r)
	case '.', '$', '<', '>', '@', ';', '"', '\'': // this is just to stop *-comments freaking out
		return CHARCODE, string(r)
	}

	return ILLEGAL, string(r)
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	for {
		if r := s.read(); r == eof {
			break
		} else if !isWhitespace(r) {
			s.unread()
			break
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

func (s *Scanner) scanStringLiteral() (tok Token, lit string) {
	buf := bytes.Buffer{}
	buf.WriteRune('"') // initial quote

	for {
		r := s.read()
		if r == eof {
			return ILLEGAL, buf.String()
		}

		buf.WriteRune(r)

		if !isCharCode(r) {
			return ILLEGAL, buf.String()
		}

		if r == '"' {
			return STRINGLITERAL, buf.String()
		}
	}
}

func (s *Scanner) read() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	if r == '\n' {
		s.lastCols[s.line] = s.col
		s.line++
		s.col = 0
	} else {
		s.col++
	}

	return r
}

func (s *Scanner) unread() {
	if s.col == 0 {
		s.line--
		s.col = s.lastCols[s.line]
	} else {
		s.col--
	}

	s.r.UnreadRune()
}
