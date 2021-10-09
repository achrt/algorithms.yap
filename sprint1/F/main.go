package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const maxCapacity = 32 * 1024

func main() {
	var (
		splited     []string
		left, right string
		result      = "True"
	)

	var re = regexp.MustCompile(`a|b|c|d|e|f|g|h|i|j|k|l|m|n|o|p|q|r|s|t|u|v|w|x|y|z|0|1|2|3|4|5|6|7|8|9`)

	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		splited = strings.Split(scanner.Text(), "")
	}

	leftPartLen := len(splited)/2 + len(splited)%2

	var leftIndex, rightIndex = 0, len(splited) - 1
	var match bool

	for i := 0; i < leftPartLen; i++ {
		left = strings.ToLower(splited[leftIndex])
		leftIndex++

		if !re.MatchString(left) {
			continue
		}

		for rightIndex >= leftPartLen {
			right = strings.ToLower(splited[rightIndex])
			rightIndex--
			if re.MatchString(right) {
				match = true
				break
			}
			match = false
		}

		if !match {
			break
		}

		if left != right {
			result = "False"
			break
		}
	}

	if err = ioutil.WriteFile("output.txt", []byte(result), 0777); err != nil {
		os.Exit(3)
	}
}
