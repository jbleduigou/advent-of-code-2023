package calibration_document

import (
	"testing"

	"github.com/jbleduigou/advent-of-code-2023/util"
)

func TestGetValidGames(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example given in challenge description",
			input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			want: 8,
		},
		{
			name:  "Real challenge",
			input: util.GetFileSilently("puzzle.txt"),
			want:  2776,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValidGames(tt.input); got != tt.want {
				t.Errorf("getValidGames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSumOfPower(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example given in challenge description",
			input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			want: 2286,
		},
		{
			name:  "Real challenge",
			input: util.GetFileSilently("puzzle.txt"),
			want:  68638,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfPower(tt.input); got != tt.want {
				t.Errorf("getSumOfPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
