package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"jonnystoten.com/mixologist/shake"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Usage: shake filename")
	}
	filename := args[0]

	log.Println("SHAKE")
	log.Println("==========")
	log.Println("LEX:")
	lex(filename)
	log.Println()

	log.Println("PARSE:")
	prog := parse(filename)
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

func lex(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := shake.NewScanner(file)

	var debug string
	for {
		lexeme := scanner.Scan()

		if lexeme.Tok == shake.EOF {
			log.Println("[EOF]")
			debug = ""
			break
		}

		if lexeme.Tok == shake.ILLEGAL {
			log.Printf("ERROR: %v (%v:%v)", lexeme.Lit, lexeme.Line, lexeme.Col)
			debug = ""
			break
		}

		if lexeme.Tok == shake.EOL {
			log.Printf("%v[EOL]", debug)
			debug = ""
		} else {
			debug += fmt.Sprintf("[%v]", lexeme.Lit)
		}
	}
}

func parse(filename string) *shake.Program {
	file, err := os.Open(filename)
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
