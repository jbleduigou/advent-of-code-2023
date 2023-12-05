package almanac

type CorrespondingTable []Mapping

type Mapping struct {
	sourceStart      int
	destinationStart int
	rangeLength      int
}

func (ct CorrespondingTable) Less(i, j int) bool {
	return ct[i].sourceStart < ct[j].sourceStart
}

func (ct CorrespondingTable) Len() int { return len(ct) }

func (ct CorrespondingTable) Swap(i, j int) { ct[i], ct[j] = ct[j], ct[i] }

func (ct *CorrespondingTable) Push(x interface{}) {
	*ct = append(*ct, x.(Mapping))
}

func (ct *CorrespondingTable) Pop() interface{} {
	old := *ct
	n := len(old)
	x := old[n-1]
	*ct = old[0 : n-1]
	return x
}
