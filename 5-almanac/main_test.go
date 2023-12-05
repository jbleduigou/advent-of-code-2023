package almanac

import (
	"reflect"
	"testing"

	"github.com/jbleduigou/advent-of-code-2023/util"
)

func TestGetLowestLocationNumberPartOne(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example given in challenge description",
			input: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
			want: 35,
		},
		{
			name:  "Real challenge",
			input: util.GetFileSilently("puzzle.txt"),
			want:  551761867,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLowestLocationNumberPartOne(tt.input); got != tt.want {
				t.Errorf("getLowestLocationNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLowestLocationNumberPartTwo(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example given in challenge description",
			input: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
			want: 46,
		},
		// {
		// 	name:  "Real challenge",
		// 	input: util.GetFileSilently("puzzle.txt"),
		// 	want:  551761867,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLowestLocationNumberPartTwo(tt.input); got != tt.want {
				t.Errorf("getLowestLocationNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSeedNumbersPartOne(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{
			name:  "Example given in challenge description",
			input: "seeds: 79 14 55 13",
			want:  []int{79, 14, 55, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSeedNumbersPartOne(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSeedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSeedNumbersPartTwo(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{
			name:  "Simple example",
			input: "seeds: 79 2 5 1",
			want:  []int{79, 80, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSeedNumbersPartTwo(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSeedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
