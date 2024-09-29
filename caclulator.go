package main  

import (  
	"fmt"
	"strconv"
	"strings" 
)  

var romanNumerals = map[string]int{  
	"I":  1, "II": 2, "III": 3, "IV": 4, "V": 5,  
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,  
}  

var arabicNumerals = map[int]string{  
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",   
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",  
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
			if !isValidRomanOperation(num1, num2) {  
				fmt.Println("Недействительная римская операция")
				continue  
			}  
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
			result := calculate(a, b, op)  
			printResult(result, true)  
		} else if isArabicNumber(num1) && isArabicNumber(num2) {  
			if !isValidArabicOperation(num1, num2) {  
				fmt.Println("Недействительная арабская операция")
				continue  
			}  
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
	if val, ok := romanNumerals[roman]; ok {  
		return val, nil  
	}  
	return 0, fmt.Errorf("Недействительное римское число: %s", roman)  
}  

func arabicToRoman(arabic int) string {  
	if val, ok := arabicNumerals[arabic]; ok {  
		return val  
	}  
	return " "
}  

func isRomanNumber(num string) bool {  
	_, ok := romanNumerals[num]  
	return ok  
}  

func isArabicNumber(num string) bool {  
	_, err := strconv.Atoi(num)  
	return err == nil  
}  

func isValidOperation(op string) bool {  
	return op == "+" || op == "-" || op == "*" || op == "/"  
}  

func isValidRomanOperation(num1, num2 string) bool {  
	return isRomanNumber(num1) && isRomanNumber(num2)  
}  

func isValidArabicOperation(num1, num2 string) bool {  
	return isArabicNumber(num1) && isArabicNumber(num2)  
}  

func printError(err string) {  
	fmt.Println(err) 
}  

func printResult(result int, isRoman bool) {  
	if isRoman {  
		fmt.Println(arabicToRoman(result))
	} else {  
		fmt.Println(result)
	}  
}