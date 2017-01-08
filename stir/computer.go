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
	Index          []mix.Address
	JumpAddress    mix.Address
	Memory         []mix.Word
	ProgramCounter int
	Overflow       bool
	Comparison     mix.Comparison
	IODevices      []IODevice
	IOWaitGroup    *sync.WaitGroup
}

func NewComputer() *Computer {
	computer := &Computer{}
	computer.Accumulator = mix.Word{}
	computer.Extension = mix.Word{}
	computer.Index = make([]mix.Address, 6)
	computer.Memory = make([]mix.Word, 4000)
	setupIODevices(computer)
	return computer
}

func setupIODevices(computer *Computer) {
	ioDir := fmt.Sprintf("%v/.stir", os.Getenv("HOME"))
	os.Mkdir(ioDir, 0755)
	computer.IODevices = make([]IODevice, 19)
	for i := 0; i < 8; i++ {
		filename := fmt.Sprintf("%v/tape%v.dat", ioDir, i)
		//os.Create(filename)
		tu := NewTapeUnit(computer, filename)
		computer.IODevices[i] = tu
	}
	for i := 8; i < 16; i++ {
		filename := fmt.Sprintf("%v/disk%v.dat", ioDir, i)
		//os.Create(filename)
		dd := NewDiskDrumUnit(computer, filename)
		computer.IODevices[i] = dd
	}
	//os.Create(fmt.Sprintf("%v/cardreader.dat", ioDir))
	cr := NewCardReader(computer, fmt.Sprintf("%v/cardreader.dat", ioDir))
	computer.IODevices[16] = cr
	os.Create(fmt.Sprintf("%v/cardwriter.dat", ioDir))
	cw := NewCardWriter(computer, fmt.Sprintf("%v/cardwriter.dat", ioDir))
	computer.IODevices[17] = cw
	os.Create(fmt.Sprintf("%v/lineprinter.dat", ioDir))
	lp := NewLinePrinter(computer, fmt.Sprintf("%v/lineprinter.dat", ioDir))
	computer.IODevices[18] = lp

	computer.IOWaitGroup = &sync.WaitGroup{}
	for _, device := range computer.IODevices {
		device.Start()
	}
}

func (c *Computer) Run() {
	c.RunInteractive(func() {})
}

func (c *Computer) RunInteractive(callback func()) {
	c.Running = true
	for c.Running {
		callback()
		c.FetchDecodeExecute()
		if c.ProgramCounter >= len(c.Memory) {
			c.Running = false
		}
	}
}

func (c *Computer) FetchDecodeExecute() {
	word := c.Memory[c.ProgramCounter]
	operation := Decode(word)
	operation.Execute(c)
	switch operation.(type) {
	case JumpOp:
	case RegisterJumpOp:
	case IOJumpOp:
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
