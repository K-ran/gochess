package main

import "github.com/K-ran/gochess/gochess"

func main() {
	board := gochess.NewBoardState()
	board.Display()
	newBoard := board.GetCopyOfState()
	newBoard.Display()
}
