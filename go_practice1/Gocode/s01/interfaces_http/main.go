package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//一定要加上http:// 否則會出問題
	resp, err := http.Get("http://google.com")
	//Golang的null叫做 "nil"
	if err != nil {
		fmt.Println("Error: ", err)
		//記得要直接exit
		os.Exit(1)
	}
	// //會看到一堆東西(像是header之類的資訊，但找不到body)
	// fmt.Println(resp)

	// bs := make([]byte, 99999)
	// // read bytes in the slice of byte
	// resp.Body.Read(bs)
	// // fmt.Printf("type is : %T ", bs)
	// //[]byte 可以直接轉換成string
	// fmt.Println(string(bs))

	//另解 直接用io API output
	//os.Stdout有overwrite Write()，所以可以做為writer interface 代入
	//流程 -> os之property(os的) -> 本身為 *File type -> File有實作write() method -> 為writer interface
	io.Copy(os.Stdout, resp.Body)
}
