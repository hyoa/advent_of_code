package helper

import (
	"bufio"
	"log"
	"os"
)

func ReadTextFile(path string, parser func(*bufio.Scanner) interface{}) interface{} {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	return parser(scanner)
}
