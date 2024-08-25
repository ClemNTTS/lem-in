package src

import (
	"fmt"
	"lem-in/src/structs"
	"lem-in/src/utils"
)

func Launch(url string) {
	//Scrap informations
	fileText := utils.GetFileText(url)
	nbAnt, start, tabRooms, end, tabLinks := utils.ScrapValues((fileText))
	if start == "" ||end == "" ||nbAnt <= 0{
		fmt.Println("ERROR: invalid data format")
		return
	} 
	tabTabRooms := utils.SplitTabRoom(tabRooms)

	//Add informations in right variable
	antHill, err := structs.CreateRooms(start, tabTabRooms, end, tabLinks)
	if err != nil{
		fmt.Println(err)
		return
	}

	//varaibles
	turn := 0

	Paths := structs.InitPaths(antHill.Rooms[antHill.StartRoom], antHill.Rooms[antHill.EndRoom])
	if len(Paths) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}

	//création de n fourmies
	totalAnts := structs.InitFourmie(nbAnt, Paths)
	arrivedAnts := 0

	for {
		//pass all tunnels.Occupied at false
		InitTunnels(totalAnts)
		turn++
		for i := 1; i <= nbAnt; i++ {
			currentAnt := totalAnts[i-1]
			if currentAnt.IsArrived {
				continue
			}
			if !currentAnt.IsArrived {
				currentAnt.ManageMoves(antHill.Rooms[antHill.StartRoom], antHill.Rooms[antHill.EndRoom])

				if currentAnt.Pos.Name == antHill.EndRoom {

					arrivedAnts++
					currentAnt.IsArrived = true
					continue
				}
			}
		}
		fmt.Println()
		if arrivedAnts == len(totalAnts) {
			fmt.Println("Finito pipo in : ", turn)
			break
		}
		fmt.Println("-------")
	}
	utils.PrintAntHill(antHill)
}

func InitTunnels(antsSlice []*structs.Ant) {
	for _, ant := range antsSlice {
		for _, tunn := range ant.Pos.Links {
			tunn.Occupied = false
		}
	}
}