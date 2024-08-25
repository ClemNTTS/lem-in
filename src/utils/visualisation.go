package utils

import (
	"fmt"
	"lem-in/src/structs"
)

func PrintAntHillRooms(antHill structs.AntHill) {
	fmt.Println("AntHill:")
	fmt.Printf("Start Room: %s\n", antHill.StartRoom)
	fmt.Printf("End Room: %s\n", antHill.EndRoom)
	fmt.Printf("Largest X : %d Largest Y : %d\n", antHill.HighestCoords[0], antHill.HighestCoords[1])
	fmt.Println("Rooms:")
	for _, room := range antHill.Rooms {
		fmt.Printf("Room: %s\n", room.Name)
		fmt.Printf("Coords: %d %d\n", room.X, room.Y)
		fmt.Printf("  Occupied: %v\n", room.Occupied)
		fmt.Print("  Links: ")
		for _, link := range room.Links {
			for _, rooom := range link.Rooms {
				fmt.Print(rooom.Name, " ")
			}
			fmt.Print("| ")
		}
		fmt.Println()
	}
}

func CreateBoard(x, y int) [][]string {

	var board [][]string

	for i := 0; i <= y; i++ {
		row := []string{}
		for j := 0; j <= x; j++ {
			row = append(row, " ")
		}
		board = append(board, row)
	}
	return board
}

func PrintBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(row)
	}
}

func PrintAntHill(antHill structs.AntHill) {

	var board [][]string
	margin := 1

	board = CreateBoard(antHill.HighestCoords[0]+margin, antHill.HighestCoords[1]+margin)

	for _, room := range antHill.Rooms {
		board[room.Y][room.X] = room.Name
	}

	fmt.Println()

	PrintBoard(board)

}
