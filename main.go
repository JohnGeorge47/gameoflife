package main

import "fmt"

//var a = [][]int{
//	{0, 0, 1, 0},
//	{0, 1, 1, 0},
//	{0, 0, 0, 0},
//	{0, 0, 0, 0},
//}

func makeNewBoard(m int,n int)[][]int{
	b:=make([][]int,m)
	for elem:=range b{
		b[elem]=make([]int,n)
	}
	return b
}

func initialBoard(b [][]int,m int,n int)[][]int{
	for i:=0;i<m ;i++  {
		for j:=0;j<n ;j++  {
			b[i][j]=0
		}
	}
	return b
}

func main() {
	boardWidth:=4
	boardHeight:=4
	initboard:=makeNewBoard(boardHeight,boardWidth)
	futureboard:=makeNewBoard(boardHeight,boardWidth)
	start:=initialBoard(initboard,boardWidth,boardHeight)
	for i := 0; i < 4; i++ {
		for j := 1; j < 4; j++ {
			futureboard = TraverseNeighbors(i, j, start, futureboard)
		}
	}
	fmt.Println(futureboard)
}

func TraverseNeighbors(m, n int, board [][]int, nextboard [][]int) [][]int {
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
