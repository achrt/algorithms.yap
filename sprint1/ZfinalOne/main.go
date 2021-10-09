package main

// https://contest.yandex.ru/contest/22450/run-report/53172611/

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const filePath = "input.txt"

func main() {
	lenght, list, err := getListFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	lastEmptyPos := -1
	var value string

	backward := func(currentIter int, lastEmptyPos int) {
		for inner := currentIter; inner > lastEmptyPos; inner-- {
			if lastEmptyPos == -1 { // значение -1 говорит нам, что первый ноль в позиции currentIter, значит нужно пройти обратно весь отрезок
				list[inner] = fmt.Sprintf("%d", currentIter-inner)
			} else { // здесь можно пройти отрезок до половины,т.к. первая половина уже была проставлена во внешнем цикле
				list[inner] = fmt.Sprintf("%v", (math.Min(float64(inner-lastEmptyPos), float64(currentIter-inner))))
				if inner <= (currentIter-lastEmptyPos)/2 {
					break
				}
			}
		}
	}

	for i := 0; i < lenght; i++ {
		value = list[i]
		if value == "0" {
			if i != 0 {
				backward(i, lastEmptyPos)
			}
			lastEmptyPos = i
		}
		list[i] = fmt.Sprintf("%d", i-lastEmptyPos)
	}

	if err = ioutil.WriteFile("output.txt", []byte(strings.Join(list[:], " ")), 0777); err != nil {
		log.Fatal(err)
	}
}

func getListFromFile(filePath string) (lenght int, list []string, err error) {
	var file *os.File
	if file, err = os.Open(filePath); err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var i int
	for scanner.Scan() {
		val := scanner.Text()
		if i == 0 {
			if lenght, err = strconv.Atoi(val); err != nil {
				return
			}
			i++
			continue
		}
		if i == 1 {
			list = make([]string, lenght)
		}
		if val == "0" {
			list[i-1] = val
		}
		i++
	}
	return
}
