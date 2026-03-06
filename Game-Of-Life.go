package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printBoard(board [][]int) {
	for _, row := range board {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	board := make([][]int, 10)
	for i := range board {
		board[i] = make([]int, 10)
	}

	board[rand.Intn(9)][rand.Intn(9)] = 1

	printBoard(board)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

}
