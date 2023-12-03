package engine_parts

import (
	"testing"

	"github.com/jbleduigou/advent-of-code-2023/util"
)

func Test_getSumOfEngineParts(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example given in challenge description",
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: 4361,
		},
		{
			name:  "Real challenge",
			input: util.GetFileSilently("puzzle.txt"),
			want:  536202,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfEngineParts(tt.input); got != tt.want {
				t.Errorf("getSumOfEngineParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSumOfGears(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example given in challenge description",
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: 467835,
		},
		{
			name:  "Real challenge",
			input: util.GetFileSilently("puzzle.txt"),
			want:  78272573,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfGears(tt.input); got != tt.want {
				t.Errorf("getSumOfGears() = %v, want %v", got, tt.want)
			}
		})
	}
}
