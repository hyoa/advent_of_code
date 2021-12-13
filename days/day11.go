package day

import (
	_file "advent_of_code/helper/file"
	"fmt"
	"strconv"
	"strings"
)

type Day11 struct {
	octopus map[position]int
}

func CreateDay11(path string) Day11 {
	r := _file.ReadTextFile(path)
	tiles := make(map[position]int)

	row := 0
	for _, l := range r {
		column := 0
		for _, octopus := range strings.Split(l, "") {
			v, _ := strconv.Atoi(octopus)
			p := position{x: row, y: column}
			tiles[p] = v
			column++
		}
		row++
	}

	return Day11{
		octopus: tiles,
	}
}

func (d Day11) GetStep1Result() int {
	octopus := make(map[position]int)
	for k, v := range d.octopus {
		octopus[k] = v
	}

	return getFlashCount(octopus, 100)
}

func getFlashCount(octopus map[position]int, iteration int) int {
	countFlash := 0

	for i := 0; i < iteration; i++ {
		toIncrement := make(map[position]int)
		flashed := make([]position, 0)

		for k := range octopus {
			toIncrement[k] = 1
		}

		for {
			if len(toIncrement) == 0 {
				break
			}

			for k, v := range toIncrement {
				if octopus[k]+v > 9 {
					countFlash++
					octopus[k] = 0
					flashed = append(flashed, k)
					adjacent := getAdjacentOctopus(octopus, k)
					for _, v := range adjacent {
						toIncrement[v] += 1
					}
				} else {
					if !existInPositionSlice(flashed, k) {
						octopus[k] += v
					}
				}
				delete(toIncrement, k)
			}
		}
	}

	return countFlash
}

func debugAsGrid(octopus map[position]int) {
	copyOctopus := make(map[position]int)

	for k, v := range octopus {
		copyOctopus[k] = v
	}

	lines := make([]string, 0)
	row := 0
	for i := 0; i < 10; i++ {
		column := 0
		line := ""
		for j := 0; j < 10; j++ {
			line += strconv.Itoa(copyOctopus[position{x: row, y: column}])
			column++
		}
		lines = append(lines, line)
		row++
	}

	fmt.Println("------ GRID ------")
	for _, l := range lines {
		fmt.Println(l)
	}
	fmt.Println("------ END GRID ------")
}

func getAdjacentOctopus(octopus map[position]int, actual position) []position {
	adjacent := make([]position, 0)
	change := []position{
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: 0, y: 1},
		{x: -1, y: -1},
		{x: -1, y: 1},
		{x: 1, y: -1},
		{x: 1, y: 1},
	}

	for _, v := range change {
		nP := position{x: actual.x + v.x, y: actual.y + v.y}
		if _, ok := octopus[nP]; ok {
			adjacent = append(adjacent, nP)
		}
	}

	return adjacent
}

func existInPositionSlice(s []position, p position) bool {
	for _, w := range s {
		if w == p {
			return true
		}
	}

	return false
}

func (d Day11) GetStep2Result() int {
	i := 1
	for {
		count := getFlashCount(d.octopus, 1)

		if count == 100 {
			break
		}
		i++
	}

	return i
}
