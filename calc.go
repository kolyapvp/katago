package main

import "fmt"

func main() {
	var a float64
	var b float64
	var c string

	fmt.Println("Калькулятор")
	fmt.Println("Введи первое число: ")
	fmt.Scan(&a)
	fmt.Println("Введи второе число: ")
	fmt.Scan(&b)
	fmt.Println("Введи знак действия (например + или -): ")
	fmt.Scan(&c)

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
	fmt.Printf("Результат: %f\n", result)
}
