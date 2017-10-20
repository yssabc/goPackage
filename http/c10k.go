package main

import (
	"fmt"
	"net"
	"log"
)

func handle(conn net.Conn) {
	fmt.Fprintf(conn, "%s", "v1.0")
	conn.Close()
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

