package stir

import "jonnystoten.com/mixologist/mix"

type CompareOp struct{ mix.Instruction }

func (op CompareOp) Execute(c *Computer) {
	word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)

	var register mix.Word
	switch {
	case op.OpCode == mix.CMPA:
		register = mix.ApplyFieldSpec(c.Accumulator, op.FieldSpec)
	case op.OpCode == mix.CMPX:
		register = mix.ApplyFieldSpec(c.Extension, op.FieldSpec)
	case mix.CMP1 <= op.OpCode && op.OpCode <= mix.CMP6:
		index := op.OpCode - mix.CMP1
		register = mix.ApplyFieldSpec(mix.NewWordFromAddress(c.Index[index]), op.FieldSpec)
	}

	wVal := word.Value()
	rVal := register.Value()

	switch {
	case rVal < wVal:
		c.Comparison = mix.Less
	case rVal == wVal:
		c.Comparison = mix.Equal
	case rVal > wVal:
		c.Comparison = mix.Greater
	}
}
