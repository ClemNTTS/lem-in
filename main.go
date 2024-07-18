package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	fileText := GetFileText(args[0])
	nbAnt, start, tabRooms, end, tabLinks := ScrapValues((fileText))
	affich(fileText)
	fmt.Println()
	fmt.Println(nbAnt)
	fmt.Println(start)
	fmt.Println(tabRooms)
	fmt.Println(end)
	fmt.Println(tabLinks)
}
