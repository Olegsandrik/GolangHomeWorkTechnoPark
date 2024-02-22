package main

import (
	"bufio"
	"strings"

	//"github.com/go-park-mail-ru/lectures/1-basics/6_is_sorted/sorted"

	calc "GolangHomeWorkTechnoPark/calc"
	//"github.com/go-park-mail-ru/lectures/1-basics/6_is_sorted/sorted"
	"log"
	"os"
	// go get "github.com/go-park-mail-ru/lectures/1-basics/6_is_sorted" для установки!
	// echo "123\n123\n321\n" | go run main.go для тестирования утилиты, показавшую на лекции - не корректно работает)
	// echo -e "123\n123\n333\n111" | go run main.go -- работает корректно)
	// go run main.go "(1+2)-3" - тестим калькулятор
)

func main() {
	var inputStrings string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStrings = scanner.Text()

	inputString := strings.Split(strings.ReplaceAll(inputStrings, " ", ""), "")

	if err := scanner.Err(); err != nil {
		log.Fatalf("input scanning failed: %s", err)
	}

	log.Println(calc.Calculator(inputString))

}
