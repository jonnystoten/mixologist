package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"

	"jonnystoten.com/mixologist/mix"
)

func main() {
	log.Println("STIR")
	log.Println("==========")
	words := []*mix.Word{}

	for {
		word := &mix.Word{}
		err := binary.Read(os.Stdin, binary.LittleEndian, word)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
			return
		}
		words = append(words, word)
	}

	for _, word := range words {
		log.Printf("%+v\n", word)
	}

	log.Println("done!")
	log.Println("==========")
	log.Println()
}
