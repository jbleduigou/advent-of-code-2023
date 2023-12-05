package almanac

type Almanac []Seed

func (a Almanac) Less(i, j int) bool {
	return a[i].location < a[j].location
}

func (a Almanac) Len() int { return len(a) }

func (a Almanac) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a *Almanac) Push(x interface{}) {
	*a = append(*a, x.(Seed))
}

func (a *Almanac) Pop() interface{} {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[0 : n-1]
	return x
}
