package shake

import (
	"fmt"
	"strconv"

	"jonnystoten.com/mixologist/mix"
)

func Assemble(program *Program) ([]mix.Instruction, error) {
	var instructions []mix.Instruction
	for _, stmt := range program.Statements {
		switch s := stmt.(type) {
		case MixStatement:
			instruction, err := assembleMixStatement(s)
			if err != nil {
				return nil, err
			}
			instructions = append(instructions, instruction)
		}
	}

	return instructions, nil
}

func assembleMixStatement(stmt MixStatement) (mix.Instruction, error) {
	opCode, ok := mix.OpCodeTable[stmt.Op]
	if !ok {
		return mix.Instruction{}, fmt.Errorf("Unknown op code: %v", stmt.Op)
	}

	instruction := mix.Instruction{OpCode: opCode}

	if stmt.Address != "" {
		address, err := strconv.Atoi(stmt.Address)
		if err != nil {
			return mix.Instruction{}, err
		}

		sign := mix.Positive
		if address < 0 {
			sign = mix.Negative
			address = address * -1
		}
		instruction.Address = mix.NewAddress(sign, uint16(address))
	}

	return instruction, nil
}
