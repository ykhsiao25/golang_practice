package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// listen -> get listener -> accept connection -> go routine to handle -> write to conn
// -> scanner -> scan the line -> do something -> conn close -> listener close
//request 只要parse就好，格式會在client那邊寫好
//但server response要自己寫好，包含response header and body
func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err2 := li.Accept()
		if err2 != nil {
			log.Fatal(err2)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//Read the request
	ln := request(conn)
	// go request(conn) //也可以 但defer conn.Close()不能放在main routine

	//Write the response
	response(conn, ln)
	// go response(conn) //也可以 但defer conn.Close()不能放在main routine
}

func request(conn net.Conn) []string {
	i := 0
	var m string
	var u string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			//parse the request action
			m = strings.Fields(ln)[0]
			u = strings.Fields(ln)[1]
			fmt.Println("***METHOD", m)
			fmt.Println("***URL", u)
			return []string{m, u}
		}
		// headers are done
		if ln == "" {
			break
		}
		i++
	}
	return []string{}
}
func response(conn net.Conn, ln []string) {
	body := `<!DOCTYPE html>
				<html lang="en">
					<head><meta charset="UTF-8">
						<title></title>
					</head>
					<body>
						<strong>Hello World</strong>
					</body>
				</html>`

	fmt.Println(ln)
	//這些都要傳過去，否則會404 (這應該算response的header)
	fmt.Fprintf(conn, "HTTP/1.1 200 OK \r\n")
	// fmt.Fprintf(conn, "Content-Length %d\r\n", len(body))
	// fmt.Fprintf(conn, "Content-Type \r\n")
	fmt.Fprintf(conn, "\r\n")
	// io.WriteString(conn, ln[0]+"\r\n")
	// io.WriteString(conn, ln[1]+"\r\n")
	fmt.Fprintf(conn, body)
}
