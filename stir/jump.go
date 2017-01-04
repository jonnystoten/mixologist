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

func jump(address int, c *Computer) {
	c.JumpAddress = mix.NewAddress(c.ProgramCounter)
	c.ProgramCounter = address
}

func jumpSaveJ(address int, c *Computer) {
	c.ProgramCounter = address
}

func jumpOnOverflow(address int, c *Computer, overflow bool) {
	if c.Overflow == overflow {
		jump(address, c)
	}
	c.Overflow = false
}

func jumpOnComparison(address int, c *Computer, comparisons ...mix.Comparison) {
	for _, comp := range comparisons {
		if c.Comparison == comp {
			jump(address, c)
			return
		}
	}
}
