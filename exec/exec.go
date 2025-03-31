package exec

import (
	"etornam.ben/polygon/lexer"
	"etornam.ben/polygon/parser"
)

func Run(input string) {
  l := lexer.New(input)
  p := parser.New(l)

  program := p.ParseProgram()
  print(program.String())
}
