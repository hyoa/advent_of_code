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
