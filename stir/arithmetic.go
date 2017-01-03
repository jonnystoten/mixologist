package stir

import "jonnystoten.com/mixologist/mix"

type AddOp struct{ mix.Instruction }

func (op AddOp) Execute(c *Computer) {
	word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)
	acc := c.Accumulator

	sum := acc.Value() + word.Value()

	// TODO: check for overflow
	result := mix.NewWord(sum)
	c.Accumulator = result
}

type SubOp struct{ mix.Instruction }

func (op SubOp) Execute(c *Computer) {
	word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)
	acc := c.Accumulator

	sum := acc.Value() - word.Value()

	// TODO: check for overflow
	result := mix.NewWord(sum)
	c.Accumulator = result
}

type MulOp struct{ mix.Instruction }

func (op MulOp) Execute(c *Computer) {
	word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)
	acc := c.Accumulator

	var sign mix.Sign
	if word.Sign == acc.Sign {
		sign = mix.Positive
	} else {
		sign = mix.Negative
	}

	sum := word.Value() * acc.Value()
	accResult := sum / 1073741824 // TODO: base on byte size
	extResult := sum % 1073741824

	c.Accumulator = mix.NewWord(accResult)
	c.Accumulator.Sign = sign
	c.Extension = mix.NewWord(extResult)
	c.Extension.Sign = sign
}
