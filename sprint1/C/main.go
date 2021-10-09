package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

const maxCapacity = 1024 * 1024
const fileName = "input.txt"

type matrix struct {
	lenght, width, coordX, coordY int
	matrix                        [][]int
}

func main() {

	matrix, err := read(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	n := matrix.getNeighbours()
	result := make([]string, len(n))

	for i := range n {
		result[i] = strconv.Itoa(n[i])
	}

	if err = ioutil.WriteFile("output.txt", []byte(strings.Join(result, " ")), 0777); err != nil {
		os.Exit(4)
	}
}

func (m *matrix) getNeighbours() (neighbours []int) {
	left := m.coordX - 1
	if left >= 0 {
		neighbours = append(neighbours, m.matrix[m.coordY][left])
	}

	right := m.coordX + 1
	if right < m.width {
		neighbours = append(neighbours, m.matrix[m.coordY][right])
	}

	up := m.coordY - 1
	if up >= 0 {
		neighbours = append(neighbours, m.matrix[up][m.coordX])
	}

	down := m.coordY + 1
	if down < m.lenght {
		neighbours = append(neighbours, m.matrix[down][m.coordX])
	}
	sort.Ints(neighbours)
	return
}

func read(fileName string) (*matrix, error) {
	var (
		file    *os.File
		splited []string
		err     error
	)

	m := &matrix{}

	file, err = os.Open(fileName)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var i, iMatrix int

	for scanner.Scan() {
		txt := scanner.Text()
		if i == 0 {
			if m.lenght, err = strconv.Atoi(strings.TrimSpace(txt)); err != nil {
				return nil, err
			}
			m.matrix = make([][]int, m.lenght)
			i++
			continue
		}

		if i == 1 {
			if m.width, err = strconv.Atoi(strings.TrimSpace(txt)); err != nil {
				return nil, err
			}
			i++
			continue
		}

		if i < m.lenght+2 {
			row := make([]int, m.width)
			splited = strings.Split(string(txt), " ")

			if len(splited) != m.width {
				return nil, fmt.Errorf("lenght of string is incorrect")
			}

			for i, rs := range splited {
				if r, err := strconv.Atoi(strings.TrimSpace(rs)); err == nil {
					row[i] = r
					continue
				}
				return nil, err
			}
			m.matrix[iMatrix] = row

			iMatrix++
			i++
			continue
		}

		if i == m.lenght+2 {
			if m.coordY, err = strconv.Atoi(strings.TrimSpace(txt)); err != nil {
				return nil, err
			}
			i++
			continue
		}

		if i == m.lenght+3 {
			if m.coordX, err = strconv.Atoi(strings.TrimSpace(txt)); err != nil {
				return nil, err
			}
			break
		}
	}
	return m, nil
}
