package day

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Day4 struct {
	draw  []int
	cards []card
}

type card struct {
	column  []line
	row     []line
	raw     string
	numbers []int
}

type line struct {
	numbers []int
}

func CreateDay4(path string) Day4 {

	content, _ := ioutil.ReadFile(path)

	inputs := strings.Split(string(content), "\n\n")

	draw := make([]int, 0)
	drawNumber := strings.Split(inputs[0], ",")
	rawCards := inputs[1:]

	for _, i := range drawNumber {
		v, _ := strconv.Atoi(i)
		draw = append(draw, v)
	}

	cards := make([]card, 0)
	for _, rawCard := range rawCards {
		card := card{
			column: make([]line, 0),
			row:    make([]line, 0),
			raw:    rawCard,
		}
		r := regexp.MustCompile(`(?P<Number>[\d]+)`)
		res := r.FindAllStringSubmatch(rawCard, -1)

		numbers := make([]int, 0)
		for _, number := range res {
			n, _ := strconv.Atoi(number[1])
			numbers = append(numbers, n)
		}

		card.numbers = numbers

		chunkSize := 5
		for i := 0; i < len(numbers); i += chunkSize {
			end := i + chunkSize

			if end > len(numbers) {
				end = len(numbers)
			}

			line := line{
				numbers: numbers[i:end],
			}

			card.row = append(card.row, line)
		}

		for i := 0; i < 5; i++ {
			column := make([]int, 0)

			for j := 0; j < len(numbers); j += 5 {
				column = append(column, numbers[i+j])
			}

			line := line{
				numbers: column,
			}

			card.column = append(card.column, line)
		}
		cards = append(cards, card)
	}

	return Day4{
		draw:  draw,
		cards: cards,
	}
}

func (d Day4) GetStep1Result() int {
	data := getWinningCards(d.draw, d.cards)
	first := data[0]
	return first.number * first.sumUnmarked
}

type winningData struct {
	number      int
	count       int
	sumUnmarked int
}

func getWinningCards(draw []int, cards []card) []winningData {
	data := make([]winningData, 0)
	for _, card := range cards {
		var winningNumber int
		pulled := make([]int, 0)

		columnsChecked := make(map[int]int)
		rowsChecked := make(map[int]int)
	out:
		for _, n := range draw {
			pulled = append(pulled, n)

			for k2, column := range card.column {
				for _, number := range column.numbers {
					if number == n {
						columnsChecked[k2] = columnsChecked[k2] + 1
					}

					if columnsChecked[k2] == 5 {
						winningNumber = n
						break out
					}
				}
			}

			for k2, row := range card.row {
				for _, number := range row.numbers {
					if number == n {
						rowsChecked[k2] = rowsChecked[k2] + 1
					}

					if rowsChecked[k2] == 5 {
						winningNumber = n
						break out
					}
				}
			}
		}

		data = append(data, winningData{
			sumUnmarked: sumUnmarked(card, pulled),
			count:       len(pulled),
			number:      winningNumber,
		})
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].count < data[j].count
	})

	return data
}

func sumUnmarked(c card, p []int) int {
	sum := 0
	for _, u := range c.numbers {
		if !existInSlice(p, u) {
			sum += u
		}
	}

	return sum
}

func (d Day4) GetStep2Result() int {
	data := getWinningCards(d.draw, d.cards)
	last := data[len(data)-1]
	return last.number * last.sumUnmarked
}

func existInSlice(s []int, i int) bool {
	for _, w := range s {
		if w == i {
			return true
		}
	}

	return false
}
