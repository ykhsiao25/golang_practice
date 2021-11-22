package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	//li *tcpListener
	li, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "\nHello from TCP server\n")
		// fmt.Fprintf(conn, "%T\n", conn) // *net.TCPConn
		fmt.Fprintln(conn, "How's your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope !")

		//若practice01 and practice04要互打 不可以用defer，不然永遠不會釋放.......
		// defer conn.Close()
		conn.Close()
	}
}
