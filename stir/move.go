package stir

import "jonnystoten.com/mixologist/mix"

type MoveOp struct{ mix.Instruction }

func (op MoveOp) Execute(c *Computer) {
	src := c.getIndexedAddressValue(op.Instruction)
	dest := c.Index[0].Value()
	num := int(op.FieldSpec)

	for i := 0; i < num; i++ {
		c.Memory[dest+i] = c.Memory[src+i]
	}

	c.Index[0] = mix.NewAddress(num)
}
