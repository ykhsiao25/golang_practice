// package main
// import "fmt"
// func main(){
// 	// var card string = "Ace of Spades"
// 	card := "Ace of Spades" //rely on the compiler to define the card type
// 	fmt.Println(card)
// }

//go global變數必須宣告但不可定義
package main

func main() {
	// cards := newCards()

	// // //注意這邊是宣告 要多注意
	// // handCards, remainCards := deal(cards, 5)

	// // handCards.print()
	// // remainCards.print()

	// // fmt.Println(cards.toString())
	// fmt.Println(cards.saveToFile("my_cards.txt"))

	// cards := newDeckFromFile()
	cards := newDeckFromFile("my_cards.txt")
	cards.shuffle()
	cards.print()
}
