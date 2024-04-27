package main
 
import (
	"fmt"
	"os"
	"os/user"

	"etornam.ben/polygon/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		println(err)
	}

	fmt.Printf("Hello %s, this is the Polygon language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout) 
}
