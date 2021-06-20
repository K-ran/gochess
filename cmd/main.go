package main

import "github.com/K-ran/gochess/gochess"

func main() {
	game := gochess.NewGameManager()
	game.DisplayCurrentBoard()
	game.MakeComputerMove()
}
