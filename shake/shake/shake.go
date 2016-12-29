package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"jonnystoten.com/mixologist/shake"
)

func main() {
	log.Println("SHAKE")
	log.Println("==========")
	log.Println("LEX:")
	lex()
	log.Println()

	log.Println("PARSE:")
	prog := parse()
	log.Println()

	log.Println("ASSEMBLE:")
	words, err := shake.Assemble(prog)
	if err != nil {
		log.Fatalln(err)
	}

	for _, word := range words {
		err = binary.Write(os.Stdout, binary.LittleEndian, word)
		if err != nil {
			log.Fatalln(err)
			break
		}
	}

	log.Println("done!")
	log.Println("==========")
	log.Println()
}

func lex() {
	file, err := os.Open("loading.mixal")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := shake.NewScanner(file)

	var debug string
	for {
		tok, lit := scanner.Scan()

		if tok == shake.EOF {
			log.Println("[EOF]")
			debug = ""
			break
		}

		if tok == shake.ILLEGAL {
			log.Println("ERROR", lit)
			debug = ""
			break
		}

		if tok == shake.EOL {
			log.Printf("%v[EOL]", debug)
			debug = ""
		} else {
			debug += fmt.Sprintf("[%v]", lit)
		}
	}
}

func parse() *shake.Program {
	file, err := os.Open("loading.mixal")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	parser := shake.NewParser(file)

	prog, err := parser.Parse()
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	log.Printf("%+v", prog)

	return prog
}
