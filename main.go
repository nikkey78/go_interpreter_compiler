package main

import (
	"console"
	"fmt"
	"monkey2/repl"
	"os"
	"os/user"
)

func main() {

	defer console.GreenReset()()

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Feel free type in commnands")
	repl.Start(os.Stdin, os.Stdout)
}
