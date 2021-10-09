package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const filePath = "input.txt"

type Stack struct {
	stack []int
	len   int
}

func (s *Stack) getMax() (int, error) {
	if s.len == 0 {
		return 0, errors.New("None")
	}
	var max = s.stack[0]
	for _, v := range s.stack {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func (s *Stack) push(val int) {
	s.stack = append(s.stack, val)
	s.len++
}

func (s *Stack) pop() error {
	max := len(s.stack)
	if max > 0 {
		s.stack = s.stack[:max-1]
		s.len--
		return nil
	}
	return errors.New("error")
}

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var i int

	var stack = &Stack{}
	for scanner.Scan() {
		val := scanner.Text()
		if i == 0 {
			i++
			continue
		}

		if val == "get_max" {
			v, err := stack.getMax()
			if err != nil {
				os.Stdout.WriteString(err.Error() + "\n")
			} else {
				os.Stdout.WriteString(strconv.Itoa(v) + "\n")
			}
			continue
		}

		if val == "pop" {
			if err := stack.pop(); err != nil {
				os.Stdout.WriteString(err.Error() + "\n")
			}
			continue
		}

		push := strings.Split(val, " ")
		v, err := strconv.Atoi(push[1])
		if err != nil {
			log.Fatal(err)
		}
		stack.push(v)
	}
}
