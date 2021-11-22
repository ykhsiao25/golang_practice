package main

import (
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

		// write to connection
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
