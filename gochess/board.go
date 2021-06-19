// This files has the structures and methods to represent a chess board
package gochess

import "fmt"

//TODO: remove this
func Display() {
	fmt.Printf("From board!")
}

const numberOfSquares int = 64

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

// represents the state of a square on board
type square struct {
	id           int    // used to represent the position of square 0-63
	piecePointer *piece // pointer to a piece lying on it, if any
}

// this structe represents the state of a piece on board
type piece struct {
	Color         colorType // color white or black
	pieceValue    pieceType // what type of piece is it
	squarePointer *square   // pointer to square it lies on, null if out of board
}

// constructor to create new piece
func NewPeiece(color colorType, value pieceType, squarePointer *square) *piece {
	newPiece := new(piece)
	newPiece.Color = color
	newPiece.pieceValue = value
	newPiece.squarePointer = squarePointer
	return newPiece
}

// this structed mantains the state of the chess board at any point in time
type BoardState struct {
	squares     [numberOfSquares]square // represents the 64 squares and their state
	WhitePieces []*piece                // reperesents white pieces on board and their states
	blackPieces []*piece                // reperesents white pieces on board and their states
}

// converts pieve value to string for display
func getStingForPiece(p *piece) string {
	if p == nil {
		return "  -  "
	}
	if p.Color == white {
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

// display the board in console
func (b *BoardState) Display() {
	rowStartIndexs := [...]int{56, 48, 40, 32, 24, 16, 8, 0}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(getStingForPiece(b.squares[rowStartIndexs[i]+j].piecePointer))
		}
		fmt.Println("")
	}
}

// initilize the board state
func (b *BoardState) init() {

	//initilize the IDs
	for i := 0; i < numberOfSquares; i++ {
		b.squares[i].id = i
	}

	//initilize the white pieces. first two rows.
	rooka1 := NewPeiece(white, rook, &b.squares[0])
	b.squares[0].piecePointer = rooka1
	b.WhitePieces = append(b.WhitePieces, rooka1)

	rookh1 := NewPeiece(white, rook, &b.squares[7])
	b.squares[7].piecePointer = rookh1
	b.WhitePieces = append(b.WhitePieces, rookh1)

	knightb1 := NewPeiece(white, knight, &b.squares[1])
	b.squares[1].piecePointer = knightb1
	b.WhitePieces = append(b.WhitePieces, knightb1)

	knightg1 := NewPeiece(white, knight, &b.squares[6])
	b.squares[6].piecePointer = knightg1
	b.WhitePieces = append(b.WhitePieces, knightg1)

	bishopc1 := NewPeiece(white, bishop, &b.squares[2])
	b.squares[2].piecePointer = bishopc1
	b.WhitePieces = append(b.WhitePieces, bishopc1)

	bishopf1 := NewPeiece(white, bishop, &b.squares[5])
	b.squares[5].piecePointer = bishopf1
	b.WhitePieces = append(b.WhitePieces, bishopf1)

	queend1 := NewPeiece(white, queen, &b.squares[3])
	b.squares[3].piecePointer = queend1
	b.WhitePieces = append(b.WhitePieces, queend1)

	kinge1 := NewPeiece(white, king, &b.squares[4])
	b.squares[4].piecePointer = kinge1
	b.WhitePieces = append(b.WhitePieces, kinge1)

	for i := 8; i < 16; i++ {
		pawn := NewPeiece(white, pawn, &b.squares[i])
		b.squares[i].piecePointer = pawn
		b.blackPieces = append(b.WhitePieces, pawn)
	}

	//initilize the black pieces. last two rows.
	rooka8 := NewPeiece(black, rook, &b.squares[56])
	b.squares[56].piecePointer = rooka8
	b.blackPieces = append(b.blackPieces, rooka8)

	rookh8 := NewPeiece(black, rook, &b.squares[63])
	b.squares[63].piecePointer = rookh8
	b.blackPieces = append(b.blackPieces, rookh8)

	knightb8 := NewPeiece(black, knight, &b.squares[57])
	b.squares[57].piecePointer = knightb8
	b.blackPieces = append(b.blackPieces, knightb8)

	knightg8 := NewPeiece(black, knight, &b.squares[62])
	b.squares[62].piecePointer = knightg8
	b.blackPieces = append(b.blackPieces, knightg8)

	bishopc8 := NewPeiece(black, bishop, &b.squares[58])
	b.squares[58].piecePointer = bishopc8
	b.blackPieces = append(b.blackPieces, bishopc8)

	bishopf8 := NewPeiece(black, bishop, &b.squares[61])
	b.squares[61].piecePointer = bishopf8
	b.blackPieces = append(b.blackPieces, bishopf8)

	queend8 := NewPeiece(black, queen, &b.squares[59])
	b.squares[59].piecePointer = queend8
	b.blackPieces = append(b.blackPieces, queend8)

	kinge8 := NewPeiece(black, king, &b.squares[60])
	b.squares[60].piecePointer = kinge8
	b.blackPieces = append(b.blackPieces, kinge8)

	for i := 48; i < 56; i++ {
		pawn := NewPeiece(black, pawn, &b.squares[i])
		b.squares[i].piecePointer = pawn
		b.blackPieces = append(b.blackPieces, pawn)
	}
}

func NewBoardState() *BoardState {
	board := new(BoardState)
	board.init()
	return board
}
