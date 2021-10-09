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
const maxCapacity = 1024

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

	var i, cost int
	var 

	mb := &MoneyBox{}

	for scanner.Scan() {
		val := scanner.Text()
		if i == 0 {
			if mb.days, err = strconv.Atoi(val); err != nil {
				log.Fatal(err)
			}
		} else if i == 1 {
			mb.amounts = strings.Split(val, " ")
		} else if i == 2 {
			if cost, err = strconv.Atoi(val); err != nil {
				log.Fatal(err)
			}
		}
		i++
	}

	ret, err := mb.binarySearch(cost, 0, mb.days-1)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.WriteString(fmt.Sprintf("%d %d", ret, ret))
}

type MoneyBox struct {
	amounts []string
	days    int
}

func (mb *MoneyBox) binarySearch(cost, left, right int) (int, error) {
	if right <= left {
		return -1, nil
	}
	index := (right + left) / 2
	val, err := strconv.Atoi(mb.amounts[left+index])
	if err != nil {
		return -1, err
	}
	if val == cost {
		return index, nil
	}
	if val > cost {
		return mb.binarySearch(cost, index, right)
	}
	
	return mb.binarySearch(cost, left, index)
}
