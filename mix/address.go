package mix

type Address struct {
	Sign  Sign
	Value [2]byte
}

func NewAddress(value int) Address {
	sign := Positive
	if value < 0 {
		sign = Negative
		value = value * -1
	}
	return Address{Sign: sign, Value: [2]byte{byte(value / 64), byte(value % 64)}}
}

func (a *Address) GetValue() int {
	return int(a.Value[0])*64 + int(a.Value[1])
}
