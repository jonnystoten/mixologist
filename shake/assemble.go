package shake

import (
	"fmt"
	"log"
	"strconv"

	"jonnystoten.com/mixologist/mix"
)

func Assemble(program *Program) ([]mix.Word, error) {
	words := make([]mix.Word, 4000)
	locationCounter := 0
	for _, stmt := range program.Statements {
		switch s := stmt.(type) {
		case MixStatement:
			instruction, err := assembleMixStatement(s)
			if err != nil {
				return nil, err
			}
			word := mix.EncodeInstruction(instruction)
			words[locationCounter] = word
			log.Printf("%v: %v", locationCounter, word)
			locationCounter++
		case OrigStatement:
			address, err := strconv.Atoi(s.Address)
			if err != nil {
				return nil, err
			}
			locationCounter = address
		case ConStatement:
			address, err := strconv.Atoi(s.Address)
			if err != nil {
				return nil, err
			}
			word := mix.NewWord(address)
			words[locationCounter] = word
			log.Printf("%v: %v", locationCounter, word)
			locationCounter++
		}
	}

	return words, nil
}

func assembleMixStatement(stmt MixStatement) (mix.Instruction, error) {
	opInfo, ok := mix.OperationTable[stmt.Op]
	if !ok {
		return mix.Instruction{}, fmt.Errorf("Unknown op code: %v", stmt.Op)
	}

	instruction := mix.Instruction{OpCode: opInfo.OpCode, FieldSpec: opInfo.DefaultFS}

	var address int
	switch aPart := stmt.APart.(type) {
	case Nothing:
		address = 0
	case Number:
		address = aPart.Value
	default:
		return mix.Instruction{}, fmt.Errorf("No value for A-part at %+v", stmt)
	}

	sign := mix.Positive
	if address < 0 {
		sign = mix.Negative
		address = address * -1
	}
	instruction.Address = mix.NewAddress(sign, uint16(address))

	return instruction, nil
}
