package stir

import "jonnystoten.com/mixologist/mix"

type JumpOp struct{ mix.Instruction }

func (op JumpOp) Execute(c *Computer) {
	address := c.getIndexedAddressValue(op.Instruction)
	switch op.FieldSpec {
	case 0: // JMP
		jump(address, c)
	case 1: // JSJ
		jumpSaveJ(address, c)
	case 2: // JOV
		jumpOnOverflow(address, c, true)
	case 3: // JNOV
		jumpOnOverflow(address, c, false)
	case 4: // JL
		jumpOnComparison(address, c, mix.Less)
	case 5: // JE
		jumpOnComparison(address, c, mix.Equal)
	case 6: // JG
		jumpOnComparison(address, c, mix.Greater)
	case 7: // JGE
		jumpOnComparison(address, c, mix.Greater, mix.Equal)
	case 8: // JNE
		jumpOnComparison(address, c, mix.Less, mix.Greater)
	case 9: // JLE
		jumpOnComparison(address, c, mix.Less, mix.Equal)
	}
}

type RegisterJumpOp struct{ mix.Instruction }

func (op RegisterJumpOp) Execute(c *Computer) {
	address := c.getIndexedAddressValue(op.Instruction)
	var register mix.Word
	switch {
	case op.OpCode == mix.JAN: // TODO: use int?
		register = c.Accumulator
	case op.OpCode == mix.JXN:
		register = c.Extension
	case mix.J1N <= op.OpCode && op.OpCode <= mix.J6N:
		index := op.OpCode - mix.J1N
		register = mix.NewWordFromAddress(c.Index[index])
	}

	value := register.Value()

	switch op.FieldSpec {
	case 0: // JAN
		conditionalJump(value < 0, address, c)
	case 1: // JAZ
		conditionalJump(value == 0, address, c)
	case 2: // JAP
		conditionalJump(value > 0, address, c)
	case 3: // JANN
		conditionalJump(value >= 0, address, c)
	case 4: // JANZ
		conditionalJump(value != 0, address, c)
	case 5: // JANP
		conditionalJump(value <= 0, address, c)
	}
}

type IOJumpOp struct{ mix.Instruction }

func (op IOJumpOp) Execute(c *Computer) {
	address := c.getIndexedAddressValue(op.Instruction)
	device := c.IODevices[op.FieldSpec]

	switch op.OpCode {
	case mix.JRED:
		conditionalJump(!device.Busy(), address, c)
	case mix.JBUS:
		conditionalJump(device.Busy(), address, c)
	}
}

func jump(address int, c *Computer) {
	c.JumpAddress = mix.NewAddress(c.ProgramCounter + 1)
	c.ProgramCounter = address
}

func noJump(c *Computer) {
	c.ProgramCounter++
}

func jumpSaveJ(address int, c *Computer) {
	c.ProgramCounter = address
}

func jumpOnOverflow(address int, c *Computer, overflow bool) {
	conditionalJump(c.Overflow == overflow, address, c)
	c.Overflow = false
}

func conditionalJump(condition bool, address int, c *Computer) {
	if condition {
		jump(address, c)
	} else {
		noJump(c)
	}
}

func jumpOnComparison(address int, c *Computer, comparisons ...mix.Comparison) {
	for _, comp := range comparisons {
		if c.Comparison == comp {
			jump(address, c)
			return
		}
	}
	noJump(c)
}
