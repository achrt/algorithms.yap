package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const inputValues = 4

var (
	file    *os.File
	splited []string
	line    []byte
	input   []int
	err     error
)

func main() {

	file, err = os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	sc := bufio.NewReader(file)
	if line, _, err = sc.ReadLine(); err != nil {
		os.Exit(2)
	}

	splited = strings.Split(strings.TrimSpace(string(line)), " ")
	if len(splited) != inputValues {
		os.Exit(3)
	}

	input = make([]int, inputValues)

	for i, v := range splited {
		if val, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
			input[i] = val
			continue
		}
		os.Exit(4)
	}

	// y = ax2 + bx + c
	// a, x, b, c

	result := input[0]*(input[1]*input[1]) + (input[2] * input[1]) + input[3]

	if err = ioutil.WriteFile("output.txt", []byte(fmt.Sprintf("%d", result)), 0777); err != nil {
		os.Exit(5)
	}
}
