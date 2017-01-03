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

func NewWordFromAddress(address Address) Word {
	return Word{
		Sign:  address.Sign,
		Bytes: [5]byte{0, 0, 0, address.Bytes[0], address.Bytes[1]},
	}
}

func (w *Word) Value() (value int) {
	for i := 0; i < 5; i++ {
		base := pow(64, i)
		value += base * int(w.Bytes[4-i])
	}
	if w.Sign == Negative {
		value *= -1
	}
	return
}

func ToggleSign(word Word) Word {
	if word.Sign == Positive {
		word.Sign = Negative
	} else {
		word.Sign = Positive
	}
	return word
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
