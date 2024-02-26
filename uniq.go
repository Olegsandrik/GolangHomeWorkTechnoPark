package main

import (
	inOut "GolangHomeWorkTechnoPark/inputOutput"
	"fmt"
)

func main() {
	// go run uniq.go -c -i -f 1 -s 1 input.txt output.txt
	// cat input.txt | go run uniq.go -c -i -f 1 -s 1 input.txt output.txt
	option, err := inOut.Scanner()
	if err != nil {
		fmt.Println(err)
		return
	}
	inOut.FileWork(option)

}
