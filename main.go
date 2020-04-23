package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8090", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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

func initialBoard(g Game) Game {
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			g.initboard[i][j] = 0
		}
	}
	return g
}

func (g Game) ResetBoard() {
	g = initialBoard(g)
}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("here")
		action(w, r)
	})
	http.HandleFunc("/", checkalive)
	log.Println("Server listening on port:", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("error while starting server", err)
	}
}

func checkalive(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("here"))
}

func action(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		start()
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func start()[][]int{
	boardWidth := 4
	boardHeight := 4
	initboard := makeNewBoard(boardHeight, boardWidth)
	futureboard := makeNewBoard(boardHeight, boardWidth)
	start := initialBoard(initboard)
	for r := 0; r < 50; r++ {
		for i := 0; i < boardHeight; i++ {
			for j := 1; j < boardWidth; j++ {
				futureboard.initboard = futureboard.TraverseNeighbors(start.initboard, futureboard, i, j)
			}
		}
		start = futureboard
	}
	return start.initboard
}

func (g Game) TraverseNeighbors(board [][]int, nextboard Game, m int, n int) [][]int {
	var aliveNeigbours int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if m+i >= 0 && n+j >= 0 && m+i < g.height && n+j < g.width {
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
