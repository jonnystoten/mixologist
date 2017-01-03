package stir

import (
	"testing"

	"jonnystoten.com/mixologist/mix"
)

func TestSTA(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.STA, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{6, 7, 8, 9, 0}},
		},
		{
			mix.Instruction{OpCode: mix.STA, FieldSpec: mix.NewFieldSpec(1, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{6, 7, 8, 9, 0}},
		},
		{
			mix.Instruction{OpCode: mix.STA, FieldSpec: mix.NewFieldSpec(5, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 2, 3, 4, 0}},
		},
		{
			mix.Instruction{OpCode: mix.STA, FieldSpec: mix.NewFieldSpec(2, 2), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 0, 3, 4, 5}},
		},
		{
			mix.Instruction{OpCode: mix.STA, FieldSpec: mix.NewFieldSpec(2, 3), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 9, 0, 4, 5}},
		},
		{
			mix.Instruction{OpCode: mix.STA, FieldSpec: mix.NewFieldSpec(0, 1), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 2, 3, 4, 5}},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Memory[2000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 2, 3, 4, 5}}
		computer.Accumulator = mix.Word{Sign: mix.Positive, Bytes: [5]byte{6, 7, 8, 9, 0}}

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Memory[2000]
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestSTX(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.STX, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{6, 7, 8, 9, 0}},
		},
		{
			mix.Instruction{OpCode: mix.STX, FieldSpec: mix.NewFieldSpec(1, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{6, 7, 8, 9, 0}},
		},
		{
			mix.Instruction{OpCode: mix.STX, FieldSpec: mix.NewFieldSpec(5, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 2, 3, 4, 0}},
		},
		{
			mix.Instruction{OpCode: mix.STX, FieldSpec: mix.NewFieldSpec(2, 2), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 0, 3, 4, 5}},
		},
		{
			mix.Instruction{OpCode: mix.STX, FieldSpec: mix.NewFieldSpec(2, 3), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 9, 0, 4, 5}},
		},
		{
			mix.Instruction{OpCode: mix.STX, FieldSpec: mix.NewFieldSpec(0, 1), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 2, 3, 4, 5}},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Memory[2000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 2, 3, 4, 5}}
		computer.Extension = mix.Word{Sign: mix.Positive, Bytes: [5]byte{6, 7, 8, 9, 0}}

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Memory[2000]
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestSTi(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		index       int
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.ST1, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(2000)},
			1,
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 6, 7}},
		},
		{
			mix.Instruction{OpCode: mix.ST2, FieldSpec: mix.NewFieldSpec(1, 5), Address: mix.NewAddress(2000)},
			2,
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 6, 7}},
		},
		{
			mix.Instruction{OpCode: mix.ST3, FieldSpec: mix.NewFieldSpec(5, 5), Address: mix.NewAddress(2000)},
			3,
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 2, 3, 4, 7}},
		},
		{
			mix.Instruction{OpCode: mix.ST4, FieldSpec: mix.NewFieldSpec(2, 2), Address: mix.NewAddress(2000)},
			4,
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 7, 3, 4, 5}},
		},
		{
			mix.Instruction{OpCode: mix.ST5, FieldSpec: mix.NewFieldSpec(2, 3), Address: mix.NewAddress(2000)},
			5,
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 6, 7, 4, 5}},
		},
		{
			mix.Instruction{OpCode: mix.ST6, FieldSpec: mix.NewFieldSpec(0, 1), Address: mix.NewAddress(2000)},
			6,
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{7, 2, 3, 4, 5}},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Memory[2000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 2, 3, 4, 5}}
		computer.Index[test.index-1] = mix.Address{Sign: mix.Positive, Bytes: [2]byte{6, 7}}

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Memory[2000]
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestSTJ(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.STJ, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 6, 7}},
		},
		{
			mix.Instruction{OpCode: mix.STJ, FieldSpec: mix.NewFieldSpec(1, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 6, 7}},
		},
		{
			mix.Instruction{OpCode: mix.STJ, FieldSpec: mix.NewFieldSpec(5, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 2, 3, 4, 7}},
		},
		{
			mix.Instruction{OpCode: mix.STJ, FieldSpec: mix.NewFieldSpec(2, 2), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 7, 3, 4, 5}},
		},
		{
			mix.Instruction{OpCode: mix.STJ, FieldSpec: mix.NewFieldSpec(2, 3), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 6, 7, 4, 5}},
		},
		{
			mix.Instruction{OpCode: mix.STJ, FieldSpec: mix.NewFieldSpec(0, 1), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{7, 2, 3, 4, 5}},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Memory[2000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 2, 3, 4, 5}}
		computer.JumpAddress = mix.Address{Sign: mix.Positive, Bytes: [2]byte{6, 7}}

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Memory[2000]
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestSTZ(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.STZ, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 0}},
		},
		{
			mix.Instruction{OpCode: mix.STZ, FieldSpec: mix.NewFieldSpec(1, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 0, 0}},
		},
		{
			mix.Instruction{OpCode: mix.STZ, FieldSpec: mix.NewFieldSpec(5, 5), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 2, 3, 4, 0}},
		},
		{
			mix.Instruction{OpCode: mix.STZ, FieldSpec: mix.NewFieldSpec(2, 2), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 0, 3, 4, 5}},
		},
		{
			mix.Instruction{OpCode: mix.STZ, FieldSpec: mix.NewFieldSpec(2, 3), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 0, 0, 4, 5}},
		},
		{
			mix.Instruction{OpCode: mix.STZ, FieldSpec: mix.NewFieldSpec(0, 1), Address: mix.NewAddress(2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 2, 3, 4, 5}},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Memory[2000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 2, 3, 4, 5}}

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Memory[2000]
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}
