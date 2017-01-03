package mix

type Instruction struct {
	Address   Address
	OpCode    OpCode
	FieldSpec byte
	IndexSpec byte
}

func EncodeInstruction(instruction Instruction) Word {
	return Word{
		Sign: instruction.Address.Sign,
		Bytes: [5]byte{
			instruction.Address.Bytes[0],
			instruction.Address.Bytes[1],
			instruction.IndexSpec,
			instruction.FieldSpec,
			byte(instruction.OpCode),
		},
	}
}

func DecodeInstruction(word Word) Instruction {
	address := Address{Sign: word.Sign}
	copy(address.Bytes[:], word.Bytes[0:2])

	return Instruction{
		Address:   address,
		IndexSpec: word.Bytes[2],
		FieldSpec: word.Bytes[3],
		OpCode:    OpCode(word.Bytes[4]),
	}
}
