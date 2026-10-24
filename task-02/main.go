package main

import "fmt"

func isPositive(num int) string {
	if num > 0 {
		return "Число положительное"
	} else if num < 0 {
		return "Число отрицательное"
	} else {
		return "У вас число 0"
	}
}

func isEven(num int) string {
	if num%2 == 0 {
		return "Число четное"
	} else {
		return "Число нечетное"
	}
}

func isDiv(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "Число делится на 3 и на 5"
	} else if num%3 == 0 {
		return "Число делится только на 3"
	} else if num%5 == 0 {
		return "Число делится только на 5"
	} else {
		return "Число не делится ни на одно из этих чисел"
	}
}

func main() {

	var inputNum int

	fmt.Println("Программа для проверки:")
	fmt.Println("1. Является ли число положительным/отрицательным?")
	fmt.Println("2. Число четное/нечетное")
	fmt.Println("3. Делится ли число на 3 и 5?")
	fmt.Print("\nВведите свое число: ")

	_, err := fmt.Scan(&inputNum)

	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return
	}

	fmt.Print("\n")
	fmt.Println("1.", isPositive(inputNum))
	fmt.Println("2.", isEven(inputNum))
	fmt.Println("3.", isDiv(inputNum))

}
