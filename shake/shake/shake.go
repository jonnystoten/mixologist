package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"jonnystoten.com/mixologist/mix"
	"jonnystoten.com/mixologist/mix/garnish"
	"jonnystoten.com/mixologist/shake"
)

func main() {
	filename := flag.String("input", "", "the input file")
	format := flag.String("format", "binary", "the output format")
	flag.Parse()

	log.Println("SHAKE")
	log.Println("==========")
	log.Println("LEX:")
	lex(*filename)
	log.Println()

	log.Println("PARSE:")
	prog := parse(*filename)
	log.Println()

	log.Println("ASSEMBLE:")
	assembler := shake.NewAssembler()
	err := assembler.Assemble(prog)
	if err != nil {
		log.Fatalln(err)
	}

	memLocs := []int{}
	for loc := range assembler.Words {
		memLocs = append(memLocs, loc)
	}
	sort.Ints(memLocs)

	for _, loc := range memLocs {
		word := assembler.Words[loc]
		if word != (mix.Word{}) {
			log.Printf("%04v: %v", loc, garnish.SprintWord(word))
		}
	}

	switch *format {
	case "binary":
		binary.Write(os.Stdout, binary.LittleEndian, uint16(assembler.ProgramStart))
		for _, loc := range memLocs {
			word := assembler.Words[loc]
			err = binary.Write(os.Stdout, binary.LittleEndian, uint16(loc))
			err = binary.Write(os.Stdout, binary.LittleEndian, word)
			if err != nil {
				log.Fatalln(err)
				break
			}
		}
	case "raw":
		memLocs := []int{}
		for loc := range assembler.Words {
			memLocs = append(memLocs, loc)
		}
		sort.Ints(memLocs)

		buf := bytes.Buffer{}
		for _, loc := range memLocs {
			word := assembler.Words[loc]
			code := mix.WordToCharCodeString(word)
			buf.WriteString(code)
			if buf.Len() == 80 {
				str := buf.String()
				buf = bytes.Buffer{}
				_, err := io.WriteString(os.Stdout, str+"\n")
				if err != nil {
					log.Fatalln(err)
					break
				}
			}
		}
		str := buf.String()
		_, err := io.WriteString(os.Stdout, str+"\n")
		if err != nil {
			log.Fatalln(err)
			break
		}
	case "deck":
		groups := makeGroups(memLocs)
		fmt.Fprintf(os.Stdout, "%v\n", loader())
		for _, group := range groups {
			buf := &bytes.Buffer{}
			fmt.Fprintf(buf, "SHAKE%v%04v", len(group), group[0])
			for _, loc := range group {
				word := assembler.Words[loc]
				value := word.Value()
				if value >= 0 {
					fmt.Fprintf(buf, "%010v", value)
				} else {
					value = -value
					fmt.Fprintf(buf, "%09v", value/10)
					lsb := value % 10
					char := mix.CharCodes.GetChar(byte(lsb + 10))
					buf.WriteRune(char)
				}
			}
			str := buf.String()
			fmt.Fprintf(os.Stdout, "%v\n", str)
		}

		fmt.Fprintf(os.Stdout, "TRANS0%04v\n", assembler.ProgramStart)
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

func makeGroups(memLocs []int) (groups [][]int) {
	var group []int
	lastLoc := memLocs[0] - 1
	for _, loc := range memLocs {
		log.Printf("loc = %v, last = %v", loc, lastLoc)
		if loc != lastLoc+1 || len(group) == 7 { // new card
			groups = append(groups, group)
			group = []int{}
		}
		group = append(group, loc)
		lastLoc = loc
	}
	groups = append(groups, group)

	for _, x := range groups {
		log.Print("(")
		for _, y := range x {
			log.Print(y)
		}
		log.Print(")")
	}
	return
}

func loader() string {
	return ` O O6 2 O6    I C O4 3 EH A  F F CF    E   EU 3 IH 0 EB   EJ  CA. 2 EU   EH Z EA
   EU 5A-H 0 EB  C U 4AEH 5AEN    E  CLU  ABG 2 EH Z EB J B. A  9    A    0`
}
