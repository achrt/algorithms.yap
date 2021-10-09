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

type Queue struct {
	head *queueElement
	tail *queueElement
	size int
}

type queueElement struct {
	data     int
	previous *queueElement
	next     *queueElement
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Get() (int, error) {
	if q.size == 0 {
		return 0, errors.New("error")
	}
	ret := q.head.data
	if q.head.next != nil {
		q.head.next.previous = nil
		q.head = q.head.next
	} else {
		q.head = nil
	}
	q.size--
	return ret, nil
}

func (q *Queue) Put(i int) {
	qn := &queueElement{
		previous: q.tail,
		data:     i,
	}
	if q.tail == nil {
		q.tail = &queueElement{}
	}
	q.tail.next = qn
	q.tail = qn

	if q.size == 0 {
		q.head = qn
	}
	q.size++
}

func (q *Queue) Size() int {
	return q.size
}

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var i = -1
	var toPrint []string
	var queue = New()

	for scanner.Scan() {
		val := scanner.Text()

		if i == -1 {
			if size, err := strconv.Atoi(val); err == nil {
				toPrint = make([]string, size) // задаем буффер для печати
			} else {
				log.Fatal(err)
			}
			i++
			continue
		}

		if val == "size" {
			toPrint[i] = strconv.Itoa(queue.Size()) + "\n"
			i++
		} else if val == "get" {
			v, err := queue.Get()
			if err != nil {
				toPrint[i] = err.Error() + "\n"
			} else {
				toPrint[i] = strconv.Itoa(v) + "\n"
			}
			i++
		} else {
			// тут, по идее, должны остаться только команды типа put -23
			push := strings.Split(val, " ")
			v, err := strconv.Atoi(push[1])
			if err != nil {
				log.Fatal(err)
			}
			queue.Put(v)
		}
	}

	if len(toPrint) > 0 {
		os.Stdout.WriteString(strings.Join(toPrint, ""))
	}
}
