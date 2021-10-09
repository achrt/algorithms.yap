package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const filePath = "input.txt"
const maxCapacity = 10000000

func main() {
	var err error

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, maxCapacity), maxCapacity)

	var i, cost, lenght int
	var mb = &MoneyBox{}

	for scanner.Scan() {
		val := scanner.Text()
		if i == 0 {
			if lenght, err = strconv.Atoi(val); err != nil {
				log.Fatal(err)
			}
		} else if i == 1 {
			mb.countByDay = strings.Split(val, " ")
		} else if i == 2 {
			if cost, err = strconv.Atoi(val); err != nil {
				log.Fatal(err)
			}
		}
		i++
	}

	one, err := mb.binarySearch(cost, 0, lenght)
	if err != nil {
		log.Fatal(err)
	}

	two, err := mb.binarySearch(cost*2, 0, lenght)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.WriteString(fmt.Sprintf("%d %d", one, two))
}

type MoneyBox struct {
	countByDay []string
}

func (mb *MoneyBox) binarySearch(cost, left, right int) (int, error) {
	var previousVal, val int
	var err error

	if right <= left {
		return -1, nil
	}
	index := (right + left) / 2
	if val, err = strconv.Atoi(mb.countByDay[index]); err != nil {
		return -1, err
	}

	if index > 0 {
		if previousVal, err = strconv.Atoi(mb.countByDay[index-1]); err != nil {
			return -1, err
		}
	}

	if val >= cost && previousVal < cost {
		return index + 1, nil
	}
	if val < cost {
		return mb.binarySearch(cost, index+1, right)
	}
	return mb.binarySearch(cost, left, index)
}
