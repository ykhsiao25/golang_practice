package main

import "fmt"

type slice []int

func main() {
	s := newSlice(10)
	s.print()
}

func newSlice(length int) slice {
	//記得要加上 {} (他才有初始化)
	s := slice{}
	for i := 0; i < 11; i++ {
		s = append(s, i)
	}
	return s
}

func (s slice) print() {
	for _, value := range s {
		if value%2 == 0 {
			fmt.Printf("%v is even\n", value)
		} else {
			fmt.Printf("%v is odd\n", value)
		}
	}
}
