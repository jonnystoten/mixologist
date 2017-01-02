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
	Statement
	Symbol  *Symbol
	Op      string
	Address string
}

type OrigStatement struct {
	Statement
	Symbol  *Symbol
	Address string
}

type ConStatement struct {
	Statement
	Symbol  *Symbol
	Address string
}

type Program struct {
	Statements []Statement
}

type Parser struct {
	s   *Scanner
	buf struct {
		len  int
		vals []Lexeme
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() (*Program, error) {
	prg := &Program{}

	for {
		if lexeme := p.scan(); lexeme.Tok == EOF {
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
	if lexeme := p.scan(); lexeme.Tok == EOL {
		return p.parseStatement()
	}
	p.unscan()

	symbol := p.parseSymbol()
	op, err := p.parseOpCode()
	if err != nil {
		return nil, err
	}

	switch op {
	case "EQU":
		return nil, fmt.Errorf("%v is unsupported", op)
	case "ORIG":
		return p.parseOrigStatement(symbol)
	case "CON":
		return p.parseConStatement(symbol)
	case "ALF":
		return nil, fmt.Errorf("%v is unsupported", op)
	case "END":
		return nil, fmt.Errorf("%v is unsupported", op)
	default:
		return p.parseMixStatement(symbol, op)
	}
}

func (p *Parser) parseMixStatement(symbol *Symbol, op string) (MixStatement, error) {
	stmt := MixStatement{Symbol: symbol, Op: op}

	if _, ok := mix.OperationTable[op]; !ok {
		return stmt, fmt.Errorf("Unknown OP code (%v)", op)
	}

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok == EOL {
		return stmt, nil
	}
	p.unscan()

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok == NUMBER {
		stmt.Address = lexeme.Lit
	}

	if lexeme := p.scan(); lexeme.Tok != EOL {
		return stmt, parseError("Expected EOL", lexeme)
	}

	return stmt, nil
}

func (p *Parser) parseOrigStatement(symbol *Symbol) (OrigStatement, error) {
	stmt := OrigStatement{Symbol: symbol}

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok == NUMBER {
		stmt.Address = lexeme.Lit
	}

	if lexeme := p.scan(); lexeme.Tok != EOL {
		return stmt, parseError("Expected EOL", lexeme)
	}

	return stmt, nil
}

func (p *Parser) parseConStatement(symbol *Symbol) (ConStatement, error) {
	stmt := ConStatement{Symbol: symbol}

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok == EOL {
		return stmt, nil
	}
	p.unscan()

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok == NUMBER {
		stmt.Address = lexeme.Lit
	}

	if lexeme := p.scan(); lexeme.Tok != EOL {
		return stmt, parseError("Expected EOL", lexeme)
	}

	return stmt, nil
}

func (p *Parser) parseSymbol() *Symbol {
	if lexeme := p.scan(); lexeme.Tok == STRING {
		return &Symbol{Name: lexeme.Lit}
	}

	p.unscan()
	return nil
}

func (p *Parser) parseOpCode() (string, error) {
	lexeme := p.scanIgnoreWhitespace()
	if lexeme.Tok != STRING {
		return "", parseError("Expected OP code", lexeme)
	}

	return lexeme.Lit, nil
}

func (p *Parser) scanIgnoreWhitespace() (lexeme Lexeme) {
	lexeme = p.scan()
	if lexeme.Tok == WS {
		lexeme = p.scan()
	}
	return
}

func (p *Parser) scan() (lexeme Lexeme) {
	if p.buf.len > 0 {
		vals := p.buf.vals
		lexeme = vals[len(vals)-p.buf.len]
		p.buf.len--
		return
	}

	lexeme = p.s.Scan()
	p.buf.vals = append(p.buf.vals, lexeme)
	return
}

func (p *Parser) unscan() {
	p.buf.len++
}

func parseError(err string, lexeme Lexeme) error {
	return fmt.Errorf("%v: %v (%v:%v)", err, lexeme.Lit, lexeme.Line, lexeme.Col)
}
