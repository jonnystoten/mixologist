package stir

import (
	"encoding/binary"
	"io"
	"log"
	"os"

	"jonnystoten.com/mixologist/mix"
)

type ioCom struct {
	rwc     string
	address int
}

type IODevice interface {
	Start()
	ReadBlock(address int)
	WriteBlock(address int)
	Control(m int)
	BlockSize() int
}

type TapeUnit struct {
	computer *Computer
	filename string
	position int
	ch       chan ioCom
}

func NewTapeUnit(computer *Computer, filename string) *TapeUnit {
	return &TapeUnit{
		computer: computer,
		filename: filename,
		ch:       make(chan ioCom)}
}

func (t *TapeUnit) Start() {
	go func() {
		for {
			select {
			case com := <-t.ch:
				//time.Sleep(1 * time.Second)
				switch com.rwc {
				case "r":
					log.Println("read recv")
					err := read(t, com.address)
					if err != nil {
						panic(err)
					}
				case "w":
					log.Println("write recv")
					err := write(t, com.address)
					if err != nil {
						panic(err)
					}
				case "c":
					log.Println("control recv")
					err := control(t, com.address)
					if err != nil {
						panic(err)
					}
				}
			default:
			}
		}
	}()
}

func (t *TapeUnit) ReadBlock(address int) {
	t.computer.IOWaitGroup.Add(1)
	log.Println("blocked trying to read")
	t.ch <- ioCom{"r", address}
	log.Println("unblocked, ready to read")
}

func (t *TapeUnit) WriteBlock(address int) {
	t.computer.IOWaitGroup.Add(1)
	log.Println("blocked trying to write")
	t.ch <- ioCom{"w", address}
	log.Println("unblocked, ready to write")
}

func (t *TapeUnit) Control(m int) {
	t.computer.IOWaitGroup.Add(1)
	log.Println("blocked trying to control")
	t.ch <- ioCom{"c", m}
	log.Println("unblocked, ready to control")
}

func (t *TapeUnit) BlockSize() int {
	return 100
}

func read(t *TapeUnit, address int) error {
	defer t.computer.IOWaitGroup.Done()
	file, err := os.Open(t.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Seek(int64(t.position), io.SeekStart)
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
	defer t.computer.IOWaitGroup.Done()
	file, err := os.OpenFile(t.filename, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	words := make([]mix.Word, t.BlockSize())
	for i := 0; i < t.BlockSize(); i++ {
		words[i] = t.computer.Memory[address+i]
	}

	file.Seek(int64(t.position), io.SeekStart)
	err = binary.Write(file, binary.LittleEndian, words)
	if err != nil {
		return err
	}

	t.position += t.BlockSize()
	return nil
}

func control(t *TapeUnit, m int) error {
	defer t.computer.IOWaitGroup.Done()

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
