package main

/*

	stringsIn1 := []string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.",
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

			if len(option.Input) > 0 {
				if len(option.Output) > 0 {
					// тут оба указаны
				}
				// тут о только ввод
			}
		//
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

		option6 := Options{F: 1} // А если полей больше, чем вообще в строчке
		stringsOut6 := Uniq(stringsIn6, option6)
		for i := 0; i < len(stringsOut6); i++ {
			log.Println(stringsOut6[i])
		}

		log.Println("Пошел 7 тест")

		option7 := Options{S: 1} // А если символов больше, чем вообще в строчке
		stringsOut7 := Uniq(stringsIn7, option7)
		for i := 0; i < len(stringsOut7); i++ {
			log.Println(stringsOut7[i])
		}
*/
