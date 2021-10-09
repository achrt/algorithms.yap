package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

const maxCapacity = 16 * 1024
const filePath = "input.txt"

func main() {
	long, short, lenOfShort, err := getSlicesFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var result string
	for i, v := range long {
		if i == lenOfShort {
			result = v
			break
		}
		if v != short[i] {
			result = v
			break
		}
	}

	if err = ioutil.WriteFile("output.txt", []byte(result), 0777); err != nil {
		os.Exit(3)
	}
}

func getSlicesFromFile(filePath string) (long []string, short []string, lenOfShort int, err error) {
	var file *os.File

	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, maxCapacity), maxCapacity)

	var lineOne, lineTwo []string

	var i int
	for scanner.Scan() {
		if i == 0 {
			lineOne = strings.Split(strings.TrimSpace(scanner.Text()), "")
		}
		if i == 1 {
			lineTwo = strings.Split(strings.TrimSpace(scanner.Text()), "")
		}
		i++
	}

	lenOne := len(lineOne)
	lenTwo := len(lineTwo)

	sort.Strings(lineOne)
	sort.Strings(lineTwo)

	if lenOne > lenTwo {
		long, short, lenOfShort = lineOne, lineTwo, lenTwo
		return
	}
	long, short, lenOfShort = lineTwo, lineOne, lenOne
	return
}
