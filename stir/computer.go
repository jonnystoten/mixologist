package stir

import "jonnystoten.com/mixologist/mix"

type Computer struct {
	Accumulator mix.Word
	Extension   mix.Word
	Index       [6]mix.Address
	Memory      [4000]mix.Word
}

func NewComputer() *Computer {
	computer := &Computer{}
	computer.Accumulator = mix.Word{}
	computer.Extension = mix.Word{}
	return computer
}

func (c *Computer) Execute(instruction *mix.Instruction) {
	switch {
	case instruction.OpCode == mix.LDA:
		c.Accumulator = mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(instruction)], instruction.FieldSpec)
	case instruction.OpCode >= mix.LD1 && instruction.OpCode <= mix.LD6:
		index := instruction.OpCode - mix.LD1
		word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(instruction)], instruction.FieldSpec)
		c.Index[index] = mix.CastAsAddress(word)
	case instruction.OpCode == mix.LDX:
		c.Extension = mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(instruction)], instruction.FieldSpec)
	}
}

func (c *Computer) getIndexedAddressValue(i *mix.Instruction) uint16 {
	value := i.Address.GetValue()
	index := i.IndexSpec
	indexValue := c.Index[index].GetValue()
	return value + indexValue
}
