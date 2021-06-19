package main

import "github.com/K-ran/gochess/gochess"

func main() {
	board := gochess.NewBoardState()
	board.Display()
	board.WhitePieces[0].Color = false
}
