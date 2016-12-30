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
		computer.Execute(test.instruction)

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
		computer.Execute(test.instruction)

		actual := computer.Extension
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestLDi(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		index       int
		expected    mix.Address
	}{
		{
			mix.Instruction{OpCode: mix.LD1, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			1,
			mix.Address{Sign: mix.Negative, Value: [2]byte{5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LD2, FieldSpec: mix.NewFieldSpec(1, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			2,
			mix.Address{Sign: mix.Positive, Value: [2]byte{5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LD3, FieldSpec: mix.NewFieldSpec(3, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			3,
			mix.Address{Sign: mix.Positive, Value: [2]byte{5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LD4, FieldSpec: mix.NewFieldSpec(0, 3), Address: mix.NewAddress(mix.Positive, 2000)},
			4,
			mix.Address{Sign: mix.Negative, Value: [2]byte{0, 0}},
		},
		{
			mix.Instruction{OpCode: mix.LD5, FieldSpec: mix.NewFieldSpec(4, 4), Address: mix.NewAddress(mix.Positive, 2000)},
			5,
			mix.Address{Sign: mix.Positive, Value: [2]byte{0, 5}},
		},
		{
			mix.Instruction{OpCode: mix.LD6, FieldSpec: mix.NewFieldSpec(0, 0), Address: mix.NewAddress(mix.Positive, 2000)},
			6,
			mix.Address{Sign: mix.Negative, Value: [2]byte{0, 0}},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Memory[2000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 5, 4}}
		computer.Execute(test.instruction)

		actual := computer.Index[test.index-1]
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestLDAN(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.LDAN, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{1, 16, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDAN, FieldSpec: mix.NewFieldSpec(1, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 16, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDAN, FieldSpec: mix.NewFieldSpec(3, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDAN, FieldSpec: mix.NewFieldSpec(0, 3), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 1, 16, 3}},
		},
		{
			mix.Instruction{OpCode: mix.LDAN, FieldSpec: mix.NewFieldSpec(4, 4), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 0, 5}},
		},
		{
			mix.Instruction{OpCode: mix.LDAN, FieldSpec: mix.NewFieldSpec(0, 0), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 0}},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Memory[2000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 16, 3, 5, 4}}
		computer.Execute(test.instruction)

		actual := computer.Accumulator
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestLDXN(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.LDXN, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{1, 16, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDXN, FieldSpec: mix.NewFieldSpec(1, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 16, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDXN, FieldSpec: mix.NewFieldSpec(3, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 3, 5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LDXN, FieldSpec: mix.NewFieldSpec(0, 3), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 1, 16, 3}},
		},
		{
			mix.Instruction{OpCode: mix.LDXN, FieldSpec: mix.NewFieldSpec(4, 4), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 0, 5}},
		},
		{
			mix.Instruction{OpCode: mix.LDXN, FieldSpec: mix.NewFieldSpec(0, 0), Address: mix.NewAddress(mix.Positive, 2000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 0}},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Memory[2000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{1, 16, 3, 5, 4}}
		computer.Execute(test.instruction)

		actual := computer.Extension
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestLDiN(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		index       int
		expected    mix.Address
	}{
		{
			mix.Instruction{OpCode: mix.LD1N, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			1,
			mix.Address{Sign: mix.Positive, Value: [2]byte{5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LD2N, FieldSpec: mix.NewFieldSpec(1, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			2,
			mix.Address{Sign: mix.Negative, Value: [2]byte{5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LD3N, FieldSpec: mix.NewFieldSpec(3, 5), Address: mix.NewAddress(mix.Positive, 2000)},
			3,
			mix.Address{Sign: mix.Negative, Value: [2]byte{5, 4}},
		},
		{
			mix.Instruction{OpCode: mix.LD4N, FieldSpec: mix.NewFieldSpec(0, 3), Address: mix.NewAddress(mix.Positive, 2000)},
			4,
			mix.Address{Sign: mix.Positive, Value: [2]byte{0, 0}},
		},
		{
			mix.Instruction{OpCode: mix.LD5N, FieldSpec: mix.NewFieldSpec(4, 4), Address: mix.NewAddress(mix.Positive, 2000)},
			5,
			mix.Address{Sign: mix.Negative, Value: [2]byte{0, 5}},
		},
		{
			mix.Instruction{OpCode: mix.LD6N, FieldSpec: mix.NewFieldSpec(0, 0), Address: mix.NewAddress(mix.Positive, 2000)},
			6,
			mix.Address{Sign: mix.Positive, Value: [2]byte{0, 0}},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Memory[2000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 5, 4}}
		computer.Execute(test.instruction)

		actual := computer.Index[test.index-1]
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}
