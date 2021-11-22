package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type thumbnail struct {
	URL           string
	Width, Height int
}
type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

type city struct {
	Latitude, Longitude float64
	City                string
}

type cities []city

func main() {
	var dst img
	s := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`
	err := json.Unmarshal([]byte(s), &dst)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dst)

	for i, v := range dst.IDs {
		fmt.Println(i, v)
	}
	fmt.Println(dst.Thumbnail.URL)

	var data2 cities
	s2 := `[{"precision":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"precision":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`
	//注意 一定要用 "&" address value
	//可以只取得要用的attributes(json)
	err = json.Unmarshal([]byte(s2), &data2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data2)
}
