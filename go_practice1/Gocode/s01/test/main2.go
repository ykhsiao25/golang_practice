package main

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	fmt.Println(r)
}
