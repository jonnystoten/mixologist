package stir

import "jonnystoten.com/mixologist/mix"

type InputOutputOp struct{ mix.Instruction }

func (op InputOutputOp) Execute(c *Computer) {
	address := c.getIndexedAddressValue(op.Instruction)
	device := c.IODevices[op.FieldSpec]

	c.IOWaitGroup.Add(1)
	device.Channel() <- ioMessage{op.OpCode, address}
}
