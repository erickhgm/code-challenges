package generateallpairs

type Pair struct {
	X int
	Y int
}

func GetPairs(n int) []Pair {
	var pairs []Pair
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			pairs = append(pairs, Pair{i, j})
		}
	}
	return pairs
}
