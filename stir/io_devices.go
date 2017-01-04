package stir

import (
	"encoding/binary"
	"log"
	"os"
	"sync"

	"jonnystoten.com/mixologist/mix"
)

type IODevice interface {
	Lock()
	ReadBlock(address int, c *Computer) error
	WriteBlock(address int, c *Computer) error
	Control(m int) error
	BlockSize() int
}

type TapeUnit struct {
	filename string
	mut      sync.Mutex
}

func (t TapeUnit) ReadBlock(address int, c *Computer) error {
	defer func() {
		log.Println("IN about to unlock")
		t.mut.Unlock()
	}()

	file, err := os.Open(t.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	words := make([]mix.Word, t.BlockSize())
	return binary.Read(file, binary.LittleEndian, words)
}

func (t TapeUnit) WriteBlock(address int, c *Computer) error {
	defer func() {
		log.Println("OUT about to unlock")
		t.mut.Unlock()
	}()

	file, err := os.Open(t.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	words := make([]mix.Word, t.BlockSize())
	for i := 0; i < t.BlockSize(); i++ {
		word := c.Memory[address+i]
		words = append(words, word)
	}
	return binary.Write(file, binary.LittleEndian, words)
}

func (t TapeUnit) Control(m int) error {
	return nil
}

func (t TapeUnit) BlockSize() int {
	return 100
}

func (t TapeUnit) Lock() {
	t.mut.Lock()
}
