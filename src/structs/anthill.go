package structs

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Room struct {
	Name     string
	X        int
	Y        int
	Occupied bool
	Links    []*Tunnel
}

type Tunnel struct {
	Occupied bool
	Rooms    []*Room
}

type AntHill struct {
	NumberOfAnts  int
	HighestCoords []int
	StartRoom     string
	EndRoom       string
	Rooms         map[string]*Room
}

func CreateRooms(start string, tabRoomSplitted [][]string, end string, tabLinks []string) (AntHill, error) {

	antHill := AntHill{
		Rooms: make(map[string]*Room),
	}

	tabStart := strings.Split(start, " ")
	tabEnd := strings.Split(end, " ")

	highestCoords := []int{0, 0}

	for i := 0; i < len(tabRoomSplitted); i++ {
		x, _ := strconv.Atoi(tabRoomSplitted[i][1])
		y, _ := strconv.Atoi(tabRoomSplitted[i][2])

		if x > highestCoords[0] {
			highestCoords[0] = x
		}
		if y > highestCoords[1] {
			highestCoords[1] = y
		}

		item := Room{
			Name:     tabRoomSplitted[i][0],
			X:        x,
			Y:        y,
			Occupied: false,
		}
		if reflect.DeepEqual(tabStart, tabRoomSplitted[i]) {
			antHill.StartRoom = item.Name
		} else if reflect.DeepEqual(tabEnd, tabRoomSplitted[i]) {
			antHill.EndRoom = item.Name
		}

		antHill.Rooms[item.Name] = &item

	}

	antHill.HighestCoords = highestCoords

	for _, value := range tabLinks {
		tmpLink := strings.Split(value, "-")
		for _, val := range tmpLink{
			if val == ""{
				return antHill, errors.New("nil room value")
			}
		}

		if len(tmpLink) == 2 {

			room1 := antHill.Rooms[tmpLink[0]]
			room2 := antHill.Rooms[tmpLink[1]]

			link := Tunnel{
				Occupied: false,
			}

			link.Rooms = append(link.Rooms, room2)
			link.Rooms = append(link.Rooms, room1)

			room1.Links = append(room1.Links, &link)
			room2.Links = append(room2.Links, &link)

			antHill.Rooms[room1.Name] = room1
			antHill.Rooms[room2.Name] = room2

		}else{
			return antHill, errors.New("number of room error")
		}
	}

	return antHill, nil
}

func InitPaths(start, end *Room) []*AntPath {
	var totalPaths []*AntPath
	queue := []*Room{start}
	parent := map[*Room]*Room{}
	

	for len(queue) > 0{
		current := queue[0]
		queue = queue[1:]


		for _,tunn := range current.Links{
			if tunn.Occupied{
				continue
			}
			for _, room := range tunn.Rooms{
				if room == end{
					var path AntPath
					path.Path = append(path.Path, end)
					for current != start{
						path.Path = append(path.Path, current)
						current = parent[current]
					}
					path.Path = append(path.Path, start)
					path.Path = ReversePathOrder(path.Path)
					path.Score = len(path.Path)
					totalPaths = append(totalPaths, &path)
				}else if room != current {
						for _, p := range totalPaths{
							if In(p.Path, parent[room]){
								parent[room] = nil
							}
						}
						if len(room.Links) == 2{
							queue = append([]*Room{room},queue...)
						}else{
							queue = append(queue, room)
						}
						if parent[room] == nil{
							parent[room] = current
						}
				}
			}
			tunn.Occupied = true
		}
	}
	return totalPaths
}

// Reverse path order so that it is in a valid state.
func ReversePathOrder(path []*Room) []*Room {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func In(t []*Room, r *Room) bool{
	if r == nil{
		return false
	}
	for _, room := range t{
		if room.Name == r.Name{
			return true
		}
	}
	return false
}