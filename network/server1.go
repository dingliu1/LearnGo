package network

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// read from the connection
	var buf = make([]byte, 50)
	log.Println("start to read from conn")
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
		} else {
			log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":9627")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept new connection error", err)
			continue
		}
		go handleConnection(conn)
	}
}
