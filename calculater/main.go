package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var convIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}
var a, b *int
var operators = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string

const (
	LOW = "Вывод ошибки, так как строка " +
		"не является математической операцией."
	HIGH = "Вывод ошибки, так как формат математической операции " +
		"не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	SCALE = "Вывод ошибки, так как используются " +
		"одновременно разные системы счисления."
	DIV = "Вывод ошибки, так как в римской системе " +
		"нет отрицательных чисел."
	ZERO  = "Вывод ошибки, так как в римской системе нет числа 0."
	RANGE = "Калькулятор умеет работать только с арабскими целыми " +
		"числами или римскими цифрами от 1 до 10 включительно"
)

func calc(s string) {
	var operator string
	var stringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)
	// Цикл в цикле. Внешним циклом пробегаемся по мапе с арифметическими операторами. Во внутреннем по получаемой на ввод строке в котором сравниваем элемент строки с ключами operators. После чего в operator присваиваем данный ключ, а в data закидываем срез строк с двумя операндами
	for idx := range operators {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				data = strings.Split(s, operator)
			}
		}
	}
	// Проверка валидность, а именно количество подданных арифметических операций
	switch {
	case len(operator) > 1:
		panic(HIGH)
	case len(operator) < 1:
		panic(LOW)
	}
	// Итерируемся по срезу строк переводим в целочисленный тип с проверкой на error и используемую систему счисления
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	// Свичем проверяем валидны ли поданные арифметические операнды и прокидываем их в глобальные переменные (указатели на тип int) операцией взятия адреса
	switch stringsFound {
	case 1:
		panic(SCALE)
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11
		if val, ok := operators[operator]; ok && errCheck == true {
			a, b = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic(RANGE)
		}
	case 2:
		for _, elem := range romans {
			if val, ok := roman[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic(RANGE)
			}
		}
		if val, ok := operators[operator]; ok {
			a, b = &romansToInt[0], &romansToInt[1]
			intToRoman(val())
		}
	}
}

// Функция конвертер в валидную римскую систему счисления. В первом цикле итерируемся по массиву (игнорируя переменую индекса) значениями которого выступают арабсий эквивалент римских чисел. Во втором цикле итерируемся по интовому параметру функции intToRoman. В третьем цикле пробегаем по мапе с ключами в виде римских чисел типа стринг и сравниваем значением по ключу с значением поданным на вход
func intToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		panic(ZERO)
	} else if romanResult < 0 {
		panic(DIV)
	}
	for romanResult > 0 {
		for _, elem := range convIntToRoman {
			for i := elem; i <= romanResult; {
				for index, value := range roman {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}

// Получаем строку на ввод. Читаем её с помощью буфера до конца самой строки (До знака переноса строки). Инициализируем переменную s. В ней вызываем встроенную функцию ReplaceAll пакета strings которая возвращает копию s без пробелов. И наконец вызываем функцию calc с одним аргументом => строкой s, предварительно преобразовав все символы в верхний регистр чтоб их не учитывать при вычислениях (Справедливо для римских чисел) и удалив пробелы из начала и конца строки
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(console, " ", "")
		calc(strings.ToUpper(strings.TrimSpace(s)))
	}
}
