package main

import (
	"fmt"
)

func main() {

	var currNum, max, min, amountEven, amountOdd, amountNums int
	sum := 0

	fmt.Println("Программа для нахождения:")
	fmt.Println("1. Сумму введеных чисел")
	fmt.Println("2. Их среднего значения")
	fmt.Println("3. Максимальное число")
	fmt.Println("4. Минимальное число")
	fmt.Println("5. Количество четных чисел")
	fmt.Println("6. Количество нечетных чисел")

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

		if i == 0 {
			min = currNum
			max = currNum
		}

		sum += currNum

		if currNum > max {
			max = currNum
		}

		if currNum < min {
			min = currNum
		}

		if currNum%2 == 0 {
			amountEven++
		} else {
			amountOdd++
		}
	}

	fmt.Print("\n")
	fmt.Println("1. Сумма введеных чисел:", sum)
	fmt.Println("2. Среднее значение", float64(sum)/float64(amountNums))
	fmt.Println("3. Максимальное число", max)
	fmt.Println("4. Минимальное число", min)
	fmt.Println("5. Количество четных чисел", amountEven)
	fmt.Println("6. Количество нечетных чисел", amountOdd)
}
