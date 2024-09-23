package main

import (
	"fmt"
)

// 1. Определение четности числа
func isEvenOrOdd(num int) string {
	if num%2 == 0 {
		return "Четное"
	}
	return "Нечетное"
}

// 2. Функция для определения знака числа
func numberSign(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	}
	return "Zero"
}

// 3. Вывод чисел от 1 до 10
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 4. Функция для получения длины строки
func stringLength(s string) int {
	return len(s)
}

// 5. Структура Rectangle и метод для вычисления площади
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 6. Функция для вычисления среднего значения двух чисел
func average(a, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	// Пример использования всех функций

	// 1. Четность числа
	var num int
	fmt.Print("Введите число для проверки четности: ")
	fmt.Scan(&num)
	fmt.Println(isEvenOrOdd(num))

	// 2. Определение знака числа
	fmt.Print("Введите число для определения знака: ")
	fmt.Scan(&num)
	fmt.Println(numberSign(num))

	// 3. Вывод чисел от 1 до 10
	fmt.Println("Числа от 1 до 10:")
	printNumbers()

	// 4. Длина строки
	var str string
	fmt.Print("Введите строку для получения её длины: ")
	fmt.Scan(&str)
	fmt.Println("Длина строки:", stringLength(str))

	// 5. Площадь прямоугольника
	rect := Rectangle{Width: 5, Height: 10}
	fmt.Println("Площадь прямоугольника:", rect.Area())

	// 6. Среднее значение двух чисел
	var a, b int
	fmt.Print("Введите два целых числа для вычисления среднего значения: ")
	fmt.Scan(&a, &b)
	fmt.Println("Среднее значение:", average(a, b))
}
