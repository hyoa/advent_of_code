package day

import (
	_file "advent_of_code/helper/file"
	"sort"
	"strings"
)

type Day10 struct {
	lines             []string
	scoreCorrupted    map[string]int
	scoreIncomplete   map[string]int
	endAutocomplete   map[string]string
	startAutocomplete map[string]string
}

func CreateDay10(path string) Day10 {
	r := _file.ReadTextFile(path)

	scoreCorrupted := make(map[string]int)
	scoreCorrupted[")"] = 3
	scoreCorrupted["]"] = 57
	scoreCorrupted["}"] = 1197
	scoreCorrupted[">"] = 25137

	scoreIncomplete := make(map[string]int)
	scoreIncomplete[")"] = 1
	scoreIncomplete["]"] = 2
	scoreIncomplete["}"] = 3
	scoreIncomplete[">"] = 4

	endCompletion := make(map[string]string)
	startCompletion := make(map[string]string)

	endCompletion[")"] = "("
	endCompletion["]"] = "["
	endCompletion["}"] = "{"
	endCompletion[">"] = "<"
	startCompletion["("] = ")"
	startCompletion["["] = "]"
	startCompletion["{"] = "}"
	startCompletion["<"] = ">"

	return Day10{
		lines:             r,
		scoreCorrupted:    scoreCorrupted,
		scoreIncomplete:   scoreIncomplete,
		endAutocomplete:   endCompletion,
		startAutocomplete: startCompletion,
	}
}

type openChar string
type closingChar string

const (
	parenthesisO  openChar = "("
	bracketO      openChar = "["
	chevronO      openChar = "<"
	curlyBracketO openChar = "{"
)

func (d Day10) GetStep1Result() int {
	_, illegals := getIllegalsLineIndexAndChar(d.lines, d.endAutocomplete)

	sum := 0
	for k, v := range illegals {
		sum += d.scoreCorrupted[k] * v
	}

	return sum
}

func getIllegalsLineIndexAndChar(lines []string, autocomplete map[string]string) ([]int, map[string]int) {
	illegals := make(map[string]int)
	index := make([]int, 0)
	for k, l := range lines {
		oC := make([]string, 0)
		for _, c := range strings.Split(l, "") {
			if isStartingChar(c) {
				oC = append(oC, c)
			} else {
				if len(oC) == 0 {
					illegals[c] += 1
					index = append(index, k)
					break
				}

				last := oC[len(oC)-1]
				if autocomplete[c] == last {
					if len(oC) > 0 {
						oC = oC[:len(oC)-1]
					}
				} else {
					illegals[c] += 1
					index = append(index, k)
					break
				}
			}

		}

	}

	return index, illegals
}

func isStartingChar(s string) bool {
	if s == string(parenthesisO) || s == string(bracketO) || s == string(chevronO) || s == string(curlyBracketO) {
		return true
	}

	return false
}

func (d Day10) GetStep2Result() int {
	illegals, _ := getIllegalsLineIndexAndChar(d.lines, d.endAutocomplete)
	incompleteLines := make([]string, 0)

	for k := range d.lines {
		found := false
		for _, v := range illegals {
			if k == v {
				found = true
				break
			}
		}

		if !found {
			incompleteLines = append(incompleteLines, d.lines[k])
		}
	}

	scores := make([]int, 0)
	for _, l := range incompleteLines {
		oC := make([]string, 0)
		for _, c := range strings.Split(l, "") {
			if isStartingChar(c) {
				oC = append(oC, c)
			} else {
				last := oC[len(oC)-1]
				if d.endAutocomplete[c] == last {
					if len(oC) > 0 {
						oC = oC[:len(oC)-1]
					}
				}
			}
		}

		for i, j := 0, len(oC)-1; i < j; i, j = i+1, j-1 {
			oC[i], oC[j] = oC[j], oC[i]
		}

		score := 0
		for _, c := range oC {
			score = score*5 + d.scoreIncomplete[d.startAutocomplete[c]]
		}
		scores = append(scores, score)

	}

	sort.Ints(scores)
	middle := len(scores) / 2
	return scores[middle]
}
