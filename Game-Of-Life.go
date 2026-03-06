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

	count := rand.Intn(20) + 3
	for i := 0; i < count; i++ {
		board[rand.Intn(10)][rand.Intn(10)] = 1
	}

	for range ticker.C {
		printBoard(board)

	}
	printBoard(board)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

}

//je vois pas dutout comment continuer, jpense que parcourir recursivement chaque case et le store dans une list / dict
// et une sorte de fonction update qui a acces a ces données et qui traite le prochain print
