package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
-- ПРИНЦИП РАБОТЫ --
Структура Deq реализована на массиве фиксированной длины с использованием двух указателей на начало и конец очереди.
Tail и Head всегда указывают на последнюю заполненную ячейку с каждой строны.
Если size == 0, то head == 0, а tail указывает на последнюю ячейку (capacity - 1).


-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Дек - структура, в которой добавление и извлечение элементов возможно с обеих сторон. Массив фиксированной длины
с двумя указателями позволяет реализовать методы pushFront(), pushBack(), popBack() и popFront() для выполнения условий.


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Добавление и удаление в дек стоят O(1), т.к. всегда есть указатель на последнюю заполненную ячейку с обеих сторон.
Параметр size также всегда поддерживаетcя в актуальном состоянии, поэтому запрос текущего размера тоже стоит O(1).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Дек занимает O(n) памяти и только на хренение элементов
*/

// https://contest.yandex.ru/contest/22781/run-report/53797002/

const filePath = "input.txt"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var i = -2
	var toPrint []string
	var deq *Deq
	var v int

	for scanner.Scan() {
		val := scanner.Text()

		if i == -2 {
			if buffSize, err := strconv.Atoi(val); err == nil {
				toPrint = make([]string, buffSize) // инициализация буффера для печати
			} else {
				log.Fatal(err)
			}
			i++
			continue
		}

		if i == -1 {
			if qsize, err := strconv.Atoi(val); err == nil {
				deq = New(qsize) // создание очереди
			} else {
				log.Fatal(err)
			}
			i++
			continue
		}

		if val == "pop_back" {

			if v, err = deq.PopBack(); err != nil {
				toPrint[i] = err.Error() + "\n"
			} else {
				toPrint[i] = strconv.Itoa(v) + "\n"
			}

		} else if val == "pop_front" {

			if v, err = deq.PopFront(); err != nil {
				toPrint[i] = err.Error() + "\n"
			} else {
				toPrint[i] = strconv.Itoa(v) + "\n"
			}

		} else {
			// тут, по условию задачи, должны остаться команды типа push_front -201 или push_back 959
			push := strings.Split(val, " ")
			if push[0] == "push_front" {

				if v, err = strconv.Atoi(push[1]); err != nil {
					log.Fatal(err)
				}
				if err = deq.PushFront(v); err != nil {
					toPrint[i] = err.Error() + "\n"
				}

			} else {

				if v, err = strconv.Atoi(push[1]); err != nil {
					log.Fatal(err)
				}
				if err = deq.PushBack(v); err != nil {
					toPrint[i] = err.Error() + "\n"
				}
			}
		}
		i++
	}

	if len(toPrint) > 0 {
		os.Stdout.WriteString(strings.Join(toPrint, ""))
	}
}

type Deq struct {
	queue    []int
	head     int
	tail     int
	capacity int
	size     int
}

func New(n int) *Deq {
	return &Deq{
		queue:    make([]int, n),
		capacity: n,
		tail:     n - 1,
	}
}

func (d *Deq) PushFront(n int) error {
	if d.size == d.capacity {
		return errors.New("error")
	}
	d.queue[d.head] = n

	d.head++
	d.size++

	if d.head == d.capacity {
		d.head = 0
	}
	return nil
}

func (d *Deq) PushBack(n int) error {
	if d.size == d.capacity {
		return errors.New("error")
	}
	d.queue[d.tail] = n

	d.tail--
	d.size++

	if d.tail < 0 {
		d.tail = d.capacity - 1
	}
	return nil
}

func (d *Deq) PopFront() (n int, err error) {
	if d.size == 0 {
		err = errors.New("error")
		return
	}
	d.head--
	if d.head < 0 {
		d.head = d.capacity - 1
	}

	n = d.queue[d.head]
	d.size--
	return
}

func (d *Deq) PopBack() (n int, err error) {
	if d.size == 0 {
		err = errors.New("error")
		return
	}
	d.tail++
	if d.tail == d.capacity {
		d.tail = 0
	}

	n = d.queue[d.tail]
	
	d.size--
	return
}
