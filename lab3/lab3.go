package main

import (
	"fmt"
	"laba/mathutils"
	"laba/stringutils"
)

func main() {
	// Часть 1: Вычисление факториала
	var number int
	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scan(&number)
	factorial := mathutils.Factorial(number)
	if factorial == -1 {
		fmt.Println("Факториал отрицательного числа не определен.")
	} else {
		fmt.Printf("Факториал %d равен %d\n", number, factorial)
	}

	// Часть 2: Переворот строки
	var str string
	fmt.Print("Введите строку для переворота: ")
	fmt.Scan(&str)
	reversedStr := stringutils.Reverse(str)
	fmt.Printf("Перевернутая строка: %s\n", reversedStr)

	// Часть 3: Создание массива из 5 целых чисел
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Массив целых чисел:", arr)

	// Часть 4: Создание среза и операции добавления и удаления элементов
	slice := arr[:]
	slice = append(slice, 6) // Добавление элемента
	fmt.Println("Срез после добавления элемента:", slice)
	slice = slice[:len(slice)-1] // Удаление последнего элемента
	fmt.Println("Срез после удаления последнего элемента:", slice)

	// Часть 5: Создание среза строк и нахождение самой длинной строки
	strings := []string{"apple", "banana", "cherry", "date"}
	longest := ""
	for _, s := range strings {
		if len(s) > len(longest) {
			longest = s
		}
	}
	fmt.Printf("Самая длинная строка: %s\n", longest)
}
