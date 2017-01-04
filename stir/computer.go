package stir

import (
	"fmt"

	"jonnystoten.com/mixologist/mix"
)

type Computer struct {
	Running        bool
	Accumulator    mix.Word
	Extension      mix.Word
	Index          [6]mix.Address
	JumpAddress    mix.Address
	Memory         [4000]mix.Word
	ProgramCounter int
	Overflow       bool
	Comparison     mix.Comparison
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
		c.FetchDecodeExecute()
	}
}

func (c *Computer) FetchDecodeExecute() {
	word := c.Memory[c.ProgramCounter]
	operation := Decode(word)
	operation.Execute(c)
	if _, isJump := operation.(JumpOp); !isJump {
		c.ProgramCounter++
	}
}

func (c *Computer) getIndexedAddressValue(i mix.Instruction) int {
	index := i.IndexSpec
	if index > 6 {
		panic(fmt.Sprintf("index spec out of range: %v", index))
	}

	value := i.Address.Value()
	if index == 0 {
		return value
	}
	indexValue := c.Index[index-1].Value()
	return value + indexValue
}
