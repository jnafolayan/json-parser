package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jnafolayan/json-parser/lexer"
	"github.com/jnafolayan/json-parser/parser"
)

func main() {
	var lex *lexer.Lexer

	raw := flag.String("raw", "", "raw json string")
	flag.Parse()

	filename := flag.Arg(0)

	// str := "\t\n"
	// for _, ru := range str {
	// 	if !isValidEscapeSequence(ru) {
	// 		fmt.Printf("illegal escape sequence: %x\n", string(ru))
	// 		break
	// 	}
	// 	if unicode.IsControl(ru) {
	// 		fmt.Printf("unescaped control character: %x\n", ru)
	// 	}

	// }

	if filename == "" && *raw == "" {
		fmt.Fprint(os.Stderr, "You must supply either a raw json string or file")
		return
	}

	if filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			return
		}
		lex = lexer.NewLexer(f)
	} else {
		lex = lexer.FromString(*raw)
	}

	parser := parser.NewParser(lex)
	parseErr := parser.Parse()
	if parseErr == nil {
		fmt.Println("JSON is valid")
		os.Exit(0)
	} else {
		fmt.Fprintf(os.Stderr, "JSON is invalid: %s\n", parseErr)
		os.Exit(1)
	}
}

// func isValidEscapeSequence(r rune) bool {
// 	switch r {
// 	case '\b', '\f', '\n', '\r', '\t':
// 		return true
// 	default:
// 		return false
// 	}
// }
