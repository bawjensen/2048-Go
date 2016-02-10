package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	board := NewBoard(4)

	fmt.Println(board.String())

	board.AddRandomTile()

	fmt.Println(board.String())
}