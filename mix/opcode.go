package mix

type OpCode byte

const (
	NOP  OpCode = 0
	ADD  OpCode = 1
	SUB  OpCode = 2
	MUL  OpCode = 3
	DIV  OpCode = 4
	NUM  OpCode = 5
	CHAR OpCode = 5
	HLT  OpCode = 5
	SLA  OpCode = 6
	SRA  OpCode = 6
	SLAX OpCode = 6
	SRAX OpCode = 6
	SLC  OpCode = 6
	SRC  OpCode = 6
	MOVE OpCode = 7
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
	STA  OpCode = 24
	ST1  OpCode = 25
	ST2  OpCode = 26
	ST3  OpCode = 27
	ST4  OpCode = 28
	ST5  OpCode = 29
	ST6  OpCode = 30
	STX  OpCode = 31
	STJ  OpCode = 32
	STZ  OpCode = 33
	JBUS OpCode = 34
	IOC  OpCode = 35
	IN   OpCode = 36
	OUT  OpCode = 37
	JRED OpCode = 38
	JMP  OpCode = 39
	JSJ  OpCode = 39
	JOV  OpCode = 39
	JNOV OpCode = 39
	JL   OpCode = 39
	JE   OpCode = 39
	JG   OpCode = 39
	JGE  OpCode = 39
	JNE  OpCode = 39
	JLE  OpCode = 39
	JAN  OpCode = 40
	JAZ  OpCode = 40
	JAP  OpCode = 40
	JANN OpCode = 40
	JANZ OpCode = 40
	JANP OpCode = 40
	J1N  OpCode = 41
	J1Z  OpCode = 41
	J1P  OpCode = 41
	J1NN OpCode = 41
	J1NZ OpCode = 41
	J1NP OpCode = 41
	J2N  OpCode = 42
	J2Z  OpCode = 42
	J2P  OpCode = 42
	J2NN OpCode = 42
	J2NZ OpCode = 42
	J2NP OpCode = 42
	J3N  OpCode = 43
	J3Z  OpCode = 43
	J3P  OpCode = 43
	J3NN OpCode = 43
	J3NZ OpCode = 43
	J3NP OpCode = 43
	J4N  OpCode = 44
	J4Z  OpCode = 44
	J4P  OpCode = 44
	J4NN OpCode = 44
	J4NZ OpCode = 44
	J4NP OpCode = 44
	J5N  OpCode = 45
	J5Z  OpCode = 45
	J5P  OpCode = 45
	J5NN OpCode = 45
	J5NZ OpCode = 45
	J5NP OpCode = 45
	J6N  OpCode = 46
	J6Z  OpCode = 46
	J6P  OpCode = 46
	J6NN OpCode = 46
	J6NZ OpCode = 46
	J6NP OpCode = 46
	JXN  OpCode = 47
	JXZ  OpCode = 47
	JXP  OpCode = 47
	JXNN OpCode = 47
	JXNZ OpCode = 47
	JXNP OpCode = 47
	INCA OpCode = 48
	DECA OpCode = 48
	ENTA OpCode = 48
	ENNA OpCode = 48
	INC1 OpCode = 49
	DEC1 OpCode = 49
	ENT1 OpCode = 49
	ENN1 OpCode = 49
	INC2 OpCode = 50
	DEC2 OpCode = 50
	ENT2 OpCode = 50
	ENN2 OpCode = 50
	INC3 OpCode = 51
	DEC3 OpCode = 51
	ENT3 OpCode = 51
	ENN3 OpCode = 51
	INC4 OpCode = 52
	DEC4 OpCode = 52
	ENT4 OpCode = 52
	ENN4 OpCode = 52
	INC5 OpCode = 53
	DEC5 OpCode = 53
	ENT5 OpCode = 53
	ENN5 OpCode = 53
	INC6 OpCode = 54
	DEC6 OpCode = 54
	ENT6 OpCode = 54
	ENN6 OpCode = 54
	INCX OpCode = 55
	DECX OpCode = 55
	ENTX OpCode = 55
	ENNX OpCode = 55
	CMPA OpCode = 56
	CMP1 OpCode = 57
	CMP2 OpCode = 58
	CMP3 OpCode = 59
	CMP4 OpCode = 60
	CMP5 OpCode = 61
	CMP6 OpCode = 62
	CMPX OpCode = 63
)

var OperationTable = map[string]struct {
	OpCode    OpCode
	DefaultFS byte
}{
	"NOP":  {NOP, 0},
	"ADD":  {ADD, 5},
	"SUB":  {SUB, 5},
	"MUL":  {MUL, 5},
	"DIV":  {DIV, 5},
	"NUM":  {NUM, 5},
	"CHAR": {CHAR, 5},
	"HLT":  {HLT, 2},
	"SLA":  {SLA, 0},
	"SRA":  {SRA, 1},
	"SLAX": {SLAX, 2},
	"SRAX": {SRAX, 3},
	"SLC":  {SLC, 4},
	"SRC":  {SRC, 5},
	"MOVE": {MOVE, 5},
	"LDA":  {LDA, 5},
	"LD1":  {LD1, 5},
	"LD2":  {LD2, 5},
	"LD3":  {LD3, 5},
	"LD4":  {LD4, 5},
	"LD5":  {LD5, 5},
	"LD6":  {LD6, 5},
	"LDX":  {LDX, 5},
	"LDAN": {LDAN, 5},
	"LD1N": {LD1N, 5},
	"LD2N": {LD2N, 5},
	"LD3N": {LD3N, 5},
	"LD4N": {LD4N, 5},
	"LD5N": {LD5N, 5},
	"LD6N": {LD6N, 5},
	"LDXN": {LDXN, 5},
	"STA":  {STA, 5},
	"ST1":  {ST1, 5},
	"ST2":  {ST2, 5},
	"ST3":  {ST3, 5},
	"ST4":  {ST4, 5},
	"ST5":  {ST5, 5},
	"ST6":  {ST6, 5},
	"STX":  {STX, 5},
	"STJ":  {STJ, 2},
	"STZ":  {STZ, 5},
	"JBUS": {JBUS, 0},
	"IOC":  {IOC, 0},
	"IN":   {IN, 0},
	"OUT":  {OUT, 0},
	"JRED": {JRED, 0},
	"JMP":  {JMP, 0},
	"JSJ":  {JSJ, 1},
	"JOV":  {JOV, 2},
	"JNOV": {JNOV, 3},
	"JL":   {JL, 4},
	"JE":   {JE, 5},
	"JG":   {JG, 6},
	"JGE":  {JGE, 7},
	"JNE":  {JNE, 8},
	"JLE":  {JLE, 9},
	"JAN":  {JAN, 0},
	"JAZ":  {JAZ, 1},
	"JAP":  {JAP, 2},
	"JANN": {JANN, 3},
	"JANZ": {JANZ, 4},
	"JANP": {JANP, 5},
	"J1N":  {J1N, 0},
	"J1Z":  {J1Z, 1},
	"J1P":  {J1P, 2},
	"J1NN": {J1NN, 3},
	"J1NZ": {J1NZ, 4},
	"J1NP": {J1NP, 5},
	"J2N":  {J2N, 0},
	"J2Z":  {J2Z, 1},
	"J2P":  {J2P, 2},
	"J2NN": {J2NN, 3},
	"J2NZ": {J2NZ, 4},
	"J2NP": {J2NP, 5},
	"J3N":  {J3N, 0},
	"J3Z":  {J3Z, 1},
	"J3P":  {J3P, 2},
	"J3NN": {J3NN, 3},
	"J3NZ": {J3NZ, 4},
	"J3NP": {J3NP, 5},
	"J4N":  {J4N, 0},
	"J4Z":  {J4Z, 1},
	"J4P":  {J4P, 2},
	"J4NN": {J4NN, 3},
	"J4NZ": {J4NZ, 4},
	"J4NP": {J4NP, 5},
	"J5N":  {J5N, 0},
	"J5Z":  {J5Z, 1},
	"J5P":  {J5P, 2},
	"J5NN": {J5NN, 3},
	"J5NZ": {J5NZ, 4},
	"J5NP": {J5NP, 5},
	"J6N":  {J6N, 0},
	"J6Z":  {J6Z, 1},
	"J6P":  {J6P, 2},
	"J6NN": {J6NN, 3},
	"J6NZ": {J6NZ, 4},
	"J6NP": {J6NP, 5},
	"JXN":  {JXN, 0},
	"JXZ":  {JXZ, 1},
	"JXP":  {JXP, 2},
	"JXNN": {JXNN, 3},
	"JXNZ": {JXNZ, 4},
	"JXNP": {JXNP, 5},
	"INCA": {INCA, 0},
	"DECA": {DECA, 1},
	"ENTA": {ENTA, 2},
	"ENNA": {ENNA, 3},
	"INC1": {INC1, 0},
	"DEC1": {DEC1, 1},
	"ENT1": {ENT1, 2},
	"ENN1": {ENN1, 3},
	"INC2": {INC2, 0},
	"DEC2": {DEC2, 1},
	"ENT2": {ENT2, 2},
	"ENN2": {ENN2, 3},
	"INC3": {INC3, 0},
	"DEC3": {DEC3, 1},
	"ENT3": {ENT3, 2},
	"ENN3": {ENN3, 3},
	"INC4": {INC4, 0},
	"DEC4": {DEC4, 1},
	"ENT4": {ENT4, 2},
	"ENN4": {ENN4, 3},
	"INC5": {INC5, 0},
	"DEC5": {DEC5, 1},
	"ENT5": {ENT5, 2},
	"ENN5": {ENN5, 3},
	"INC6": {INC6, 0},
	"DEC6": {DEC6, 1},
	"ENT6": {ENT6, 2},
	"ENN6": {ENN6, 3},
	"INCX": {INCX, 0},
	"DECX": {DECX, 1},
	"ENTX": {ENTX, 2},
	"ENNX": {ENNX, 3},
	"CMPA": {CMPA, 5},
	"CMP1": {CMP1, 5},
	"CMP2": {CMP2, 5},
	"CMP3": {CMP3, 5},
	"CMP4": {CMP4, 5},
	"CMP5": {CMP5, 5},
	"CMP6": {CMP6, 5},
	"CMPX": {CMPX, 5},
}
