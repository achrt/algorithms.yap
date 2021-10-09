package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const maxCapacity = 16 * 1024

func main() {
	var (
		termOne, termTwo int
		result           []string
	)

	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, maxCapacity), maxCapacity)

	var i int
	for scanner.Scan() {
		if i == 1 {
			var tmp string
			for _, v := range strings.Split(strings.TrimSpace(scanner.Text()), " ") {
				tmp += v
			}
			if termOne, err = strconv.Atoi(tmp); err != nil {
				os.Exit(2)
			}

		}
		if i == 2 {
			termTwo, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				os.Exit(3)
			}
		}
		i++
	}
	res := strconv.Itoa(termOne + termTwo)
	result = make([]string, len(res))
	for i, r := range res {
		result[i] = string(r)
	}

	if err = ioutil.WriteFile("output.txt", []byte([]byte(strings.Join(result[:], " "))), 0777); err != nil {
		os.Exit(3)
	}
}
