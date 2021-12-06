package day

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Day6 struct {
	data []int
}

func CreateDay6(path string) Day6 {
	content, _ := ioutil.ReadFile(path)

	data := make([]int, 0)
	for _, v := range strings.Split(string(content), ",") {
		n, _ := strconv.Atoi(v)
		data = append(data, n)
	}

	return Day6{
		data: data,
	}
}

func (d Day6) GetStep1Result() int {
	return getNbFishAfterNDays(80, d.data)
}

func getNbFishAfterNDays(nbDays int, data []int) int {
	ages := make(map[int]int)

	for k := range data {
		ages[data[k]]++
	}

	for i := 0; i < nbDays; i++ {
		copyAges := make(map[int]int)

		for k, v := range ages {
			copyAges[k] = v
		}

		for a, n := range ages {
			if a == 0 {
				copyAges[0] -= n
				copyAges[8] += n
				copyAges[6] += n
			} else {
				copyAges[a] -= n
				copyAges[a-1] += n
			}
		}

		ages = copyAges
	}

	sum := 0
	for _, v := range ages {
		sum += v
	}

	return sum
}

func (d Day6) GetStep2Result() int {
	return getNbFishAfterNDays(256, d.data)
}
