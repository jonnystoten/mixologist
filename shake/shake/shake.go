package main

import (
	"fmt"
	"log"
	"os"

	"jonnystoten.com/mixologist/shake"
)

func main() {
	fmt.Println("LEX:")
	lex()
	fmt.Println()

	fmt.Println("PARSE:")
	prog := parse()
	fmt.Println()

	fmt.Println("ASSEMBLE:")
	instructions, err := shake.Assemble(prog)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", instructions)
}

func lex() {
	file, err := os.Open("programM.mixal")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := shake.NewScanner(file)

	for {
		tok, lit := scanner.Scan()
		if tok == shake.EOF {
			fmt.Println("[EOF]")
			break
		}

		if tok == shake.ILLEGAL {
			fmt.Println("ERROR", lit)
			break
		}

		if tok == shake.EOL {
			fmt.Println("[EOL]")
		} else {
			fmt.Printf("[%v]", lit)
		}
	}
}

func parse() *shake.Program {
	file, err := os.Open("programM.mixal")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	parser := shake.NewParser(file)

	prog, err := parser.Parse()
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Printf("%+v\n", prog)

	return prog
}
