package stir

import (
	"encoding/binary"
	"io"
	"log"
	"os"

	"jonnystoten.com/mixologist/mix"
)

var wordSizeOnDisk = binary.Size(&mix.Word{})

type ioMessage struct {
	op mix.OpCode
	m  int
}

type IODevice interface {
	Start()
	Channel() chan<- ioMessage
	BlockSize() int
	Busy() bool
}

type TapeUnit struct {
	computer *Computer
	filename string
	position int
	busy     bool
	ch       chan ioMessage
}

func NewTapeUnit(computer *Computer, filename string) *TapeUnit {
	return &TapeUnit{
		computer: computer,
		filename: filename,
		ch:       make(chan ioMessage)}
}

func (t *TapeUnit) Start() {
	go func() {
		for message := range t.ch {
			t.busy = true

			var err error
			switch message.op {
			case mix.IN:
				log.Println("read recv")
				err = read(t, message.m)
			case mix.OUT:
				log.Println("write recv")
				err = write(t, message.m)
			case mix.IOC:
				log.Println("control recv")
				err = control(t, message.m)
			}
			if err != nil {
				panic(err)
			}
		}
	}()
}

func (t *TapeUnit) Channel() chan<- ioMessage {
	return t.ch
}

func (t *TapeUnit) Busy() bool {
	return t.busy
}

func (t *TapeUnit) BlockSize() int {
	return 100
}

func read(t *TapeUnit, address int) error {
	defer func() {
		t.computer.IOWaitGroup.Done()
		t.busy = false
	}()

	file, err := os.Open(t.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Seek(int64(t.position*wordSizeOnDisk), io.SeekStart)
	words := make([]mix.Word, t.BlockSize())
	err = binary.Read(file, binary.LittleEndian, words)
	if err != nil {
		return err
	}

	for i := 0; i < t.BlockSize(); i++ {
		t.computer.Memory[address+i] = words[i]
	}

	t.position += t.BlockSize()
	return nil
}

func write(t *TapeUnit, address int) error {
	defer func() {
		t.computer.IOWaitGroup.Done()
		t.busy = false
	}()

	file, err := os.OpenFile(t.filename, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	words := make([]mix.Word, t.BlockSize())
	for i := 0; i < t.BlockSize(); i++ {
		words[i] = t.computer.Memory[address+i]
	}

	file.Seek(int64(t.position*wordSizeOnDisk), io.SeekStart)
	err = binary.Write(file, binary.LittleEndian, words)
	if err != nil {
		return err
	}

	t.position += t.BlockSize()
	return nil
}

func control(t *TapeUnit, m int) error {
	defer func() {
		t.computer.IOWaitGroup.Done()
		t.busy = false
	}()

	if m == 0 {
		t.position = 0
		return nil
	}

	delta := m * t.BlockSize()
	newPos := t.position + delta
	if newPos < 0 {
		t.position = 0
	} else {
		t.position = newPos
	}

	return nil
}
