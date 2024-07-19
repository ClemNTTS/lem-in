package main

import (
	"fmt"
	"os"
	"strconv"
)

var TFourmies []Fourmie

func main() {
	InitLinks()
	if len(os.Args) == 1{
		return
	}
	n,_ := strconv.Atoi(os.Args[1])
	turn := 0
	
	//création de n fourmies
	for f := 0 ; f < n; f++{
		TFourmies = append(TFourmies, InitFourmie())
	}
		
	for {
		for i := range TFourmies {
				TFourmies[i].ManageMoves()
				fmt.Println(TFourmies[i].Name + " : " + TFourmies[i].Pos.Name)
		}
		fmt.Println("-------")
		turn++

		if TFourmies[len(TFourmies)-1].Pos.Name == "end"{
			fmt.Println(turn)
			break
		}
	}
}

type Room struct {
	Name     string
	Occupied bool
	Links    []*Room
}