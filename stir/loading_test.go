package stir

import (
	"testing"

	"jonnystoten.com/mixologist/mix"
)

func TestLDA(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.LDA, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 16, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDA, FieldSpec: mix.NewFieldSpec(1, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{1, 16, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDA, FieldSpec: mix.NewFieldSpec(3, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDA, FieldSpec: mix.NewFieldSpec(0, 3), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 1, 16, 3}},
		},
		{
			mix.Instruction{OpCode: mix.LDA, FieldSpec: mix.NewFieldSpec(4, 4), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 5}},
		},
		{
			mix.Instruction{OpCode: mix.LDA, FieldSpec: mix.NewFieldSpec(0, 0), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 0, 0}},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Memory[2000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 16, 3, 5, 4}}
		computer.Execute(&test.instruction)

		actual := computer.Accumulator
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestLDX(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.LDX, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 16, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDX, FieldSpec: mix.NewFieldSpec(1, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{1, 16, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDX, FieldSpec: mix.NewFieldSpec(3, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDX, FieldSpec: mix.NewFieldSpec(0, 3), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 1, 16, 3}},
		},
		{
			mix.Instruction{OpCode: mix.LDX, FieldSpec: mix.NewFieldSpec(4, 4), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 5}},
		},
		{
			mix.Instruction{OpCode: mix.LDX, FieldSpec: mix.NewFieldSpec(0, 0), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 0, 0}},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Memory[2000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 16, 3, 5, 4}}
		computer.Execute(&test.instruction)

		actual := computer.Extension
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}
