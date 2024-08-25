package structs

type AntPath struct {
	Path  []*Room
	Score int
}

func SortPathByScore(t []*AntPath) {
	for i := 0; i < len(t); i++ {
		for j := i + 1; j < len(t); j++ {
			if t[i].Score > t[j].Score {
				t[i], t[j] = t[j], t[i]
			}
		}
	}
}