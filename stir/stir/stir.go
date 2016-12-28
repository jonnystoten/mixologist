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

	log.Printf("A: %v", computer.Accumulator)
	log.Printf("X: %v", computer.Extension)
	log.Printf("I1: %v", computer.Index[0])
	log.Printf("I2: %v", computer.Index[1])
	log.Printf("I3: %v", computer.Index[2])
	log.Printf("I4: %v", computer.Index[3])
	log.Printf("I5: %v", computer.Index[4])
	log.Printf("I6: %v", computer.Index[5])

	log.Println("done!")
	log.Println("==========")
	log.Println()
}
