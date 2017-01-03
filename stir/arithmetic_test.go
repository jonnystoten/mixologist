package stir

import (
	"testing"

	"jonnystoten.com/mixologist/mix"
)

func TestADD(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.ADD, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{20, 54, 6, 3, 8}}, // +[1334][6][200]
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Accumulator = mix.Word{Sign: mix.Positive, Bytes: [5]byte{19, 18, 1, 2, 22}} // +[1234][1][150]
		computer.Memory[1000] = mix.Word{Sign: mix.Positive, Bytes: [5]byte{1, 36, 5, 0, 50}} // +[100][5][50]

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Accumulator
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestSUB(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.SUB, FieldSpec: mix.NewFieldSpec(0, 5), Address: mix.NewAddress(1000)},
			mix.Word{Sign: mix.Positive, Bytes: [5]byte{11, 62, 2, 21, 55}}, // +[766][149][?]
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Accumulator = mix.Word{Sign: mix.Negative, Bytes: [5]byte{19, 18, 0, 0, 9}}   // -[1234][0][0][9]
		computer.Memory[1000] = mix.Word{Sign: mix.Negative, Bytes: [5]byte{31, 16, 2, 22, 0}} // -[2000][150][0]

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Accumulator
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}
