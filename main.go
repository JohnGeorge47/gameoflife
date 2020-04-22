package main

import "fmt"

var a = [][]int{
	{1, 1, 1, 0},
	{0, 1, 1, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

func main() {
	var futureboard [4][4]int
	for i := 0; i < 4; i++ {
		for j := 1; j < 4; j++ {
			futureboard = TraverseNeighbors(i, j, a, futureboard)
		}
	}
	fmt.Println(futureboard)
}

func TraverseNeighbors(m, n int, board [][]int, nextboard [4][4]int) [4][4]int {
	var aliveNeigbours int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if m+i >= 0 && n+j >= 0 && m+i < 4 && n+j < 4 {
				aliveNeigbours = board[m+i][n+j] + aliveNeigbours
			}
		}
	}
	aliveNeigbours = aliveNeigbours - board[m][n]
	if board[m][n] == 0 && aliveNeigbours == 3 {
		nextboard[m][n] = 1
	} else if board[m][n] == 1 && aliveNeigbours >= 4 {
		nextboard[m][n] = 0
	} else if board[m][n] == 1 && aliveNeigbours < 2 {
		nextboard[m][n] = 0
	} else {
		nextboard[m][n] = board[m][n]
	}
	return nextboard
}
