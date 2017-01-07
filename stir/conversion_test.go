package stir

import (
	"testing"

	"jonnystoten.com/mixologist/mix"
)

func TestNUMAndCHAR(t *testing.T) {
	computer := NewComputer()
	computer.Accumulator = mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 31, 32, 39}}
	computer.Extension = mix.Word{Sign: mix.Positive, Bytes: [5]byte{37, 57, 47, 30, 30}}

	operation := NewOperation(mix.Instruction{OpCode: mix.NUM, FieldSpec: 0, Address: mix.NewAddress(0)})
	operation.Execute(computer)

	if computer.Accumulator != mix.NewWord(-12977700) || computer.Extension != (mix.Word{Sign: mix.Positive, Bytes: [5]byte{37, 57, 47, 30, 30}}) {
		t.Errorf("After NUM expected rA = -12977700 and rX unchanged, got rA = %v and rX = %v", computer.Accumulator, computer.Extension)
	}

	operation = NewOperation(mix.Instruction{OpCode: mix.INCA, FieldSpec: 0, Address: mix.NewAddress(1)})
	operation.Execute(computer)

	if computer.Accumulator != mix.NewWord(-12977699) || computer.Extension != (mix.Word{Sign: mix.Positive, Bytes: [5]byte{37, 57, 47, 30, 30}}) {
		t.Errorf("After INCA expected rA = -12977699 and rX unchanged, got rA = %v and rX = %v", computer.Accumulator, computer.Extension)
	}

	operation = NewOperation(mix.Instruction{OpCode: mix.CHAR, FieldSpec: 1, Address: mix.NewAddress(0)})
	operation.Execute(computer)

	if computer.Accumulator != (mix.Word{Sign: mix.Negative, Bytes: [5]byte{30, 30, 31, 32, 39}}) || computer.Extension != (mix.Word{Sign: mix.Positive, Bytes: [5]byte{37, 37, 36, 39, 39}}) {
		t.Errorf("After CHAR expected rA = [30, 30, 31, 32, 39] and rX [37, 37, 36, 39, 39], got rA = %v and rX = %v", computer.Accumulator, computer.Extension)
	}
}
