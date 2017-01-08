package stir

import "jonnystoten.com/mixologist/mix"

type ShiftOp struct{ mix.Instruction }

func (op ShiftOp) Execute(c *Computer) {
	m := c.getIndexedAddressValue(op.Instruction)

	if op.FieldSpec%2 == 1 { // SRA, SRAX, SRC
		m = -m
	}

	switch op.FieldSpec {
	case 0, 1:
		bytes := c.Accumulator.Bytes[:]
		newBytes := shift(bytes, m)
		copy(c.Accumulator.Bytes[:], newBytes)
	case 2, 3:
		bytes := append(c.Accumulator.Bytes[:], c.Extension.Bytes[:]...)
		newBytes := shift(bytes, m)
		copy(c.Accumulator.Bytes[:], newBytes[:5])
		copy(c.Extension.Bytes[:], newBytes[5:])
	case 4, 5:
		bytes := append(c.Accumulator.Bytes[:], c.Extension.Bytes[:]...)
		newBytes := shiftCircle(bytes, m)
		copy(c.Accumulator.Bytes[:], newBytes[:5])
		copy(c.Extension.Bytes[:], newBytes[5:])
	}
}

func shift(bytes []byte, m int) []byte {
	size := len(bytes)
	newBytes := make([]byte, size)
	for i := 0; i < size; i++ {
		j := i + m
		if j < 0 || j >= size {
			newBytes[i] = 0
		} else {
			newBytes[i] = bytes[j]
		}
	}
	return newBytes
}

func shiftCircle(bytes []byte, m int) []byte {
	size := len(bytes)
	newBytes := make([]byte, size)
	for i := 0; i < size; i++ {
		j := i + m%size
		switch {
		case j < 0:
			newBytes[i] = bytes[j+size]
		case j >= size:
			newBytes[i] = bytes[j-size]
		default:
			newBytes[i] = bytes[j]
		}
	}
	return newBytes
}
