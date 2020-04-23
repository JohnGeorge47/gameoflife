package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Game struct {
	initboard [][]int
	height    int
	width     int
}

func makeNewBoard(m int, n int) Game {
	b := make([][]int, m)
	for elem := range b {
		b[elem] = make([]int, n)
	}
	board := Game{
		initboard: b,
		height:    m,
		width:     n,
	}
	return board
}

func initialBoard(g Game) Game{
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			g.initboard[i][j] = 0
		}
	}
	return g
}

func (g Game)ResetBoard(){
	g=initialBoard(g)
}

func main() {
	boardWidth := 4
	boardHeight := 4
	initboard := makeNewBoard(boardHeight, boardWidth)
	futureboard := makeNewBoard(boardHeight, boardWidth)
	start := initialBoard(initboard)

	for r := 0; r < 50; r++ {
		for i := 0; i < boardHeight; i++ {
			for j := 1; j < boardWidth; j++ {
				futureboard.initboard = futureboard.TraverseNeighbors(start.initboard, futureboard)
			}
		}
		start = futureboard
	}
	fmt.Println(futureboard)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		action(w,r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err!=nil{
		log.Fatal("error while starting server",err)
	}
}

func action( w http.ResponseWriter,r *http.Request){
	for  {
		c, err := upgrader.Upgrade(w, r, nil)
		if err!=nil{
			log.Fatal(err)
			return
		}
		defer c.Close()
		mt,message,err:=c.ReadMessage()
        fmt.Println(mt,message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (g Game) TraverseNeighbors(board [][]int, nextboard Game) [][]int {
	m := nextboard.height
	n := nextboard.width
	var aliveNeigbours int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if +i >= 0 && n+j >= 0 && m+i < 4 && n+j < 4 {
				aliveNeigbours = board[m+i][n+j] + aliveNeigbours
			}
		}
	}
	aliveNeigbours = aliveNeigbours - board[m][n]
	if board[m][n] == 0 && aliveNeigbours == 3 {
		nextboard.initboard[m][n] = 1
	} else if board[m][n] == 1 && aliveNeigbours >= 4 {
		nextboard.initboard[m][n] = 0
	} else if board[m][n] == 1 && aliveNeigbours < 2 {
		nextboard.initboard[m][n] = 0
	} else {
		nextboard.initboard[m][n] = board[m][n]
	}
	return nextboard.initboard
}
