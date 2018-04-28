package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func acceptLoopForSlaves(l net.Listener) {

	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("New Slave Registered!!")

	}
}
func acceptLoopForClients(l net.Listener) {
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("New connection found!")

		var (
			buf = make([]byte, 1024)
			r   = bufio.NewReader(c)
			// w   = bufio.NewWriter(c)
		)

		n, err := r.Read(buf)
		data := string(buf[:n])

		log.Printf("Received: %s", data)
		c.Write([]byte(data))
		log.Printf("Sent: %s", data)

		//go listenConnection(c)
	}
}

func main() {

	fmt.Println("The server is listening on Port 3000 for Slaves")

	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The server is listening on Port 8080 for Clients")

	listener2, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go acceptLoopForSlaves(listener)
	acceptLoopForClients(listener2) // run in the main goroutine

}
