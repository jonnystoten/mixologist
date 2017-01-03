package stir

import "jonnystoten.com/mixologist/mix"

type AddOp struct{ mix.Instruction }

func (op AddOp) Execute(c *Computer) {
	word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)
	acc := c.Accumulator

	sum := word.Value() + acc.Value()

	// TODO: check for overflow
	result := mix.NewWord(sum)
	c.Accumulator = result
}

type SubOp struct{ mix.Instruction }

func (op SubOp) Execute(c *Computer) {
	word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)
	acc := c.Accumulator

	sum := word.Value() - acc.Value()

	// TODO: check for overflow
	result := mix.NewWord(sum)
	c.Accumulator = result
}
