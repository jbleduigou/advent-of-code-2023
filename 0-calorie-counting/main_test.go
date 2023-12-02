package calorie_counting

import (
	"testing"

	"github.com/jbleduigou/advent-of-code-2023/util"
)

var example = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestGetElfWithMaxCalories(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example given in challenge description",
			input: example,
			want:  24000,
		},
		{
			name:  "Real challenge",
			input: util.GetFileSilently("puzzle.txt"),
			want:  70509,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getElfWithMaxCalories(tt.input); got != tt.want {
				t.Errorf("getMaxCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTop3ElvesWithMaxCalories(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example given in challenge description",
			input: example,
			want:  45000,
		},
		{
			name:  "Real challenge",
			input: util.GetFileSilently("puzzle.txt"),
			want:  208567,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTopElvesWithMaxCalories(tt.input, 3); got != tt.want {
				t.Errorf("getMaxCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}