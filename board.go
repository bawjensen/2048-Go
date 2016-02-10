package main

import (
	// "fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Board struct {
	size int
	board [][]Tile
}

func NewBoard(size int) (board *Board) {
	boardArray := make([][]Tile, size)

	for i := 0; i < size; i++ {
		boardArray[i] = make([]Tile, size)
	}

	board = &Board{ size: size, board: boardArray }

	return board
}

func (self *Board) String() (buffer string) {
	buffer = ""

	for i := 0; i < self.size; i++ {
		for j := 0; j < self.size; j++ {
			value := self.board[i][j]
			buffer += strings.Repeat(" ", 6 - DigitLength(int(value)))
			buffer += strconv.Itoa(int(value))
		}
		buffer += "\n\n"
	}

	return buffer
}

func (self *Board) AddRandomTile() {
	randI := rand.Int() % 4
	randJ := rand.Int() % 4

	self.board[randI][randJ] = RandomTile()
}