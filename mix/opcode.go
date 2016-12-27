package mix

type OpCode byte

const (
	HLT OpCode = 5
	LDA        = 8
	LD1        = 9
	LD2        = 10
	LD3        = 11
	LD4        = 12
	LD5        = 13
	LD6        = 14
	LDX        = 15
)

var OpCodeTable = map[string]OpCode{
	"HLT": HLT,
	"LDA": LDA,
	"LD1": LD1,
	"LD2": LD2,
	"LD3": LD3,
	"LD4": LD4,
	"LD5": LD5,
	"LD6": LD6,
	"LDX": LDX,
}
