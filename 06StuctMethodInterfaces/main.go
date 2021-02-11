package main

import "fmt"

func main() {
	data := Perimeter(10.00, 25.00)
	rectangle := Rectangle{25.00, 30.00}
	rectangle2 := Rectangle{25.00, 30.00}
	data2 := PerimeterRectangle(rectangle)
	fmt.Println(data)
	fmt.Println(data2)
	fmt.Println(rectangle2.getArea())

}

// Perimeter retrun value of perimeter
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// Rectangle type of Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// PerimeterRectangle return value of PerimeterRectangle
func PerimeterRectangle(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Width)
}

func (r Rectangle) getArea() float64 {
	return 2 * (r.Width + r.Height)
}
