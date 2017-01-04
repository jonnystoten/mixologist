package stir

import (
	"log"

	"jonnystoten.com/mixologist/mix"
)

type InputOp struct{ mix.Instruction }

func (op InputOp) Execute(c *Computer) {
	log.Println("IN")
	address := c.getIndexedAddressValue(op.Instruction)
	device := c.IODevices[op.FieldSpec]

	c.IOWaitGroup.Add(1)
	device.Lock()
	log.Println("IN after lock")
	go func() {
		err := device.ReadBlock(address, c)
		if err != nil {
			panic("IN " + err.Error())
		}
		c.IOWaitGroup.Done()
	}()
}

type OutputOp struct{ mix.Instruction }

func (op OutputOp) Execute(c *Computer) {
	log.Println("OUT")
	address := c.getIndexedAddressValue(op.Instruction)
	device := c.IODevices[op.FieldSpec]

	c.IOWaitGroup.Add(1)
	device.Lock()
	log.Println("OUT after lock")
	go func() {
		err := device.WriteBlock(address, c)
		if err != nil {
			panic("OUT " + err.Error())
		}
		c.IOWaitGroup.Done()
	}()
}
