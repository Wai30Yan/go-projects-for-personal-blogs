package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func startServer(gb *GameBoard) {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Print(err)
	}

	defer listener.Close()

	conn, _ := listener.Accept()
	fmt.Println("Game is started...")
	for {
		netData, _ := bufio.NewReader(conn).ReadString('\n')
		temp := strings.TrimSpace(netData)
		fmt.Println("client: " + temp)
		_, err := strconv.Atoi(temp);

		if err != nil {
			conn.Write([]byte("please use 2 digit integer\n"))
			log.Print("not an integer...")
		}

		num, _ := strconv.Atoi(temp)
		row := int(num / 10) - 1
		col := int(num % 10) - 1

		(*gb)[row][col] = 1

		score := checkWin(gb, conn)

		if score == 3 {
			break
		}

		fmt.Print(gb, score)

		response, _ := json.Marshal(gb)

		conn.Write([]byte(string(response)+"\n"))

	}

}

func checkWin(gb *GameBoard, conn net.Conn) int {
	var score int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*gb)[i][j] == 1 {
				score++
			}
		}
		if score == 3 {
			conn.Write([]byte("You won!!!\n"))
		} else {
			score = 0
		}
	}
	return score
}
