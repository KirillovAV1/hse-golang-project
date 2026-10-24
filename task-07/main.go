package main

import "fmt"

func sumByCondition(num int) int {
	sum := 0

	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%2 == 0 {
			sum += i
		}
	}

	return sum
}

func main() {

	var inputNum int

	fmt.Println("Программа для подсчета суммы чисел из последовательности от 1 до N с условиями:")
	fmt.Println("1. Число является четным")
	fmt.Println("2. Число делится на 3")
	fmt.Print("\nВведите число N: ")

	_, err := fmt.Scan(&inputNum)

	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return
	} else if inputNum <= 0 {
		fmt.Println("Введите положительное число")
		return
	}

	fmt.Print("\n")
	sum := sumByCondition(inputNum)

	if sum > 0 {
		fmt.Println("Итоговая сумма:", sum)
	} else {
		fmt.Println("Не нашлось чисел, удовлетворяющих условиям")
	}
}
