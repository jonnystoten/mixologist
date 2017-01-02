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

func NewWord(value int) Word {
	sign := Positive
	if value < 0 {
		sign = Negative
		value *= -1
	}

	bytes := [5]byte{}
	for i := 0; i < 5; i++ {
		maxVal := pow(64, 4-i)
		b := byte(value / maxVal)
		bytes[i] = b
		value %= maxVal
	}

	return Word{Sign: sign, Bytes: bytes}
}

func ToggleSign(word Word) Word {
	if word.Sign == Positive {
		word.Sign = Negative
	} else {
		word.Sign = Positive
	}
	return word
}

func EncodeInstruction(instruction Instruction) Word {
	return Word{
		Sign: instruction.Address.Sign,
		Bytes: [5]byte{
			instruction.Address.Value[0],
			instruction.Address.Value[1],
			instruction.IndexSpec,
			instruction.FieldSpec,
			byte(instruction.OpCode),
		},
	}
}

func DecodeInstruction(word Word) Instruction {
	address := Address{Sign: word.Sign}
	copy(address.Value[:], word.Bytes[0:2])

	return Instruction{
		Address:   address,
		IndexSpec: word.Bytes[2],
		FieldSpec: word.Bytes[3],
		OpCode:    OpCode(word.Bytes[4]),
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

func NewFieldSpec(left, right byte) byte {
	return left*8 + right
}

func DecodeFieldSpec(fieldSpec byte) (left, right byte) {
	left = fieldSpec / 8
	right = fieldSpec % 8
	return
}

func CastAsAddress(word Word) Address {
	return Address{
		Sign:  word.Sign,
		Value: [2]byte{word.Bytes[3], word.Bytes[4]},
	}
}

func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
