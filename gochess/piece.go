package gochess

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

//enum to represent color of piece
type colorType bool

const (
	white colorType = true
	black colorType = false
)

// this structe represents the state of a piece on board
type piece struct {
	color         colorType // color white or black
	pieceValue    pieceType // what type of piece is it
	squarePointer *square   // pointer to square it lies on, null if out of board
}

// constructor to create new piece
func NewPeiece(color colorType, value pieceType, squarePointer *square) *piece {
	newPiece := new(piece)
	newPiece.color = color
	newPiece.pieceValue = value
	newPiece.squarePointer = squarePointer
	return newPiece
}

// converts pieve value to string for display
func getStingForPiece(p *piece) string {
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
