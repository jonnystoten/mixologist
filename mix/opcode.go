package mix

type OpCode byte

const (
	HLT OpCode = 5
	LDA        = 8
	LDX        = 15
)

var OpCodeTable = map[string]OpCode{
	"HLT": HLT,
	"LDA": LDA,
	"LDX": LDX,
}
