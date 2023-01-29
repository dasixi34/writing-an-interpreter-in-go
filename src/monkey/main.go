package main

import (
	"fmt"
	"monkey/cli"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		cli.Start(filename, os.Stdout)
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
