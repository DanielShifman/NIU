package main

import (
	"Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the NIU REPL!\n", usr.Username)
	fmt.Printf("Enter commands:\n")
	repl.Start(os.Stdin, os.Stdout)
}
