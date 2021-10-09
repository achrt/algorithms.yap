package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

const filePath = "input.txt"

func main() {
	os.Stdout.WriteString(isCorrectBracketSeq())
}

func isCorrectBracketSeq() string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var square, braces, round int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	result := "False"

	var previous string
	var prevOpen, isEmpty bool

	var re = regexp.MustCompile(`\(\)|\[\]|\{\}`)

	for scanner.Scan() {

		val := scanner.Text()

		switch val {
		case "{":
			prevOpen = true
			braces++
		case "}":
			if prevOpen && !re.MatchString(previous+val) {
				return result
			}
			prevOpen = false
			braces--
		case "(":
			prevOpen = true
			round++
		case ")":
			if prevOpen && !re.MatchString(previous+val) {
				return result
			}
			prevOpen = false
			round--
		case "[":
			prevOpen = true
			square++
		case "]":
			if prevOpen && !re.MatchString(previous+val) {
				return result
			}
			prevOpen = false
			square--
		default:
			if square == 0 && braces == 0 && round == 0 { // иногда в тестах есть 2 строки, одна со скобками, вторая пустая
				isEmpty = true
			}
		}

		// если количество закрывающих скобок превосходит открывающие, то можно выходить
		if square < 0 || braces < 0 || round < 0 {
			return result
		}

		previous = val

	}

	if (square == 0 && braces == 0 && round == 0) || isEmpty {
		result = "True"
	}
	return result
}
