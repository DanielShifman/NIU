package main

import (
	"Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

const SPLASH = `
 ___ _  _____  __ __
|   | ||_   _||  |  |
| | | | _| |_ |  |  |
|_|___||_____||_____|REPL
`

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(SPLASH)
	fmt.Printf("\n\nHello %s! Welcome to the NIU REPL!\n", usr.Username)
	fmt.Printf("Enter commands:\n")
	repl.Start(os.Stdin, os.Stdout)
}
