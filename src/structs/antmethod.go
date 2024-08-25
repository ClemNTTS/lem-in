package structs

import "fmt"

func (f *Ant) ManageMoves(start, end *Room) {
	if f.Pos.Name == end.Name {
		return
	}

	for i, r := range f.Path {
		if f.Pos.Name == r.Name {
			t := FindTunnel(f, f.Path[i+1])
			if !f.Path[i+1].Occupied && !t.Occupied {
				if f.Pos.Name != start.Name {
					f.Pos.Occupied = false
				}
				f.Pos = f.Path[i+1]
				if f.Pos.Name != end.Name {
					f.Pos.Occupied = true
				}
				t.Occupied = true
				fmt.Print("L", f.Name, "-", f.Pos.Name, " ")
				return
			}
		}
	}
}

func ShowPath(path []*Room) string {
	if path == nil {
		return "nil"
	}
	str := ""
	for _, room := range path {
		str += room.Name + " "
	}
	return str
}

func FindTunnel(f *Ant, targetRoom *Room) *Tunnel {
	for _, tunn := range f.Pos.Links {
		for _, room := range tunn.Rooms {
			if room.Name == targetRoom.Name {
				return tunn
			}
		}
	}
	return nil
}
