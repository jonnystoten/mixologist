package shake

import (
	"fmt"
	"io"
)

type Statement interface {
}

type MixStatement struct {
	Op      string
	Address string
}

type Program struct {
	Statements []Statement
}

type Parser struct {
	s   *Scanner
	buf struct {
		full bool
		tok  Token
		lit  string
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() (*Program, error) {
	prg := &Program{}

	for {
		if tok, _ := p.scan(); tok == EOF {
			break
		}
		p.unscan()
		s, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		prg.Statements = append(prg.Statements, s)
	}

	return prg, nil
}

func (p *Parser) parseStatement() (Statement, error) {
	stmt := MixStatement{}

	tok, lit := p.scan()
	if tok != WS {
		return nil, fmt.Errorf("Symbols not supported yet (%v, %v)", tok, lit)
	}

	tok, lit = p.scan()
	if tok != STRING {
		return nil, fmt.Errorf("Expected OP code (%v, %v)", tok, lit)
	}

	stmt.Op = lit

	tok, lit = p.scan()
	if tok == EOL {
		return stmt, nil
	}
	p.unscan()

	tok, lit = p.scan()
	if tok != WS {
		return nil, fmt.Errorf("Expected WS (%v, %v)", tok, lit)
	}

	tok, lit = p.scan()
	if tok == NUMBER {
		stmt.Address = lit
	}

	tok, _ = p.scan()
	if tok != EOL {
		return nil, fmt.Errorf("Expected EOL (%v, %v)", tok, lit)
	}

	return stmt, nil
}

func (p *Parser) scan() (tok Token, lit string) {
	if p.buf.full {
		p.buf.full = false
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *Parser) unscan() {
	p.buf.full = true
}
