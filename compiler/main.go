package main

import (
	"compiler/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the CLI!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}