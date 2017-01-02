package mix

type Instruction struct {
	Address   Address
	OpCode    OpCode
	FieldSpec byte
	IndexSpec byte
}
