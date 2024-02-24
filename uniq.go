package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Options - Структура флагов
type Options struct {
	C      bool // false
	D      bool
	U      bool
	I      bool
	F      int // 0
	S      int
	Input  string // ""
	Output string
}

// Uniq - Функция уникализации
func Uniq(stringsIn []string, option Options) []string {
	var (
		stringsOut        []string
		stringsOutWithC   []string
		stringsOutWithD   []string
		stringsOutWithU   []string
		stringsOutWithout []string
	)

	numChars := option.S // работает правильно, только если преобразовывать в rune?
	numFields := option.F

	// Соблюдаем CDU или его отсутствие и все остальные флаги

	counterUniq := 0
	currentString := stringsIn[0]
	counterUniq++
	stringsOutWithout = append(stringsOut, currentString)
	stringsOutWithC = append(stringsOut, currentString)
	for i := 1; i < len(stringsIn); i++ {
		currentPointer := 0
		stringPointer := 0
		// Выполним условие флага f
		if numFields > 0 {
			countFieldCurrent := 0
			for j := 0; j < len(currentString)-1; j++ {
				if currentString[j] != ' ' && currentString[j+1] == ' ' {
					countFieldCurrent++
					if countFieldCurrent == numFields {
						currentPointer = j + 1 // указывает на байт!
						break
					}
				}
			}

			countFieldString := 0
			for k := 0; k < len(stringsIn[i])-1; k++ {
				if stringsIn[i][k] != ' ' && stringsIn[i][k+1] == ' ' {
					countFieldString++
					if countFieldString == numFields {
						stringPointer = k + 1 // указывает на байт!
						break
					}
				}
			}
		}

		// корректировка, если -f и -s одновременно
		if numChars != 0 && numFields != 0 {
			// Не учитывая пробел-разделитель после последнего поля. - из условия задачки
			// Я так понял, что нужно выкинуть лишь 1 пробел, а остальные значимыми считаются
			currentPointer++
			stringPointer++
		}

		if option.I {
			// сначала обрезаем field, а потом обрезаем char
			if strings.ToLower(string(([]rune(currentString[currentPointer:]))[numChars:])) == strings.ToLower(string(([]rune(stringsIn[i][stringPointer:]))[numChars:])) {
				counterUniq++
				continue
			} else {
				currentString = stringsIn[i]
				stringsOutWithout = append(stringsOutWithout, currentString)
				stringsOutWithC = append(stringsOutWithC, strconv.Itoa(counterUniq))
				stringsOutWithC = append(stringsOutWithC, currentString)
				counterUniq = 1
			}
		} else {
			// Учитывать регистр
			if string(([]rune(currentString[currentPointer:]))[numChars:]) == (string(([]rune(stringsIn[i][stringPointer:]))[numChars:])) {
				counterUniq++
				continue
			} else {
				currentString = stringsIn[i]
				stringsOutWithout = append(stringsOutWithout, currentString)
				stringsOutWithC = append(stringsOutWithC, strconv.Itoa(counterUniq))
				stringsOutWithC = append(stringsOutWithC, currentString)
				counterUniq = 1
			}
		}

	}

	stringsOutWithC = append(stringsOutWithC, strconv.Itoa(counterUniq))
	for i := 1; i < len(stringsOutWithC); {
		currentCount, _ := strconv.Atoi(stringsOutWithC[i])
		if currentCount > 1 {
			stringsOutWithD = append(stringsOutWithD, stringsOutWithC[i-1])
		} else {
			stringsOutWithU = append(stringsOutWithU, stringsOutWithC[i-1])
		}
		i += 2
	}

	// теперь у нас есть 4 массива для каждого из вариантов CDU или его отсутствие и все остальные флаги

	if option.C {
		// Подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.
		result := []string{}
		for i := 1; i < len(stringsOutWithC); {
			result = append(result, stringsOutWithC[i]+" "+stringsOutWithC[i-1])
			i += 2
		}
		stringsOut = result
	} else if option.D {
		// Вывести только те строки, которые повторились во входных данных.
		stringsOut = stringsOutWithD
	} else if option.U {
		// Вывести только те строки, которые не повторились во входных данных.
		stringsOut = stringsOutWithU
	} else {
		stringsOut = stringsOutWithout
	}

	return stringsOut
}

func main() {
	// go run uniq.go -c -i -f 1 -s 1 input.txt output.txt
	// cat input.txt | go run uniq.go -c -i -f 1 -s 1 input.txt output.txt

	condition := os.Args[1:]
	option := Options{}
	var err error
	for i := 0; i < len(condition); i++ {
		switch condition[i] {
		case "-c":
			if option.U || option.D {
				err = errors.New("нельзя использовать флаг -c и флаг -u или -d")
				log.Fatal(err)
			}
			option.C = true
		case "-d":
			if option.C || option.U {
				err = errors.New("нельзя использовать флаг -d и флаг -u или -c")
				log.Fatal(err)
			}
			option.D = true
		case "-u":
			if option.C || option.D {
				err = errors.New("нельзя использовать флаг -u и флаг -d или -c")
				log.Fatal(err)
			}
			option.U = true
		case "-i":
			option.I = true
		case "-f":
			if i+1 >= len(condition) {
				err = errors.New("пропущено число после флага -f")
				log.Fatal(err)
			}
			option.F, err = strconv.Atoi(condition[i+1])
			if err != nil {
				err = errors.New("пропущено число после флага -f")
				log.Fatal(err)
			}
			i++
		case "-s":
			if i+1 >= len(condition) {
				err = errors.New("пропущено число после флага -s")
				log.Fatal(err)
			}
			option.S, err = strconv.Atoi(condition[i+1])
			if err != nil {
				err = errors.New("пропущено число после флага -s")
				log.Fatal(err)
			}
			i++
		default:
			if option.Input == "" {
				option.Input = condition[i]
			} else if option.Output == "" {
				option.Output = condition[i]
			} else {
				err = errors.New("лишний аргумент" + condition[i])
				log.Fatal(err)
			}
		}
	}
	fileWork(option)
}

// Работа с вводом и выводом
func fileWork(option Options) {
	inputFile := option.Input
	outputFile := option.Output
	var readFrom io.Reader
	var writeTo io.Writer
	stringsIn := []string{}

	if option.Input != "" { // Либо файлик ввели как параметр
		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		readFrom = file
	} else { // Либо данные нужно считать
		readFrom = os.Stdin
	}

	// читаем данные
	scanner := bufio.NewScanner(readFrom)
	for scanner.Scan() {
		stringsIn = append(stringsIn, scanner.Text())
	}

	stringsOut := Uniq(stringsIn, option)

	if outputFile != "" { // Либо файлик ввели как параметр
		file2, err := os.Create(outputFile)
		if err != nil {
			log.Fatal(err)
		}
		writeTo = file2
		defer file2.Close()
	} else { // Либо выводим в консоль
		writeTo = os.Stdout
	}
	for i := 0; i < len(stringsOut); i++ {
		fmt.Fprintln(writeTo, stringsOut[i])
	}
}
