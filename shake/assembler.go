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
		symbol := stmt.Symbol()
		if symbol != nil {
			a.symbolTable[symbol.Name] = a.locationCounter
		}
		switch s := stmt.(type) {
		case MixStatement:
			instruction, err := a.assembleMixStatement(s)
			if err != nil {
				return nil, err
			}
			word := mix.EncodeInstruction(instruction)
			a.words[a.locationCounter] = word
			log.Printf("%v: %v", a.locationCounter, word)
			a.locationCounter++
		case OrigStatement:
			address := a.getValue(s.Address)
			a.locationCounter = address
		case ConStatement:
			address := a.getValue(s.Address)
			word := mix.NewWord(address)
			a.words[a.locationCounter] = word
			log.Printf("%v: %v", a.locationCounter, word)
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

	instruction := mix.Instruction{OpCode: opInfo.OpCode, FieldSpec: opInfo.DefaultFS}

	address := a.getValue(stmt.APart)
	sign := mix.Positive
	if address < 0 {
		sign = mix.Negative
		address = address * -1
	}
	instruction.Address = mix.NewAddress(sign, uint16(address))

	return instruction, nil
}

func (a *Assembler) addFutureRef(name string) {
	refs, ok := a.futureRefTable[name]
	if !ok {
		a.futureRefTable[name] = []int{a.locationCounter}
	} else {
		a.futureRefTable[name] = append(refs, a.locationCounter)
	}
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
	return 2
}

func (a *Assembler) visitLiteralConstant(literal LiteralConstant) int {
	_ = a.getValue(literal.Value)
	// TODO: add future ref
	return 0
}

func (a *Assembler) assembleWValue(wValue WValue) mix.Word {
	return mix.NewWord(a.visitWValue(wValue))
}
