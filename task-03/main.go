package main

import "fmt"

func avgAndSum(num int) (int, float64) {
	sum := 0

	for i := 1; i <= num; i++ {
		sum += i
	}

	avg := float64(sum) / float64(num)

	return sum, avg
}

func main() {

	var inputNum int

	fmt.Println("Программа для подсчета:")
	fmt.Println("1. Суммы всех чисел от 1 до N")
	fmt.Println("2. Среднее значение чисел")
	fmt.Print("\nВведите свое число N: ")

	_, err := fmt.Scan(&inputNum)

	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return
	} else if inputNum <= 0 {
		fmt.Println("Введите положительное число")
		return
	}

	fmt.Print("\n")
	sum, avg := avgAndSum(inputNum)
	fmt.Println("1.", sum)
	fmt.Println("2.", avg)
}
