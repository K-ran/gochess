package gochess

// represents the state of a square on board
type square struct {
	id           int    // used to represent the position of square 0-63
	piecePointer *piece // pointer to a piece lying on it, if any
}
