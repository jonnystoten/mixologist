package stir

import "jonnystoten.com/mixologist/mix"

type InputOutputOp struct{ mix.Instruction }

func (op InputOutputOp) Execute(c *Computer) {
	address := c.getIndexedAddressValue(op.Instruction)
	device := c.IODevices[op.FieldSpec]

	device.WaitReady()
	device.SetBusy()
	device.Channel() <- ioMessage{op.OpCode, address}
}
