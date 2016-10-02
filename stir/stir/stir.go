package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"jonnystoten.com/mixologist/mix"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	words := []*mix.Word{}

	for {
		word := &mix.Word{}
		err = binary.Read(file, binary.LittleEndian, word)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		words = append(words, word)
	}

	for _, word := range words {
		fmt.Printf("%+v\n", word)
	}
}
