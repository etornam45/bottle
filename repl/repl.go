package repl

import (
	"bufio"
	"fmt"
	"io"

	"etornam.ben/polygon/lexer"
  "etornam.ben/polygon/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

    if line == "exit" {
      println("Exiting interpreter..... bye")
      break
    }

		l := lexer.New(line)
    p := parser.New(l)
    
    program := p.ParseProgram()

		fmt.Println(program.String())
	}
}
