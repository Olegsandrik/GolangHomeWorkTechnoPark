package main

import (
	"bufio"
	//"github.com/go-park-mail-ru/lectures/1-basics/6_is_sorted/sorted"

	sorted "GolangHomeWorkTechnoPark/calc"
	//"github.com/go-park-mail-ru/lectures/1-basics/6_is_sorted/sorted"
	"log"
	"os"

	// go get "github.com/go-park-mail-ru/lectures/1-basics/6_is_sorted" для установки!
	// echo "123\n123\n321\n" | go run main.go для тестирования утилиты
)

func main() {
	var inputStrings []string

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		inputStrings = append(inputStrings, in.Text())
	}

	if err := in.Err(); err != nil {
		log.Fatalf("input scanning failed: %s", err)
	}

	if err := sorted.Check(inputStrings); err != nil {
		log.Fatalf("sorted check failed: %s", err)
	}

	log.Println("strings are sorted")

}
