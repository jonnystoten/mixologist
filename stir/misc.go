package stir

import "jonnystoten.com/mixologist/mix"

type NoOp struct{ mix.Instruction }

func (op NoOp) Execute(c *Computer) {
}
