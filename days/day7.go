package day

import (
	_file "advent_of_code/helper/file"
	"math"
	"strconv"
	"strings"
)

type Day7 struct {
	positions []int
}

func CreateDay7(path string) Day7 {
	data := _file.GetFileContent(path)
	positionsAsString := strings.Split(data, ",")

	positions := make([]int, 0)

	for _, d := range positionsAsString {
		v, _ := strconv.Atoi(d)
		positions = append(positions, v)
	}

	return Day7{
		positions: positions,
	}
}

func (d Day7) GetStep1Result() int {
	return getCheapestCost(d.positions, false)
}

func getCheapestCost(positions []int, increment bool) int {
	directionCost := make(map[int]int)

	for i := 1; i <= len(positions); i++ {
		cost := 0
		for _, n := range positions {
			c := int(math.Abs(float64(n) - float64(i)))

			if increment {
				cost += c * (c + 1) / 2
			} else {
				cost += c
			}
		}

		directionCost[i] = cost
	}

	less := 0
	for _, v := range directionCost {
		if less == 0 {
			less = v
		} else if v < less {
			less = v
		}
	}

	return less
}

func (d Day7) GetStep2Result() int {
	return getCheapestCost(d.positions, true)
}
