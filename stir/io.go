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

	device.ReadBlock(address)
}

type OutputOp struct{ mix.Instruction }

func (op OutputOp) Execute(c *Computer) {
	log.Println("OUT")
	address := c.getIndexedAddressValue(op.Instruction)
	device := c.IODevices[op.FieldSpec]

	device.WriteBlock(address)
}

type IOControlOp struct{ mix.Instruction }

func (op IOControlOp) Execute(c *Computer) {
	log.Println("IOC")
	address := c.getIndexedAddressValue(op.Instruction)
	device := c.IODevices[op.FieldSpec]

	device.Control(address)
}
