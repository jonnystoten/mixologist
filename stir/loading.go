package stir

import "jonnystoten.com/mixologist/mix"

type LoadOp struct{ mix.Instruction }

func (op LoadOp) Execute(c *Computer) {
	switch {
	case op.OpCode == mix.LDA:
		c.Accumulator = mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)
	case mix.LD1 <= op.OpCode && op.OpCode <= mix.LD6:
		index := op.OpCode - mix.LD1
		word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)
		c.Index[index] = mix.CastAsAddress(word)
	case op.OpCode == mix.LDX:
		c.Extension = mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)
	case op.OpCode == mix.LDAN:
		word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)
		word = mix.ToggleSign(word)
		c.Accumulator = word
	case mix.LD1N <= op.OpCode && op.OpCode <= mix.LD6N:
		index := op.OpCode - mix.LD1N
		word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)
		word = mix.ToggleSign(word)
		c.Index[index] = mix.CastAsAddress(word)
	case op.OpCode == mix.LDXN:
		word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)
		word = mix.ToggleSign(word)
		c.Extension = word
	}
}
