package main

import (
	"fmt"
	"strings"
)

// Функция для расчета среднего возраста
func averageAge(ages map[string]int) float64 {
	total := 0
	for _, age := range ages {
		total += age
	}
	return float64(total) / float64(len(ages))
}

// Функция для удаления записи по имени
func deletePerson(ages map[string]int, name string) {
	delete(ages, name)
}

func main() {
	// Создание карты с именами и возрастами
	ages := make(map[string]int)

	// Добавление нового человека
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	// Вывод всех записей на экран
	fmt.Println("Записи о людях:")
	for name, age := range ages {
		fmt.Printf("%s: %d лет\n", name, age)
	}

	// Вычисление и вывод среднего возраста
	avg := averageAge(ages)
	fmt.Printf("Средний возраст: %.2f лет\n", avg)

	// Удаление записи по имени
	var nameToDelete string
	fmt.Print("Введите имя для удаления: ")
	fmt.Scan(&nameToDelete)
	deletePerson(ages, nameToDelete)

	// Вывод всех записей после удаления
	fmt.Println("Записи о людях после удаления:")
	for name, age := range ages {
		fmt.Printf("%s: %d лет\n", name, age)
	}

	// Считывание строки и вывод в верхнем регистре
	var inputString string
	fmt.Print("Введите строку: ")
	fmt.Scan(&inputString)
	fmt.Println("В верхнем регистре:", strings.ToUpper(inputString))

	// Считывание чисел и вывод их суммы
	var numberCount int
	fmt.Print("Сколько чисел вы хотите ввести? ")
	fmt.Scan(&numberCount)

	sum := 0
	for i := 0; i < numberCount; i++ {
		var number int
		fmt.Print("Введите число: ")
		fmt.Scan(&number)
		sum += number
	}
	fmt.Printf("Сумма введенных чисел: %d\n", sum)

	// Считывание массива целых чисел и вывод их в обратном порядке
	var arraySize int
	fmt.Print("Введите размер массива целых чисел: ")
	fmt.Scan(&arraySize)

	array := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		fmt.Print("Введите число: ")
		fmt.Scan(&array[i])
	}

	fmt.Print("Числа в обратном порядке: ")
	for i := arraySize - 1; i >= 0; i-- {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}
