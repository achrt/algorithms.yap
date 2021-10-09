package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const maxCapacity = 10050
const filePath = "input.txt"

func main() {
	var result = "True"

	num, err := getDataFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if num == 1 {
			break
		}
		if num%4 == 0 {
			num = num / 4
		} else {
			result = "False"
			break
		}
	}

	if err = ioutil.WriteFile("output.txt", []byte(result), 0777); err != nil {
		os.Exit(3)
	}
}

func getDataFromFile(filePath string) (i int, err error) {
	var file *os.File

	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, maxCapacity), maxCapacity)

	for scanner.Scan() {
		i, err = strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
	}
	return
}
