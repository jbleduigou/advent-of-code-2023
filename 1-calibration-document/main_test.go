package calibration_document

import (
	"testing"

	"github.com/jbleduigou/advent-of-code-2023/util"
)

func TestSumOfCalibrationValues(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int32
	}{
		{
			name: "Example given in challenge description",
			input: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			want: 142,
		},
		{
			name:  "Real challenge",
			input: util.GetFileSilently("puzzle.txt"),
			want:  55488,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfCalibrationValues(tt.input); got != tt.want {
				t.Errorf("sumOfCalibrationValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOfCalibrationValuesWithLetters(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int32
	}{
		{
			name: "Example given in challenge description",
			input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			want: 281,
		},
		{
			name:  "Real challenge",
			input: util.GetFileSilently("puzzle.txt"),
			want:  55614,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfCalibrationValuesWithLetters(tt.input); got != tt.want {
				t.Errorf("sumOfCalibrationValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
