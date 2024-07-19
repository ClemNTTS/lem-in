package main

import (
	"strconv"
)

var nb_fourmies int
var start Room = Room{Name: "Start", Occupied: true}
var r1 Room = Room{Name: "r1", Occupied: false}
var r2 Room = Room{Name: "r2", Occupied: false}
var r3 Room = Room{Name: "r3", Occupied: false}
var r4 Room = Room{Name: "r4", Occupied: false}
var end Room = Room{Name: "end", Occupied: false}

var path []*Room = []*Room{&start,&r1, &r3, &end}

type Fourmie struct {
	Name string
	Pos  *Room
	Path []*Room
}

func InitLinks(){
	start.Links = []*Room{&r2}
	r1.Links =  []*Room{&start, &r3}
	r2.Links = []*Room{&start,&r3,&r4}
	r3.Links = []*Room{&r1,&r2,&end}
	r4.Links = []*Room{&r3,&end}
}

func InitFourmie() Fourmie {
	nb_fourmies++
	return Fourmie{Name: "F" + strconv.Itoa(nb_fourmies), Pos: &start, Path: path}
}

func (f *Fourmie) ManageMoves(){
	if f.Pos.Name == "end"{
		return
	}



	for i, r := range f.Path{
		if f.Pos.Name == r.Name{
			if !f.Path[i+1].Occupied{
				if f.Pos.Name != "start"{
					f.Pos.Occupied = false
				}
				f.Pos = f.Path[i+1]
				if f.Pos.Name != "end"{
					f.Pos.Occupied = true
				}
				return
			}else{
				for i := range f.Pos.Links{
					if !f.Pos.Links[i].Occupied{
						if f.Pos.Name != "start"{
							f.Pos.Occupied = false
						}
						f.Pos = f.Pos.Links[i]
						if f.Pos.Name != "end"{
							f.Pos.Occupied = true
						}
					}
				}
				return
			}
		}
		
	}

	for _, elem := range f.Pos.Links{
		if !elem.Occupied{
			if f.Pos.Name != "start"{
				f.Pos.Occupied = false
			}
			f.Pos = elem
			if f.Pos.Name != "end"{
				f.Pos.Occupied = true
			}
		}
	}
}
