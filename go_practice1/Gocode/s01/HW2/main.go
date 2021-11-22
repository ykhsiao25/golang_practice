package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}
type circle struct {
	radius float64
}

//pointer receiver
func (t *triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

//type receiver
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

//error 本身interface就不可以作為receiver
// func (sh shape) f() {}

func (c circle) getArea() float64 {
	fmt.Println(c)
	return 3.14 * c.radius * c.radius
}

// func printArea(sh shape) {
// 	fmt.Println(sh.getArea())
// }
func (t *triangle) printArea() {
	fmt.Println(t.getArea())
}
func main() {
	tri := triangle{2, 3}
	// squ := square{2}
	(&tri).printArea()
	// printArea(&squ)

	// c1 := circle{5}
	// c2 := &circle{6}
	// fmt.Println((&c1).getArea())
	// fmt.Println(c2.getArea())

}
