package stir

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"bufio"

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

	// time.Sleep(2 * time.Second) // artificial slow down of I/O

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

	file.Seek(int64(t.position*t.BlockSize()*wordSizeOnDisk), io.SeekStart)

	err = binary.Read(file, binary.LittleEndian, words)
	if err != nil {
		return err
	}

	t.position++
	return nil
}

func (t *TapeUnit) Write(words []mix.Word) error {
	file, err := os.OpenFile(t.filename, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Seek(int64(t.position*t.BlockSize()*wordSizeOnDisk), io.SeekStart)
	err = binary.Write(file, binary.LittleEndian, words)
	if err != nil {
		return err
	}

	t.position++
	return nil
}

func (t *TapeUnit) Control(delta int) error {
	if delta == 0 {
		t.position = 0
		return nil
	}

	newPos := t.position + delta
	if newPos < 0 {
		t.position = 0
	} else {
		t.position = newPos
	}

	return nil
}

type DiskDrumUnit struct {
	computer *Computer
	filename string
	busy     bool
	ch       chan ioMessage
}

func NewDiskDrumUnit(computer *Computer, filename string) *DiskDrumUnit {
	return &DiskDrumUnit{
		computer: computer,
		filename: filename,
		ch:       make(chan ioMessage)}
}

func (dd *DiskDrumUnit) Start() {
	go func() {
		for message := range dd.ch {
			ioAction(dd, message)
		}
	}()
}

func (dd *DiskDrumUnit) Channel() chan<- ioMessage {
	return dd.ch
}

func (dd *DiskDrumUnit) Computer() *Computer {
	return dd.computer
}

func (dd *DiskDrumUnit) Busy() bool {
	return dd.busy
}

func (dd *DiskDrumUnit) SetBusy() {
	dd.busy = true
}

func (dd *DiskDrumUnit) SetReady() {
	dd.busy = false
	dd.computer.IOWaitGroup.Done()
}

func (dd *DiskDrumUnit) BlockSize() int {
	return 100
}

func (dd *DiskDrumUnit) Read(words []mix.Word) error {
	file, err := os.Open(dd.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	position := dd.computer.Extension.Value()
	file.Seek(int64(position*dd.BlockSize()*wordSizeOnDisk), io.SeekStart)

	err = binary.Read(file, binary.LittleEndian, words)
	if err != nil {
		return err
	}

	return nil
}

func (dd *DiskDrumUnit) Write(words []mix.Word) error {
	file, err := os.OpenFile(dd.filename, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	position := dd.computer.Extension.Value()
	file.Seek(int64(position*dd.BlockSize()*wordSizeOnDisk), io.SeekStart)

	err = binary.Write(file, binary.LittleEndian, words)
	if err != nil {
		return err
	}

	return nil
}

func (dd *DiskDrumUnit) Control(m int) error {
	// this should seek to new position, but we just
	// seek on every read/write anyway so nothing to do
	return nil
}

type CardReader struct {
	computer *Computer
	filename string
	position int
	busy     bool
	ch       chan ioMessage
}

func NewCardReader(computer *Computer, filename string) *CardReader {
	return &CardReader{
		computer: computer,
		filename: filename,
		ch:       make(chan ioMessage)}
}

func (cr *CardReader) Start() {
	go func() {
		for message := range cr.ch {
			ioAction(cr, message)
		}
	}()
}

func (cr *CardReader) Channel() chan<- ioMessage {
	return cr.ch
}

func (cr *CardReader) Computer() *Computer {
	return cr.computer
}

func (cr *CardReader) Busy() bool {
	return cr.busy
}

func (cr *CardReader) SetBusy() {
	cr.busy = true
}

func (cr *CardReader) SetReady() {
	cr.busy = false
	cr.computer.IOWaitGroup.Done()
}

func (cr *CardReader) BlockSize() int {
	return 16
}

func (cr *CardReader) Read(words []mix.Word) error {
	file, err := os.Open(cr.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; i <= cr.position; i++ {
		scanner.Scan()
	}

	line := scanner.Text()
	for i := 0; i < cr.BlockSize(); i++ {
		var block string
		if len(line) >= 5 {
			block, line = line[:5], line[5:]
		} else {
			block, line = spacePad(line), ""
		}
		words[i] = mix.NewWordFromCharCode(block)
	}

	cr.position++
	return nil
}

type CardWriter struct {
	computer *Computer
	filename string
	busy     bool
	ch       chan ioMessage
}

func NewCardWriter(computer *Computer, filename string) *CardWriter {
	return &CardWriter{
		computer: computer,
		filename: filename,
		ch:       make(chan ioMessage)}
}

func (cw *CardWriter) Start() {
	go func() {
		for message := range cw.ch {
			ioAction(cw, message)
		}
	}()
}

func (cw *CardWriter) Channel() chan<- ioMessage {
	return cw.ch
}

func (cw *CardWriter) Computer() *Computer {
	return cw.computer
}

func (cw *CardWriter) Busy() bool {
	return cw.busy
}

func (cw *CardWriter) SetBusy() {
	cw.busy = true
}

func (cw *CardWriter) SetReady() {
	cw.busy = false
	cw.computer.IOWaitGroup.Done()
}

func (cw *CardWriter) BlockSize() int {
	return 16
}

func (cw *CardWriter) Write(words []mix.Word) error {
	file, err := os.OpenFile(cw.filename, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := bytes.Buffer{}
	for _, word := range words {
		str := mix.WordToCharCodeString(word)
		buf.WriteString(str)
	}
	buf.WriteRune('\n')

	_, err = file.WriteString(buf.String())
	if err != nil {
		return err
	}

	return nil
}

func spacePad(str string) string {
	bytes := make([]byte, 5)
	for i := 0; i < 5; i++ {
		if len(str) > i {
			bytes[i] = str[i]
		} else {
			bytes[i] = '_'
		}
	}
	return string(bytes)
}
