package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"net"
	"time"
	hackson0605 "LearningSample/hackathon"
	_ "encoding/xml"
)

func test() {

	log.Println("start connection...")
	conn, err := net.DialTimeout("tcp", ":9627", 30*time.Second)
	if err != nil {
		log.Println("connection error:", err)
		return
	}
	defer conn.Close()

	var buf bytes.Buffer
	for i := 0; ; i++ {
		if i == math.MaxInt8 {
			i = 0
			log.Println("reset")
		}
		buf.WriteString("Client Hello ")
		buf.WriteString(string(i))
		send := buf.Bytes()
		n, err := conn.Write(send)
		if err != nil {
			log.Println("write error:", err)
		} else {
			log.Printf("write % bytes, content is %s\n", n, send[:n])
		}
		buf.Reset()
		time.Sleep(time.Second * 5)
	}
}

func worker(die chan bool, index int) {
	fmt.Println("Begin: This is Worker:", index)
	for {
		select {
		//case xx：
		//做事的分支
		case <-die:
			fmt.Println("Done: This is Worker:", index)
			return
		}
	}
}

func generator(strings chan string) {
	strings <- "Five hour's New York jet lag"
	fmt.Println("here")
	strings <- "and Cayce Pollard wakes in Camden Town"
	strings <- "to the dire and ever-decreasing circles"
	strings <- "of disrupted circadian rhythm."
	close(strings)
}

func newUniqueIDService() <-chan string {
	id := make(chan string)
	go func() {
		var counter int64 = 0
		for {
			id <- fmt.Sprintf("%x", counter)
			counter += 1
		}

	}()
	return id
}

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func readBlock() {

	conn, err := net.DialTimeout("tcp", ":9627", 30*time.Second)
	if err != nil {
		log.Println("connection error:", err)
		return
	}
	defer conn.Close()

	for {
		// read from the connection
		time.Sleep(2 * time.Second)
		var buf = make([]byte, 60000)
		log.Println("start to read from conn")
		conn.SetReadDeadline(time.Now().Add(time.Hour))
		n, err := conn.Read(buf)

		if err != nil {
			log.Printf("conn read %d bytes,  error: %s", n, err)

			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				continue
			}
		}

		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}




func netconfClient() {

	//conn, err := net.DialTimeout("tcp", ":9627", 30*time.Second)
	conn, err := net.DialTimeout("tcp", "135.242.212.247:830", 30*time.Second)
	if err != nil {
		log.Println("connection error:", err)
		return
	}
	defer conn.Close()


	var transport = hackson0605.NewTransportFromConn(conn)

	log.Println("start to send hello message")
	// 1. send hello reply to client, use the same session id and capabilities
	transport.SendHello(&hackson0605.HelloMessage{
		SessionID:    12345,
		Capabilities: hackson0605.DefaultCapabilities,
	})
	log.Println("start to receive hello message")
	// 2. receive <hello> from client
	serverHello,err := transport.ReceiveHello()
	if err!=nil {
		fmt.Errorf(err.Error())
		return
	}

	fmt.Println("receive server hello")
	fmt.Println(serverHello.SessionID)
	fmt.Println(serverHello.Capabilities)
	fmt.Println(serverHello)

}

func main2() {

	//fmt.Println(sample)
	//for i := 0; i < len(sample); i++ {
	//	fmt.Printf("%x", sample[i])
	//}
	//fmt.Println()
	//fmt.Printf("%x\n", sample)
	//fmt.Printf("% x\n", sample)

	//go hackathon.StartServer()

	netconfClient()

	time.Sleep(time.Hour)

}
