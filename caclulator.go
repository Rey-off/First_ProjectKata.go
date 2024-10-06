package main

import (
	"fmt"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I":  1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"XX": 20, "XXX": 30, "XL": 40, "L": 50,
	"LX": 60, "LXX": 70, "LXXX": 80, "XC": 90, "C": 100,
}

var arabicNumerals = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	20: "XX", 30: "XXX", 40: "XL", 50: "L",
	60: "LX", 70: "LXX", 80: "LXXX", 90: "XC", 100: "C",
}

func main() {
	for {
		fmt.Print("Введите выражение: ")
		var input string
		fmt.Scanln(&input)

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Неправильный ввод")
			continue
		}

		num1, op, num2 := parts[0], parts[1], parts[2]

		if isRomanNumber(num1) && isRomanNumber(num2) {
			a, err := romanToArabic(num1)
			if err != nil {
				fmt.Println(err)
				continue
			}
			b, err := romanToArabic(num2)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if !isValidRomanOperation(a, b) {
				fmt.Println("Недействительная римская операция")
				continue
			}

			result := calculate(a, b, op)
			printResult(result, true)
		} else if isArabicNumber(num1) && isArabicNumber(num2) {
			a, err := strconv.Atoi(num1)
			if err != nil {
				fmt.Println(err)
				continue
			}
			b, err := strconv.Atoi(num2)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if !isValidArabicOperation(num1, num2) {
				fmt.Println("Недействительная арабская операция")
				continue
			}

			result := calculate(a, b, op)
			printResult(result, false)
		} else {
			fmt.Println("Неправильный ввод")
		}
	}
}

func calculate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return 0
		}
		return a / b
	default:
		fmt.Println("Недопустимая операция")
		return 0
	}
}

func romanToArabic(roman string) (int, error) {
	total := 0
	prevValue := 0
	for _, char := range roman {
		value, exists := romanNumerals[string(char)]
		if !exists {
			return 0, fmt.Errorf("Недействительное римское число: %s", roman)
		}
		if value > prevValue {
			total += value - 2*prevValue
		} else {
			total += value
		}
		prevValue = value
	}
	return total, nil
}

func arabicToRoman(arabic int) string {
	if arabic <= 0 || arabic > 100 {
		return ""
	}
	for i := 100; i >= 1; i-- {
		if roman, ok := arabicNumerals[i]; ok && arabic == i {
			return roman
		}
	}
	return ""
}

func isRomanNumber(num string) bool {
	_, ok := romanNumerals[num]
	return ok || len(num) < 4   
}

func isArabicNumber(num string) bool {
	_, err := strconv.Atoi(num)
	return err == nil
}

func isValidOperation(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func isValidRomanOperation(num1, num2 int) bool {
	return num1 >= 1 && num2 >= 1
}

func isValidArabicOperation(num1, num2 string) bool {
	a, _ := strconv.Atoi(num1)
	b, _ := strconv.Atoi(num2)
	return a >= 1 && b >= 1
}

func printResult(result int, isRoman bool) {
	if isRoman {
		if result <= 0 {
			fmt.Println("Результат: недопустимое римское число")
		} else {
			fmt.Println("Результат:", arabicToRoman(result))
		}
	} else {
		fmt.Println("Результат:", result)
	}
}
