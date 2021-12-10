package day

import (
	"advent_of_code/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	type test struct {
		d     domain.Day
		step1 int
		step2 int
	}

	tests := []test{
		{d: CreateDay1("../inputs/day1/input_test.txt"), step1: 7, step2: 5},
		{d: CreateDay2("../inputs/day2/input_test.txt"), step1: 150, step2: 900},
		{d: CreateDay3("../inputs/day3/input_test.txt"), step1: 198, step2: 230},
		{d: CreateDay4("../inputs/day4/input_test.txt"), step1: 4512, step2: 1924},
		{d: CreateDay5("../inputs/day5/input_test.txt"), step1: 5, step2: 12},
		{d: CreateDay6("../inputs/day6/input_test.txt"), step1: 5934, step2: 26984457539},
		{d: CreateDay7("../inputs/day7/input_test.txt"), step1: 37, step2: 168},
		{d: CreateDay8("../inputs/day8/input_test.txt"), step1: 26, step2: 61229},
		{d: CreateDay9("../inputs/day9/input_test.txt"), step1: 15, step2: 1134},
		{d: CreateDay10("../inputs/day10/input_test.txt"), step1: 26397, step2: 288957},
	}

	for _, tc := range tests {
		if tc.step1 != 0 {
			assert.Equal(t, tc.step1, tc.d.GetStep1Result())
		}

		if tc.step2 != 0 {
			assert.Equal(t, tc.step2, tc.d.GetStep2Result())
		}
	}
}
