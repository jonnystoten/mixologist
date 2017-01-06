package shake

import (
	"fmt"

	"jonnystoten.com/mixologist/mix"
)

type Assembler struct {
	Words           []mix.Word
	ProgramStart    int
	locationCounter int
	symbolTable     map[string]int
	futureRefTable  map[string][]int
}

func NewAssembler() *Assembler {
	return &Assembler{
		Words:          make([]mix.Word, 4000),
		symbolTable:    make(map[string]int),
		futureRefTable: make(map[string][]int),
	}
}

func (a *Assembler) Assemble(program *Program) error {
	for _, stmt := range program.Statements {
		switch stmt := stmt.(type) {
		case MixStatement:
			a.dealWithNormalSymbolDecl(stmt, a.locationCounter)
			instruction, err := a.assembleMixStatement(stmt)
			if err != nil {
				return err
			}
			word := mix.EncodeInstruction(instruction)
			a.Words[a.locationCounter] = word
			a.dealWithLocalSymbolDecl(stmt, a.locationCounter)
			a.locationCounter++
		case OrigStatement:
			a.dealWithNormalSymbolDecl(stmt, a.locationCounter)
			address := a.getValue(stmt.Address)
			a.dealWithLocalSymbolDecl(stmt, a.locationCounter)
			a.locationCounter = address
		case EquStatement:
			address := a.getValue(stmt.Address)
			a.dealWithNormalSymbolDecl(stmt, address) // TODO: need to add a
			a.dealWithLocalSymbolDecl(stmt, address)  // special case for locals?
		case ConStatement:
			a.dealWithNormalSymbolDecl(stmt, a.locationCounter)
			address := a.getValue(stmt.Address)
			word := mix.NewWord(address)
			a.Words[a.locationCounter] = word
			a.dealWithLocalSymbolDecl(stmt, a.locationCounter)
			a.locationCounter++
		case AlfStatement:
			a.dealWithNormalSymbolDecl(stmt, a.locationCounter)
			word := a.assembleAlfStatement(stmt)
			a.Words[a.locationCounter] = word
			a.dealWithLocalSymbolDecl(stmt, a.locationCounter)
			a.locationCounter++
		case EndStatement:
			a.dealWithNormalSymbolDecl(stmt, a.locationCounter)
			address := a.getValue(stmt.Address)
			a.ProgramStart = address
			a.dealWithLocalSymbolDecl(stmt, a.locationCounter)
			break
		}
	}

	return nil
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

	switch stmt.IndexPart.(type) {
	case Nothing:
		instruction.IndexSpec = 0
	default:
		instruction.IndexSpec = byte(a.getValue(stmt.IndexPart))
	}

	return instruction, nil
}

func (a *Assembler) dealWithNormalSymbolDecl(stmt Statement, value int) {
	if symbol := stmt.Symbol(); symbol != nil && !symbol.IsLocal() {
		a.symbolTable[symbol.InternalName()] = value
		a.fixupFutureRefs(symbol.InternalName(), false)
	}
}

func (a *Assembler) dealWithLocalSymbolDecl(stmt Statement, value int) {
	if symbol := stmt.Symbol(); symbol != nil && symbol.IsLocalDecl() {
		a.symbolTable[symbol.InternalName()] = value
		a.fixupFutureRefs(symbol.InternalName(), true)
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

func (a *Assembler) fixupFutureRefs(name string, notCurrent bool) {
	refs, ok := a.futureRefTable[name]
	if !ok {
		return
	}

	target := a.symbolTable[name]

	var keep bool
	for _, ref := range refs {
		if notCurrent && ref == a.locationCounter {
			keep = true
			continue
		}
		address := mix.NewAddress(target)
		a.Words[ref].Sign = address.Sign
		a.Words[ref].Bytes[0] = address.Bytes[0]
		a.Words[ref].Bytes[1] = address.Bytes[1]
	}

	if !keep {
		delete(a.futureRefTable, name)
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
	value, existing := a.symbolTable[symbol.InternalName()]
	if !existing || symbol.IsLocalForwardRef() {
		a.addFutureRef(symbol.InternalName())
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
	charcode := stmt.CharCode
	inner := charcode[1 : len(charcode)-1]
	return mix.NewWordFromCharCode(inner)
}

func (a *Assembler) assembleWValue(wValue WValue) mix.Word {
	return mix.NewWord(a.visitWValue(wValue))
}
