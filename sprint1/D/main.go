package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const maxCapacity = 512 * 1024
const filePath = "input.txt"

func main() {
	var result int

	lenght, data, err := getDataFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	if lenght == 1 {
		result = 1
	} else {

		for i, v := range data {
			if i == 0 && v > data[i+1] {
				result += 1
			}
			if i == lenght-1 && v > data[i-1] {
				result += 1
			}

			if i != 0 && i != lenght-1 {
				if v > data[i-1] && v > data[i+1] {
					result += 1
				}
			}
		}

	}

	if err = ioutil.WriteFile("output.txt", []byte(fmt.Sprintf("%d", result)), 0777); err != nil {
		os.Exit(3)
	}
}

func getDataFromFile(filePath string) (lenght int, data []int, err error) {
	var file *os.File

	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, maxCapacity), maxCapacity)

	var i int
	for scanner.Scan() {
		if i == 0 {
			lenght, err = strconv.Atoi(scanner.Text())
			if err != nil {
				return
			}
		}
		if i == 1 {
			splitted := strings.Split(strings.TrimSpace(string(scanner.Text())), " ")
			data = make([]int, len(splitted))
			for i, sp := range splitted {
				if data[i], err = strconv.Atoi(sp); err != nil {
					return
				}
			}
		}
		i++
	}
	return
}
