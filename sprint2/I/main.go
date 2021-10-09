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

func main() {
	// v := time.Now().UnixNano()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var i = -2
	var toPrint []string
	var queue *MyQueueSized

	for scanner.Scan() {
		val := scanner.Text()

		if i == -2 {
			buff, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			toPrint = make([]string, buff)
			i++
			continue
		}

		if i == -1 {
			if size, err := strconv.Atoi(val); err == nil {
				queue = New(size)
			} else {
				log.Fatal(err)
			}
			i++
			continue
		}

		if val == "pop" {
			v, err := queue.Pop()
			if err != nil {
				toPrint[i] = err.Error() + "\n"
			} else {
				toPrint[i] = strconv.Itoa(v) + "\n"
			}
		} else if val == "peek" {
			v, err := queue.Peek()
			if err != nil {
				toPrint[i] = err.Error() + "\n"
			} else {
				toPrint[i] = strconv.Itoa(v) + "\n"
			}
		} else if val == "size" {
			toPrint[i] = strconv.Itoa(queue.Size()) + "\n"
		} else {
			// тут, по идее, должны остаться только команды типа push 5
			push := strings.Split(val, " ")
			v, err := strconv.Atoi(push[1])
			if err != nil {
				log.Fatal(err)
			}
			if err = queue.Push(v); err != nil {
				toPrint[i] = err.Error() + "\n"
			}

		}
		i++
	}

	if len(toPrint) > 0 {
		os.Stdout.WriteString(strings.Join(toPrint, ""))
	}
	// fmt.Println(time.Now().UnixNano() - v)
}

type MyQueueSized struct {
	queue    []int
	capacity int
	tail     int
	head     int
	size     int
}

func New(size int) *MyQueueSized {
	return &MyQueueSized{
		queue:    make([]int, size),
		capacity: size,
	}
}

func (q *MyQueueSized) Push(i int) error {
	if q.size == q.capacity {
		return errors.New("error")
	}
	q.queue[q.tail] = i
	q.tail++
	q.size++

	if q.tail == q.capacity {
		q.tail = 0
	}
	return nil
}

func (q *MyQueueSized) Pop() (int, error) {
	if q.size == 0 {
		return 0, errors.New("None")
	}

	i := q.queue[q.head]
	q.queue[q.head] = 0

	q.head++
	q.size--

	if q.tail == q.capacity {
		q.tail = 0
	}

	if q.head == q.capacity {
		q.head = 0
	}

	return i, nil
}

func (q *MyQueueSized) Peek() (int, error) {
	if q.size == 0 {
		return 0, errors.New("None")
	}
	return q.queue[q.head], nil
}

func (q *MyQueueSized) Size() int {
	return q.size
}
