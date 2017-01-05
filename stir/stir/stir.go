package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"jonnystoten.com/mixologist/mix"
	"jonnystoten.com/mixologist/stir"
)

func main() {
	log.Println("STIR")
	log.Println("==========")
	words := []mix.Word{}
	var start uint16
	err := binary.Read(os.Stdin, binary.LittleEndian, &start)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		word := mix.Word{}
		err := binary.Read(os.Stdin, binary.LittleEndian, &word)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
			return
		}
		words = append(words, word)
	}

	computer := stir.NewComputer()
	copy(computer.Memory[:], words)
	computer.ProgramCounter = int(start)

	log.Printf("Ready to GO... (starting at %v)", computer.ProgramCounter)

	computer.Run()

	log.Println("done, waiting for remaining IO...")
	computer.IOWaitGroup.Wait()

	log.Println("done!")
	log.Println("==========")

	log.Printf("rA: %v", sprintWord(computer.Accumulator))
	log.Printf("rX: %v", sprintWord(computer.Extension))
	log.Printf("rI1: %v", sprintAddress(computer.Index[0]))
	log.Printf("rI2: %v", sprintAddress(computer.Index[1]))
	log.Printf("rI3: %v", sprintAddress(computer.Index[2]))
	log.Printf("rI4: %v", sprintAddress(computer.Index[3]))
	log.Printf("rI5: %v", sprintAddress(computer.Index[4]))
	log.Printf("rI6: %v", sprintAddress(computer.Index[5]))
	log.Printf("rJ: %v", sprintAddress(computer.JumpAddress))
	var overflow string
	switch computer.Overflow {
	case true:
		overflow = "✔"
	case false:
		overflow = "✘"
	}
	log.Printf("Overflow: %v", overflow)
	var comparison string
	switch computer.Comparison {
	case mix.Less:
		comparison = "L"
	case mix.Equal:
		comparison = "E"
	case mix.Greater:
		comparison = "G"
	}
	log.Printf("Comparison: %v", comparison)

	for i, word := range computer.Memory {
		if word != (mix.Word{}) {
			log.Printf("M %04v: %v", i, sprintWord(word))
		}
	}

	log.Println()
}

func sprintWord(w mix.Word) string {
	return sprintSignedBytes(w.Sign, w.Bytes[:], w.Value())
}

func sprintAddress(a mix.Address) string {
	return sprintSignedBytes(a.Sign, a.Bytes[:], a.Value())
}

func sprintSignedBytes(sign mix.Sign, bytes []byte, value int) string {
	if value < 0 {
		value = -value
	}
	format := fmt.Sprintf("%%v %%v (%%0%vv)", len(bytes)*2)
	return fmt.Sprintf(format, sprintSign(sign), sprintBytes(bytes), value)
}

func sprintSign(sign mix.Sign) string {
	switch sign {
	case mix.Positive:
		return "+"
	case mix.Negative:
		return "-"
	default:
		panic("invalid value for Sign")
	}
}

func sprintBytes(bytes []byte) string {
	var res string
	for _, b := range bytes {
		res += fmt.Sprintf("%02v ", b)
	}

	return res
}
