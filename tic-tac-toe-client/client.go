package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8080")

	if err != nil {
		log.Print(err)
	}
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text + "\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)

		if message == "You won!!!" {
			fmt.Println("TCP client exiting...")
			break
		}
	}
}