package mix

type Address struct {
	Sign  Sign
	Value [2]byte
}

func NewAddress(sign Sign, value uint16) Address {
	return Address{Sign: sign, Value: [2]byte{byte(value / 64), byte(value % 64)}}
}

func (a *Address) GetValue() uint16 {
	return uint16(a.Value[0])*64 + uint16(a.Value[1])
}

type Instruction struct {
	Address   Address
	OpCode    OpCode
	FieldSpec byte
	IndexSpec byte
}

func NewFieldSpec(left, right byte) byte {
	return left*8 + right
}

func DecodeFieldSpec(fieldSpec byte) (left, right byte) {
	left = fieldSpec / 8
	right = fieldSpec % 8
	return
}
