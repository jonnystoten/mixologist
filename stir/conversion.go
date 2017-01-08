package stir

import "jonnystoten.com/mixologist/mix"

type NumOp struct{ mix.Instruction }

func (op NumOp) Execute(c *Computer) {
	var result int
	for i := 0; i < 5; i++ {
		accB := int(c.Accumulator.Bytes[5-1-i])
		extB := int(c.Extension.Bytes[5-1-i])

		result += (accB % 10) * mix.Pow(10, 5+i)
		result += (extB % 10) * mix.Pow(10, i)
	}

	//log.Printf("NUM result = %v", result)

	sign := c.Accumulator.Sign
	if mix.FitsInWord(result) {
		c.Accumulator = mix.NewWord(result)
	} else {
		c.Overflow = true
		c.Accumulator = mix.NewWordWithOverflow(result)
	}
	c.Accumulator.Sign = sign
}

type CharOp struct{ mix.Instruction }

func (op CharOp) Execute(c *Computer) {
	value := c.Accumulator.Value()
	if value < 0 {
		value = -value
	}

	for i := 0; i < 10; i++ {
		b := byte(value%10 + 30)
		value /= 10
		if i < 5 {
			c.Extension.Bytes[5-i-1] = b
		} else {
			c.Accumulator.Bytes[10-i-1] = b
		}
	}
}
