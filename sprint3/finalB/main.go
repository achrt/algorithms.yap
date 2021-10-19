package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
-- ПРИНЦИП РАБОТЫ --
В исходном масиве A выбирается опорный элемент, осуществляется проход по массиву слева на право
со сравнением элементов. Если левый элемент больше правого, производится замена левого элемента на опорный.
По завершению прохода возвращается опорный элемент и функция выполняется снова для новых отрезков.
Базовым случаем считается указание на отрезок из 0 или 1 элементов.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Алгоритм рекурсивно упорядочивает массив, разбивая его на части и упорядочивая элементы с
использованием опорного, отправляя в левую часть элементы меньше опорного и в правую часть
элементы больше опорного. Таким образом, когда рекурсия дойдет до базового случая, все элементы будут
проверены и отсортированы.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
В худшем случае, при подаче на вход отсортированного массива или массива где все элементы равны
сортировка будет выполняться за квадратичное время. Худшим случаем также можно считать выпадение
pivot на каждой итерации равным наименьшему или наибольшему значению.

В среднем, если pivot будет делить массив на части длиной примерно от 25% до 75%, то глубина рекурсии
не будет превышать O(log n) + на каждом уровне рекурсии выполняется не более O(n) операций, то средняя
сложность получится в сумме O(n log n)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Пространственная сложность O(1), т.к. значения в исходном массиве меняются местами без использования
доп памяти.
*/

// https://contest.yandex.ru/contest/23815/run-report/55014790/

const filePath = "input.txt"
const maxCapacity = 100000

func main() {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, maxCapacity), maxCapacity)

	var i int
	var rowData []string

	list := CompetitionList{}

	for scanner.Scan() {
		val := scanner.Text()
		if i == 0 {
			list.CompetitorsCount, err = strconv.Atoi(val)
			if err != nil {
				log.Fatal()
			}
			list.Competitors = make([]Competitor, list.CompetitorsCount)
			i++
			continue
		}

		if i > list.CompetitorsCount {
			break
		}

		rowData = strings.Split(val, " ")
		c := Competitor{Nickname: rowData[0]}
		c.Score, err = strconv.Atoi(rowData[1])
		if err != nil {
			log.Fatal(err)
		}
		c.Penalty, err = strconv.Atoi(rowData[2])
		if err != nil {
			log.Fatal(err)
		}

		list.Competitors[i-1] = c
		i++
	}

	list.QuickSort(0, list.Len()-1)
	list.PrintOut(1000)
}

type Competitor struct {
	Nickname string
	Score    int
	Penalty  int
}

type CompetitionList struct {
	CompetitorsCount int
	Competitors      []Competitor
	counter          int
}

func (c *CompetitionList) Len() int { return c.CompetitorsCount }

func (c *CompetitionList) Less(i, j int) bool {
	if c.Competitors[i].Score != c.Competitors[j].Score {
		return c.Competitors[i].Score < c.Competitors[j].Score
	}
	if c.Competitors[i].Penalty != c.Competitors[j].Penalty {
		return c.Competitors[i].Penalty > c.Competitors[j].Penalty
	}
	return c.Competitors[i].Nickname >= c.Competitors[j].Nickname
}

func (c *CompetitionList) Swap(i, j int) {
	c.Competitors[i], c.Competitors[j] = c.Competitors[j], c.Competitors[i]
}

func (c *CompetitionList) QuickSort(left, right int) {
	if left >= right {
		return
	}

	p := c.partition(left, right)
	if p > 0 {
		c.QuickSort(left, p-1)
	}
	if p+1 < c.Len() {
		c.QuickSort(p+1, right)
	}
}

func (c *CompetitionList) partition(left int, right int) int {
	c.Swap(left, right)
	p := left
	for i := left; i < right; i++ {
		if !c.Less(i, right) {
			c.Swap(i, p)
			p++
		}
	}
	c.Swap(p, right)
	return p
}

func (c *CompetitionList) PrintOut(chunkSize int) {
	left, right := 0, chunkSize
	for {
		if right >= c.Len() {
			right = c.Len()
		}

		buff := make([]string, chunkSize)
		var counter int
		for i := left; i < right; i++ {
			buff[counter] = c.Competitors[i].Nickname + "\n"
			counter++
		}

		os.Stdout.WriteString(strings.Join(buff, ""))
		if right == c.Len() {
			break
		}
		left, right = right, right+chunkSize
	}
}
