package mix

type Sign uint8

const (
	Positive Sign = iota
	Negative
)

type Word struct {
	Sign  Sign
	Bytes [5]byte
}

func DecodeInstruction(word *Word) *Instruction {
	address := Address{Sign: word.Sign}
	copy(address.Value[:], word.Bytes[0:1])

	return &Instruction{
		Address:   address,
		IndexSpec: word.Bytes[2],
		FieldSpec: word.Bytes[3],
		OpCode:    word.Bytes[4],
	}
}

func ApplyFieldSpec(word Word, fieldSpec byte) Word {
	newWord := Word{}
	newWord.Sign = Positive

	left, right := DecodeFieldSpec(fieldSpec)
	if left == 0 {
		newWord.Sign = word.Sign
		left = 1
	}

	bytes := []byte{}
	for i := left; i <= right; i++ {
		bytes = append(bytes, word.Bytes[i-1])
	}

	for len(bytes) < 5 {
		bytes = append([]byte{0}, bytes...)
	}

	copy(newWord.Bytes[:], bytes)

	return newWord
}
