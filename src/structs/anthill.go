package src

type Room struct {
	Name     string
	Occupied bool
	Links    []Room
}
