package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//print 出 $0 ~ $N (file, par1, ...)
	fmt.Println(os.Args)
	//File implements both Writer{} and Reader{} interface
	f, err := os.Open(os.Args[1])
	if err != nil {
		//或是直接使用log.Fatal(err) log package
		fmt.Println("Error is :", err)
		os.Exit(1)
	}
	//Both value and pointer type can be parameters
	io.Copy(os.Stdout, f)
}
