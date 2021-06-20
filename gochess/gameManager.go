package gochess

// top level class to store the game state and history
type gameManager struct {
	boards []*BoardState // stores all moves/state at each move of a game
	turn   colorType     // who plays next
}

func NewGameManager() *gameManager {
	newGame := new(gameManager)
	board := NewBoardState()
	newGame.boards = append(newGame.boards, board)
	newGame.turn = white
	return newGame
}

// displays current position of the board
func (game *gameManager) DisplayCurrentBoard() {
	game.boards[len(game.boards)-1].Display()
}

// called if a user wants to make a move
func (game *gameManager) MakeComputerMove() {
	nextBoards := game.boards[len(game.boards)-1].GetNextStates(game.turn)

	// just for debugging
	for _, b := range nextBoards {
		b.Display()
	}

}

// plays the next best move for the turn colorType
func (game *gameManager) nextStep() {
	//TODO
}
