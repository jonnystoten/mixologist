package stir

import "jonnystoten.com/mixologist/mix"

type StoreOp struct{ mix.Instruction }

func (op StoreOp) Execute(c *Computer) {
	var register mix.Word
	switch {
	case op.OpCode == mix.STA:
		register = c.Accumulator
	case op.OpCode == mix.STX:
		register = c.Extension
	case mix.ST1 <= op.OpCode && op.OpCode <= mix.ST6:
		index := op.OpCode - mix.ST1
		register = mix.NewWordFromAddress(c.Index[index])
	case op.OpCode == mix.STJ:
		register = mix.NewWordFromAddress(c.JumpAddress)
	case op.OpCode == mix.STZ:
		register = mix.Word{}
	}

	address := c.getIndexedAddressValue(op.Instruction)
	fsLeft, fsRight := mix.DecodeFieldSpec(op.FieldSpec)
	numBytes := int(fsRight-fsLeft) + 1
	if numBytes > 0 && fsLeft == 0 {
		numBytes--
	}

	bytes := GetBytesToStore(register, numBytes)
	word := c.Memory[address]
	if fsLeft == 0 {
		word.Sign = register.Sign
		fsLeft++
	}
	for i := 0; i < numBytes; i++ {
		word.Bytes[int(fsLeft)+i-1] = bytes[i]
	}

	c.Memory[address] = word
}

func GetBytesToStore(register mix.Word, count int) []byte {
	offset := 5 - count
	bytes := make([]byte, count)
	for index := 0; index < count; index++ {
		bytes[index] = register.Bytes[index+offset]
	}
	return bytes
}
