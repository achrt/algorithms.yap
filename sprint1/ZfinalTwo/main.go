package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/22450/run-report/53172328/

const filePath = "input.txt"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var i, press, result int
	list := map[string]int{}

	for scanner.Scan() {
		val := scanner.Text()
		if i == 0 {
			if press, err = strconv.Atoi(val); err != nil {
				log.Fatal(err)
			}
			i++
			continue
		}
		if val == "\n" || val == "." {
			continue
		}
		list[val] += 1
	}
	for _, count := range list {
		if press*2 >= count {
			result += 1
		}
	}

	if err = ioutil.WriteFile("output.txt", []byte(strconv.Itoa(result)), 0777); err != nil {
		log.Fatal(err)
	}
}
