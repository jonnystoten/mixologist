package stir

import (
	"testing"

	"jonnystoten.com/mixologist/mix"
)

func TestShifts(t *testing.T) {
	computer := NewComputer()
	computer.Accumulator = mix.Word{Sign: mix.Positive, Bytes: [5]byte{1, 2, 3, 4, 5}}
	computer.Extension = mix.Word{Sign: mix.Negative, Bytes: [5]byte{6, 7, 8, 9, 10}}

	operation := NewOperation(mix.Instruction{OpCode: mix.SRAX, FieldSpec: 3, Address: mix.NewAddress(1)})
	operation.Execute(computer)

	expectedA := mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 1, 2, 3, 4}}
	expectedX := mix.Word{Sign: mix.Negative, Bytes: [5]byte{5, 6, 7, 8, 9}}
	if computer.Accumulator != expectedA || computer.Extension != expectedX {
		t.Errorf("After SRAX 1 expected rA = %v and rX = %v, got rA = %v and rX = %v", expectedA, expectedX, computer.Accumulator, computer.Extension)
	}

	operation = NewOperation(mix.Instruction{OpCode: mix.SLA, FieldSpec: 0, Address: mix.NewAddress(2)})
	operation.Execute(computer)

	expectedA = mix.Word{Sign: mix.Positive, Bytes: [5]byte{2, 3, 4, 0, 0}}
	expectedX = mix.Word{Sign: mix.Negative, Bytes: [5]byte{5, 6, 7, 8, 9}}
	if computer.Accumulator != expectedA || computer.Extension != expectedX {
		t.Errorf("After SLA 2 expected rA = %v and rX = %v, got rA = %v and rX = %v", expectedA, expectedX, computer.Accumulator, computer.Extension)
	}

	operation = NewOperation(mix.Instruction{OpCode: mix.SRC, FieldSpec: 5, Address: mix.NewAddress(4)})
	operation.Execute(computer)

	expectedA = mix.Word{Sign: mix.Positive, Bytes: [5]byte{6, 7, 8, 9, 2}}
	expectedX = mix.Word{Sign: mix.Negative, Bytes: [5]byte{3, 4, 0, 0, 5}}
	if computer.Accumulator != expectedA || computer.Extension != expectedX {
		t.Errorf("After SRC 4 expected rA = %v and rX = %v, got rA = %v and rX = %v", expectedA, expectedX, computer.Accumulator, computer.Extension)
	}

	operation = NewOperation(mix.Instruction{OpCode: mix.SRA, FieldSpec: 1, Address: mix.NewAddress(2)})
	operation.Execute(computer)

	expectedA = mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 6, 7, 8}}
	expectedX = mix.Word{Sign: mix.Negative, Bytes: [5]byte{3, 4, 0, 0, 5}}
	if computer.Accumulator != expectedA || computer.Extension != expectedX {
		t.Errorf("After SRA 2 expected rA = %v and rX = %v, got rA = %v and rX = %v", expectedA, expectedX, computer.Accumulator, computer.Extension)
	}

	operation = NewOperation(mix.Instruction{OpCode: mix.SLC, FieldSpec: 4, Address: mix.NewAddress(501)})
	operation.Execute(computer)

	expectedA = mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 6, 7, 8, 3}}
	expectedX = mix.Word{Sign: mix.Negative, Bytes: [5]byte{4, 0, 0, 5, 0}}
	if computer.Accumulator != expectedA || computer.Extension != expectedX {
		t.Errorf("After SLC 501 expected rA = %v and rX = %v, got rA = %v and rX = %v", expectedA, expectedX, computer.Accumulator, computer.Extension)
	}

	operation = NewOperation(mix.Instruction{OpCode: mix.SLAX, FieldSpec: 2, Address: mix.NewAddress(100)})
	operation.Execute(computer)

	expectedA = mix.Word{Sign: mix.Positive, Bytes: [5]byte{0, 0, 0, 0, 0}}
	expectedX = mix.Word{Sign: mix.Negative, Bytes: [5]byte{0, 0, 0, 0, 0}}
	if computer.Accumulator != expectedA || computer.Extension != expectedX {
		t.Errorf("After SLAX 100 expected rA = %v and rX = %v, got rA = %v and rX = %v", expectedA, expectedX, computer.Accumulator, computer.Extension)
	}
}
