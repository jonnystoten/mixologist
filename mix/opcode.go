package mix

type OpCode byte

const (
	HLT  OpCode = 5
	LDA  OpCode = 8
	LD1  OpCode = 9
	LD2  OpCode = 10
	LD3  OpCode = 11
	LD4  OpCode = 12
	LD5  OpCode = 13
	LD6  OpCode = 14
	LDX  OpCode = 15
	LDAN OpCode = 16
	LD1N OpCode = 17
	LD2N OpCode = 18
	LD3N OpCode = 19
	LD4N OpCode = 20
	LD5N OpCode = 21
	LD6N OpCode = 22
	LDXN OpCode = 23
)

var OperationTable = map[string]struct {
	OpCode    OpCode
	DefaultFS byte
}{
	"HLT":  {HLT, 2},
	"LDA":  {LDA, 5},
	"LD1":  {LD1, 5},
	"LD2":  {LD2, 5},
	"LD3":  {LD3, 5},
	"LD4":  {LD4, 5},
	"LD5":  {LD5, 5},
	"LD6":  {LD6, 5},
	"LDX":  {LDX, 5},
	"LDAN": {LDAN, 5},
}
