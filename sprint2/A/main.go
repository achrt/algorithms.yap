package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const filePath = "input.txt"
const filePathOut = "output.txt"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var i, lenght, width int
	var counter, size int
	var keepNext = 1

	var res []int

	for scanner.Scan() {
		val := scanner.Text()

		if i == 0 {
			if lenght, err = strconv.Atoi(val); err != nil {
				log.Fatal(err)
			}
			i++
			continue
		}
		if i == 1 {
			if width, err = strconv.Atoi(val); err != nil {
				log.Fatal(err)
			}
			i++
			continue
		}
		if i == 2 {
			size = width * lenght
			res = make([]int, size)
			i++
		}

		if counter >= size {
			counter = keepNext
			keepNext += 1
		}

		v, _ := strconv.Atoi(val)
		res[counter] = v
		counter += lenght

	}

	var f *os.File
	if f, err = os.OpenFile(filePathOut, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777); err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var rowCounter = 0

	buffSize := lenght
	buffer := make([]string, buffSize)
	buffCounter := 0

	for i := 0; i < size; i++ {
		rowCounter++
		val := strconv.Itoa(res[i])
		if rowCounter >= lenght {
			val += "\n"
			rowCounter = 0
		} else {
			val += " "
		}
		buffer[buffCounter] = val
		buffCounter++

		if buffCounter == buffSize || i == size-1 {
			if _, err = f.WriteString(strings.Join(buffer, "")); err != nil {
				log.Fatal(err)
			}
			buffCounter = 0
		}

	}
}

// Alloc = 24 MiB	TotalAlloc = 24 MiB	Sys = 70 MiB	NumGC = 3
// Alloc = 156 MiB	TotalAlloc = 2528 MiB	Sys = 203 MiB	NumGC = 36
