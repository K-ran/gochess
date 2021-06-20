package gochess

// represents the state of a square on board
type square struct {
	id           int    // used to represent the position of square 0-63
	piecePointer *piece // pointer to a piece lying on it, if any
}

type direction int

const (
	left direction = iota
	right
	top
	bottom
)

// A few helper functions
// checks if the square id is left border of board
func isLeftBorder(position int) bool {
	leftBorderIndex := [...]int{56, 48, 40, 32, 24, 16, 8, 0}

	for _, v := range leftBorderIndex {
		if v == position {
			return true
		}
	}
	return false
}

// checks if the square id is right border of board
func isRightBorder(position int) bool {
	rightBorderIndex := [...]int{63, 55, 47, 39, 31, 23, 15, 7}

	for _, v := range rightBorderIndex {
		if v == position {
			return true
		}
	}
	return false
}

// checks if the square id is top border of board
func isTopBorder(position int) bool {
	topBorderIndex := [...]int{56, 57, 58, 59, 60, 61, 62, 63}

	for _, v := range topBorderIndex {
		if v == position {
			return true
		}
	}
	return false
}

// checks if the square id is bottom border of board
func isBottomBorder(position int) bool {
	bottomBorderIndex := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	for _, v := range bottomBorderIndex {
		if v == position {
			return true
		}
	}
	return false
}

// returns the adjacent index based on direction. If out of board, returns -1
func getAdjecentSquare(currentSquareId int, dir direction) int {

	if currentSquareId == -1 {
		return -1
	}

	//first some exception cases
	switch {
	case isLeftBorder(currentSquareId) && dir == left:
		fallthrough
	case isRightBorder(currentSquareId) && dir == right:
		fallthrough
	case isBottomBorder(currentSquareId) && dir == bottom:
		fallthrough
	case isTopBorder(currentSquareId) && dir == top:
		return -1
	}

	switch dir {
	case left:
		return currentSquareId - 1
	case right:
		return currentSquareId + 1
	case bottom:
		return currentSquareId - 8
	case top:
		return currentSquareId + 8
	}
	return -1
}
