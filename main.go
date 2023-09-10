package main

import (
	"fmt"
	"os"
)

var rus = [2][33]string{
	{"А", "Б", "В", "Г", "Д", "Е", "Ё", "Ж", "З", "И", "Й", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Х", "Ц", "Ч", "Ш", "Щ", "Ъ", "Ы", "Ь", "Э", "Ю", "Я"},
	{"а", "б", "в", "г", "д", "е", "ё", "ж", "з", "и", "й", "к", "л", "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ъ", "ы", "ь", "э", "ю", "я"},
}

// Проверка каждого введённого слова на отсутствие некорректных символов. Допускаются только буквы русского алфавита и знак '='.
func checkWord(n []string) bool {
	if len(n) == 0 {
		return false
	}
	flag1 := 0
	for _, symbol := range n {
		for j := 0; j < 2; j++ {
			for k := 0; k < 33; k++ {
				if symbol == "=" {
					return true
				} else if symbol == rus[j][k] {
					flag1++
					continue
				} else {
					continue
				}
			}
		}
	}
	if flag1 == len(n) {
		return true
	} else {
		return false
	}
}

func main() {

	var word string       // Принимает значение каждого введённого слова.
	var allWords []string // Содержит все введённые слова для формирования результата.
	var toDo []string     // Служит для хранения всех символов всех слов.
	var result int

	for {
		fmt.Println("Введите слово или введите '=' для расчёта:")
		_, err := fmt.Scanln(&word)
		if err != nil {
			fmt.Println("Ошибка ввода!")
			return
		}

		// Преобразования каждого слова из набора byte в набор string.
		var wordInSlice []string // Область видимости бесконечного цикла for позволяет хранить в этой переменной только одно обрабатываемое слово.
		for _, b := range word {
			wordInSlice = append(wordInSlice, string(b))
		}

		flag2 := checkWord(wordInSlice)

		if flag2 == true {

			if word != "=" {
				allWords = append(allWords, word)
			}

			toDo = append(toDo, wordInSlice...)
			switch word {
			case "=":
				{
					for _, symbol := range toDo {
						for j := 0; j < 2; j++ {
							for k := 0; k < 33; k++ {
								if symbol == rus[j][k] {
									result = result + (k + 1)
								}
							}
						}
					}
					for i := 0; i < len(allWords); i++ {
						fmt.Printf(allWords[i])
						if i != len(allWords)-1 {
							fmt.Printf(" + ")
						}
					}
					fmt.Printf(" = %d", result)
					os.Exit(0)
				}
			}
		} else if flag2 == false {
			fmt.Println("Ошибка ввода! Символы некорректны. Введённое слово не учтёно.")
		}
	}
}
