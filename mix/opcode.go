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

var OperationTable = map[string]struct {
	OpCode    OpCode
	DefaultFS byte
}{
	"HLT": {HLT, 2},
	"LDA": {LDA, 5},
	"LD1": {LD1, 5},
	"LD2": {LD2, 5},
	"LD3": {LD3, 5},
	"LD4": {LD4, 5},
	"LD5": {LD5, 5},
	"LD6": {LD6, 5},
	"LDX": {LDX, 5},
}
