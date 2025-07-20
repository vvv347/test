package main

import (
	"fmt"
	"math"
)

// создаем интерфейс Shape - поверхность
type Shape interface {
	Area() float64
	Perimeter() float64
}

// создаем структуру Square - квадрат
type Square struct {
	X float64
}

// создаём метод Area() для работы с данными экз. структуры Square
func (s Square) Area() float64 {
	return s.X * s.X
}

// создаём метод Perimeter() для работы с данными экз. структуры Square
func (s Square) Perimeter() float64 {
	return 4 * s.X
}

// создаем структуру Circle
type Circle struct {
	Radius float64
}

// создаём метод Area() для работы с данными экз. структуры Circle
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// создаём метод Perimeter() для работы с данными экз. структуры Circle
func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

// создаём функцию Calculate, аргументом является переменная типа интерфейса Shape (экземпляры структур)
// функция выводит результаты действий методов интерфейса Shape
func Calculate(x Shape) {
	fmt.Printf("Area: %f,\nPerimeter: %f\n\n", x.Area(), x.Perimeter())
}

func main() {
	//создаём в качестве переменных (var) экз. стр. Square и Circle и заполняем их поля и
	s := Square{X: 20}
	c := Circle{Radius: 10}
	//вводм в качестве аргумента функции экз. стр. (var)
	Calculate(s)
	Calculate(c)
}
