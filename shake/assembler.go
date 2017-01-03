package shake

import (
	"fmt"
	"log"

	"jonnystoten.com/mixologist/mix"
)

type Assembler struct {
	words           []mix.Word
	locationCounter int
	symbolTable     map[string]int
	futureRefTable  map[string][]int
}

func NewAssembler() *Assembler {
	return &Assembler{
		words:          make([]mix.Word, 4000),
		symbolTable:    make(map[string]int),
		futureRefTable: make(map[string][]int),
	}
}

func (a *Assembler) Assemble(program *Program) ([]mix.Word, error) {
	for _, stmt := range program.Statements {
		switch stmt := stmt.(type) {
		case MixStatement:
			a.addSymbol(stmt, a.locationCounter)
			instruction, err := a.assembleMixStatement(stmt)
			if err != nil {
				return nil, err
			}
			word := mix.EncodeInstruction(instruction)
			a.words[a.locationCounter] = word
			a.locationCounter++
		case OrigStatement:
			a.addSymbol(stmt, a.locationCounter)
			address := a.getValue(stmt.Address)
			log.Printf("address: %v", address)
			a.locationCounter = address
		case EquStatement:
			address := a.getValue(stmt.Address)
			a.addSymbol(stmt, address)
		case ConStatement:
			a.addSymbol(stmt, a.locationCounter)
			address := a.getValue(stmt.Address)
			word := mix.NewWord(address)
			a.words[a.locationCounter] = word
			a.locationCounter++
		case AlfStatement:
			word := a.assembleAlfStatement(stmt)
			a.words[a.locationCounter] = word
			a.locationCounter++
		}
	}

	return a.words, nil
}

func (a *Assembler) assembleMixStatement(stmt MixStatement) (mix.Instruction, error) {
	opInfo, ok := mix.OperationTable[stmt.Op]
	if !ok {
		return mix.Instruction{}, fmt.Errorf("Unknown op code: %v", stmt.Op)
	}

	instruction := mix.Instruction{OpCode: opInfo.OpCode}

	address := a.getValue(stmt.APart)
	instruction.Address = mix.NewAddress(address)

	switch stmt.FPart.(type) {
	case Nothing:
		instruction.FieldSpec = opInfo.DefaultFS
	default:
		instruction.FieldSpec = byte(a.getValue(stmt.FPart))
	}

	return instruction, nil
}

func (a *Assembler) addSymbol(stmt Statement, value int) {
	symbol := stmt.Symbol()
	if symbol != nil {
		a.symbolTable[symbol.Name] = value
		a.fixupFutureRefs(symbol.Name)
	}
}

func (a *Assembler) addFutureRef(name string) {
	refs, ok := a.futureRefTable[name]
	if !ok {
		a.futureRefTable[name] = []int{a.locationCounter}
	} else {
		a.futureRefTable[name] = append(refs, a.locationCounter)
	}
}

func (a *Assembler) fixupFutureRefs(name string) {
	log.Printf("Fixing up future refs to '%v'", name)
	refs, ok := a.futureRefTable[name]
	if !ok {
		return
	}

	target := a.symbolTable[name]

	for _, ref := range refs {
		log.Printf("Got a ref to '%v' at %v", name, ref)
		address := mix.NewAddress(target)
		a.words[ref].Sign = address.Sign
		a.words[ref].Bytes[0] = address.Value[0]
		a.words[ref].Bytes[1] = address.Value[1]
	}

	delete(a.futureRefTable, name)
}

func (a *Assembler) getValue(node Node) int {
	return node.Accept(a)
}

func (a *Assembler) visitNothing(nothing Nothing) int {
	return 0
}

func (a *Assembler) visitNumber(number Number) int {
	return number.Value
}

func (a *Assembler) visitSymbol(symbol Symbol) int {
	value, ok := a.symbolTable[symbol.Name]
	if !ok {
		log.Printf("Adding future ref to '%v'", symbol.Name)
		a.addFutureRef(symbol.Name)
		return 0
	}
	return value
}

func (a *Assembler) visitAsterisk(asterisk Asterisk) int {
	return a.locationCounter
}

func (a *Assembler) visitExpression(expression Expression) int {
	var left int
	if expression.Left != nil {
		left = a.getValue(*expression.Left)
	} else {
		left = 0
	}

	right := a.getValue(expression.Right)
	switch expression.Operator {
	case PLUS:
		return left + right
	case MINUS:
		return left - right
	case ASTERISK:
		return left * right
	case DIVIDE:
		return left / right
	case SHIFTDIVIDE:
		panic("unsupported // operator!") // TODO: how to do this?
	case FIELDSIGN:
		return 8*left + right
	default:
		panic("unknown binary operator!")
	}
}

func (a *Assembler) visitWValue(wVal WValue) int {
	return a.getValue(wVal.Parts[0].Exp) // TODO: make this work properly
}

func (a *Assembler) visitLiteralConstant(literal LiteralConstant) int {
	_ = a.getValue(literal.Value)
	// TODO: add future ref
	return 0
}

func (a *Assembler) assembleAlfStatement(stmt AlfStatement) mix.Word {
	word := mix.Word{}

	charcode := stmt.CharCode
	inner := charcode[1 : len(charcode)-1]

	index := 0 // can't use index from range because unicode
	for _, char := range inner {
		log.Printf("%v: %v", index, char)
		code := mix.CharCodeTable[char]
		word.Bytes[index] = code
		index++
	}

	return word
}

func (a *Assembler) assembleWValue(wValue WValue) mix.Word {
	return mix.NewWord(a.visitWValue(wValue))
}
