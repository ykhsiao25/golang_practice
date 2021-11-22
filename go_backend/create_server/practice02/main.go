package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//listen port -> accept port -> do something -> get resource back
func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err2 := li.Accept()
		if err2 != nil {
			log.Fatal(err)
			continue
		}
		// because we use go routine here, so Close() in handle()
		go handle(conn)

	}
}
func handle(conn net.Conn) {
	// return false when eof or an error
	scanner := bufio.NewScanner(conn)

	//token 由 type Splitfunc規範
	// scanner有個property叫做split (SplitFunc type)
	// 而NewScanner() 的 default split 就是SplitLines() .
	// SplitLines() 就是碰到 '\n'會停下，這樣scanner就知道要讀到哪了
	// for each token, it'll do a loop until it hits no more token
	// but Conn is an open stream, so never end
	for scanner.Scan() {
		//這邊應該是因為預設split() 為讀取一行，所以才覺得output為一行
		ln := scanner.Text()
		//os.Stdout
		fmt.Println(ln)
		fmt.Println("A line")

		// to client
		//注意 如果用browser加入這行，defer conn會close()，但如果用telnet則不會
		fmt.Fprintf(conn, "I heard what you say: %s\n", ln)
	}

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	defer conn.Close()
	fmt.Println("Code got here ! ")
}
