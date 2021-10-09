package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const maxCapacity = 128 * 1024

func main() {
	var (
		splited []string
		longest string
		lenght  int
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
			i++
			continue
		}
		if i == 1 {
			splited = strings.Split(strings.TrimSpace(scanner.Text()), " ")
			break
		}
	}

	for _, w := range splited {
		wLen := len(w)
		if wLen > lenght {
			longest = w
			lenght = wLen
		}
	}

	if err = ioutil.WriteFile("output.txt", []byte(fmt.Sprintf("%s \n%d", longest, lenght)), 0777); err != nil {
		os.Exit(3)
	}
}
