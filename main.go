package main

import "fmt"

func main() {
	InitLinks()
	fmt.Println(InitFourmie().Name)
}

type Room struct {
	Name     string
	Occupied bool
	Links    []Room
}