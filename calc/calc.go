package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

// 4***100
// 4**-100 - ошибка, но не корректно обработана
// 4+-100 - ошибка, но не корректно обработана
// 4+(-100) считает
func main() {
	var inputString string
	inputString = os.Args[1]
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	result, err := Calculator(stringToCalc)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(result)
	}
}

type Stack []string

func (s *Stack) Push(value string) {
	*s = append(*s, value)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value
}

func Calculator(inputSting []string) (int, error) {
	inputSting = unionBigNumbers(inputSting)
	answer, err := parseAll(inputSting)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

// Функция, позволяющая корректно разбить входные данные на отдельные строки, чтобы в последствии работать со слайсами.
func unionBigNumbers(stringWithoutBigNumbers []string) []string {
	var (
		stringWithBigNumbers []string
		intermediateString   string
	)
	for i := 0; i < len(stringWithoutBigNumbers); i++ {
		_, err := strconv.Atoi(stringWithoutBigNumbers[i])
		if err == nil {
			intermediateString += stringWithoutBigNumbers[i]
			for j := i + 1; j < len(stringWithoutBigNumbers); j++ {
				_, err = strconv.Atoi(stringWithoutBigNumbers[j])
				if err == nil {
					intermediateString += stringWithoutBigNumbers[j]
				} else {
					stringWithBigNumbers = append(stringWithBigNumbers, intermediateString)
					i = j - 1
					break
				}
				if j == len(stringWithoutBigNumbers)-1 {
					stringWithoutBigNumbers = append(stringWithBigNumbers, stringWithoutBigNumbers[j])
					stringWithBigNumbers = append(stringWithBigNumbers, intermediateString)
					i = j
				}
			}
		} else {
			stringWithBigNumbers = append(stringWithBigNumbers, stringWithoutBigNumbers[i])
			intermediateString = ""
		}
	}
	if len(intermediateString) == 1 {
		stringWithBigNumbers = append(stringWithBigNumbers, intermediateString)
	}
	//stringWithBigNumbers = append([]string{"("}, append(stringWithBigNumbers, ")")...)
	return stringWithBigNumbers
}

// Функция бинарных операций
func parseBinOp(stringToCalc []string) int {
	var (
		result = 0
	)

	switch {

	case stringToCalc[1] == "+":
		firstBinOp, err := strconv.Atoi(stringToCalc[0])
		if err != nil {
			log.Fatal(err)
		}
		secondBinOp, err := strconv.Atoi(stringToCalc[2])
		if err != nil {
			log.Fatal(err)
		}
		result += firstBinOp + secondBinOp
	case stringToCalc[1] == "-":
		firstBinOp, err := strconv.Atoi(stringToCalc[0])
		if err != nil {
			log.Fatal(err)
		}
		secondBinOp, err := strconv.Atoi(stringToCalc[2])
		if err != nil {
			log.Fatal(err)
		}
		result += firstBinOp - secondBinOp
	case stringToCalc[1] == "/": // наверное лишнее, так как тут считаются только + и -
		firstBinOp, err := strconv.Atoi(stringToCalc[0])
		if err != nil {
			log.Fatal(err)
		}
		secondBinOp, err := strconv.Atoi(stringToCalc[2])
		if err != nil {
			log.Fatal(err)
		}
		result += firstBinOp / secondBinOp
	case stringToCalc[1] == "*":
		firstBinOp, err := strconv.Atoi(stringToCalc[0])
		if err != nil {
			log.Fatal(err)
		}
		secondBinOp, err := strconv.Atoi(stringToCalc[2])
		if err != nil {
			log.Fatal(err)
		}
		result = result + firstBinOp*secondBinOp
	}

	return result
}

// Основная функция преобразования строки в ответ
func parseAll(stringToCalc []string) (int, error) {
	var (
		result = 0
	)
	if stringToCalc[0] == "-" { // значит требуется вычислить отрицательное число
		stringToCalc[0] = stringToCalc[0] + stringToCalc[1]
	}
	// Этот цикл разбивает данные на скобки, рекурсивно заходя в каждую из подскобок и вычисляя значение в ней
	for i := 0; i < len(stringToCalc); i++ {
		if stringToCalc[i] == "(" {
			openIndex := i
			openCount := 1
			closeIndex := i
			for j := i + 1; j < len(stringToCalc); j++ {
				if stringToCalc[j] == "(" {
					openCount++
				}
				if stringToCalc[j] == ")" {
					openCount--
					if openCount == 0 {
						closeIndex = j
						break
					}
				}
			}
			tmp, err := parseAll(stringToCalc[openIndex+1 : closeIndex])
			if err != nil {
				return 0, err
			}
			intermediateString := strconv.Itoa(tmp)
			leftPart := stringToCalc[:openIndex]
			leftPart = append(leftPart, intermediateString)
			rightPart := stringToCalc[closeIndex+1:]
			stringToCalc = append(leftPart, rightPart...)
		}
	}

	// Цикл, который позволяет пройтись по слайсу и вычислить значения внутри него относительно умножения и деления
	var stack Stack
	stack.Push(stringToCalc[0])
	for i := 1; i < len(stringToCalc)-1; {
		if stringToCalc[i] == "+" || stringToCalc[i] == "-" {
			stack.Push(stringToCalc[i])
			stack.Push(stringToCalc[i+1])
		} else {
			firstBinOp, _ := strconv.Atoi(stack.Pop())
			secondBinOp, _ := strconv.Atoi(stringToCalc[i+1])
			if stringToCalc[i] == "*" {
				if stringToCalc[i+1] == "-" {
					Err := errors.New("ошибка синтаксиса! Добавьте, пожалуйста, унарные минусы в скобки, так гласят правила математики")
					return 0, Err
				}
				stack.Push(strconv.Itoa(firstBinOp * secondBinOp))
			} else if stringToCalc[i] == "/" {
				if stringToCalc[i+1] == "-" {
					Err := errors.New("ошибка синтаксиса! Добавьте, пожалуйста, унарные минусы в скобки, так гласят правила математики")
					return 0, Err
				}
				stack.Push(strconv.Itoa(firstBinOp / secondBinOp))
			} else {
				Err := errors.New("ошибка синтаксиса! Добавьте, пожалуйста, унарные минусы в скобки, так гласят правила математики")
				return 0, Err
			}
		}
		i += 2
	}
	stringToCalc = stack

	// Цикл, который позволяет пройтись по слайсу и линейно вычислить значения внутри него относительно сложения и вычитания
	for 1 < len(stringToCalc) {
		intermediateResult := strconv.Itoa(parseBinOp(stringToCalc[0:3]))
		rightPart := make([]string, len(stringToCalc)-2)
		if len(rightPart) == 0 {
			break
		}
		copy(rightPart[1:], stringToCalc[3:])
		rightPart[0] = intermediateResult
		stringToCalc = rightPart
	}
	result, _ = strconv.Atoi(stringToCalc[0])
	return result, nil
}
