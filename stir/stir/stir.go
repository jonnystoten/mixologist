package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"io"
	"log"
	"os"

	"jonnystoten.com/mixologist/mix"
	"jonnystoten.com/mixologist/mix/garnish"
	"jonnystoten.com/mixologist/stir"
)

func main() {
	interactive := flag.Bool("interactive", false, "whether to launch in interactive mode")
	format := flag.String("format", "deck", "the input format")
	flag.Parse()

	log.Println("STIR")
	log.Println("==========")

	computer := stir.NewComputer()

	switch *format {
	case "deck":
		instruction := mix.Instruction{OpCode: mix.IN, FieldSpec: 16, Address: mix.NewAddress(0)}
		operation := stir.InputOutputOp{Instruction: instruction}
		operation.Execute(computer)
		computer.IOWaitGroup.Wait()
		computer.ProgramCounter = 0
		computer.JumpAddress = mix.NewAddress(0)
	case "binary":
		words := make([]mix.Word, 4000)
		var start uint16
		err := binary.Read(os.Stdin, binary.LittleEndian, &start)
		if err != nil {
			log.Fatalln(err)
		}

		for {
			var loc uint16
			err := binary.Read(os.Stdin, binary.LittleEndian, &loc)
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalln(err)
				return
			}
			word := mix.Word{}
			err = binary.Read(os.Stdin, binary.LittleEndian, &word)
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalln(err)
				return
			}
			words[loc] = word
		}

		computer.Memory = words
		computer.ProgramCounter = int(start)
	}

	log.Printf("Ready to go... (starting at %v)", computer.ProgramCounter)

	broke := false
	if *interactive {
		computer.RunInteractive(func() {
			if computer.ProgramCounter == 12 && computer.Accumulator.Value() == 0 {
				// if loading finished
				broke = true
			}
			if broke {
				PrintState(computer)
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
			}
		})
	} else {
		computer.Run()
	}

	log.Println("done, waiting for remaining IO...")
	computer.IOWaitGroup.Wait()

	log.Println("done!")
	log.Println("==========")
	PrintState(computer)
}

func PrintState(computer *stir.Computer) {
	log.Printf("PC: %v", computer.ProgramCounter)
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
