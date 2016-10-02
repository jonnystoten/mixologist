package stir

import "jonnystoten.com/mixologist/mix"

type Computer struct {
	Accumulator mix.Word
	Extension   mix.Word
	Memory      [4000]mix.Word
}

func NewComputer() *Computer {
	computer := &Computer{}
	computer.Accumulator = mix.Word{}
	computer.Extension = mix.Word{}
	return computer
}

func (c *Computer) Execute(instruction *mix.Instruction) {
	switch instruction.OpCode {
	case mix.LDA:
		c.Accumulator = mix.ApplyFieldSpec(c.Memory[instruction.Address.GetValue()], instruction.FieldSpec)
	case mix.LDX:
		c.Extension = mix.ApplyFieldSpec(c.Memory[instruction.Address.GetValue()], instruction.FieldSpec)
	}
}
