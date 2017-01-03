package mix

type Address struct {
	Sign  Sign
	Bytes [2]byte
}

func NewAddress(value int) Address {
	sign := Positive
	if value < 0 {
		sign = Negative
		value = value * -1
	}
	return Address{Sign: sign, Bytes: [2]byte{byte(value / 64), byte(value % 64)}}
}

func (a *Address) Value() (value int) {
	value = int(a.Bytes[0])*64 + int(a.Bytes[1])
	if a.Sign == Negative {
		value *= -1
	}
	return
}

func CastAsAddress(word Word) Address {
	return Address{
		Sign:  word.Sign,
		Bytes: [2]byte{word.Bytes[3], word.Bytes[4]},
	}
}
