package main

import (
	"fmt"
)

func printTable(num int) {

	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			fmt.Printf("%4d |", i*j)
		}
		fmt.Println()
	}

}

func main() {

	var inputNum int

	fmt.Println("Программа для вывода таблицы умножения:")
	fmt.Println("Введите число N (1 <= N <= 10)")
	fmt.Print("\nВведите число N: ")

	_, err := fmt.Scan(&inputNum)

	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return
	} else if inputNum < 1 || inputNum > 10 {
		fmt.Println("Введите число в пределах 1 <= N <= 10")
		return
	}

	fmt.Print("\n")
	printTable(inputNum)

}
