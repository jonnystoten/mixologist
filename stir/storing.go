package stir

import "jonnystoten.com/mixologist/mix"

type StoreOp struct{ mix.Instruction }

func (op StoreOp) Execute(c *Computer) {
}
