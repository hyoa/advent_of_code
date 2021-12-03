package day

import (
	_file "advent_of_code/helper/file"
	"bufio"
	"log"
	"regexp"
	"strconv"
)

type direction string

type movement struct {
	direction direction
	number    int
}

const (
	up      direction = "up"
	down    direction = "down"
	forward direction = "forward"
)

type Day2 struct {
	movements []movement
}

func CreateDay2(path string) Day2 {
	data := _file.ReadTextFile(path, func(s *bufio.Scanner) interface{} {
		directions := make([]string, 0)
		for s.Scan() {
			directions = append(directions, s.Text())
		}

		return directions
	})

	dataOk, ok := data.([]string)

	directions := make([]movement, 0)
	if ok {
		for k := range dataOk {
			r := regexp.MustCompile(`(?P<Direction>[\w]+) (?P<Value>[\d])`)

			res := r.FindStringSubmatch(dataOk[k])
			n, _ := strconv.Atoi(res[2])
			directions = append(directions, movement{
				direction: direction(res[1]),
				number:    n,
			})
		}
	}

	return Day2{
		movements: directions,
	}
}

func (d Day2) GetStep1Result() int {
	x := 0
	y := 0

	for _, movement := range d.movements {
		switch movement.direction {
		case up:
			y = y - movement.number
		case down:
			y = y + movement.number
		case forward:
			x = x + movement.number
		default:
			log.Fatal("Unknown movement")
		}
	}

	return x * y
}

func (d Day2) GetStep2Result() int {
	aim := 0
	x := 0
	y := 0

	for _, movement := range d.movements {
		switch movement.direction {
		case up:
			aim = aim - movement.number
		case down:
			aim = aim + movement.number
		case forward:
			x = x + movement.number
			y = y + (movement.number * aim)
		default:
			log.Fatal("Unknown movement")
		}
	}

	return x * y
}
