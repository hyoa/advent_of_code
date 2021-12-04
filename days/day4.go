package day

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Day4 struct {
	draw  []int
	cards []card
}

type card struct {
	column []line
	row    []line
	raw    string
}

type line struct {
	foundNumber    []int
	notFoundNumber []int
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

		chunkSize := 5
		for i := 0; i < len(numbers); i += chunkSize {
			end := i + chunkSize

			if end > len(numbers) {
				end = len(numbers)
			}

			line := line{
				notFoundNumber: numbers[i:end],
			}

			card.row = append(card.row, line)
		}

		for i := 0; i < 5; i++ {
			column := make([]int, 0)

			for j := 0; j < len(numbers); j += 5 {
				column = append(column, numbers[i+j])
			}

			line := line{
				notFoundNumber: column,
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
	c, n := getWinningCard(d.draw, d.cards, "first")
	return sumUnmarked(d.cards[c], n) * n
}

func getWinningCard(draw []int, cards []card, position string) (int, int) {
	fmt.Println("nb cards", len(cards))
	cardsWinning := make([]int, 0)
	var lastN int
	var lastK1 int
	for _, n := range draw {
		fmt.Println(n)
		if len(cardsWinning) == len(cards) {
			fmt.Println(lastK1, lastN)
			return lastK1, lastN
		}

		for k1, card := range cards {
			fmt.Println(cardsWinning)
			if existInSlice(cardsWinning, k1) {
				continue
			}

			for k2, column := range card.column {
				for k3, number := range column.notFoundNumber {
					if number == n {
						cards[k1].column[k2].notFoundNumber = remove(cards[k1].column[k2].notFoundNumber, k3)
						cards[k1].column[k2].foundNumber = append(cards[k1].column[k2].foundNumber, n)
					}

					if len(cards[k1].column[k2].foundNumber) == 5 {
						fmt.Println("winning: ", n)
						cardsWinning = append(cardsWinning, k1)
						fmt.Println("cards", cardsWinning)
						if position == "first" {
							return k1, n
						} else {
							lastN = n
							lastK1 = k1
							continue
						}
					}
				}
			}

			for k2, row := range card.row {
				for k3, number := range row.notFoundNumber {
					if number == n {
						cards[k1].row[k2].notFoundNumber = remove(cards[k1].row[k2].notFoundNumber, k3)
						cards[k1].row[k2].foundNumber = append(cards[k1].row[k2].foundNumber, n)
					}

					if len(cards[k1].row[k2].foundNumber) == 5 {
						cardsWinning = append(cardsWinning, k1)
						if position == "first" {
							return k1, n
						} else {
							lastN = n
							lastK1 = k1
							continue
						}
					}
				}
			}
		}
	}

	fmt.Println(lastN)
	return lastK1, lastN
}

func existInSlice(s []int, i int) bool {
	for _, w := range s {
		if w == i {
			return true
		}
	}

	return false
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func sumUnmarked(c card, n int) int {
	unmarked := make([]int, 0)
	for _, p := range c.column {
		unmarked = append(unmarked, p.notFoundNumber...)
	}

	for _, p := range c.row {
		unmarked = append(unmarked, p.notFoundNumber...)
	}

	unmarked = unique(unmarked)
	sum := 0
	for _, u := range unmarked {
		if u != n {
			sum += u
		}
	}

	return sum
}

func (d Day4) GetStep2Result() int {
	c, n := getWinningCard(d.draw, d.cards, "last")
	// fmt.Printf("%#v\n", d.cards[c])
	// fmt.Println(n)
	return sumUnmarked(d.cards[c], n) * n
}
