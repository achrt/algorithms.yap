package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const filePath = "input.txt"
const maxCapacity = 1024

func main() {

	var err error

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, maxCapacity), maxCapacity)

	p := Pad{
		numpad:       pad,
		combinations: []string{},
	}

	var input string

	for scanner.Scan() {
		input = scanner.Text()
	}

	var digits = make([]int, len(input))
	for i, n := range input {
		num, err := strconv.Atoi(string(n))
		if err != nil {
			log.Fatal(err)
		}
		digits[i] = num
	}

	p.DFS(digits, "")

	os.Stdout.WriteString(strings.Join(p.combinations, " "))

}

var pad = [][]string{
	0: {},
	1: {},
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

type Pad struct {
	numpad       [][]string
	combinations []string
}

func (p *Pad) DFS(digits []int, buff string) {
	if len(buff) == len(digits) {
		p.combinations = append(p.combinations, buff)
		return
	}

	for _, letter := range p.numpad[digits[len(buff)]] {
		p.DFS(digits, buff+letter)
	}
}
