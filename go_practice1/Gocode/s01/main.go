package main

import "fmt"

type jojo struct {
	j bool
}
type gogo struct {
	jo jojo
	g  int
}

func main() {
	var j jojo
	var g gogo
	fmt.Println(j)
	fmt.Println(g)
}
