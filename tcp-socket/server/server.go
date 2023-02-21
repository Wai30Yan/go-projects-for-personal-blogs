package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
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
	conn, err := listener.Accept()
	for {
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
        conn.Write([]byte("You sent "+ temp + " at " + time.Now().Format("15:04:05\n")))
	}
}

func handleConnection(conn net.Conn) {
	// adding the for loop in main func
}