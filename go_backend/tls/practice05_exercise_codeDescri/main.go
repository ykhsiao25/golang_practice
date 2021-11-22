package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type code struct {
	Code    int    `json:"Code"`
	Descrip string `json:"Descrip"`
}
type codes []code

func main() {
	var data codes
	s := `[{"Code":200,"Descrip":"StausOK"},{"Code":301,"Descrip":"StatusMovedPermanently"},
	{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},
	{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},
	{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},
	{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},
	{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},
	{"Code":500,"Descrip":"StatusInternalServerError"}]`
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, c := range data {
		fmt.Println(c.Code, c.Descrip)
	}
}
