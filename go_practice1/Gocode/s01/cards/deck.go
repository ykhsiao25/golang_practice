package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of dck
// which is a slice of strings
//like C++ #typedef a_type b_type
type deck []string

func newCards() deck {
	cards := deck{}
	cardSuits := deck{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := deck{"Ace", "Two", "Three"}
	//這個range是像python 一樣 range(int) 產生一個list (也就是一個iterator)
	for _, suit := range cardSuits { //注意如果沒用到的變數，用 "_" 代替，否則會噴error
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// (d deck) is a receiver (not return value or function name)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// it can be cards deck(but it may cause confusion when we talk about the projects, so we use 'd' instead)
// you can also use (d deck) in receiver type to declare the func, but we just practice func here
// the last () is a return type set (return set just type needed)
// receiver還是parameter 要看func 在call的時候，會不會造成混淆 (如: cards.deal(5) 就會像是從cards內隨便取五張return，且cards也會少五張，但跟我們要做的事不同)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//deck to []string
	// []string(d)

	//[]string to string (use "," to combine them)
	return strings.Join([]string(d), ",")
}

//func WriteFile(filename string, data []byte, perm fs.FileMode) error
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	//means failing
	if err != nil {
		//Option1 - Print the error and return a call to newDeck()
		//Option2 - Print the error and entirely quit the program
		fmt.Println("Error: ", err)
		// if not zero, means some failed
		os.Exit(1)
	}

	//slice of type byte can be converted to string directly
	// []byte to string
	s := string(bs)

	// string to []string
	return deck(strings.Split(s, ","))

}

func (d deck) shuffle() {
	//real random 寫法
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for index := range d {
		//因為seed相同，所以這每次序列都一樣
		// newPosition := rand.Intn(len(d) - 1)

		newPosition := r.Intn(len(d) - 1)

		//swap 新寫法
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
