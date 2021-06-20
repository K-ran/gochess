package gochess

// structure to hold a move
type move struct {
	from int // square index 0-63
	to   int // square index 0-63
}

// enum to represent piece type
type pieceType int8

const (
	pawn pieceType = iota
	bishop
	knight
	rook
	queen
	king
)

// enum to represent color of piece
type colorType bool

const (
	white colorType = true
	black colorType = false
)

// this structe represents the state of a piece on board
type piece struct {
	color            colorType // color white or black
	pieceValue       pieceType // what type of piece is it
	squarePointer    *square   // pointer to square it lies on, null if out of board
	pseudoLegalMoves []int     // stores pseudo legal moves at any point of time
}

// constructor to create new piece
func NewPeiece(color colorType, value pieceType, squarePointer *square) *piece {
	newPiece := new(piece)
	newPiece.color = color
	newPiece.pieceValue = value
	newPiece.squarePointer = squarePointer
	return newPiece
}

// implementing stringer interface for piece class
func (p *piece) String() string {
	if p == nil {
		return "  -  "
	}
	if p.color == white {
		switch p.pieceValue {
		case pawn:
			return "  P  "
		case rook:
			return "  R  "
		case knight:
			return "  N  "
		case bishop:
			return "  B  "
		case queen:
			return "  Q  "
		case king:
			return "  K  "
		}
	} else {
		switch p.pieceValue {
		case pawn:
			return " _P_ "
		case rook:
			return " _R_ "
		case knight:
			return " _N_ "
		case bishop:
			return " _B_ "
		case queen:
			return " _Q_ "
		case king:
			return " _K_ "
		}
	}
	return "  -  "
}

// return the squareIds for legal pawn moves
func (p *piece) pseudoLegalPawnMoves(b *BoardState) []int {
	moves := []int{}
	currentPosition := p.squarePointer.id
	if p.color == white {
		upperSquare := getAdjecentSquare(currentPosition, top)

		// if we are on the 2nd rank
		if isBottomBorder(getAdjecentSquare(currentPosition, bottom)) {
			upperUpperSquare := getAdjecentSquare(upperSquare, top)
			if b.isSquareEmpty(upperSquare) &&
				b.isSquareEmpty(upperUpperSquare) {
				moves = append(moves, upperUpperSquare)
			}
		}

		// check if one step is possible
		if b.isSquareEmpty(upperSquare) {
			moves = append(moves, upperSquare)
		}

		//check for captures
		topLeftSquare := getAdjecentSquare(upperSquare, left)
		topRightSquare := getAdjecentSquare(upperSquare, right)

		if topLeftSquare != -1 {
			if isOpposite, isNonEmpty := b.isSquareOppositeColor(topLeftSquare, p.color); isOpposite && isNonEmpty {
				moves = append(moves, topLeftSquare)
			}
		}

		if topRightSquare != -1 {
			if isOpposite, isNonEmpty := b.isSquareOppositeColor(topRightSquare, p.color); isOpposite && isNonEmpty {
				moves = append(moves, topRightSquare)
			}
		}

		//TODO: Check for en passant
		//TODO: Check for pins
		//TODO: pawn promotion

	} else {
		bottomSquare := getAdjecentSquare(currentPosition, bottom)
		// if we are on the 7nd rank
		if isTopBorder(getAdjecentSquare(currentPosition, top)) {
			bottomBottomSquare := getAdjecentSquare(bottomSquare, bottom)
			if b.isSquareEmpty(bottomSquare) &&
				b.isSquareEmpty(bottomBottomSquare) {
				moves = append(moves, bottomBottomSquare)
			}
		}

		// check if one step is possible
		if b.isSquareEmpty(bottomSquare) {
			moves = append(moves, bottomSquare)
		}

		//check for captures
		bottomLeftSquare := getAdjecentSquare(bottomSquare, left)
		bottomRightSquare := getAdjecentSquare(bottomSquare, right)
		if bottomLeftSquare != -1 {
			if isOpposite, isNonEmpty := b.isSquareOppositeColor(bottomLeftSquare, p.color); isOpposite && isNonEmpty {
				moves = append(moves, bottomLeftSquare)
			}
		}
		if bottomRightSquare != -1 {
			if isOpposite, isNonEmpty := b.isSquareOppositeColor(bottomRightSquare, p.color); isOpposite && isNonEmpty {
				moves = append(moves, bottomRightSquare)
			}
		}
		//TODO: Check for en passant
		//TODO: Check for pins
		//TODO: pawn promotion
	}
	return moves
}

func (p *piece) calculatePseudoLegalMoves(b *BoardState) {
	switch p.pieceValue {
	case pawn:
		p.pseudoLegalMoves = p.pseudoLegalPawnMoves(b)
		//TODO: Add pseudo legal moves for other piece types
	default:
		p.pseudoLegalMoves = make([]int, 0)
	}
}
