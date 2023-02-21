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
	turn := true

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

		fmt.Println("turn: ", turn)

		win := play(gb, temp, &turn)

		fmt.Print(gb, win)
		
		response, _ := json.Marshal(gb)

		if win {
			message := append(response, []byte("You Won!!!\n")...)
			conn.Write(message)
			break
		}
		conn.Write([]byte(string(response)+"\n"))
	}
}

func play(gb *GameBoard, temp string, turn *bool) bool {
	num, _ := strconv.Atoi(temp)
	row := int(num / 10) - 1
	col := int(num % 10) - 1

	if *turn {
		(*gb)[row][col] = 1
		*turn = false
		return checkWin(gb, 3, 1)
	} else {
		(*gb)[row][col] = 2
		*turn = true
		return checkWin(gb, 6, 2)
	}
}

func checkWin(gb *GameBoard, winScore, point int) bool {
	r := rowCheckWin(gb, winScore, point)
	c := colCheckWin(gb, winScore, point)
	d := diagCheckWin(gb, winScore, point)

	if r || c || d {
		return true
	}

	return false
}

func colCheckWin(gb *GameBoard, winScore, point int) bool {
    for col := 0; col < 3; col++ {
        var score int
        for row := 0; row < 3; row++ {
			if (*gb)[row][col] == point {
				score += (*gb)[row][col]
			}
        }
        if score == winScore {
			return true
		}
    }
	return false
}

func rowCheckWin(gb *GameBoard, winScore, point int) bool {
	var score int
	for i := range (*gb) {
		for j := range (*gb)[i] {
			if (*gb)[i][j] == point {
				score += (*gb)[i][j]
			}
		}
		if score == winScore {
			return true
		}
		score = 0
	}
	return false
}

func diagCheckWin(gb *GameBoard, winScore, point int) bool {
	var score int
	if (*gb)[0][2] == point && (*gb)[1][1] == point && (*gb)[2][0] == point {
		score = (*gb)[0][2] + (*gb)[1][1] + (*gb)[2][0]
		fmt.Print("diagScore: ", score, "\n")
		if score == winScore {
			return true
		}

	}

	score = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j && (*gb)[i][j] == point {
				score += (*gb)[i][j]
			}
			if score == winScore {
				return true
			}
		}
	}
	return false
}
