package stir

import (
	"fmt"

	"jonnystoten.com/mixologist/mix"
)

type Operation interface {
	Execute(c *Computer)
}

func Decode(word mix.Word) Operation {
	instruction := mix.DecodeInstruction(word)
	return NewOperation(instruction)
}

func NewOperation(instruction mix.Instruction) Operation {
	switch {
	case instruction.OpCode == mix.NOP:
		return NoOp{instruction}
	case instruction.OpCode == mix.ADD:
		return AddOp{instruction}
	case instruction.OpCode == mix.SUB:
		return SubOp{instruction}
	case instruction.OpCode == mix.HLT && instruction.FieldSpec == 2:
		return HaltOp{instruction}
	case mix.LDA <= instruction.OpCode && instruction.OpCode <= mix.LDXN:
		return LoadOp{instruction}
	case mix.STA <= instruction.OpCode && instruction.OpCode <= mix.STZ:
		return StoreOp{instruction}
	default:
		panic(fmt.Sprintf("unimplemented op code %v! Full instruction: %+v", instruction.OpCode, instruction))
	}
}
