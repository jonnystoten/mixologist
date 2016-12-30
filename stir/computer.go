package stir

import "jonnystoten.com/mixologist/mix"

type Computer struct {
	Running        bool
	Accumulator    mix.Word
	Extension      mix.Word
	Index          [6]mix.Address
	Memory         [4000]mix.Word
	ProgramCounter int
}

func NewComputer() *Computer {
	computer := &Computer{}
	computer.Accumulator = mix.Word{}
	computer.Extension = mix.Word{}
	return computer
}

func (c *Computer) Run() {
	c.Running = true
	for c.Running {
		word := c.Memory[c.ProgramCounter]
		instruction := mix.DecodeInstruction(word)
		c.Execute(instruction)
		c.ProgramCounter++ // TODO: will this screw up jumps?
	}
}

func (c *Computer) Execute(instruction mix.Instruction) {
	switch {
	case instruction.OpCode == mix.HLT:
		c.Running = false
	case instruction.OpCode == mix.LDA:
		c.Accumulator = mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(instruction)], instruction.FieldSpec)
	case mix.LD1 <= instruction.OpCode && instruction.OpCode <= mix.LD6:
		index := instruction.OpCode - mix.LD1
		word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(instruction)], instruction.FieldSpec)
		c.Index[index] = mix.CastAsAddress(word)
	case instruction.OpCode == mix.LDX:
		c.Extension = mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(instruction)], instruction.FieldSpec)
	case instruction.OpCode == mix.LDAN:
		word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(instruction)], instruction.FieldSpec)
		word = mix.ToggleSign(word)
		c.Accumulator = word
	case mix.LD1N <= instruction.OpCode && instruction.OpCode <= mix.LD6N:
		index := instruction.OpCode - mix.LD1N
		word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(instruction)], instruction.FieldSpec)
		word = mix.ToggleSign(word)
		c.Index[index] = mix.CastAsAddress(word)
	case instruction.OpCode == mix.LDXN:
		word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(instruction)], instruction.FieldSpec)
		word = mix.ToggleSign(word)
		c.Extension = word
	}
}

func (c *Computer) getIndexedAddressValue(i mix.Instruction) uint16 {
	value := i.Address.GetValue()
	index := i.IndexSpec
	indexValue := c.Index[index].GetValue()
	return value + indexValue
}
