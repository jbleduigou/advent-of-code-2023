package scratch_cards

import (
	"fmt"
	"math"
	"strings"
)

type ScratchCard struct {
	winningNumbers map[string]interface{}
	pickedNumbers  []string
	numberMatching int
	points         int
}

func ParseScratchCard(i int, raw string) ScratchCard {
	clean := strings.Replace(raw, fmt.Sprintf("Card %v: ", i), "", -1)
	splits := strings.Split(clean, " | ")
	// first populate the set of winning numbers
	winningNumbers := make(map[string]interface{}, 0)
	for _, v := range strings.Split(splits[0], " ") {
		if v != "" {
			winningNumbers[v] = struct{}{}
		}
	}
	// then populate the array of picked numbers
	pickedNumbers := make([]string, 0)
	for _, v := range strings.Split(splits[1], " ") {
		if v != "" {
			pickedNumbers = append(pickedNumbers, v)
		}
	}
	card := ScratchCard{winningNumbers: winningNumbers, pickedNumbers: pickedNumbers}
	card.calculatePoints()
	return card
}

func (c *ScratchCard) calculatePoints() {
	for _, number := range c.pickedNumbers {
		_, found := c.winningNumbers[number]
		if found {
			c.numberMatching++
		}
	}
	if c.numberMatching == 0 {
		c.points = 0
	}
	c.points = int(math.Pow(2.0, float64(c.numberMatching-1)))
}

func getSumOfPoints(input string) int {
	pile := make([]ScratchCard, 0)
	for i, line := range strings.Split(input, "\n") {
		card := ParseScratchCard(i+1, line)
		pile = append(pile, card)
	}
	sum := 0
	for _, card := range pile {
		sum = sum + card.points
	}
	return sum
}

func getNumberOfScratchCards(input string) int {
	pile := make([]ScratchCard, 0)
	amount := make(map[int]int)
	for i, line := range strings.Split(input, "\n") {
		card := ParseScratchCard(i+1, line)
		pile = append(pile, card)
		amount[i] = 1
	}
	sum := 0
	for i, card := range pile {
		for j := 0; j < card.numberMatching; j++ {
			current := amount[i+j+1]
			amount[i+j+1] = current + amount[i]
		}
		sum = sum + amount[i]
	}
	return sum
}
