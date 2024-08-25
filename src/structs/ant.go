package structs

import "strconv"

type Ant struct {
	Name      string
	IsArrived bool
	Pos       *Room
	Path      []*Room
}

func InitFourmie(n int, paths []*AntPath) []*Ant {
	allAnts := []*Ant{}
	for i := 1; i <= n; i++ {
		ant := Ant{
			Name: strconv.Itoa(i),
			IsArrived: false,
			Pos: paths[0].Path[0],
			Path: ChoosePath(paths),
		}
		allAnts = append(allAnts, &ant)
	}
	return allAnts
}

func ChoosePath(paths []*AntPath) []*Room{
	min := paths[0]
	for _, p := range paths{
		if p.Score < min.Score{
			min = p
		}
	}
	min.Score++
	return min.Path
}