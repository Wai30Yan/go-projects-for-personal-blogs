package main


type GameBoard [][]int

func main() {
	gb := &GameBoard{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	startServer(gb)
}