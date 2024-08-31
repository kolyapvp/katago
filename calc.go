package main

import "fmt"

func checkA(a string) int {
	switch a {
	case "I", "1":
		return 1
	case "II", "2":
		return 2
	case "III", "3":
		return 3
	case "IV", "4":
		return 4
	case "V", "5":
		return 5
	case "VI", "6":
		return 6
	case "VII", "7":
		return 7
	case "VIII", "8":
		return 8
	case "IX", "9":
		return 9
	case "X", "10":
		return 10
	default:
		return 0
	}
}

func main() {
	var a string
	var b string
	var c string

	fmt.Println("Калькулятор")
	fmt.Println("Введи выражение (например, 3 + 4): ")
	fmt.Scan(&a, &c, &b)
	numA := checkA(a)
	numB := checkA(b)
	if numA < 1 || numA > 10 || numB < 1 || numB > 10 {
		fmt.Println("Ошибка: числа должны быть в диапазоне от 1 до 10 включительно")
		return
	}

	var result float64
	switch c {
	case "+":
		result = float64(numA + numB)
	case "-":
		result = float64(numA - numB)
	case "*":
		result = float64(numA * numB)
	case "/":
		if numB != 0 {
			result = float64(numA) / float64(numB)
		} else {
			fmt.Println("Ошибка: Деление на ноль!")
			return
		}
	default:
		fmt.Println("Некорректный знак действия!")
		return
	}

	fmt.Printf("Результат: %.1f\n", result)
}
