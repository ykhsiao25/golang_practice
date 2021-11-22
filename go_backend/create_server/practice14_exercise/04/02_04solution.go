package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer conn.Close()

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			if ln == "" {
				fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
				break
			}
			fmt.Println(ln)
			fmt.Println("A line")
		}

		// we never get here
		// we have an open stream connection
		// how does the above reader know when it's done?
		// write to connection
		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		io.WriteString(conn, "\r\n")

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}

//if no limit
// type hotdog int

// func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "I see you connected.")
// }
// func main() {
// 	var d hotdog
// 	http.ListenAndServe(":8080", d)
// }
