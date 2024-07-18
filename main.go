package main

import "fmt"

var TFourmies []Fourmie
var turn int = 0

func main() {
	InitLinks()
	for range 5 {
		TFourmies = append(TFourmies, InitFourmie())
	}
	for i := 0; i<= 3; i++{
		for _, f := range TFourmies {
			ManageMoves(f)
		}
		turn++
	}
	fmt.Println(turn)
}

type Room struct {
	Name     string
	Occupied bool
	Links    []Room
}