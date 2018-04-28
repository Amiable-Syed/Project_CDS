package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	//Checking Length of CommandLineArguments
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	searchText := os.Args[1] //Password To search
	server := os.Args[2]     //server port

	fmt.Println("Connecting to the server at port 3000")

	conn, err := net.Dial("tcp", server)

	// defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	conn.Write([]byte(searchText))
	log.Printf("Sent: %s", searchText)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Received and Exiting: %s", buff[:n])

}
