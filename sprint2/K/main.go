package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const filePath = "input.txt"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		if v, err := strconv.Atoi(scanner.Text()); err == nil {
			os.Stdout.WriteString(strconv.Itoa(commits(v)))
		}
	}
}

func commits(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return commits(n-1) + commits(n-2)
}
