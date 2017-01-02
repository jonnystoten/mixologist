package shake

import (
	"fmt"
	"io"

	"strconv"

	"jonnystoten.com/mixologist/mix"
)

type Node interface {
}

type Nothing struct{}

type Plus struct{}
type Minus struct{}
type Asterisk struct{}

type Operator struct {
	Type string
}

type Symbol struct {
	Name string
}

type Number struct {
	Value int
}

type LiteralConstant struct {
	Value Node
}

type Expression struct {
	Left     *Node
	Operator Node
	Right    Node
}

type Statement interface {
	Node
}

type MixStatement struct {
	Statement
	Symbol    *Symbol
	Op        string
	APart     Node
	IndexPart Node
	FPart     Node
}

type EquStatement struct {
	Statement
	Symbol  *Symbol
	Address Node
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

type EndStatement struct {
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
		return p.parseEquStatement(symbol)
	case "ORIG":
		return p.parseOrigStatement(symbol)
	case "CON":
		return p.parseConStatement(symbol)
	case "ALF":
		return nil, fmt.Errorf("%v is unsupported", op)
	case "END":
		return p.parseEndStatement(symbol)
	default:
		return p.parseMixStatement(symbol, op)
	}
}

func (p *Parser) parseMixStatement(symbol *Symbol, op string) (MixStatement, error) {
	stmt := MixStatement{Symbol: symbol, Op: op}

	if _, ok := mix.OperationTable[op]; !ok {
		return stmt, fmt.Errorf("Unknown OP code (%v)", op)
	}

	p.swallowWhitespace()

	aPart, err := p.parseAPart()
	if err != nil {
		return stmt, err
	}
	stmt.APart = aPart

	indexPart, err := p.parseIndexPart()
	if err != nil {
		return stmt, err
	}
	stmt.IndexPart = indexPart

	fPart, err := p.parseFPart()
	if err != nil {
		return stmt, err
	}
	stmt.FPart = fPart

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok != EOL {
		return stmt, parseError("Expected EOL", lexeme)
	}

	return stmt, nil
}

func (p *Parser) parseAPart() (Node, error) {
	exp, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	if exp != nil {
		return exp, nil
	}

	if quote := p.scan(); quote.Tok == LITERALQUOTE {
		exp, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if exp == nil {
			return nil, parseError("expected expression after literal quote", quote)
		}

		if endQuote := p.scan(); endQuote.Tok != LITERALQUOTE {
			return nil, parseError("expected closing literal quote", endQuote)
		}

		return LiteralConstant{exp}, nil
	}
	p.unscan()

	return Nothing{}, nil
}

func (p *Parser) parseIndexPart() (Node, error) {
	if comma := p.scan(); comma.Tok == COMMA {
		exp, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if exp == nil {
			return nil, parseError("expected expression after comma", comma)
		}
		return exp, nil
	}
	p.unscan()

	return Nothing{}, nil
}

func (p *Parser) parseFPart() (Node, error) {
	if lparen := p.scan(); lparen.Tok == LPAREN {
		exp, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if exp == nil {
			return nil, parseError("expected expression after lparen", lparen)
		}
		if rparen := p.scan(); rparen.Tok != RPAREN {
			return nil, parseError("expected rparen", rparen)
		}
		return exp, nil
	}
	p.unscan()

	return Nothing{}, nil
}

func (p *Parser) parseExpression() (Node, error) {
	var exp Node

	atom := p.parseAtom()
	if atom != nil {
		exp = atom
	} else {
		lexeme := p.scan()
		if lexeme.Tok == PLUS || lexeme.Tok == MINUS {
			atom := p.parseAtom()
			if atom == nil {
				p.unscan()
				return nil, parseError(fmt.Sprintf("Expected atom after %v", lexeme.Lit), lexeme)
			}

			var operator Node
			if lexeme.Tok == PLUS {
				operator = Plus{}
			} else {
				operator = Minus{}
			}

			exp = Expression{Operator: operator, Right: atom}
		} else {
			p.unscan()
			return nil, nil
		}
	}

	tail, err := p.parseExpressionTail(exp)
	if err != nil {
		return nil, err
	}
	if tail != nil {
		return tail, nil
	}

	return exp, nil
}

func (p *Parser) parseExpressionTail(head Node) (Node, error) {
	lexeme := p.scan()
	switch lexeme.Tok {
	case PLUS, MINUS, ASTERISK, DIVIDE, SHIFTDIVIDE, FIELDSIGN:
		atom := p.parseAtom()
		if atom == nil {
			p.unscan()
			return nil, parseError(fmt.Sprintf("Expected atom after %v", lexeme.Lit), lexeme)
		}
		exp := Expression{Left: &head, Operator: Plus{}, Right: atom} // TODO: hard-coded plus
		full, err := p.parseExpressionTail(exp)
		if err != nil {
			p.unscan()
			return nil, err
		}
		if full != nil {
			return full, nil
		}
		return exp, nil
	default:
		p.unscan()
		return nil, nil
	}
}

func (p *Parser) parseAtom() Node {
	lexeme := p.scan()
	if lexeme.Tok == NUMBER {
		value, _ := strconv.Atoi(lexeme.Lit) // cannot error
		return Number{value}
	}
	if lexeme.Tok == STRING {
		return Symbol{lexeme.Lit}
	}
	if lexeme.Tok == ASTERISK {
		return Asterisk{}
	}

	p.unscan()
	return nil
}

func (p *Parser) parseEquStatement(symbol *Symbol) (EquStatement, error) {
	stmt := EquStatement{Symbol: symbol}

	p.swallowWhitespace()

	exp, err := p.parseExpression()
	if err != nil {
		return stmt, err
	}

	if exp == nil {
		return stmt, fmt.Errorf("Expected expression (%v:%v)", p.s.line, p.s.col)
	}

	stmt.Address = exp

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok != EOL {
		return stmt, parseError("Expected EOL", lexeme)
	}

	return stmt, nil
}

func (p *Parser) parseOrigStatement(symbol *Symbol) (OrigStatement, error) {
	stmt := OrigStatement{Symbol: symbol}

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok == NUMBER {
		stmt.Address = lexeme.Lit
	}

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok != EOL {
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

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok != EOL {
		return stmt, parseError("Expected EOL", lexeme)
	}

	return stmt, nil
}

func (p *Parser) parseEndStatement(symbol *Symbol) (EndStatement, error) {
	stmt := EndStatement{Symbol: symbol}

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok == EOL {
		return stmt, nil
	}
	p.unscan()

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok == NUMBER {
		stmt.Address = lexeme.Lit
	}

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok != EOL {
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

func (p *Parser) swallowWhitespace() {
	lexeme := p.scan()
	if lexeme.Tok != WS {
		p.unscan()
	}
	return
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
