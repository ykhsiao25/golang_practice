package main

import "fmt"

type m map[string]int

func main() {
	// map宣告
	// key string, and value int
	colors := m{
		"red":   0,
		"bule":  1,
		"white": 2,
	}
	// for k, v := range colors {
	// 	fmt.Println(k, v)
	// }
	colors.printMap()
	fmt.Println(colors)
}

//注意 map type要加上key and value type (非常容易忘記，要注意!!!!)
func (m1 m) printMap() {
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
