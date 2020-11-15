package main

import (
	"MonkeyLanguage/src/repl"
	"fmt"
	"os"
	user2 "os/user"
)

func main(){
	user, err := user2.Current()
	if err != nil{
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey Programming Language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
