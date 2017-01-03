package stir

import "jonnystoten.com/mixologist/mix"

type HaltOp struct{ mix.Instruction }

func (op HaltOp) Execute(c *Computer) {
	c.Running = false
}
