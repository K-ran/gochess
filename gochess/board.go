// This files has the structures and methods to represent a chess board
package gochess

import "fmt"

const numberOfSquares int = 64

// this structed mantains the state of the chess board at any point in time
type BoardState struct {
	squares     [numberOfSquares]square // represents the 64 squares and their state
	whitePieces []*piece                // reperesents white pieces on board and their states
	blackPieces []*piece                // reperesents white pieces on board and their states
}

// display the board in console
func (b *BoardState) Display() {
	rowStartIndexs := [...]int{56, 48, 40, 32, 24, 16, 8, 0}
	fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(getStingForPiece(b.squares[rowStartIndexs[i]+j].piecePointer))
		}
		fmt.Println("")
	}
	fmt.Print("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
}

// initilize the board state with a new game position
func (b *BoardState) init() {

	//initilize the IDs
	for i := 0; i < numberOfSquares; i++ {
		b.squares[i].id = i
	}

	//initilize the white pieces. first two rows.
	rooka1 := NewPeiece(white, rook, &b.squares[0])
	b.squares[0].piecePointer = rooka1
	b.whitePieces = append(b.whitePieces, rooka1)

	rookh1 := NewPeiece(white, rook, &b.squares[7])
	b.squares[7].piecePointer = rookh1
	b.whitePieces = append(b.whitePieces, rookh1)

	knightb1 := NewPeiece(white, knight, &b.squares[1])
	b.squares[1].piecePointer = knightb1
	b.whitePieces = append(b.whitePieces, knightb1)

	knightg1 := NewPeiece(white, knight, &b.squares[6])
	b.squares[6].piecePointer = knightg1
	b.whitePieces = append(b.whitePieces, knightg1)

	bishopc1 := NewPeiece(white, bishop, &b.squares[2])
	b.squares[2].piecePointer = bishopc1
	b.whitePieces = append(b.whitePieces, bishopc1)

	bishopf1 := NewPeiece(white, bishop, &b.squares[5])
	b.squares[5].piecePointer = bishopf1
	b.whitePieces = append(b.whitePieces, bishopf1)

	queend1 := NewPeiece(white, queen, &b.squares[3])
	b.squares[3].piecePointer = queend1
	b.whitePieces = append(b.whitePieces, queend1)

	kinge1 := NewPeiece(white, king, &b.squares[4])
	b.squares[4].piecePointer = kinge1
	b.whitePieces = append(b.whitePieces, kinge1)

	for i := 8; i < 16; i++ {
		pawn := NewPeiece(white, pawn, &b.squares[i])
		b.squares[i].piecePointer = pawn
		b.blackPieces = append(b.whitePieces, pawn)
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

// return a deep copy of self
func (b *BoardState) GetCopyOfState() *BoardState {
	newBoardCopy := *b
	newBoardCopy.whitePieces = make([]*piece, 0)
	newBoardCopy.blackPieces = make([]*piece, 0)

	//TODO: can this be parallalized?
	for i := 0; i < 64; i++ {
		oldPiece := b.squares[i].piecePointer
		if oldPiece != nil {
			newPiece := NewPeiece(oldPiece.color, oldPiece.pieceValue, &newBoardCopy.squares[i])
			newBoardCopy.squares[i].piecePointer = newPiece
			if newPiece.color == white {
				newBoardCopy.whitePieces = append(newBoardCopy.whitePieces, newPiece)
			} else {
				newBoardCopy.blackPieces = append(newBoardCopy.blackPieces, newPiece)
			}
		} else {
			newBoardCopy.squares[i].piecePointer = nil
		}
	}
	return &newBoardCopy
}

// returns a new board state with the move made
func (b *BoardState) MakeNextMove(initSquate int, finalSquare int) *BoardState {
	newState := b.GetCopyOfState()
	//TODO: make a move here in the new state
	return newState
}

// returns the next board states with nextPlay colors turn
func (b *BoardState) GetNextStates(nextPlay colorType) []BoardState {
	nextStates := []BoardState{}

	if nextPlay == white {
		//iterate over all white pieces wp // this could trigger a go routine
		//for each wp, iterate over all legal moves // this could further trigger a go routine
	} else {

	}
	return nextStates
}

func NewBoardState() *BoardState {
	board := new(BoardState)
	board.init()
	return board
}
