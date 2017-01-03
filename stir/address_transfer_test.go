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

func TestINCA(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.INCA, FieldSpec: 0, Address: mix.NewAddress(2000)},
			mix.NewWord(3000),
		},
		{
			mix.Instruction{OpCode: mix.INCA, FieldSpec: 0, Address: mix.NewAddress(-2000)},
			mix.NewWord(-1000),
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Accumulator = mix.NewWord(1000)

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Accumulator
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestINCX(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.INCX, FieldSpec: 0, Address: mix.NewAddress(2000)},
			mix.NewWord(3000),
		},
		{
			mix.Instruction{OpCode: mix.INCX, FieldSpec: 0, Address: mix.NewAddress(-2000)},
			mix.NewWord(-1000),
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Extension = mix.NewWord(1000)

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Extension
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestINCi(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		index       int
		expected    mix.Address
	}{
		{
			mix.Instruction{OpCode: mix.INC1, FieldSpec: 0, Address: mix.NewAddress(2000)},
			1,
			mix.NewAddress(3000),
		},
		{
			mix.Instruction{OpCode: mix.INC2, FieldSpec: 0, Address: mix.NewAddress(-2000)},
			2,
			mix.NewAddress(-1000),
		},
		{
			mix.Instruction{OpCode: mix.INC3, FieldSpec: 0, Address: mix.NewAddress(100)},
			3,
			mix.NewAddress(1100),
		},
		{
			mix.Instruction{OpCode: mix.INC4, FieldSpec: 0, Address: mix.NewAddress(750)},
			4,
			mix.NewAddress(1750),
		},
		{
			mix.Instruction{OpCode: mix.INC5, FieldSpec: 0, Address: mix.NewAddress(0)},
			5,
			mix.NewAddress(1000),
		},
		{
			mix.Instruction{OpCode: mix.INC6, FieldSpec: 0, Address: mix.NewAddress(-1000)},
			6,
			mix.NewAddress(0),
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Index[test.index-1] = mix.NewAddress(1000)

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Index[test.index-1]
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestDECA(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.DECA, FieldSpec: 1, Address: mix.NewAddress(2000)},
			mix.NewWord(-1000),
		},
		{
			mix.Instruction{OpCode: mix.DECA, FieldSpec: 1, Address: mix.NewAddress(-2000)},
			mix.NewWord(3000),
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Accumulator = mix.NewWord(1000)

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Accumulator
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestDECX(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		expected    mix.Word
	}{
		{
			mix.Instruction{OpCode: mix.DECX, FieldSpec: 1, Address: mix.NewAddress(2000)},
			mix.NewWord(-1000),
		},
		{
			mix.Instruction{OpCode: mix.DECX, FieldSpec: 1, Address: mix.NewAddress(-2000)},
			mix.NewWord(3000),
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Extension = mix.NewWord(1000)

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Extension
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}

func TestDECi(t *testing.T) {
	tests := []struct {
		instruction mix.Instruction
		index       int
		expected    mix.Address
	}{
		{
			mix.Instruction{OpCode: mix.DEC1, FieldSpec: 1, Address: mix.NewAddress(2000)},
			1,
			mix.NewAddress(-1000),
		},
		{
			mix.Instruction{OpCode: mix.DEC2, FieldSpec: 1, Address: mix.NewAddress(-2000)},
			2,
			mix.NewAddress(3000),
		},
		{
			mix.Instruction{OpCode: mix.DEC3, FieldSpec: 1, Address: mix.NewAddress(100)},
			3,
			mix.NewAddress(900),
		},
		{
			mix.Instruction{OpCode: mix.DEC4, FieldSpec: 1, Address: mix.NewAddress(750)},
			4,
			mix.NewAddress(250),
		},
		{
			mix.Instruction{OpCode: mix.DEC5, FieldSpec: 1, Address: mix.NewAddress(0)},
			5,
			mix.NewAddress(1000),
		},
		{
			mix.Instruction{OpCode: mix.DEC6, FieldSpec: 1, Address: mix.NewAddress(1000)},
			6,
			mix.NewAddress(0),
		},
	}

	for _, test := range tests {
		computer := NewComputer()
		computer.Index[test.index-1] = mix.NewAddress(1000)

		operation := NewOperation(test.instruction)
		operation.Execute(computer)

		actual := computer.Index[test.index-1]
		if actual != test.expected {
			t.Errorf("Instruction: %+v: expected %+v, actual %+v", test.instruction, test.expected, actual)
		}
	}
}
