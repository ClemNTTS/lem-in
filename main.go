package main

import (
	"lem-in/src"
	"os"
)

func main() {
	//Manage Args
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	src.Launch(args[0])	
}