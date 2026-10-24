package main

import (
	"fmt"
)

func main() {

	var amount, currNum, min, max int

	fmt.Println("Программа для нахождения:")
	fmt.Println("1. Максимума из последовательности чисел")
	fmt.Println("2. Минимума")
	fmt.Print("\nВведите  N (количество чисел): ")

	_, err := fmt.Scan(&amount)

	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return
	} else if amount <= 0 {
		fmt.Println("Последовательность должна иметь минимум 1 число")
		return
	}

	fmt.Printf("Введите %d чисел/числа:\n", amount)
	for i := 0; i < amount; i++ {
		_, err := fmt.Scan(&currNum)
		if err != nil {
			fmt.Println("Некорректный ввод:", err)
			return
		}

		if i == 0 {
			min = currNum
			max = currNum
			continue
		}

		if currNum > max {
			max = currNum
		}

		if currNum < min {
			min = currNum
		}
	}

	fmt.Print("\n")
	fmt.Println("1. Максимальное число:", max)
	fmt.Println("2. Минимальное число:", min)

}
