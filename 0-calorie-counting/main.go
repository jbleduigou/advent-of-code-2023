package calorie_counting

import (
	"container/heap"
	"strconv"
	"strings"
)

// See https://adventofcode.com/2022/day/1

type Elf struct {
	calories int
}

type Elves []Elf

func (e Elves) Less(i, j int) bool {
	return e[i].calories > e[j].calories
}

func (e Elves) Len() int { return len(e) }

func (e Elves) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

func (e *Elves) Push(x interface{}) {
	*e = append(*e, x.(Elf))
}

func (e *Elves) Pop() interface{} {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

func getElfWithMaxCalories(input string) int {
	// init the max heap
	elves := &Elves{}
	heap.Init(elves)

	total := 0

	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			heap.Push(elves, Elf{calories: total})
			total = 0
		} else {
			cal, _ := strconv.Atoi(v)
			total = total + cal
		}
	}

	// pop the first element
	e := heap.Pop(elves).(Elf)

	return e.calories
}

func getTopElvesWithMaxCalories(input string, nbElves int) int {
	// init the max heap
	elves := &Elves{}
	heap.Init(elves)

	tmp := 0

	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			heap.Push(elves, Elf{calories: tmp})
			tmp = 0
		} else {
			cal, _ := strconv.Atoi(v)
			tmp = tmp + cal
		}
	}

	calories := 0
	for i := 0; i < nbElves; i++ {
		// pop the first element
		e := heap.Pop(elves).(Elf)
		calories = calories + e.calories
	}

	return calories
}
