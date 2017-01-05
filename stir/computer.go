package stir

import (
	"fmt"
	"os"
	"sync"

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
	IODevices      [20]IODevice
	IOWaitGroup    *sync.WaitGroup
}

func NewComputer() *Computer {
	computer := &Computer{}
	computer.Accumulator = mix.Word{}
	computer.Extension = mix.Word{}
	computer.IODevices = [20]IODevice{}
	for i := 0; i < 8; i++ {
		filename := fmt.Sprintf("tape%v.dat", i)
		os.Create(filename)
		tu := NewTapeUnit(computer, filename)
		tu.Start()
		computer.IODevices[i] = tu
	}
	computer.IOWaitGroup = &sync.WaitGroup{}
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
	switch operation.(type) {
	case JumpOp:
	case RegisterJumpOp:
	default:
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
