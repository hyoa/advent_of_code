package day

import (
	_file "advent_of_code/helper/file"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Day5 struct {
	vents []vent
}

type position struct {
	x int
	y int
}

type vent struct {
	start position
	end   position
}

func CreateDay5(path string) Day5 {
	r := _file.ReadTextFile(path)

	regexp := regexp.MustCompile(`(?P<x1>[\d]+),(?P<y1>[\d]+) -> (?P<x2>[\d]+),(?P<y2>[\d]+)`)

	vents := make([]vent, 0)
	for _, input := range r {
		res := regexp.FindAllStringSubmatch(input, -1)
		for _, v := range res {
			x1, _ := strconv.Atoi(v[1])
			y1, _ := strconv.Atoi(v[2])
			x2, _ := strconv.Atoi(v[3])
			y2, _ := strconv.Atoi(v[4])

			vents = append(vents, vent{
				start: position{x: x1, y: y1},
				end:   position{x: x2, y: y2},
			})
		}
	}

	return Day5{
		vents: vents,
	}
}

func (d Day5) GetStep1Result() int {
	filteredVents := make([]vent, 0)

	for _, v := range d.vents {
		if v.start.x == v.end.x || v.start.y == v.end.y {
			filteredVents = append(filteredVents, v)
		}
	}

	return getNbTooMuchOverlap(filteredVents)
}

func getVentTiles(v vent) []string {
	positions := make([]string, 0)
	currentPosition := v.start
	positions = append(positions, strconv.Itoa(currentPosition.x)+","+strconv.Itoa(currentPosition.y))

	var movX int
	if v.end.x > v.start.x {
		movX = 1
	} else if v.end.x < v.start.x {
		movX = -1
	} else {
		movX = 0
	}

	var movY int
	if v.end.y > v.start.y {
		movY = 1
	} else if v.end.y < v.start.y {
		movY = -1
	} else {
		movY = 0
	}

	i := 0
	for {
		i++
		currentPosition.x = currentPosition.x + movX
		currentPosition.y = currentPosition.y + movY
		key := strconv.Itoa(currentPosition.x) + "," + strconv.Itoa(currentPosition.y)
		positions = append(positions, key)

		if i > 1000 {
			fmt.Println("Infinite loop detected")
			os.Exit(1)
		}

		if v.end.x == currentPosition.x && v.end.y == currentPosition.y {
			break
		}
	}

	return positions
}

func getNbTooMuchOverlap(vents []vent) int {
	positions := make(map[string]int)

	for _, v := range vents {
		tiles := getVentTiles(v)

		for _, t := range tiles {
			positions[t] = positions[t] + 1
		}
	}

	count := 0
	for _, p := range positions {
		if p >= 2 {
			count++
		}
	}

	return count
}

func (d Day5) GetStep2Result() int {
	return getNbTooMuchOverlap(d.vents)
}
