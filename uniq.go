package main

import (
	"log"
	"strconv"
)

// запуск echo -e "небо\nоблака\nоблака\nоблака\nрозы\nрозы\n" | uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]
/*
-с - подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.

-d - вывести только те строки, которые повторились во входных данных.

-u - вывести только те строки, которые не повторились во входных данных.

-f num_fields - не учитывать первые num_fields полей в строке. Полем в строке является непустой набор символов отделённый пробелом.

-s num_chars - не учитывать первые num_chars символов в строке. При использовании вместе с параметром -f учитываются первые символы после num_fields полей (не учитывая пробел-разделитель после последнего поля).

-i - не учитывать регистр букв.
*/

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

	if option.I {
		// Не учитывать регистр букв
		// strings.ToLower()
	} else {
		// Учитывать регистр букв.
		// Соблюдаем CDU или его отсутствие

		counterUniq := 0
		currentString := stringsIn[0]
		counterUniq++
		stringsOutWithout = append(stringsOut, currentString)
		stringsOutWithC = append(stringsOut, currentString)
		for i := 1; i < len(stringsIn); i++ {
			if currentString == stringsIn[i] {
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

		// теперь у нас есть 4 массива для каждого из вариантов CDU или его отсутствие

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
	}

	if option.F > 0 {
		// Не учитывать первые num_fields полей в строке. Полем в строке является непустой набор символов отделённый пробелом.
	}

	if option.S > 0 {
		// Не учитывать первые num_chars символов в строке. При использовании вместе с параметром
		// -f учитываются первые символы после num_fields полей (не учитывая пробел-разделитель после последнего поля).
	}

	if len(option.Input) > 0 {
		if len(option.Output) > 0 {
			// тут оба указаны
		}
		// тут о только ввод
	}

	return stringsOut
}
func main() {
	stringsIn := []string{"небо", "НЕБО", "облака", "облака", "облака", "розы", "розы"}
	option := Options{} // заполняем в соответствии с параметрами
	stringsOut := Uniq(stringsIn, option)

	for i := 0; i < len(stringsOut); i++ {
		log.Println(stringsOut[i])
	}
}
