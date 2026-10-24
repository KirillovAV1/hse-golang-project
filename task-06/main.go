package main

import (
	"fmt"
)

func main() {

	var amountNums, amountPositive, amountNegative, amountZeros, currNum int

	fmt.Println("Программа для нахождения:")
	fmt.Println("1. Количество положительных чисел")
	fmt.Println("2. Количество отрицательных чисел")
	fmt.Println("3. Количество нулей")
	fmt.Print("\nВведите  N (количество чисел): ")

	_, err := fmt.Scan(&amountNums)

	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return
	} else if amountNums <= 0 {
		fmt.Println("Последовательность должна иметь минимум 1 число")
		return
	}

	fmt.Printf("Введите %d чисел/числа:\n", amountNums)
	for i := 0; i < amountNums; i++ {
		_, err := fmt.Scan(&currNum)
		if err != nil {
			fmt.Println("Некорректный ввод:", err)
			return
		}

		if currNum > 0 {
			amountPositive++
		} else if currNum < 0 {
			amountNegative++
		} else {
			amountZeros++
		}
	}

	fmt.Print("\n")
	fmt.Println("1. Количество положительных чисел", amountPositive)
	fmt.Println("2. Количество отрицательных чисел", amountNegative)
	fmt.Println("3. Количество нулей:", amountZeros)

}
