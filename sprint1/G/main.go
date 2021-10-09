package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const maxCapacity = 1024

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, maxCapacity), maxCapacity)

	var num int

	for scanner.Scan() {
		if num, err = strconv.Atoi(strings.TrimSpace(scanner.Text())); err != nil {
			os.Exit(2)
		}
	}

	var binary string

	for num > 0 {
		binary = fmt.Sprintf("%d", num%2) + binary
		num /= 2
	}

	if err = ioutil.WriteFile("output.txt", []byte(binary), 0777); err != nil {
		os.Exit(3)
	}
}
