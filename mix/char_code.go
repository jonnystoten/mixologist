package mix

import "bytes"

type CharCodeMap struct {
	runeToByteTable map[rune]byte
	byteToRuneTable map[byte]rune
}

func NewCharCodeMap() *CharCodeMap {
	ccm := &CharCodeMap{
		runeToByteTable: make(map[rune]byte),
		byteToRuneTable: make(map[byte]rune),
	}
	addCharCode(ccm, ' ', 0)
	addCharCode(ccm, 'A', 1)
	addCharCode(ccm, 'B', 2)
	addCharCode(ccm, 'C', 3)
	addCharCode(ccm, 'D', 4)
	addCharCode(ccm, 'E', 5)
	addCharCode(ccm, 'F', 6)
	addCharCode(ccm, 'G', 7)
	addCharCode(ccm, 'H', 8)
	addCharCode(ccm, 'I', 9)
	addCharCode(ccm, '∆', 10)
	addCharCode(ccm, 'J', 11)
	addCharCode(ccm, 'K', 12)
	addCharCode(ccm, 'L', 13)
	addCharCode(ccm, 'M', 14)
	addCharCode(ccm, 'N', 15)
	addCharCode(ccm, 'O', 16)
	addCharCode(ccm, 'P', 17)
	addCharCode(ccm, 'Q', 18)
	addCharCode(ccm, 'R', 19)
	addCharCode(ccm, '∑', 20)
	addCharCode(ccm, '∏', 21)
	addCharCode(ccm, 'S', 22)
	addCharCode(ccm, 'T', 23)
	addCharCode(ccm, 'U', 24)
	addCharCode(ccm, 'V', 25)
	addCharCode(ccm, 'W', 26)
	addCharCode(ccm, 'X', 27)
	addCharCode(ccm, 'Y', 28)
	addCharCode(ccm, 'Z', 29)
	addCharCode(ccm, '0', 30)
	addCharCode(ccm, '1', 31)
	addCharCode(ccm, '2', 32)
	addCharCode(ccm, '3', 33)
	addCharCode(ccm, '4', 34)
	addCharCode(ccm, '5', 35)
	addCharCode(ccm, '6', 36)
	addCharCode(ccm, '7', 37)
	addCharCode(ccm, '8', 38)
	addCharCode(ccm, '9', 39)
	addCharCode(ccm, '.', 40)
	addCharCode(ccm, ',', 41)
	addCharCode(ccm, '(', 42)
	addCharCode(ccm, ')', 43)
	addCharCode(ccm, '+', 44)
	addCharCode(ccm, '-', 45)
	addCharCode(ccm, '*', 46)
	addCharCode(ccm, '/', 47)
	addCharCode(ccm, '=', 48)
	addCharCode(ccm, '$', 49)
	addCharCode(ccm, '<', 50)
	addCharCode(ccm, '>', 51)
	addCharCode(ccm, '@', 52)
	addCharCode(ccm, ';', 53)
	addCharCode(ccm, ':', 54)
	addCharCode(ccm, '\'', 55)
	return ccm
}

var charCodes = NewCharCodeMap()

func (ccm *CharCodeMap) GetCode(r rune) byte {
	return ccm.runeToByteTable[r]
}

func (ccm *CharCodeMap) GetChar(b byte) rune {
	return ccm.byteToRuneTable[b]
}

func addCharCode(ccm *CharCodeMap, r rune, b byte) {
	ccm.runeToByteTable[r] = b
	ccm.byteToRuneTable[b] = r
}

func NewWordFromCharCode(str string) (word Word) {
	index := 0 // can't use index from range because unicode
	for _, char := range str {
		code := charCodes.GetCode(char)
		word.Bytes[index] = code
		index++
	}
	return
}

func WordToCharCodeString(word Word) string {
	buf := bytes.Buffer{}
	for _, code := range word.Bytes {
		char := charCodes.GetChar(code)
		buf.WriteRune(char)
	}

	return buf.String()
}
