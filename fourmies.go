package main

import "strconv"

var nb_fourmies int
var start Room = Room{Name: "Start", Occupied: false}
var r1 Room = Room{Name: "r1", Occupied: false}
var r2 Room = Room{Name: "r2", Occupied: false}
var r3 Room = Room{Name: "r3", Occupied: false}
var r4 Room = Room{Name: "r4", Occupied: false}
var end Room = Room{Name: "end", Occupied: false}

var path []Room = []Room{r1, r3, end}

type Fourmie struct {
	Name string
	Pos  Room
	Path []Room
}

func InitLinks(){
	r1.Links =  []Room{start, r3}
	r2.Links = []Room{start,r3,r4}
	r3.Links = []Room{r1,r2,end}
	r4.Links = []Room{r3,end}
}

func InitFourmie() Fourmie {
	nb_fourmies++
	return Fourmie{Name: "F" + strconv.Itoa(nb_fourmies), Pos: start, Path: path}
}

func ManageMoves(f Fourmie){
	
}
