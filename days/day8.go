package day

import (
	_file "advent_of_code/helper/file"
	"sort"
	"strconv"
	"strings"
)

type Day8 struct {
	entries []entry
}

type entry struct {
	patterns []string
	output   []string
}

func CreateDay8(path string) Day8 {
	r := _file.ReadTextFile(path)

	entries := make([]entry, 0)
	for _, d := range r {
		split := strings.Split(d, " | ")
		entries = append(entries, entry{
			patterns: strings.Split(split[0], " "),
			output:   strings.Split(split[1], " "),
		})
	}

	return Day8{
		entries: entries,
	}
}

func (d Day8) GetStep1Result() int {
	c := 0
	for _, e := range d.entries {

		_, t := getSimpleDigits(e.output)
		c += t
	}
	return c
}

func getSimpleDigits(e []string) (map[int]string, int) {
	digits := make(map[int]string)
	c := 0
	for _, s := range e {
		var k int
		switch len(s) {
		case 2:
			k = 1
		case 3:
			k = 7
		case 4:
			k = 4
		case 7:
			k = 8
		}

		if k != 0 {
			digits[k] = s
			c++
		}
	}

	return digits, c
}

func (d Day8) GetStep2Result() int {
	count := 0
	for _, e := range d.entries {
		binSync := getBinariesForEntry(e.patterns)

		var countString string
		for _, o := range e.output {

			n := binSync[sortString(o)]
			countString += strconv.Itoa(n)
		}

		v, _ := strconv.Atoi(countString)
		count += v
	}

	return count
}

func subPatternIsInPattern(sp string, p string) bool {
	for _, v := range sp {
		if !strings.Contains(p, string(v)) {
			return false
		}
	}

	return true
}

func getBinariesForEntry(patterns []string) map[string]int {
	digits := make(map[int]string)
	sD, _ := getSimpleDigits(patterns)

	for k, v := range sD {
		digits[k] = v
	}

	for _, pattern := range patterns {
		l := len(pattern)

		if l == 6 {
			if !subPatternIsInPattern(digits[1], pattern) {
				digits[6] = pattern
			} else if subPatternIsInPattern(digits[4], pattern) {
				digits[9] = pattern
			} else {
				digits[0] = pattern
			}
		}
	}

	availableBinaries := []int{2, 3, 5}
out:
	for {
		for _, pattern := range patterns {
			l := len(pattern)
			if l == 5 {
				if subPatternIsInPattern(pattern, digits[9]) && !subPatternIsInPattern(digits[1], pattern) {
					digits[5] = pattern
					availableBinaries = removeFromSlice(availableBinaries, 5)
					continue
				}

				if subPatternIsInPattern(digits[1], pattern) {
					digits[3] = pattern
					availableBinaries = removeFromSlice(availableBinaries, 3)
					continue
				}

				if len(availableBinaries) == 1 {
					digits[2] = pattern
					break out
				}
			}
		}
	}

	binSync := make(map[string]int)
	for k, v := range digits {
		binSync[sortString(v)] = k
	}

	return binSync
}

func sortString(s string) string {
	a := strings.Split(s, "")
	sort.Strings(a)
	return strings.Join(a, "")
}

func removeFromSlice(slice []int, i int) []int {
	newSlice := make([]int, 0)

	for _, v := range slice {
		if v != i {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}
