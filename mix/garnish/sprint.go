package garnish

import (
	"fmt"

	"jonnystoten.com/mixologist/mix"
)

func SprintWord(w mix.Word) string {
	return fmt.Sprintf("%v [%v]  |% -5v|", SprintSignedBytes(w.Sign, w.Bytes[:], w.Value()), SprintInstruction(w), mix.WordToCharCodeString(w))
}

func SprintAddress(a mix.Address) string {
	return SprintSignedBytes(a.Sign, a.Bytes[:], a.Value())
}

func SprintSignedBytes(sign mix.Sign, bytes []byte, value int) string {
	if value < 0 {
		value = -value
	}
	format := fmt.Sprintf("%%v %%v (%%0%vv)", len(bytes)*2)
	return fmt.Sprintf(format, SprintSign(sign), SprintBytes(bytes), value)
}

func SprintInstruction(w mix.Word) string {
	address := w.Bytes[:2]
	addressVal := int(address[0])*64 + int(address[1])
	return fmt.Sprintf("%v %v %v", SprintSign(w.Sign), fmt.Sprintf("%04v", addressVal), SprintBytes(w.Bytes[2:]))
}

func SprintSign(sign mix.Sign) string {
	switch sign {
	case mix.Positive:
		return "+"
	case mix.Negative:
		return "-"
	default:
		panic("invalid value for Sign")
	}
}

func SprintBytes(bytes []byte) string {
	var res string
	for _, b := range bytes {
		res += fmt.Sprintf("%02v ", b)
	}

	return res[:len(res)-1]
}
