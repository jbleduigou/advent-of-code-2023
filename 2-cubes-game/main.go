package calibration_document

import (
	"fmt"
	"strings"
)

type Game struct {
	id       int
	maximums map[string]int
}

func (g Game) isValid(limits map[string]int) bool {
	for name, number := range limits {
		maximum, ok := g.maximums[name]
		if !ok {
			return false
		}
		if maximum > number {
			return false
		}
	}
	return true
}

func ParseGame(id int, line string) Game {
	maximums := make(map[string]int)
	clean := strings.Replace(line, fmt.Sprintf("Game %v: ", id), "", -1)
	for _, revealed := range strings.Split(clean, "; ") {
		for _, colour := range strings.Split(revealed, ", ") {
			var name string
			var nb int
			fmt.Sscanf(colour, "%d %s", &nb, &name)
			checkMaximum(maximums, name, nb)
		}
	}
	return Game{id: id, maximums: maximums}
}

func checkMaximum(maximums map[string]int, name string, nb int) {
	current, ok := maximums[name]
	if !ok {
		maximums[name] = nb
	} else {
		if current < nb {
			maximums[name] = nb
		}
	}
}

func getValidGames(input string) int {
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0
	for id, line := range strings.Split(input, "\n") {
		g := ParseGame(id+1, line)
		if g.isValid(limits) {
			sum = sum + g.id
		}
	}
	return sum
}

func getSumOfPower(input string) int {
	sum := 0
	for id, line := range strings.Split(input, "\n") {
		g := ParseGame(id+1, line)
		sum = sum + g.maximums["red"]*g.maximums["green"]*g.maximums["blue"]
	}
	return sum
}
