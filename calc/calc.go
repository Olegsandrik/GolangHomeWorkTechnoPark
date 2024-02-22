package calc

import (
	"log"
	"strconv"
)

func Calculator(stringToCalc []string) int {
	stringToCalc = unionBigNumbers(stringToCalc)
	log.Println(stringToCalc)
	return 0 //parseAll(stringToCalc)
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
	stringWithBigNumbers = append([]string{"("}, append(stringWithBigNumbers, ")")...)
	return stringWithBigNumbers
}

// Надевает на * и на / скобки, для того, чтобы вычисления были корректны с математической точки зрения.
func prioritetBinOp(stringToCalc []string) []string {
	for i := 0; i < len(stringToCalc); i++ {
		if stringToCalc[i] == "*" || stringToCalc[i] == "/" {

		}
	}
	return nil
}

// Надевает на * и на / скобки, для того, чтобы вычисления были корректны с математической точки зрения.
func prioritetBinOpV2(stringToCalc []string) []string {
	for i := 0; i < len(stringToCalc)-1; i++ {
		if stringToCalc[i] == "*" || stringToCalc[i] == "/" {
			openIndexPrioritet := i
			closeIndexPrioritet := i
			for j := i; j < len(stringToCalc); {
				if stringToCalc[j] == "*" || stringToCalc[j] == "/" {
					continue
				}
				if stringToCalc[j] == "+" || stringToCalc[j] == "-" {
					closeIndexPrioritet = j
					break
				}
				if stringToCalc[j] == "(" {
					openIndexBracket := j
					openCountBracket := 1
					closeIndexBracket := j
					for k := j; k < len(stringToCalc); k++ {
						if stringToCalc[k] == "(" {
							openCountBracket++
						}
						if stringToCalc[k] == ")" {
							openCountBracket--
							if openCountBracket == 0 {
								closeIndexBracket = k
								break
							}
						}
					}
					intermediateString := prioritetBinOpV2(stringToCalc[openIndexBracket+1 : closeIndexBracket])
					leftPart := stringToCalc[:openIndexBracket]
					leftPart = append(leftPart, intermediateString...)
					rightPart := stringToCalc[closeIndexBracket+1:]
					stringToCalc = append(leftPart, rightPart...)
				}
				j += 2
			}
			if openIndexPrioritet == closeIndexPrioritet {
				leftPart := stringToCalc[:openIndexPrioritet]
				leftPart = append([]string{"("}, leftPart...)
				rightPart := stringToCalc[closeIndexPrioritet+1:]
				stringToCalc = append(leftPart, rightPart...)
				stringToCalc = append(stringToCalc, ")")
			}
		}

	}
	return stringToCalc
}

// Функция бинарных операций
func parseBinOp(stringToCalc []string) int {
	var (
		result = 0
	)

	switch {
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
	case stringToCalc[1] == "/":
		firstBinOp, err := strconv.Atoi(stringToCalc[0])
		if err != nil {
			log.Fatal(err)
		}
		secondBinOp, err := strconv.Atoi(stringToCalc[2])
		if err != nil {
			log.Fatal(err)
		}
		result += firstBinOp / secondBinOp
	}

	return result
}

// Основная функция преобразования строки в ответ
func parseAll(stringToCalc []string) int {
	var (
		result = 0
	)

	// этот цикл разбивает данные на скобки, рекурсивно заходя в каждые подскобки и вычисляя там значения подскобки
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
			intermediateString := strconv.Itoa(parseAll(stringToCalc[openIndex+1 : closeIndex]))
			leftPart := stringToCalc[:openIndex]
			leftPart = append(leftPart, intermediateString)
			rightPart := stringToCalc[closeIndex+1:]
			stringToCalc = append(leftPart, rightPart...)
		}
	}

	// Цикл, который позволяет пройтись по слайсу и линейно вычислить значения внутри него
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
	return result
}
