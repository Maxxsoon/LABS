package main

import (
	"fmt"
	"math"
)

// Структура Person
type Person struct {
	name string
	age  int
}

// Метод для вывода информации о человеке
func (p Person) Info() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

// Метод birthday для структуры Person
func (p *Person) birthday() {
	p.age++
}

// Структура Circle
type Circle struct {
	radius float64
}

// Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Интерфейс Shape
type Shape interface {
	Area() float64
}

// Структура Rectangle
type Rectangle struct {
	width, height float64
}

// Метод для вычисления площади прямоугольника
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Функция, которая принимает срез интерфейсов Shape и выводит площадь каждого объекта
func printAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Area: %.2fn", shape.Area())
	}
}

// Интерфейс Stringer
type Stringer interface {
	String() string
}

// Структура Book
type Book struct {
	title  string
	author string
}

// Реализация метода String для структуры Book
func (b Book) String() string {
	return fmt.Sprintf("Book Title: %s, Author: %s", b.title, b.author)
}

func main() {
	// Пример использования структуры Person
	person := Person{name: "Alice", age: 30}
	fmt.Println(person.Info())
	person.birthday()
	fmt.Println(person.Info())

	// Пример использования структур Circle и Rectangle
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 6}

	shapes := []Shape{circle, rectangle}
	printAreas(shapes)

	// Пример использования структуры Book и интерфейса Stringer
	book := Book{title: "Go Programming", author: "John Doe"}
	fmt.Println(book.String())
}
