package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// listen -> get listener -> accept connection -> go routine to handle -> write to conn
// -> scanner -> scan the line -> do something -> conn close -> listener close
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
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	// if true, then still have something and no error
	for scanner.Scan() {
		ln := scanner.Text()
		// same as split(" ") 	去空白
		fs := strings.Fields(ln)
		fmt.Println(fs)
		if len(fs) < 1 {
			continue
		}
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\r\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintf(conn, "Expected value!\r\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintf(conn, "INVALID COMMAND "+fs[0]+"\r\n")
			continue
		}
	}
}
