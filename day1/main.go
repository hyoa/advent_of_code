package main

import (
	_file "advent_of_code/helper/file"
	"bufio"
	"fmt"
	"strconv"
)

func main() {

	r := _file.ReadTextFile("input.txt", func(s *bufio.Scanner) interface{} {
		data := make([]int, 0)

		for s.Scan() {
			v, _ := strconv.Atoi(s.Text())
			data = append(data, v)
		}

		return data
	})

	rOk, ok := r.([]int)

	if ok {
		input1(rOk)
		input2(rOk)
	}
}

func input1(data []int) {

	countHigher := 0

	for k := range data {
		if k != 0 && data[k] > data[k-1] {
			countHigher++
		}
	}

	fmt.Println("step1 ", countHigher)
}

func input2(data []int) {
	sum := make([]int, 0)

	for k := range data {
		if k+2 < len(data) {
			sum = append(sum, data[k]+data[k+1]+data[k+2])
		}
	}

	countHigher := 0
	for k := range sum {
		if k != 0 && sum[k] > sum[k-1] {
			countHigher++
		}
	}

	fmt.Println("step2 ", countHigher)
}
