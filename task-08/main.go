package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	numRoot := int(math.Sqrt(float64(num)))
	for i := 2; i <= numRoot; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	var inputNum int

	fmt.Println("Программа для проверки является ли число простым")
	fmt.Print("\nВведите число: ")

	_, err := fmt.Scan(&inputNum)

	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return
	}

	fmt.Print("\n")

	if isPrime(inputNum) {
		fmt.Printf("Число %d - простое\n", inputNum)
	} else {
		fmt.Printf("Число %d - не простое\n", inputNum)
	}
}
