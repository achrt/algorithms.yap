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
		first, second []string
		result        []string
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
		if i == 0 {
			first = strings.Split(strings.TrimSpace(scanner.Text()), "")
			i++
		}
		if i == 1 {
			second = strings.Split(strings.TrimSpace(scanner.Text()), "")
		}
	}

	if len(first) == 0 || len(second) == 0 {
		os.Exit(2)
	}

	// дополняем пустыми строками меньший слайс
	difference := len(first) - len(second)
	if difference > 0 {
		second = feedSlice(second, difference)
	} else if difference < 0 {
		first = feedSlice(first, difference)
	}

	resultIndex := len(first)
	rightIndex := resultIndex - 1

	result = make([]string, resultIndex+1)
	var termOne, termTwo, carryover int

	for {
		if termOne, err = getTerm(first[rightIndex]); err != nil {
			os.Exit(3)
		}

		if termTwo, err = getTerm((second[rightIndex])); err != nil {
			os.Exit(4)
		}

		res := termOne + termTwo + carryover
		if res < 2 {
			result[resultIndex] = strconv.Itoa(res)
			carryover = 0
		} else if res == 2 {
			result[resultIndex] = "0"
			carryover = 1
		} else {
			result[resultIndex] = "1"
			carryover = 1
		}

		resultIndex--

		if resultIndex == 0 {
			if carryover > 0 {
				result[resultIndex] = strconv.Itoa(carryover)
			}
			break
		}

		rightIndex--
	}

	if err = ioutil.WriteFile("output.txt", []byte([]byte(strings.Join(result[:], ""))), 0777); err != nil {
		os.Exit(3)
	}
}

func feedSlice(slice []string, difference int) []string {
	if difference < 0 {
		difference = -difference
	}
	return append(make([]string, difference), slice...)
}

func getTerm(t string) (int, error) {
	if t == "" {
		return 0, nil
	}
	return strconv.Atoi(t)
}
