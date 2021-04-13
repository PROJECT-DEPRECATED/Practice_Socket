package main

import (
	"io"
	"log"
	"net"
)

func main() {
	TCPSocket()
}

func TCPSocket() {
	l, err := net.Listen("tcp", ":8000")
	errorCheck(err)
	defer l.Close()

	for {
		conn, err := l.Accept()
		errorCheck(err)
		defer conn.Close()

		go ConnHandler(conn)
	}
}

func ConnHandler(conn net.Conn) {
	recvBuffer := make([]byte, 4096)

	for {
		n, err := conn.Read(recvBuffer)
		if err != nil {
			if io.EOF == err {
				log.Fatal(err)
				return
			}

			log.Fatal(err)
			return
		}

		if 0 < n {
			data := recvBuffer[:n]
			log.Println(string(data))
			_, err = conn.Write(data[:n])
			errorCheck(err)
		}
	}
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
