package main

import (
	"fmt"
	"os"
	"os/user"

	"etornam.ben/polygon/exec"
	"etornam.ben/polygon/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		println(err)
	}
	if len(os.Args) <= 1 {
		fmt.Printf("Hello %s, this is the Bottle language!\n", user.Username)
		repl.Start(os.Stdin, os.Stdout)
	} else {
		filename := os.Args[1]
		code, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		exec.Run(string(code))
	}
}
