package main

import (
	"fmt"
	"math/rand"
)

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func inputNumber() (float64, error) {
	var num float64
	_, err := fmt.Scan(&num)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func initBalance() float64 {
	fmt.Println("Введите ваш стартовый баланс:")
	balance, err := inputNumber()
	if err != nil || balance < 0 {
		fmt.Println("Некорректный ввод")
		fmt.Println("Вам будет сгенерирован случайный баланс")
		return float64(randomInt(500, 1000))
	}

	return balance
}

func deposit(balance float64) float64 {
	fmt.Println("Введите сумму пополнения: ")
	amountMoney, err := inputNumber()

	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return balance
	}

	if amountMoney <= 0 {
		fmt.Println("Пополнение должно быть положительным")
		return balance
	}

	return balance + amountMoney
}

func withdraw(balance float64) float64 {
	fmt.Println("Введите сумму снятия: ")
	amountMoney, err := inputNumber()

	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return balance
	}

	if amountMoney <= 0 {
		fmt.Println("Сумма снятия должна быть положительной")
		return balance
	}

	if balance < amountMoney {
		fmt.Printf("Недостаточно средств. Не хватает %.2f рублей\n", amountMoney-balance)
		return balance
	}

	return balance - amountMoney
}

func main() {

	balance := initBalance()

	var action int

	for {
		fmt.Println("\nДоступные операции:")
		fmt.Println("1. Посмотреть баланс")
		fmt.Println("2. Пополнить баланс")
		fmt.Println("3. Снять средства")
		fmt.Println("0. Выйти из приложения")

		fmt.Println("\nВведите номер команды:")
		_, err := fmt.Scan(&action)

		if err != nil {
			fmt.Println("Некорректный ввод команды:", err)
			return
		}

		switch action {
		case 1:
			fmt.Println("\n----------------------")
			fmt.Printf("Текущий баланс: %.2f рублей\n", balance)
			fmt.Println("----------------------")
		case 2:
			balance = deposit(balance)
		case 3:
			balance = withdraw(balance)
		case 0:
			return
		default:
			fmt.Println("Неизвестная команда")
		}

	}

}
