package main

import (
	_days "advent_of_code/days"
	"fmt"
)

func main() {
	d := _days.CreateDay4("inputs/day4/input.txt")

	fmt.Println(d.GetStep1Result(), d.GetStep2Result())
}
