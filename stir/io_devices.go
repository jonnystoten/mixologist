package stir

import (
	"encoding/binary"
	"fmt"
	"io"
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
	Busy() bool
	SetBusy()
	SetReady()
	BlockSize() int
	Computer() *Computer
	Channel() chan<- ioMessage
}

type InputDevice interface {
	IODevice
	Read(words []mix.Word) error
}

type OutputDevice interface {
	IODevice
	Write(words []mix.Word) error
}

type ControllableDevice interface {
	IODevice
	Control(m int) error
}

func ioAction(d IODevice, message ioMessage) {
	d.SetBusy()
	defer d.SetReady()

	var err error
	switch message.op {
	case mix.IN:
		inputD, ok := d.(InputDevice)
		if !ok {
			err = fmt.Errorf("IN used on non-input device")
		} else {
			err = read(inputD, message.m, inputD.Computer())
		}
	case mix.OUT:
		outputD, ok := d.(OutputDevice)
		if !ok {
			err = fmt.Errorf("IN used on non-input device")
		} else {
			err = write(outputD, message.m, outputD.Computer())
		}
	case mix.IOC:
		controlD, ok := d.(ControllableDevice)
		if !ok {
			err = fmt.Errorf("IOC used on non-controllable device")
		} else {
			err = controlD.Control(message.m)
		}
	}
	if err != nil {
		panic(err)
	}
}

func read(d InputDevice, address int, c *Computer) error {
	words := make([]mix.Word, d.BlockSize())
	err := d.Read(words)
	if err != nil {
		return err
	}

	for i := 0; i < d.BlockSize(); i++ {
		c.Memory[address+i] = words[i]
	}

	return nil
}

func write(d OutputDevice, address int, c *Computer) error {
	words := make([]mix.Word, d.BlockSize())
	for i := 0; i < d.BlockSize(); i++ {
		words[i] = c.Memory[address+i]
	}

	err := d.Write(words)
	if err != nil {
		return err
	}

	return nil
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
			ioAction(t, message)
		}
	}()
}

func (t *TapeUnit) Channel() chan<- ioMessage {
	return t.ch
}

func (t *TapeUnit) Computer() *Computer {
	return t.computer
}

func (t *TapeUnit) Busy() bool {
	return t.busy
}

func (t *TapeUnit) SetBusy() {
	t.busy = true
}

func (t *TapeUnit) SetReady() {
	t.busy = false
	t.computer.IOWaitGroup.Done()
}

func (t *TapeUnit) BlockSize() int {
	return 100
}

func (t *TapeUnit) Read(words []mix.Word) error {
	file, err := os.Open(t.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Seek(int64(t.position*wordSizeOnDisk), io.SeekStart)

	err = binary.Read(file, binary.LittleEndian, words)
	if err != nil {
		return err
	}

	t.position += t.BlockSize()
	return nil
}

func (t *TapeUnit) Write(words []mix.Word) error {
	file, err := os.OpenFile(t.filename, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Seek(int64(t.position*wordSizeOnDisk), io.SeekStart)
	err = binary.Write(file, binary.LittleEndian, words)
	if err != nil {
		return err
	}

	t.position += t.BlockSize()
	return nil
}

func (t *TapeUnit) Control(m int) error {
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
