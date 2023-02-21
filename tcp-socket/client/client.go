package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}
	connect := arguments[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	for {
		// reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}