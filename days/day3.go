package day

import (
	_file "advent_of_code/helper/file"
	"strconv"
)

type Day3 struct {
	binaries []string
}

func CreateDay3(path string) Day3 {
	return Day3{
		binaries: _file.ReadTextFile(path),
	}
}

type diagValues struct {
	gamma   string
	epsilon string
}

func (d Day3) GetStep1Result() int {
	size := len(d.binaries[0])

	diag := diagValues{
		gamma:   "",
		epsilon: "",
	}

	for i := 0; i < size; i++ {
		count1 := 0
		count0 := 0
		for _, b := range d.binaries {
			if b[i:i+1] == "1" {
				count1++
			} else {
				count0++
			}
		}

		if count1 > count0 {
			diag = diagValues{
				gamma:   diag.gamma + "1",
				epsilon: diag.epsilon + "0",
			}
		} else {
			diag = diagValues{
				gamma:   diag.gamma + "0",
				epsilon: diag.epsilon + "1",
			}
		}
	}

	intGamma, _ := strconv.ParseInt(diag.gamma, 2, 64)
	intEpsilon, _ := strconv.ParseInt(diag.epsilon, 2, 64)
	return int(intGamma) * int(intEpsilon)
}

func (d Day3) GetStep2Result() int {
	o2, _ := strconv.ParseInt(filterOutBinaries(d.binaries, "most", 0), 2, 64)
	co2, _ := strconv.ParseInt(filterOutBinaries(d.binaries, "least", 0), 2, 64)
	return int(o2) * int(co2)
}

func filterOutBinaries(binaries []string, criteria string, position int) string {

	startWith1 := make([]string, 0)
	startWith0 := make([]string, 0)

	for _, binary := range binaries {
		if binary[position:position+1] == "1" {
			startWith1 = append(startWith1, binary)
		} else {
			startWith0 = append(startWith0, binary)
		}
	}

	var keep []string
	if criteria == "most" {
		if len(startWith1) >= len(startWith0) {
			keep = startWith1
		} else {
			keep = startWith0
		}
	} else {
		if len(startWith0) <= len(startWith1) {
			keep = startWith0
		} else {
			keep = startWith1
		}
	}

	if len(keep) > 1 {
		return filterOutBinaries(keep, criteria, position+1)
	}

	return keep[0]
}
