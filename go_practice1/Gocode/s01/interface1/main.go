package main

import "fmt"

type bot interface {
	getGreeting() string //只要其他type or struct call的方法完全符合(含return)介面所定義之方法，就可暫時轉為interface type
	// getGreeting(int) string //only type, no var name
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

//若receiver在func內完全沒用到，可以寫receiver的type就好 (不用寫var名)
func (englishBot) getGreeting() string {
	return "hi there!"
}
func (spanishBot) getGreeting() string {
	return "hola"
}

//其他struct or type 有用到介面所定義的方法的話
//那call interface method的時候，他們就自動暫定轉為interface type去call method
func printGreeting(b bot) {
	fmt.Printf("%T", b)
	fmt.Println(b.getGreeting())
}

// //go 沒有像java overloading (他會變成重複定義func name) 也沒有generic type
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// // func printGreeting(sb.spanishBot) {
// // 	fmt.Println(sb.getGreeting)
// // }
