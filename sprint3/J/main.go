package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const filePath = "input.txt"
const maxCapacity = 1024 * 20

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

	var i, lenght int

	var nums []string
	var digits []int

	for scanner.Scan() {
		val := scanner.Text()
		if i == 0 {
			if lenght, err = strconv.Atoi(val); err != nil {
				log.Fatal(err)
			}
		} else if i == 1 {
			nums = strings.Split(val, " ")
		}
		i++
	}

	digits = make([]int, len(nums))
	for i, n := range nums {
		if digits[i], err = strconv.Atoi(n); err != nil {
			log.Fatal(err)
		}
	}

	os.Stdout.WriteString(strings.Join(BubbleSort(digits, lenght), "\n"))
}

func BubbleSort(digits []int, lenght int) []string {
	var i, f int
	result := make([]string, lenght)
	for {
		f = 0
		for i := 0; i < lenght-1; i++ {
			if digits[i] > digits[i+1] {
				store := digits[i]
				digits[i], digits[i+1] = digits[i+1], store
				f = 1
			}
		}

		if f == 0 {
			if i == 0 {
				result[i] = strings.Join(convert(digits), " ")
			}
			break
		}
		result[i] = strings.Join(convert(digits), " ")
		i++
	}
	return result
}

func convert(digits []int) (nums []string) {
	nums = make([]string, len(digits))
	for i, g := range digits {
		nums[i] = strconv.Itoa(g)
	}
	return
}
