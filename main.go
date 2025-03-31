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
	filename := os.Args[1]
	code, err := os.ReadFile(filename)
	if len(os.Args) > 1 {
		exec.Run(string(code))
	} else {

		fmt.Printf("Hello %s, this is the Bottle language!\n", user.Username)
		repl.Start(os.Stdin, os.Stdout)
	}
}
