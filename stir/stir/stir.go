package main

import (
	"encoding/binary"
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

	computer.Run()

	// for _, word := range words {
	// 	log.Printf("%+v\n", word)
	// }

	log.Printf("rA: %v", computer.Accumulator)
	log.Printf("rX: %v", computer.Extension)
	log.Printf("rI1: %v", computer.Index[0])
	log.Printf("rI2: %v", computer.Index[1])
	log.Printf("rI3: %v", computer.Index[2])
	log.Printf("rI4: %v", computer.Index[3])
	log.Printf("rI5: %v", computer.Index[4])
	log.Printf("rI6: %v", computer.Index[5])
	log.Printf("rJ: %v", computer.JumpAddress)
	log.Printf("Overflow: %v", computer.Overflow)
	log.Printf("Comparison: %v", computer.Comparison)

	log.Println("done!")
	log.Println("==========")
	log.Println()
}
