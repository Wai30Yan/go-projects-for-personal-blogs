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
		fmt.Println(" client: " + temp)
		_, err := strconv.Atoi(temp);

		if err != nil {
			conn.Write([]byte("please use 2 digit integer\n"))
			log.Print("not an integer...")
		}

		num, _ := strconv.Atoi(temp)
		row := int(num / 10) - 1
		col := int(num % 10) - 1

		(*gb)[row][col] = 1

		rwin := rowCheckWin(gb)
		cwin := colCheckWin(gb)
		dwin := diagCheckWin(gb)

		fmt.Print(gb, rwin, cwin, dwin)
		
		response, _ := json.Marshal(gb)

		if rwin || cwin || dwin {
			message := append(response, []byte("You Won!!!\n")...)
			conn.Write(message)
			break
		}



		conn.Write([]byte(string(response)+"\n"))

	}

}

func colCheckWin(gb *GameBoard) bool {
    for col := 0; col < 3; col++ {
        var score int
        for row := 0; row < 3; row++ {
            score += (*gb)[row][col]
        }
        if score == 3 {
			return true
		}
    }
	return false
}

func rowCheckWin(gb *GameBoard) bool {
	var score int
	for i := range (*gb) {
		for j := range (*gb)[i] {
			if (*gb)[i][j] == 1 {
				score += (*gb)[i][j]
			}
		}
		if score == 3 {
			return true
		}
		score = 0
	}
	return false
}

func diagCheckWin(gb *GameBoard) bool {
	for i := 0; i < 3; i++ {
		var score int
		if i == 0 {
			score = (*gb)[i][2] + (*gb)[1][1] + (*gb)[2][i]
			fmt.Print("diagScore: ", score, "\n")
			if score == 3 {
				return true
			}
		}
		for j := 0; j < 3; j++ {
			if i == j {
				score += (*gb)[i][j]
			}
			if score == 3 {
				return true
			}
		}
	}
	return false
}
