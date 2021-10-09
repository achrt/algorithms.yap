package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const maxCapacity = 20
const filePath = "input.txt"

var primeNumbers = getPrimes(1000000)

func main() {
	num, err := getNumFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	_, result := factorisation(num, []string{})

	if err = ioutil.WriteFile("output.txt", []byte(strings.Join(result[:], " ")), 0777); err != nil {
		os.Exit(3)
	}
}

func factorisation(i int, f []string) (int, []string) {
	var quotient = i

	for _, prime := range primeNumbers {
		remainder := i % prime
		if remainder == 0 {
			f = append(f, strconv.Itoa(prime))
			quotient = i / prime
			break
		}
	}
	if quotient == i {
		if quotient != 1 {
			f = append(f, strconv.Itoa(quotient))
		}
		return i, f
	}
	return factorisation(quotient, f)
}

func getNumFromFile(filePath string) (num int, err error) {
	var file *os.File

	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, maxCapacity), maxCapacity)

	for scanner.Scan() {
		num, err = strconv.Atoi(scanner.Text())
	}
	return
}

func getPrimes(num int) (primes []int) {
	crossout := make([]bool, num)
	for i := 2; i < num; i++ {
		if crossout[i] {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < num; k += i {
			crossout[k] = true
		}
	}
	return
}
