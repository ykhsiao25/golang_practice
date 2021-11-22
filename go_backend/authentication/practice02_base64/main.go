package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "2123"

	//法一
	// s_std := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/" //後面這就encodeStd
	// // 按照何種標準之寫法
	// s_base := base64.NewEncoding(s_std).EncodeToString([]byte(s))
	// fmt.Println(s_base)

	s_base := base64.StdEncoding.EncodeToString([]byte(s))
	s2, err := base64.StdEncoding.DecodeString(s_base)
	if err != nil {
		log.Fatalln(err)
	}
	//True
	fmt.Println(string(s2) == s)
}
