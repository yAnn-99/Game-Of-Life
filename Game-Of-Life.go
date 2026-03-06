package main

import (
	"fmt"
	"time"
)

var board = makeBoard(10)

func makeBoard(size int) [][]int {
	b := make([][]int, size)
	for i := range b {
		b[i] = make([]int, size)
	}
	return b
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, val := range row {
			if val == 1 {
				fmt.Print(" 1 ")
			} else {
				fmt.Print(" 0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func resetBoard(board [][]int) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			board[row][col] = 0
		}
	}
}

func rez(row int, col int) {
	board[row][col] = 1
}

func kill(row int, col int) {
	board[row][col] = 0
}

func get(row int, col int) int {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return 0
	}
	return board[row][col]
}
func countNeighbors(row int, col int) int {
	count := 0

	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			count += get(row+dr, col+dc)
		}
	}
	return count
}

func nextgen() {

	next := makeBoard(len(board))

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			neighbors := countNeighbors(row, col)
			alive := board[row][col] == 1

			if alive {
				// Rule 1:
				if neighbors == 2 || neighbors == 3 {
					next[row][col] = 1
				}

			} else {
				// Rule 3
				if neighbors == 3 {
					next[row][col] = 1
				}
			}
		}
	}

	resetBoard(board)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			board[row][col] = next[row][col]
		}
	}
}

func main() {
	// base pattern, pcque flemme de faire un random au debut
	rez(0, 1)
	rez(1, 2)
	rez(2, 0)
	rez(2, 1)
	rez(2, 2)

	printBoard(board)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	generation := 1
	for range ticker.C {
		nextgen()
		printBoard(board)
		generation++

		if generation > 20 { // modif le 20 si ouloir plus ou moins de gen
			break
		}
	}
}

// toute les secondes ca print dans le terminal la prochaine generation
// le base pattern est predefini dns la suite de rez au debut du main
//
// il me manquait des fonctions et c'est pour ca que ca marchais pas
//
// main n'est que l'executable c'est similaire au jeux d'escape en 1ere en prog orienté objet
//
// rez vie
// kill meurt
// get ignore les extremités (merci romain)
//
// next determine comment sera la prochaine gen et utilise les 3 au dessus
// pour determiner ca selon les regles du game of life
//
// la rule 1 et 3 dans le code font en meme temps la rule 2 puisque reste pareil dans la matrice
//
// dans count neibhbors le gros truc avec dr et dc sert a ignorer le bloc en question
// printboard est ce qui print dans le terminal, on peut chnager 1 et 0 pour des symboles si on le veut
//
// makeboard est ce qui s'occupe de creer la matrice
