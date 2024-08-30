package main

import "fmt"

func main() {
	var a float64
	var b float64
	var c string

	fmt.Println("Калькулятор")
	fmt.Println("Введи выражение (например, 3 + 4): ")
	fmt.Scan(&a, &c, &b)
	if a < 1 || a > 10 || b < 1 || b > 10 {
		fmt.Println("Ошибка: числа должны быть в диапазоне от 1 до 10 включительно")
		return
	}

	var result float64
	switch c {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b != 0 {
			result = a / b
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
