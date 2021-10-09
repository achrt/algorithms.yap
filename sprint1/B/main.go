package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	file      *os.File
	splited   []string
	line      []byte
	remainder int
	result    string
	err       error
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
	remainder = -1
	result = "WIN"

	for _, v := range splited {
		var val int
		if val, err = strconv.Atoi(strings.TrimSpace(v)); err != nil {
			os.Exit(3)
		}

		r := val % 2
		if r < 0 {
			r = -r
		}

		if remainder != -1 && r != remainder {
			result = "FAIL"
			break
		}
		remainder = r
	}

	if err = ioutil.WriteFile("output.txt", []byte(result), 0777); err != nil {
		os.Exit(4)
	}
}
