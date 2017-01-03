package stir

import "jonnystoten.com/mixologist/mix"

type EnterOp struct{ mix.Instruction }

func (op EnterOp) Execute(c *Computer) {
	value := c.getIndexedAddressValue(op.Instruction)
	word := mix.NewWord(value)
	if value == 0 {
		word.Sign = op.Address.Sign
	}

	if op.FieldSpec == 3 { // enter negative
		word = mix.ToggleSign(word)
	}

	switch {
	case op.OpCode == mix.ENTA || op.OpCode == mix.ENNA:
		c.Accumulator = word
	case op.OpCode == mix.ENTX || op.OpCode == mix.ENNX:
		c.Extension = word
	case mix.ENT1 <= op.OpCode && op.OpCode <= mix.ENT6:
		index := op.OpCode - mix.ENT1
		c.Index[index] = mix.CastAsAddress(word)
	case mix.ENN1 <= op.OpCode && op.OpCode <= mix.ENN6:
		index := op.OpCode - mix.ENN1
		c.Index[index] = mix.CastAsAddress(word)
	}
}
