// package main

// import (
// 	"bufio"
// 	"bytes"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// const filePath = "_input.txt"

// type StackMax struct {
// 	stack  []int
// 	len    int
// 	maxVal int
// }

// func newStack(size int) *StackMax {
// 	return &StackMax{}
// }

// func (s *StackMax) getMax() (int, error) {
// 	if s.len == 0 {
// 		return 0, errors.New("None")
// 	}
// 	return s.maxVal, nil
// }

// func (s *StackMax) setMax() {
// 	if s.len == 0 {
// 		s.maxVal = 0
// 		return
// 	}
// 	s.maxVal = s.stack[0]
// 	for _, v := range s.stack {
// 		if v > s.maxVal {
// 			s.maxVal = v
// 		}
// 	}
// }

// func (s *StackMax) push(val int) {
// 	if s.len == 0 {
// 		s.maxVal = val
// 	} else if val > s.maxVal {
// 		s.maxVal = val
// 	}
// 	s.stack = append(s.stack, val)
// 	s.len++
// }

// func (s *StackMax) pop() error {
// 	if s.len > 0 {
// 		s.stack = s.stack[:s.len-1]
// 		s.len--
// 		s.setMax()
// 		return nil
// 	}
// 	return errors.New("error")
// }

// func main() {
// 	v := time.Now().UnixNano()

// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)

// 	var i, buff = -1, 10000

// 	toPrint := make([][]byte, buff)

// 	var stack *StackMax
// 	for scanner.Scan() {

// 		if i == buff {
// 			os.Stdout.Write(bytes.Join(toPrint, []byte("")))
// 			i = 0
// 			toPrint = make([][]byte, buff)
// 		}

// 		val := scanner.Text()
// 		if i == -1 {
// 			v, err := strconv.Atoi(val)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			stack = newStack(v)
// 			i++
// 			continue
// 		}

// 		if val == "get_max" {
// 			v, err := stack.getMax()
// 			if err != nil {
// 				toPrint[i] = []byte(err.Error() + "\n")
// 			} else {
// 				toPrint[i] = []byte(strconv.Itoa(v) + "\n")
// 			}
// 			i++
// 			continue
// 		}

// 		if val == "pop" {
// 			if err := stack.pop(); err != nil {
// 				toPrint[i] = []byte(err.Error() + "\n")
// 			}
// 			i++
// 			continue
// 		}

// 		push := strings.Split(val, " ")
// 		v, err := strconv.Atoi(push[1])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		stack.push(v)
// 	}

// 	if len(toPrint) > 0 {
// 		os.Stdout.Write(bytes.Join(toPrint, []byte("")))
// 	}

// 	fmt.Println(time.Now().UnixNano() - v)
// }
