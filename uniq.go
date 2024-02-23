package main

import (
	"log"
	"strconv"
	"strings"
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

	numChars := option.S

	// Не учитывать регистр букв
	// strings.ToLower()

	// Учитывать регистр букв.
	// Соблюдаем CDU или его отсутствие

	counterUniq := 0
	currentString := stringsIn[0]
	counterUniq++
	stringsOutWithout = append(stringsOut, currentString)
	stringsOutWithC = append(stringsOut, currentString)
	for i := 1; i < len(stringsIn); i++ {
		if option.I {
			// Не учитывать регистр
			if strings.ToLower(currentString[numChars:]) == strings.ToLower(stringsIn[i][numChars:]) {
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
			if currentString[numChars:] == stringsIn[i][numChars:] {
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

// Uniq - Функция уникализации
func BadUniq(stringsIn []string, option Options) []string {
	var (
		stringsOut        []string
		stringsOutWithC   []string
		stringsOutWithD   []string
		stringsOutWithU   []string
		stringsOutWithout []string
	)

	numChars := option.S
	/*
		numFields := option.F
		currentPointer := 0
		stringPointer := 0
		//*/

	// Соблюдаем CDU или его отсутствие

	counterUniq := 0
	currentString := stringsIn[0]
	counterUniq++
	stringsOutWithout = append(stringsOut, currentString)
	stringsOutWithC = append(stringsOut, currentString)
	for i := 1; i < len(stringsIn); i++ {
		// Вычислим ограничение по numFields для обеих строчек
		/*
			countField1 := 0
			for j := 0; j < len(currentString)-1; j++ {

				if numFields == 0 {
					break
				}
				if currentString[j] != ' ' {
					if currentString[j+1] == ' ' {
						countField1++
						if countField1 == numFields {
							currentPointer = j
							break
						}
					}
				}
			}

			countField2 := 0
			for j := 0; j < len(stringsIn[i])-1; j++ {

				if numFields == 0 {
					break
				}
				if stringsIn[i][j] != ' ' {
					if stringsIn[i][j+1] == ' ' {
						countField2++
						if countField2 == numFields {
							stringPointer = j
							break
						}
					}
				}
			}
			currentString = currentString[currentPointer:]
			stringsIn = stringsIn[stringPointer:]

			//*/
		if option.I {
			// Не учитывать регистр букв
			if strings.ToLower(currentString[numChars:]) == strings.ToLower(stringsIn[i][numChars:]) {
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
			// Учитывать регистр букв.

			if currentString[0:] == stringsIn[i][numChars:] {
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

	// теперь у нас есть 4 массива для каждого из вариантов CDU или его отсутствие, с учетом флагов i, s, f
	if option.C {
		result := []string{}
		for i := 1; i < len(stringsOutWithC); {
			result = append(result, stringsOutWithC[i]+" "+stringsOutWithC[i-1])
			i += 2
		}
		stringsOut = result
	} else if option.D {
		stringsOut = stringsOutWithD
	} else if option.U {
		stringsOut = stringsOutWithU
	} else {
		stringsOut = stringsOutWithout
	}

	return stringsOut
}

func main() {
	stringsIn1 := []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
		"Thanks.", "I love music of Kartik.", "I love music of Kartik."} // Обычный
	stringsIn2 := []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
		"Thanks.", "I love music of Kartik.", "I love music of Kartik."} // С флагом C
	stringsIn3 := []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
		"Thanks.", "I love music of Kartik.", "I love music of Kartik."} // С флагом D
	stringsIn4 := []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
		"Thanks.", "I love music of Kartik.", "I love music of Kartik."} // С флагом U
	stringsIn5 := []string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", " ", "I love MuSIC of Kartik.", "I love music of kartik.",
		"Thanks.", "I love music of kartik.", "I love MuSIC of Kartik."} // С флагом I
	stringsIn6 := []string{"We love music.", "I love music.", "They love music.", " ", "I love music of Kartik.", "We love music of Kartik.",
		"Thanks"} // С флагом F

	stringsIn7 := []string{"I love music.", "A love music.", "C love music.", " ", "I love music of Kartik.", "We love music of Kartik.",
		"Thanks."} // С флагом S
	// вывод и ввод нашей утилиты
	/*
		if len(option.Input) > 0 {
			if len(option.Output) > 0 {
				// тут оба указаны
			}
			// тут о только ввод
		}
	*/
	option1 := Options{}
	stringsOut1 := Uniq(stringsIn1, option1)
	for i := 0; i < len(stringsOut1); i++ {
		log.Println(stringsOut1[i])
	}

	log.Println("Пошел 2 тест")

	option2 := Options{C: true}
	stringsOut2 := Uniq(stringsIn2, option2)
	for i := 0; i < len(stringsOut2); i++ {
		log.Println(stringsOut2[i])
	}

	log.Println("Пошел 3 тест")

	option3 := Options{D: true}
	stringsOut3 := Uniq(stringsIn3, option3)
	for i := 0; i < len(stringsOut3); i++ {
		log.Println(stringsOut3[i])
	}

	log.Println("Пошел 4 тест")

	option4 := Options{U: true}
	stringsOut4 := Uniq(stringsIn4, option4)
	for i := 0; i < len(stringsOut4); i++ {
		log.Println(stringsOut4[i])
	}

	log.Println("Пошел 5 тест")

	option5 := Options{I: true}
	stringsOut5 := Uniq(stringsIn5, option5)
	for i := 0; i < len(stringsOut5); i++ {
		log.Println(stringsOut5[i])
	}

	log.Println("Пошел 6 тест")

	option6 := Options{F: 1}
	stringsOut6 := Uniq(stringsIn6, option6)
	for i := 0; i < len(stringsOut6); i++ {
		log.Println(stringsOut6[i])
	}

	log.Println("Пошел 7 тест")

	option7 := Options{S: 1}
	stringsOut7 := Uniq(stringsIn7, option7)
	for i := 0; i < len(stringsOut7); i++ {
		log.Println(stringsOut7[i])
	}

}
