package main

import (
	"fmt"
	"time"
)

// Функция для вычисления суммы и разности двух чисел с плавающей запятой
func calculate(a, b float64) (sum float64, difference float64) {
	sum = a + b
	difference = a - b
	return
}

// Функция для вычисления среднего значения трех чисел
func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {
	// 1. Вывод текущего времени и даты
	currentTime := time.Now()
	fmt.Println("Текущее время и дата:", currentTime)

	// 2. Создание переменных различных типов
	var integerVar int = 42
	var floatVar float64 = 3.145613
	var stringVar string = "Hello, Go!"
	var boolVar bool = true

	// 3. Краткая форма объявления переменных
	intVar := 10
	floatVar2 := 2.5
	stringVar2 := "Привет"
	boolVar2 := false

	// Вывод переменных на экран
	fmt.Println("Переменные:")
	fmt.Println("int:", integerVar, intVar)
	fmt.Println("float64:", floatVar, floatVar2)
	fmt.Println("string:", stringVar, stringVar2)
	fmt.Println("bool:", boolVar, boolVar2)

	// 4. Арифметические операции с двумя целыми числами
	a, b := 8, 3
	fmt.Println("Арифметические операции:")
	fmt.Printf("Сумма %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Разность %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Произведение %d * %d = %d\n", a, b, a*b)
	fmt.Printf("Частное %d / %d = %.2f\n", a, b, float64(a)/float64(b))

	// 5. Вычисление суммы и разности двух чисел с плавающей запятой
	sum, difference := calculate(5.5, 2.0)
	fmt.Printf("Сумма: %.2f, Разность: %.2f\n", sum, difference)

	// 6. Вычисление среднего значения трех чисел
	avg := average(10.0, 20.0, 30.0)
	fmt.Printf("Среднее значение: %.2f\n", avg)
}
