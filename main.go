package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/benmch/chip-8-go/internal/evaluator"
	"github.com/benmch/chip-8-go/internal/lexer"
	"github.com/benmch/chip-8-go/internal/object"
	"github.com/benmch/chip-8-go/internal/parser"
	"github.com/benmch/chip-8-go/internal/repl"
)

func main() {
	if len(os.Args) == 2 {
		file, err := os.ReadFile(os.Args[1])

		if err != nil {
			panic("Unable to read file")
		}

		input := string(file)
		l := lexer.New(input)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) > 0 {
			fmt.Printf("Parsing error!\n")

			for _, err := range p.Errors() {
				fmt.Printf("Error: %s\n", err)
			}

			return
		}

		env := object.NewEnvironment()

		val := evaluator.Eval(program, env)

		if val.Type() == object.ERROR_OBJ {
			fmt.Printf("Error: %s\n", val.Inspect())
		}

		return
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
