package stir

import "jonnystoten.com/mixologist/mix"

type LoadOp struct{ mix.Instruction }

func (op LoadOp) Execute(c *Computer) {
	word := mix.ApplyFieldSpec(c.Memory[c.getIndexedAddressValue(op.Instruction)], op.FieldSpec)

	if mix.LDAN <= op.OpCode && op.OpCode <= mix.LDXN {
		word = mix.ToggleSign(word)
	}

	switch {
	case op.OpCode == mix.LDA || op.OpCode == mix.LDAN:
		c.Accumulator = word
	case op.OpCode == mix.LDX || op.OpCode == mix.LDXN:
		c.Extension = word
	case mix.LD1 <= op.OpCode && op.OpCode <= mix.LD6:
		index := op.OpCode - mix.LD1
		c.Index[index] = mix.CastAsAddress(word)
	case mix.LD1N <= op.OpCode && op.OpCode <= mix.LD6N:
		index := op.OpCode - mix.LD1N
		c.Index[index] = mix.CastAsAddress(word)
	}
}
