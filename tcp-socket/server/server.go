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
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := ":" + arguments[1]
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()

		defer conn.Close()

		if err != nil {
			fmt.Println(err)
			return
		}

		netData, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
            fmt.Println(err)
            return
        }

		temp := strings.TrimSpace(netData)

		if temp == "STOP" {
            break
        }

		fmt.Println("client: " + temp)
        response := "hello from server\n"
        conn.Write([]byte(string(response)))
	}
}

func handleConnection(conn net.Conn) {
	// adding the for loop in main func
}