package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
		}
		// because we use go routine here, so Close() in handle()
		go handle(conn)

	}
}
func handle(conn net.Conn) {
	// err := conn.SetDeadline(time.Now().Add(10 * time.Microsecond))
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatal("CONN TIMEOUT")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		//這邊應該是因為預設split() 為讀取一行，所以才覺得output為一行
		ln := scanner.Text()
		//os.Stdout(server side)
		fmt.Println(ln)
		// to client (client side)
		fmt.Fprintf(conn, "I heard what you say: %s\n", ln)
	}
	defer conn.Close()
	fmt.Println("Code got here ! ")
}
