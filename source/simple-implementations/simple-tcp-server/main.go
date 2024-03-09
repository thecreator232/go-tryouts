package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "5000"
	CONN_TYPE = "tcp"
)

func main() {

	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer l.Close()

	fmt.Println("listening on ", CONN_PORT)

	for {

		conn, err := l.Accept()
		fmt.Println("Accepted new connection")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		handleIncmoing(conn)

	}
}

func handleIncmoing(conn net.Conn) {
	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(buffer)
	conn.Write([]byte("test"))
	conn.Close()
}
