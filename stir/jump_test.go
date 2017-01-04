package stir

import (
	"testing"

	"jonnystoten.com/mixologist/mix"
)

func TestJMP(t *testing.T) {
	computer := NewComputer()
	computer.ProgramCounter = 100
	instruction := mix.Instruction{OpCode: mix.JMP, FieldSpec: 0, Address: mix.NewAddress(1000)}
	computer.Memory[100] = mix.EncodeInstruction(instruction)

	computer.FetchDecodeExecute()

	actualPC := computer.ProgramCounter
	actualJ := computer.JumpAddress
	if actualPC != 1000 {
		t.Errorf("Expected PC to be %+v, actual %+v", 1000, actualPC)
	}
	if actualJ != mix.NewAddress(100) {
		t.Errorf("Expected rJ to be %+v, actual %+v", mix.NewAddress(100), actualJ)
	}
}

func TestJSJ(t *testing.T) {
	computer := NewComputer()
	computer.ProgramCounter = 100
	computer.JumpAddress = mix.NewAddress(50)
	instruction := mix.Instruction{OpCode: mix.JSJ, FieldSpec: 1, Address: mix.NewAddress(1000)}
	computer.Memory[100] = mix.EncodeInstruction(instruction)

	computer.FetchDecodeExecute()

	actualPC := computer.ProgramCounter
	actualJ := computer.JumpAddress
	if actualPC != 1000 {
		t.Errorf("Expected PC to be %+v, actual %+v", 1000, actualPC)
	}
	if actualJ != mix.NewAddress(50) {
		t.Errorf("Expected rJ to be %+v, actual %+v", mix.NewAddress(50), actualJ)
	}
}

func TestJOV(t *testing.T) {
	tests := []struct {
		overflow bool
		expected int
	}{
		{
			true,
			1000,
		},
		{
			false,
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Overflow = test.overflow
		instruction := mix.Instruction{OpCode: mix.JOV, FieldSpec: 2, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}

		if computer.Overflow {
			t.Error("Expected overflow to be switched off after JOV")
		}
	}
}

func TestJNOV(t *testing.T) {
	tests := []struct {
		overflow bool
		expected int
	}{
		{
			true,
			100,
		},
		{
			false,
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Overflow = test.overflow
		instruction := mix.Instruction{OpCode: mix.JNOV, FieldSpec: 3, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}

		if computer.Overflow {
			t.Error("Expected overflow to be switched off after JNOV")
		}
	}
}

func TestJL(t *testing.T) {
	tests := []struct {
		comparision mix.Comparison
		expected    int
	}{
		{
			mix.Less,
			1000,
		},
		{
			mix.Equal,
			100,
		},
		{
			mix.Greater,
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Comparison = test.comparision
		instruction := mix.Instruction{OpCode: mix.JL, FieldSpec: 4, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJE(t *testing.T) {
	tests := []struct {
		comparision mix.Comparison
		expected    int
	}{
		{
			mix.Less,
			100,
		},
		{
			mix.Equal,
			1000,
		},
		{
			mix.Greater,
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Comparison = test.comparision
		instruction := mix.Instruction{OpCode: mix.JE, FieldSpec: 5, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJG(t *testing.T) {
	tests := []struct {
		comparision mix.Comparison
		expected    int
	}{
		{
			mix.Less,
			100,
		},
		{
			mix.Equal,
			100,
		},
		{
			mix.Greater,
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Comparison = test.comparision
		instruction := mix.Instruction{OpCode: mix.JG, FieldSpec: 6, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJGE(t *testing.T) {
	tests := []struct {
		comparision mix.Comparison
		expected    int
	}{
		{
			mix.Less,
			100,
		},
		{
			mix.Equal,
			1000,
		},
		{
			mix.Greater,
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Comparison = test.comparision
		instruction := mix.Instruction{OpCode: mix.JGE, FieldSpec: 7, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJNE(t *testing.T) {
	tests := []struct {
		comparision mix.Comparison
		expected    int
	}{
		{
			mix.Less,
			1000,
		},
		{
			mix.Equal,
			100,
		},
		{
			mix.Greater,
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Comparison = test.comparision
		instruction := mix.Instruction{OpCode: mix.JNE, FieldSpec: 8, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJLE(t *testing.T) {
	tests := []struct {
		comparision mix.Comparison
		expected    int
	}{
		{
			mix.Less,
			1000,
		},
		{
			mix.Equal,
			1000,
		},
		{
			mix.Greater,
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Comparison = test.comparision
		instruction := mix.Instruction{OpCode: mix.JLE, FieldSpec: 9, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJAN(t *testing.T) {
	tests := []struct {
		acc      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			100,
		},
		{
			mix.NewWord(-5000),
			1000,
		},
		{
			mix.NewWord(0),
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Accumulator = test.acc
		instruction := mix.Instruction{OpCode: mix.JAN, FieldSpec: 0, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJAZ(t *testing.T) {
	tests := []struct {
		acc      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			100,
		},
		{
			mix.NewWord(-5000),
			100,
		},
		{
			mix.NewWord(0),
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Accumulator = test.acc
		instruction := mix.Instruction{OpCode: mix.JAZ, FieldSpec: 1, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJAP(t *testing.T) {
	tests := []struct {
		acc      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			1000,
		},
		{
			mix.NewWord(-5000),
			100,
		},
		{
			mix.NewWord(0),
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Accumulator = test.acc
		instruction := mix.Instruction{OpCode: mix.JAP, FieldSpec: 2, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJANN(t *testing.T) {
	tests := []struct {
		acc      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			1000,
		},
		{
			mix.NewWord(-5000),
			100,
		},
		{
			mix.NewWord(0),
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Accumulator = test.acc
		instruction := mix.Instruction{OpCode: mix.JANN, FieldSpec: 3, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJANZ(t *testing.T) {
	tests := []struct {
		acc      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			1000,
		},
		{
			mix.NewWord(-5000),
			1000,
		},
		{
			mix.NewWord(0),
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Accumulator = test.acc
		instruction := mix.Instruction{OpCode: mix.JANZ, FieldSpec: 4, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJANP(t *testing.T) {
	tests := []struct {
		acc      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			100,
		},
		{
			mix.NewWord(-5000),
			1000,
		},
		{
			mix.NewWord(0),
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Accumulator = test.acc
		instruction := mix.Instruction{OpCode: mix.JANP, FieldSpec: 5, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJXN(t *testing.T) {
	tests := []struct {
		ext      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			100,
		},
		{
			mix.NewWord(-5000),
			1000,
		},
		{
			mix.NewWord(0),
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Extension = test.ext
		instruction := mix.Instruction{OpCode: mix.JXN, FieldSpec: 0, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJXZ(t *testing.T) {
	tests := []struct {
		ext      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			100,
		},
		{
			mix.NewWord(-5000),
			100,
		},
		{
			mix.NewWord(0),
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Extension = test.ext
		instruction := mix.Instruction{OpCode: mix.JXZ, FieldSpec: 1, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJXP(t *testing.T) {
	tests := []struct {
		ext      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			1000,
		},
		{
			mix.NewWord(-5000),
			100,
		},
		{
			mix.NewWord(0),
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Extension = test.ext
		instruction := mix.Instruction{OpCode: mix.JXP, FieldSpec: 2, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJXNN(t *testing.T) {
	tests := []struct {
		ext      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			1000,
		},
		{
			mix.NewWord(-5000),
			100,
		},
		{
			mix.NewWord(0),
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Extension = test.ext
		instruction := mix.Instruction{OpCode: mix.JXNN, FieldSpec: 3, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJXNZ(t *testing.T) {
	tests := []struct {
		ext      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			1000,
		},
		{
			mix.NewWord(-5000),
			1000,
		},
		{
			mix.NewWord(0),
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Extension = test.ext
		instruction := mix.Instruction{OpCode: mix.JXNZ, FieldSpec: 4, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJXNP(t *testing.T) {
	tests := []struct {
		ext      mix.Word
		expected int
	}{
		{
			mix.NewWord(5000),
			100,
		},
		{
			mix.NewWord(-5000),
			1000,
		},
		{
			mix.NewWord(0),
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Extension = test.ext
		instruction := mix.Instruction{OpCode: mix.JXNP, FieldSpec: 5, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJiN(t *testing.T) {
	tests := []struct {
		idx      mix.Address
		expected int
	}{
		{
			mix.NewAddress(5000),
			100,
		},
		{
			mix.NewAddress(-5000),
			1000,
		},
		{
			mix.NewAddress(0),
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Index[0] = test.idx
		instruction := mix.Instruction{OpCode: mix.J1N, FieldSpec: 0, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJiZ(t *testing.T) {
	tests := []struct {
		idx      mix.Address
		expected int
	}{
		{
			mix.NewAddress(5000),
			100,
		},
		{
			mix.NewAddress(-5000),
			100,
		},
		{
			mix.NewAddress(0),
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Index[1] = test.idx
		instruction := mix.Instruction{OpCode: mix.J2Z, FieldSpec: 1, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJiP(t *testing.T) {
	tests := []struct {
		idx      mix.Address
		expected int
	}{
		{
			mix.NewAddress(5000),
			1000,
		},
		{
			mix.NewAddress(-5000),
			100,
		},
		{
			mix.NewAddress(0),
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Index[2] = test.idx
		instruction := mix.Instruction{OpCode: mix.J3P, FieldSpec: 2, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJiNN(t *testing.T) {
	tests := []struct {
		idx      mix.Address
		expected int
	}{
		{
			mix.NewAddress(5000),
			1000,
		},
		{
			mix.NewAddress(-5000),
			100,
		},
		{
			mix.NewAddress(0),
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Index[3] = test.idx
		instruction := mix.Instruction{OpCode: mix.J4NN, FieldSpec: 3, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJiNZ(t *testing.T) {
	tests := []struct {
		idx      mix.Address
		expected int
	}{
		{
			mix.NewAddress(5000),
			1000,
		},
		{
			mix.NewAddress(-5000),
			1000,
		},
		{
			mix.NewAddress(0),
			100,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Index[4] = test.idx
		instruction := mix.Instruction{OpCode: mix.J5NZ, FieldSpec: 4, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}

func TestJiNP(t *testing.T) {
	tests := []struct {
		idx      mix.Address
		expected int
	}{
		{
			mix.NewAddress(5000),
			100,
		},
		{
			mix.NewAddress(-5000),
			1000,
		},
		{
			mix.NewAddress(0),
			1000,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.ProgramCounter = 100
		computer.Index[5] = test.idx
		instruction := mix.Instruction{OpCode: mix.J6NP, FieldSpec: 5, Address: mix.NewAddress(1000)}
		computer.Memory[100] = mix.EncodeInstruction(instruction)

		computer.FetchDecodeExecute()

		actual := computer.ProgramCounter
		if actual != test.expected {
			t.Errorf("Expected PC to be %+v, actual %+v", test.expected, actual)
		}
	}
}
