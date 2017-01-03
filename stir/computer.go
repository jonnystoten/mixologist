package stir

import "jonnystoten.com/mixologist/mix"

type Computer struct {
	Running        bool
	Accumulator    mix.Word
	Extension      mix.Word
	Index          [6]mix.Address
	JumpAddress    mix.Address
	Memory         [4000]mix.Word
	ProgramCounter int
	Overflow       bool
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
		operation := Decode(word)
		operation.Execute(c)
		c.ProgramCounter++ // TODO: will this screw up jumps?
	}
}

func (c *Computer) getIndexedAddressValue(i mix.Instruction) int {
	value := i.Address.Value()
	index := i.IndexSpec
	if index == 0 {
		return value
	}
	indexValue := c.Index[index-1].Value()
	return value + indexValue
}
