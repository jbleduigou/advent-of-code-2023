package almanac

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

func getLowestLocationNumberPartOne(input string) int {
	return getLowestLocationNumber(input, parseSeedNumbersPartOne)
}

func getLowestLocationNumberPartTwo(input string) int {
	return getLowestLocationNumber(input, parseSeedNumbersPartTwo)
}

func getLowestLocationNumber(input string, parseSeedNumbers func(input string) []int) int {
	splits := strings.Split(input, "\n\n")

	seedNumbers := parseSeedNumbers(splits[0])

	seedToSoil := parseMap(splits[1])

	soilToFertilizer := parseMap(splits[2])

	fertilizerToWater := parseMap(splits[3])

	waterToLight := parseMap(splits[4])

	lightToTemperature := parseMap(splits[5])

	temperatureToHumidity := parseMap(splits[6])

	humidityToLocation := parseMap(splits[7])

	// init the min heap
	a := &Almanac{}
	heap.Init(a)

	for _, nb := range seedNumbers {
		soil := getCorrespondingValue(seedToSoil, nb)
		fert := getCorrespondingValue(soilToFertilizer, soil)
		water := getCorrespondingValue(fertilizerToWater, fert)
		light := getCorrespondingValue(waterToLight, water)
		temp := getCorrespondingValue(lightToTemperature, light)
		humid := getCorrespondingValue(temperatureToHumidity, temp)
		loc := getCorrespondingValue(humidityToLocation, humid)
		heap.Push(a, Seed{
			number:      nb,
			soil:        soil,
			fertilizer:  fert,
			water:       water,
			light:       light,
			temperature: temp,
			humidity:    humid,
			location:    loc,
		})
	}

	// pop the first element
	seed2 := heap.Pop(a).(Seed)
	seed := seed2

	return seed.location
}

func parseSeedNumbersPartOne(input string) []int {
	numbers := make([]int, 0)

	clean := strings.Replace(input, "seeds: ", "", 1)
	splits := strings.Split(clean, " ")
	for _, str := range splits {
		nb, _ := strconv.Atoi(str)
		numbers = append(numbers, nb)
	}
	return numbers
}

func parseSeedNumbersPartTwo(input string) []int {
	numbers := make([]int, 0)

	clean := strings.Replace(input, "seeds: ", "", 1)
	splits := strings.Split(clean, " ")
	for i, _ := range splits {
		if i%2 == 0 {
			start, _ := strconv.Atoi(splits[i])
			length, _ := strconv.Atoi(splits[i+1])
			for j := start; j < start+length; j++ {
				numbers = append(numbers, j)
			}
		}
	}
	return numbers
}

func parseMap(input string) *CorrespondingTable {
	//// init the min heap
	output := &CorrespondingTable{}
	heap.Init(output)

	splits := strings.Split(input, "\n")
	for nb, line := range splits {
		if nb == 0 {
			continue
		}
		var destinationStart, sourceStart, rangeLength int
		fmt.Sscanf(line, "%d %d %d", &destinationStart, &sourceStart, &rangeLength)
		heap.Push(output, Mapping{
			sourceStart:      sourceStart,
			destinationStart: destinationStart,
			rangeLength:      rangeLength,
		})
	}
	return output
}

func getCorrespondingValue(ct *CorrespondingTable, key int) int {
	for _, m := range *ct {
		if m.sourceStart <= key && key < m.sourceStart+m.rangeLength {
			delta := key - m.sourceStart
			return m.destinationStart + delta
		}
	}
	return key
}
