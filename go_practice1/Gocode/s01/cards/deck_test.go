// 注意命名方式
// 檔案名: <funcName>_test.go
// func名: TestFuncName()
package main

import (
	"os"
	"testing"
)

// TestFunc() 必須大寫開頭
//go mod init deck_test.go (加上這行可以解決 go: go.mod file not found in current directory or any parent directory; see 'go help modules')
// 只要test一個func要按上面的run test
func TestNewDeck(t *testing.T) {
	cards := newCards()
	if len(cards) != 12 {
		t.Errorf("Expected deck length of 16, but got %v", len(cards))
	}
	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", cards[0])
	}
	if cards[len(cards)-1] != "Three of Clubs" {
		t.Errorf("Expected last card of Three of Clubs, but got %v", cards[len(cards)-1])
	}
}

//務必注意Test後必須大寫，否則他連跑都不跑(且不會有error)
func TestNewCardsAndnewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	cards := newCards()
	cards.saveToFile("_decktesting")

	loadedcards := newDeckFromFile("_decktesting")
	if len(loadedcards) != len(cards) {
		t.Errorf("Expected %v but got %v !!!", len(cards), len(loadedcards))
	}

	os.Remove("_decktesting")
}
