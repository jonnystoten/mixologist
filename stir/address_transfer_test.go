package stir

import (
	"testing"

	"jonnystoten.com/mixologist/mix"
)

func TestENTA(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.ENTA, FieldSpec: 2, Address: mix.NewAddress(2000)},
			mix.NewWord(2000),
		},
		{
			mix.Instruction{OpCode: mix.ENTA, FieldSpec: 2, Address: mix.NewAddress(-2000)},
			mix.NewWord(-2000),
		},
		{
			mix.Instruction{OpCode: mix.ENTA, FieldSpec: 2, Address: mix.NewAddress(2000), IndexSpec: 1},
			mix.NewWord(2100),
		},
		{
			mix.Instruction{OpCode: mix.ENTA, FieldSpec: 2, Address: mix.NewAddress(0), IndexSpec: 2},
			mix.NewWord(0),
		},
		{
			mix.Instruction{OpCode: mix.ENTA, FieldSpec: 2, Address: mix.NewAddress(0), IndexSpec: 3},
			mix.NewWord(0),
		},
		{
			mix.Instruction{OpCode: mix.ENTA, FieldSpec: 2, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 2},
			mix.Word{Sign: mix.Negative},
		},
		{
			mix.Instruction{OpCode: mix.ENTA, FieldSpec: 2, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 3},
			mix.Word{Sign: mix.Negative},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Index[0] = mix.NewAddress(100)
		computer.Index[1] = mix.Address{}                   // 0
		computer.Index[2] = mix.Address{Sign: mix.Negative} // -0

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Accumulator
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestENTX(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.ENTX, FieldSpec: 2, Address: mix.NewAddress(2000)},
			mix.NewWord(2000),
		},
		{
			mix.Instruction{OpCode: mix.ENTX, FieldSpec: 2, Address: mix.NewAddress(-2000)},
			mix.NewWord(-2000),
		},
		{
			mix.Instruction{OpCode: mix.ENTX, FieldSpec: 2, Address: mix.NewAddress(2000), IndexSpec: 1},
			mix.NewWord(2100),
		},
		{
			mix.Instruction{OpCode: mix.ENTX, FieldSpec: 2, Address: mix.NewAddress(0), IndexSpec: 2},
			mix.NewWord(0),
		},
		{
			mix.Instruction{OpCode: mix.ENTX, FieldSpec: 2, Address: mix.NewAddress(0), IndexSpec: 3},
			mix.NewWord(0),
		},
		{
			mix.Instruction{OpCode: mix.ENTX, FieldSpec: 2, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 2},
			mix.Word{Sign: mix.Negative},
		},
		{
			mix.Instruction{OpCode: mix.ENTX, FieldSpec: 2, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 3},
			mix.Word{Sign: mix.Negative},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Index[0] = mix.NewAddress(100)
		computer.Index[1] = mix.Address{}                   // 0
		computer.Index[2] = mix.Address{Sign: mix.Negative} // -0

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Extension
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestENTi(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		index       int
		expected    mix.Address
	}{
		{
			mix.Instruction{OpCode: mix.ENT1, FieldSpec: 2, Address: mix.NewAddress(2000)},
			1,
			mix.NewAddress(2000),
		},
		{
			mix.Instruction{OpCode: mix.ENT2, FieldSpec: 2, Address: mix.NewAddress(2000), IndexSpec: 1},
			2,
			mix.NewAddress(2100),
		},
		{
			mix.Instruction{OpCode: mix.ENT3, FieldSpec: 2, Address: mix.NewAddress(0), IndexSpec: 2},
			3,
			mix.NewAddress(0),
		},
		{
			mix.Instruction{OpCode: mix.ENT4, FieldSpec: 2, Address: mix.NewAddress(0), IndexSpec: 3},
			4,
			mix.NewAddress(0),
		},
		{
			mix.Instruction{OpCode: mix.ENT5, FieldSpec: 2, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 2},
			5,
			mix.Address{Sign: mix.Negative},
		},
		{
			mix.Instruction{OpCode: mix.ENT6, FieldSpec: 2, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 3},
			6,
			mix.Address{Sign: mix.Negative},
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Index[0] = mix.NewAddress(100)
		computer.Index[1] = mix.Address{}                   // 0
		computer.Index[2] = mix.Address{Sign: mix.Negative} // -0

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Index[test.index-1]
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestENNA(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.ENNA, FieldSpec: 3, Address: mix.NewAddress(2000)},
			mix.NewWord(-2000),
		},
		{
			mix.Instruction{OpCode: mix.ENNA, FieldSpec: 3, Address: mix.NewAddress(-2000)},
			mix.NewWord(2000),
		},
		{
			mix.Instruction{OpCode: mix.ENNA, FieldSpec: 3, Address: mix.NewAddress(2000), IndexSpec: 1},
			mix.NewWord(-2100),
		},
		{
			mix.Instruction{OpCode: mix.ENNA, FieldSpec: 3, Address: mix.NewAddress(0), IndexSpec: 2},
			mix.Word{Sign: mix.Negative},
		},
		{
			mix.Instruction{OpCode: mix.ENNA, FieldSpec: 3, Address: mix.NewAddress(0), IndexSpec: 3},
			mix.Word{Sign: mix.Negative},
		},
		{
			mix.Instruction{OpCode: mix.ENNA, FieldSpec: 3, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 2},
			mix.NewWord(0),
		},
		{
			mix.Instruction{OpCode: mix.ENNA, FieldSpec: 3, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 3},
			mix.NewWord(0),
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Index[0] = mix.NewAddress(100)
		computer.Index[1] = mix.Address{}                   // 0
		computer.Index[2] = mix.Address{Sign: mix.Negative} // -0

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Accumulator
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestENNX(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.ENNX, FieldSpec: 3, Address: mix.NewAddress(2000)},
			mix.NewWord(-2000),
		},
		{
			mix.Instruction{OpCode: mix.ENNX, FieldSpec: 3, Address: mix.NewAddress(-2000)},
			mix.NewWord(2000),
		},
		{
			mix.Instruction{OpCode: mix.ENNX, FieldSpec: 3, Address: mix.NewAddress(2000), IndexSpec: 1},
			mix.NewWord(-2100),
		},
		{
			mix.Instruction{OpCode: mix.ENNX, FieldSpec: 3, Address: mix.NewAddress(0), IndexSpec: 2},
			mix.Word{Sign: mix.Negative},
		},
		{
			mix.Instruction{OpCode: mix.ENNX, FieldSpec: 3, Address: mix.NewAddress(0), IndexSpec: 3},
			mix.Word{Sign: mix.Negative},
		},
		{
			mix.Instruction{OpCode: mix.ENNX, FieldSpec: 3, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 2},
			mix.NewWord(0),
		},
		{
			mix.Instruction{OpCode: mix.ENNX, FieldSpec: 3, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 3},
			mix.NewWord(0),
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Index[0] = mix.NewAddress(100)
		computer.Index[1] = mix.Address{}                   // 0
		computer.Index[2] = mix.Address{Sign: mix.Negative} // -0

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Extension
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestENNi(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		index       int
		expected    mix.Address
	}{
		{
			mix.Instruction{OpCode: mix.ENN1, FieldSpec: 3, Address: mix.NewAddress(2000)},
			1,
			mix.NewAddress(-2000),
		},
		{
			mix.Instruction{OpCode: mix.ENN2, FieldSpec: 3, Address: mix.NewAddress(2000), IndexSpec: 1},
			2,
			mix.NewAddress(-2100),
		},
		{
			mix.Instruction{OpCode: mix.ENN3, FieldSpec: 3, Address: mix.NewAddress(0), IndexSpec: 2},
			3,
			mix.Address{Sign: mix.Negative},
		},
		{
			mix.Instruction{OpCode: mix.ENN4, FieldSpec: 3, Address: mix.NewAddress(0), IndexSpec: 3},
			4,
			mix.Address{Sign: mix.Negative},
		},
		{
			mix.Instruction{OpCode: mix.ENN5, FieldSpec: 3, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 2},
			5,
			mix.NewAddress(0),
		},
		{
			mix.Instruction{OpCode: mix.ENN6, FieldSpec: 3, Address: mix.Address{Sign: mix.Negative}, IndexSpec: 3},
			6,
			mix.NewAddress(0),
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Index[0] = mix.NewAddress(100)
		computer.Index[1] = mix.Address{}                   // 0
		computer.Index[2] = mix.Address{Sign: mix.Negative} // -0

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Index[test.index-1]
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}
