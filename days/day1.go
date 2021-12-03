package day

import (
	_file "advent_of_code/helper/file"
	"bufio"
	"strconv"
)

type Day1 struct {
	data []int
}

func CreateDay1(path string) Day1 {
	r := _file.ReadTextFile(path, func(s *bufio.Scanner) interface{} {
		data := make([]int, 0)

		for s.Scan() {
			v, _ := strconv.Atoi(s.Text())
			data = append(data, v)
		}

		return data
	})
	rOk, ok := r.([]int)

	if ok {
		return Day1{
			data: rOk,
		}
	}

	return Day1{}
}

func (d Day1) GetStep1Result() int {

	countHigher := 0

	for k := range d.data {
		if k != 0 && d.data[k] > d.data[k-1] {
			countHigher++
		}
	}

	return countHigher
}

func (d Day1) GetStep2Result() int {
	sum := make([]int, 0)

	for k := range d.data {
		if k+2 < len(d.data) {
			sum = append(sum, d.data[k]+d.data[k+1]+d.data[k+2])
		}
	}

	countHigher := 0
	for k := range sum {
		if k != 0 && sum[k] > sum[k-1] {
			countHigher++
		}
	}

	return countHigher
}
