package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"regexp"

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

	for i, word := range assembler.Words {
		if word != (mix.Word{}) {
			log.Printf("%04v: %v", i, garnish.SprintWord(word))
		}
	}

	switch *format {
	case "binary":
		binary.Write(os.Stdout, binary.LittleEndian, uint16(assembler.ProgramStart))
		for _, word := range assembler.Words {
			err = binary.Write(os.Stdout, binary.LittleEndian, word)
			if err != nil {
				log.Fatalln(err)
				break
			}
		}
	case "raw":
		buf := bytes.Buffer{}
		for _, word := range assembler.Words {
			code := mix.WordToCharCodeString(word)
			buf.WriteString(code)
			if buf.Len() == 80 {
				str := buf.String()
				buf = bytes.Buffer{}
				emptyLine, _ := regexp.MatchString(`\s{80}`, str)
				if !emptyLine {
					_, err := io.WriteString(os.Stdout, str+"\n")
					if err != nil {
						log.Fatalln(err)
						break
					}
				}
			}
		}
	case "deck":
		// TODO: assemble only words from the source, then
		// group into card sized groups and output together
		buf := bytes.Buffer{}
		//buf.WriteString("SHAKE7") // TODO: don't hard-code 7
		for _, word := range assembler.Words {
			value := word.Value()
			if value >= 0 {
				fmt.Fprintf(&buf, "%010v", value)
			} else {
				value = -value
				fmt.Fprintf(&buf, "%09v", value/10)
				lsb := value % 10
				char := mix.CharCodes.GetChar(byte(lsb + 10))
				buf.WriteRune(char)
			}
			if buf.Len() == 80 {
				str := buf.String()
				buf = bytes.Buffer{}
				emptyLine, _ := regexp.MatchString(`0{80}`, str)
				if !emptyLine {
					_, err := fmt.Fprintf(os.Stdout, "%v\n", str)
					if err != nil {
						log.Fatalln(err)
						break
					}
				}
			}
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
