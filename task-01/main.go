package main

import "fmt"

func main() {

	var numFirst, numSecond float64
	var action string

	fmt.Println("Поддерживаемые операции: +, -, *, /")
	fmt.Println("Пример ввода: 10.4 + 23")
	fmt.Print("Введите выражение: ")

	_, err := fmt.Scan(&numFirst, &action, &numSecond)
	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return
	}

	switch action {
	case "+":
		fmt.Println("Результат:", numFirst+numSecond)
	case "-":
		fmt.Println("Результат:", numFirst-numSecond)
	case "*":
		fmt.Println("Результат:", numFirst*numSecond)
	case "/":
		if numSecond == 0 {
			fmt.Println("Делить на 0 - нельзя")
			return
		}

		fmt.Println("Результат:", numFirst/numSecond)
		return
	default:
		fmt.Println("Неизвестная операция")
	}
}
