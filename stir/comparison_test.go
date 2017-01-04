package stir

import (
	"testing"

	"jonnystoten.com/mixologist/mix"
)

func TestCMPA(t *testing.T) {
	tests := []struct {
		accBefore   mix.Word
		instruction mix.Instruction
		expected    mix.Comparison
	}{
		{
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 10}},
			mix.Instruction{OpCode: mix.CMPA, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Equal,
		},
		{
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 5}},
			mix.Instruction{OpCode: mix.CMPA, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Less,
		},
		{
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 15}},
			mix.Instruction{OpCode: mix.CMPA, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Greater,
		},
		{
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 0, 0}},
			mix.Instruction{OpCode: mix.CMPA, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(0)},
			mix.Equal,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Accumulator = test.accBefore
		computer.Memory[0] = mix.Word{}
		computer.Memory[1000] = mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 10}}

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Comparison
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestCMPX(t *testing.T) {
	tests := []struct {
		extBefore   mix.Word
		instruction mix.Instruction
		expected    mix.Comparison
	}{
		{
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 10}},
			mix.Instruction{OpCode: mix.CMPX, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Equal,
		},
		{
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 5}},
			mix.Instruction{OpCode: mix.CMPX, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Less,
		},
		{
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 15}},
			mix.Instruction{OpCode: mix.CMPX, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Greater,
		},
		{
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 0, 0}},
			mix.Instruction{OpCode: mix.CMPX, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(0)},
			mix.Equal,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Extension = test.extBefore
		computer.Memory[0] = mix.Word{}
		computer.Memory[1000] = mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 10}}

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Comparison
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestCMPi(t *testing.T) {
	tests := []struct {
		idxBefore   mix.Address
		index       int
		instruction mix.Instruction
		expected    mix.Comparison
	}{
		{
			mix.Address{Sign: mix.Positive, Bytes: [2]byte{0, 10}},
			1,
			mix.Instruction{OpCode: mix.CMP1, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Equal,
		},
		{
			mix.Address{Sign: mix.Positive, Bytes: [2]byte{0, 5}},
			2,
			mix.Instruction{OpCode: mix.CMP2, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Less,
		},
		{
			mix.Address{Sign: mix.Positive, Bytes: [2]byte{0, 15}},
			3,
			mix.Instruction{OpCode: mix.CMP3, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Greater,
		},
		{
			mix.Address{Sign: mix.Negative, Bytes: [2]byte{0, 0}},
			4,
			mix.Instruction{OpCode: mix.CMP4, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(0)},
			mix.Equal,
		},
		{
			mix.Address{Sign: mix.Positive, Bytes: [2]byte{0, 5}},
			5,
			mix.Instruction{OpCode: mix.CMP5, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Less,
		},
		{
			mix.Address{Sign: mix.Positive, Bytes: [2]byte{0, 15}},
			6,
			mix.Instruction{OpCode: mix.CMP6, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Greater,
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Index[test.index-1] = test.idxBefore
		computer.Memory[0] = mix.Word{}
		computer.Memory[1000] = mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 10}}

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Comparison
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

// TODO: tests for CMPi
