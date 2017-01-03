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

type IncreaseOp struct{ mix.Instruction }

func (op IncreaseOp) Execute(c *Computer) {
	value := c.getIndexedAddressValue(op.Instruction)
	word := mix.NewWord(value)

	// TODO: is there a better way to do this?
	switch {
	case op.OpCode == mix.INCA:
		acc := c.Accumulator
		sum := acc.Value() + word.Value()
		c.Accumulator = mix.NewWord(sum)
	case op.OpCode == mix.INCX:
		ext := c.Extension
		sum := ext.Value() + word.Value()
		c.Extension = mix.NewWord(sum)
	case mix.INC1 <= op.OpCode && op.OpCode <= mix.INC6:
		index := op.OpCode - mix.INC1
		i := c.Index[index]
		sum := i.Value() + word.Value()
		c.Index[index] = mix.NewAddress(sum)
	}

	// TODO: check for overflow
}

// TODO: maybe merge this with IncreaseOp
type DecreaseOp struct{ mix.Instruction }

func (op DecreaseOp) Execute(c *Computer) {
	value := c.getIndexedAddressValue(op.Instruction)
	word := mix.NewWord(value)

	// TODO: is there a better way to do this?
	switch {
	case op.OpCode == mix.DECA:
		acc := c.Accumulator
		sum := acc.Value() - word.Value()
		c.Accumulator = mix.NewWord(sum)
	case op.OpCode == mix.DECX:
		ext := c.Extension
		sum := ext.Value() - word.Value()
		c.Extension = mix.NewWord(sum)
	case mix.DEC1 <= op.OpCode && op.OpCode <= mix.DEC6:
		index := op.OpCode - mix.DEC1
		i := c.Index[index]
		sum := i.Value() - word.Value()
		c.Index[index] = mix.NewAddress(sum)
	}

	// TODO: check for overflow
}
