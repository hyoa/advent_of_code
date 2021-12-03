package helper

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadTextFileWithParser(path string, parser func(*bufio.Scanner) interface{}) interface{} {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	return parser(scanner)
}

func ReadTextFile(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	strings := make([]string, 0)

	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings
}

func ReadTextFileLinesAsInt(path string) []int {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := make([]int, 0)

	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, v)
	}

	return numbers
}
