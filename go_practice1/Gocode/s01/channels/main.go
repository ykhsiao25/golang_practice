package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		//go 的 http module必須有protocol 存在(這也是為啥前面需要有http://
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		fmt.Println("initial")
		go checkLink(link, c)
	}
	// Could be var <- c and Prtinln(var) or directly use its return value
	// for main routine, it's a kind of blocking(receiving messages from a channel is a blocking thing)
	// fmt.Println(<-c) //雖然有很多go routine，但如果只有一個blocking for main routine, then end the program

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	//How can you make sure that we only receivce a number of messages equals to requests
	// for {
	// 	//go 知道 channel知道會output string
	// 	//這邊不加上go就是把上一個for loop的channel都輸出而已，但若加上go就是不斷把channel拿出的string再input到channel(所以是不斷循環)
	// 	// <-c = receiving a value through the channel is still a blocking operation
	// 	// c<-<something> sending a value through the channel is still a blocking operation
	// 	go checkLink(<-c, c)
	// }

	//透過 := and range, 我們可以視為main routine 正等待c channel 取出value 並 assign到 l (可視為短暫的blocking)
	// whenever a value comes out of it, assign the value to l
	for l := range c {
		// sleep如果放在這不夠恰當，因為許多go routine可能已經完成，卻也在等
		// 雖然message不會消失，但會queued up or lined up
		// time.Sleep(5 * time.Second) //1
		// go checkLink(l, c)

		fmt.Println("llllllllllllllllllllllll", l)
		// // use function literal(用 function literal的原因，純粹是想go routine 去執行多行code而已)
		// // //如果這樣寫的話，go routine可能會取得同一個l，而不是改過的l
		// go func() {
		// 	time.Sleep(5 * time.Second) //3
		// 	fmt.Println("for looooop")
		// 	checkLink(l, c)
		// }()

		// //正確寫法 (這樣go routine就會複製一份)
		// // use function literal(用 function literal的原因，純粹是想go routine 去執行多行code而已)
		go func(link string) {
			fmt.Println("for looooop")
			//放在goo routine內的 pause, 是可以確保只有在checkLink()之前會pause
			time.Sleep(5 * time.Second) //3

			checkLink(link, c)
		}(l) //同for loop 的 l (for loop 的var, link則是for loop內的block 之var)
	}
}

func checkLink(link string, c chan string) {
	//這邊是fetch link 之前 "go routine"先等五秒 (和上面的結果可能是一樣的，但涵義不同，這邊是go routine的任務含pause)
	// 但我們不會期望每個 go routine 都給我們暫停5秒才做，肯定是期望他立即執行，且馬上給我fetch a link
	// time.Sleep(5 * time.Second) // 2
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		//這邊加上return 是確保這個func後面或其他地方不會再被執行(事實上因為void 所以return 也不會有東西)

		c <- link
		return
	}
	fmt.Println("this is check check")
	fmt.Println(link, " is up")
	c <- link
}
