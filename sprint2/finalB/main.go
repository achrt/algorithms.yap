package main

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"strconv"
)

/*
-- ПРИНЦИП РАБОТЫ --
Чтение входной строки происходит слева на право, для хранения операндов используется стек.
Если значение - знак арифметической операции, то из стека извлекаются два значения. Результат операции помещается на
вершину стека. По завершении работы в стеке остается один элемент - значение выражения. 

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Стек хранит данные в порядке их добавления, тем самым гарантируя корректность вычислений.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Алгоритм вычисления арифметического выражения в постфиксной нотации имеет линейную сложность - O(n). 
т.к. для вычисления значения выражения вся строка должна быть обработана.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Затраты на память константные O(1), задаются через maxCapacity при создании стека

*/

// https://contest.yandex.ru/contest/22781/run-report/53807698/

type Stack struct {
	capacity int
	size     int
	stack    []int
	tail     int
}

const maxCapacity = 500

func New() Stack {
	return Stack{
		capacity: maxCapacity,
		stack:    make([]int, maxCapacity),
	}
}

func (s *Stack) Push(n int) (err error) {
	if s.size == s.capacity {
		err = errors.New("stack overflow")
		return
	}
	s.stack[s.tail] = n
	s.tail++
	s.size++
	return
}

func (s *Stack) Pop() (n int, err error) {
	if s.size == 0 {
		err = errors.New("empty stack")
		return
	}
	n = s.stack[s.tail-1]
	s.tail--
	s.size--
	return
}

func (s *Stack) GetOperands() (o1, o2 int, err error) {
	if o2, err = s.Pop(); err != nil {
		return
	}
	if o1, err = s.Pop(); err != nil {
		return
	}
	return
}

const filePath = "input.txt"

func main() {
	var err error

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var v, o1, o2, quotient int
	var stack = New()

	for scanner.Scan() {
		val := scanner.Text()

		if val == "+" {
			if o1, o2, err = stack.GetOperands(); err != nil {
				log.Fatal(err)
			}
			if err = stack.Push(o1 + o2); err != nil {
				log.Fatal(err)
			}

		} else if val == "-" {
			if o1, o2, err = stack.GetOperands(); err != nil {
				log.Fatal(err)
			}
			if err = stack.Push(o1 - o2); err != nil {
				log.Fatal(err)
			}
		} else if val == "*" {
			if o1, o2, err = stack.GetOperands(); err != nil {
				log.Fatal(err)
			}
			if err = stack.Push(o1 * o2); err != nil {
				log.Fatal(err)
			}

		} else if val == "/" {
			if o1, o2, err = stack.GetOperands(); err != nil {
				log.Fatal(err)
			}

			if o1 < 0 {
				quotient = int(math.Floor(float64(o1) / float64(o2)))
			} else {
				quotient = o1 / o2
			}

			if err = stack.Push(quotient); err != nil {
				log.Fatal(err)
			}
		} else {
			if v, err = strconv.Atoi(val); err != nil {
				log.Fatal(err)
			}
			if err = stack.Push(v); err != nil {
				log.Fatal(err)
			}
		}
	}

	val, err := stack.Pop()
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.WriteString(strconv.Itoa(val))
}
