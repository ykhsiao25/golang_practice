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

//conn 都要帶著走，不然傳不回browser
func handle(conn net.Conn) {
	defer conn.Close()
	//Read the request
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			//在request 插入router 去找對應的webapp
			mux(conn, ln)
		}
		// headers are done(request are done)
		if ln == "" {
			break
		}
		i++
	}
}
func mux(conn net.Conn, req_line string) {
	met := strings.Fields(req_line)[0]
	url := strings.Fields(req_line)[1]
	fmt.Println("***METHOD", met)
	fmt.Println("***URI", url)
	//看是哪種方法 and url 去做不同的response
	if met == "GET" && url == "/" {
		index(conn)
	}
	if met == "GET" && url == "/about" {
		about(conn)
	}
	if met == "GET" && url == "/contact" {
		contact(conn)
	}
	if met == "GET" && url == "/apply" {
		apply(conn)
	}
	if met == "POST" && url == "/applyProcess" {
		applyProcess(conn)
	}
}

// a href 每個都要寫，否則點他render出來的畫面那些東西會不見
// 預設action 應該是GET
func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>ABOUT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>CONTACT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

//action指的就是打過去的url會到哪執行(你可以把它看成call API，打這個url，就會執行哪個API)
func apply(conn net.Conn) {
	//為什麼 form post 要寫在這？ 因為執行完他就會跳出button(寫在applyprocess() 他永遠跳不出來)
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
	<strong>APPLY</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	<form method="POST" action="/applyProcess">
	<input type="submit" value="applyProcess">
	</form>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
