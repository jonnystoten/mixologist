package shake

import (
	"fmt"
	"io"

	"strconv"

	"jonnystoten.com/mixologist/mix"
)

type Node interface {
	Accept(NodeVisitor) int
}

type NodeVisitor interface {
	visitNothing(nothing Nothing) int
	visitNumber(number Number) int
	visitSymbol(symbol Symbol) int
	visitAsterisk(asterisk Asterisk) int
	visitExpression(expression Expression) int
	visitWValue(wValue WValue) int
	visitLiteralConstant(literal LiteralConstant) int
}

type Nothing struct{}

func (n Nothing) Accept(v NodeVisitor) int {
	return v.visitNothing(n)
}

type Asterisk struct{}

func (a Asterisk) Accept(v NodeVisitor) int {
	return v.visitAsterisk(a)
}

type Symbol struct {
	Name string
}

func (s Symbol) Accept(v NodeVisitor) int {
	return v.visitSymbol(s)
}

type Number struct {
	Value int
}

func (n Number) Accept(v NodeVisitor) int {
	return v.visitNumber(n)
}

type LiteralConstant struct {
	Value Node
}

func (lc LiteralConstant) Accept(v NodeVisitor) int {
	return v.visitLiteralConstant(lc)
}

type Expression struct {
	Left     *Node
	Operator Token
	Right    Node
}

func (e Expression) Accept(v NodeVisitor) int {
	return v.visitExpression(e)
}

type WValue struct {
	Parts []WValuePart
}

func (w WValue) Accept(v NodeVisitor) int {
	return v.visitWValue(w)
}

type WValuePart struct {
	Exp   Node
	FPart Node
}

type Statement interface {
	Symbol() *Symbol
}

type MixStatement struct {
	symbol    *Symbol
	Op        string
	APart     Node
	IndexPart Node
	FPart     Node
}

type EquStatement struct {
	symbol  *Symbol
	Address Node
}

type OrigStatement struct {
	symbol  *Symbol
	Address Node
}

type ConStatement struct {
	symbol  *Symbol
	Address Node
}

type AlfStatement struct {
	symbol   *Symbol
	CharCode string
}

type EndStatement struct {
	symbol  *Symbol
	Address Node
}

func (s MixStatement) Symbol() *Symbol  { return s.symbol }
func (s EquStatement) Symbol() *Symbol  { return s.symbol }
func (s OrigStatement) Symbol() *Symbol { return s.symbol }
func (s ConStatement) Symbol() *Symbol  { return s.symbol }
func (s AlfStatement) Symbol() *Symbol  { return s.symbol }
func (s EndStatement) Symbol() *Symbol  { return s.symbol }

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
		return p.parseAlfStatement(symbol)
	case "END":
		return p.parseEndStatement(symbol)
	default:
		return p.parseMixStatement(symbol, op)
	}
}

func (p *Parser) parseMixStatement(symbol *Symbol, op string) (MixStatement, error) {
	stmt := MixStatement{symbol: symbol, Op: op}

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

func (p *Parser) parseWValue() (Node, error) {
	part, err := p.parseWValuePart()
	if err != nil {
		return nil, err
	}
	if part == nil {
		return nil, nil
	}

	parts := []WValuePart{*part}

	return p.parseWValueTail(parts)
}

func (p *Parser) parseWValuePart() (*WValuePart, error) {
	exp, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	if exp == nil {
		return nil, nil
	}
	fPart, err := p.parseFPart()
	if err != nil {
		return nil, err
	}
	if fPart == nil {
		return nil, nil
	}

	return &WValuePart{exp, fPart}, nil
}

func (p *Parser) parseWValueTail(head []WValuePart) (Node, error) {
	nextPart, err := p.parseWValuePart()
	if err != nil {
		return nil, err
	}
	if nextPart == nil {
		return WValue{Parts: head}, nil
	}

	return p.parseWValueTail(append(head, *nextPart))
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

			exp = Expression{Operator: lexeme.Tok, Right: atom}
		} else {
			p.unscan()
			return nil, nil
		}
	}

	return p.parseExpressionTail(exp)
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
		exp := Expression{Left: &head, Operator: lexeme.Tok, Right: atom}
		return p.parseExpressionTail(exp)
	default:
		p.unscan()
		return head, nil
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
	stmt := EquStatement{symbol: symbol}

	p.swallowWhitespace()

	wval, err := p.parseWValue()
	if err != nil {
		return stmt, err
	}

	if wval == nil {
		return stmt, fmt.Errorf("Expected W-value (%v:%v)", p.s.line, p.s.col)
	}

	stmt.Address = wval

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok != EOL {
		return stmt, parseError("Expected EOL", lexeme)
	}

	return stmt, nil
}

func (p *Parser) parseOrigStatement(symbol *Symbol) (OrigStatement, error) {
	stmt := OrigStatement{symbol: symbol}

	p.swallowWhitespace()

	wval, err := p.parseWValue()
	if err != nil {
		return stmt, err
	}

	if wval == nil {
		return stmt, fmt.Errorf("Expected W-value (%v:%v)", p.s.line, p.s.col)
	}

	stmt.Address = wval

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok != EOL {
		return stmt, parseError("Expected EOL", lexeme)
	}

	return stmt, nil
}

func (p *Parser) parseConStatement(symbol *Symbol) (ConStatement, error) {
	stmt := ConStatement{symbol: symbol}

	p.swallowWhitespace()

	wval, err := p.parseWValue()
	if err != nil {
		return stmt, err
	}

	if wval == nil {
		return stmt, fmt.Errorf("Expected W-value (%v:%v)", p.s.line, p.s.col)
	}

	stmt.Address = wval

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok != EOL {
		return stmt, parseError("Expected EOL", lexeme)
	}

	return stmt, nil
}

func (p *Parser) parseAlfStatement(symbol *Symbol) (AlfStatement, error) {
	stmt := AlfStatement{symbol: symbol}

	p.swallowWhitespace()

	lexeme := p.scan()
	if lexeme.Tok != STRINGLITERAL {
		return stmt, parseError("Expected string literal", lexeme)
	}

	stmt.CharCode = lexeme.Lit

	if lexeme := p.scanIgnoreWhitespace(); lexeme.Tok != EOL {
		return stmt, parseError("Expected EOL", lexeme)
	}

	return stmt, nil
}

func (p *Parser) parseEndStatement(symbol *Symbol) (EndStatement, error) {
	stmt := EndStatement{symbol: symbol}

	p.swallowWhitespace()

	wval, err := p.parseWValue()
	if err != nil {
		return stmt, err
	}

	if wval == nil {
		return stmt, fmt.Errorf("Expected W-value (%v:%v)", p.s.line, p.s.col)
	}

	stmt.Address = wval

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
