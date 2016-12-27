package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"jonnystoten.com/mixologist/mix"
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
	instructions, err := shake.Assemble(prog)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", instructions)

	for _, instruction := range instructions {
		word := mix.Word{
			Sign: instruction.Address.Sign,
			Bytes: [5]byte{
				instruction.Address.Value[0],
				instruction.Address.Value[1],
				instruction.IndexSpec,
				instruction.FieldSpec,
				instruction.OpCode,
			},
		}
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

	log.Printf("%+v\n", prog)

	return prog
}
