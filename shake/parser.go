package shake

import (
	"fmt"
	"io"

	"jonnystoten.com/mixologist/mix"
)

type Node interface {
}

type Symbol struct {
	Name string
}

type Statement interface {
	Node
}

type MixStatement struct {
	Symbol  *Symbol
	Op      string
	Address string
}

type Program struct {
	Statements []Statement
}

type pToken struct {
	tok Token
	lit string
}

type Parser struct {
	s   *Scanner
	buf struct {
		len  int
		vals []pToken
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

	symbol := p.parseSymbol()
	stmt.Symbol = symbol

	op, err := p.parseOpCode()
	if err != nil {
		return nil, err
	}
	stmt.Op = op

	if tok, _ := p.scanIgnoreWhitespace(); tok == EOL {
		return stmt, nil
	}
	p.unscan()

	if tok, lit := p.scanIgnoreWhitespace(); tok == NUMBER {
		stmt.Address = lit
	}

	if tok, lit := p.scan(); tok != EOL {
		return nil, fmt.Errorf("Expected EOL (%v, %v)", tok, lit)
	}

	return stmt, nil
}

func (p *Parser) parseSymbol() *Symbol {
	if tok, lit := p.scan(); tok == STRING {
		return &Symbol{Name: lit}
	}

	p.unscan()
	return nil
}

func (p *Parser) parseOpCode() (string, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != STRING {
		return "", fmt.Errorf("Expected OP code (%v, %v)", tok, lit)
	}

	if _, ok := mix.OperationTable[lit]; !ok {
		return "", fmt.Errorf("Unknown OP code (%v)", lit)
	}

	return lit, nil
}

func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) scan() (tok Token, lit string) {
	if p.buf.len > 0 {
		vals := p.buf.vals
		pt := vals[len(vals)-p.buf.len]
		p.buf.len--
		return pt.tok, pt.lit
	}

	tok, lit = p.s.Scan()
	p.buf.vals = append(p.buf.vals, pToken{tok, lit})

	return
}

func (p *Parser) unscan() {
	p.buf.len++
}
