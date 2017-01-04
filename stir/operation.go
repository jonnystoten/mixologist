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
	case instruction.OpCode == mix.MUL:
		return MulOp{instruction}
	case instruction.OpCode == mix.DIV:
		return DivOp{instruction}
	case instruction.OpCode == mix.HLT && instruction.FieldSpec == 2:
		return HaltOp{instruction}
	case mix.LDA <= instruction.OpCode && instruction.OpCode <= mix.LDXN:
		return LoadOp{instruction}
	case mix.STA <= instruction.OpCode && instruction.OpCode <= mix.STZ:
		return StoreOp{instruction}
	case mix.ENTA <= instruction.OpCode && instruction.OpCode <= mix.ENTX && (instruction.FieldSpec == 2 || instruction.FieldSpec == 3):
		return EnterOp{instruction}
	case mix.INCA <= instruction.OpCode && instruction.OpCode <= mix.INCX && instruction.FieldSpec == 0:
		return IncreaseOp{instruction}
	case mix.DECA <= instruction.OpCode && instruction.OpCode <= mix.DECX && instruction.FieldSpec == 1:
		return DecreaseOp{instruction}
	case mix.CMPA <= instruction.OpCode && instruction.OpCode <= mix.CMPX:
		return CompareOp{instruction}
	default:
		panic(fmt.Sprintf("unimplemented op code %v! Full instruction: %+v", instruction.OpCode, instruction))
	}
}
