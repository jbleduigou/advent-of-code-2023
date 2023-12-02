package calibration_document

import (
	"strings"
)

type CalibrationValue struct {
	digits []int32
}

func ParseCalibrationValue(input string) CalibrationValue {
	digits := make([]int32, 0)
	for _, c := range input {
		if c > 47 && c <= 57 {
			digits = append(digits, c-48)
		}
	}
	return CalibrationValue{digits: digits}
}

func sumOfCalibrationValues(input string) int32 {
	var sum int32 = 0
	for _, v := range strings.Split(input, "\n") {
		cv := ParseCalibrationValue(v)
		sum += 10 * cv.digits[0]
		sum += cv.digits[len(cv.digits)-1]
	}
	return sum
}

var replacements = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func sumOfCalibrationValuesWithLetters(input string) int32 {
	var sum int32 = 0
	for _, line := range strings.Split(input, "\n") {
		for key, value := range replacements {
			line = strings.Replace(line, key, value, -1)
		}
		cv := ParseCalibrationValue(line)
		sum += 10 * cv.digits[0]
		sum += cv.digits[len(cv.digits)-1]
	}
	return sum
}
