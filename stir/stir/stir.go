package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"

	"jonnystoten.com/mixologist/mix"
	"jonnystoten.com/mixologist/mix/garnish"
	"jonnystoten.com/mixologist/stir"
)

func main() {
	// format := flag.String("format", "dump", "the output format")
	// flag.Parse()

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

	log.Printf("rA: %v", garnish.SprintWord(computer.Accumulator))
	log.Printf("rX: %v", garnish.SprintWord(computer.Extension))
	log.Printf("rI1: %v", garnish.SprintAddress(computer.Index[0]))
	log.Printf("rI2: %v", garnish.SprintAddress(computer.Index[1]))
	log.Printf("rI3: %v", garnish.SprintAddress(computer.Index[2]))
	log.Printf("rI4: %v", garnish.SprintAddress(computer.Index[3]))
	log.Printf("rI5: %v", garnish.SprintAddress(computer.Index[4]))
	log.Printf("rI6: %v", garnish.SprintAddress(computer.Index[5]))
	log.Printf("rJ: %v", garnish.SprintAddress(computer.JumpAddress))
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
			log.Printf("M %04v: %v", i, garnish.SprintWord(word))
		}
	}

	log.Println()
}
