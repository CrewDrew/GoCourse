package main

import (
	"errors"
	"fmt"
	"math"
)

// MaxCoordinateValue is max value for coordinates of chess board
const MaxCoordinateValue int = 8

// ChessBoardCell : coordinates of one cell of chess board
type ChessBoardCell struct {
	x, y int
}

// ChessBoard is 2 dimensions array
type ChessBoard [MaxCoordinateValue][MaxCoordinateValue]int

// Init ChessBoard
func (cb *ChessBoard) Init() {
	for i := 0; i < MaxCoordinateValue; i++ {
		for j := 0; j < MaxCoordinateValue; j++ {
			cb[i][j] = j + 1
		}
	}
}

// ChessPiece interface to chess figures
type ChessPiece interface {
	IsAvailablePosition(ChessBoardCell) (bool, error)
}

//Pawn : chess fugure
type Pawn struct {
	position ChessBoardCell
}

// IsAvailablePosition return available cell for figure
func (p Pawn) IsAvailablePosition(newPosition ChessBoardCell) (bool, error) {

	if (p.position == ChessBoardCell{x: 0, y: 0}) {
		return false, errors.New("Невозможно получить доступные ячейки для фигуры вне доски")
	}

	if math.Abs(float64(p.position.y-newPosition.y)) <= 1.0 {
		return true, nil
	}

	return false, errors.New("Выбранная позиция недоступна")
}

// SetPosition set the position to Pawn
func (p *Pawn) SetPosition(newPosition ChessBoardCell) error {
	if newPosition.x > MaxCoordinateValue {
		return errors.New("Координата X выходит за пределы доски")
	}
	if newPosition.y > MaxCoordinateValue {
		return errors.New("Координата Y выходит за пределы доски")
	}

	p.position = newPosition
	return nil
}

// GetPosition return current chess position
func (p Pawn) GetPosition() ChessBoardCell {
	return p.position
}

// Knight : chess fugure
type Knight struct {
	position ChessBoardCell
}

// IsAvailablePosition return available cell for figure
func (p Knight) IsAvailablePosition(newPosition ChessBoardCell) (bool, error) {

	if (p.position == ChessBoardCell{x: 0, y: 0}) {
		return false, errors.New("Невозможно получить доступные ячейки для фигуры вне доски")
	}

	if math.Abs(float64((p.position.x-newPosition.x)*(p.position.y-newPosition.y))) == 2.0 {
		return true, nil
	}

	return false, errors.New("Выбранная позиция недоступна")
}

// SetPosition set the position to Pawn
func (p *Knight) SetPosition(newPosition ChessBoardCell) {
	p.position = newPosition
}

// GetPosition return current chess position
func (p Knight) GetPosition() ChessBoardCell {
	return p.position
}

// AvailablePiecePositions return all available ositions for piece
func AvailablePiecePositions(p ChessPiece, cb ChessBoard) []ChessBoardCell {

	var avalablePositions []ChessBoardCell

	for i := 0; i < len(cb); i++ {
		for j := 0; j < len(cb[i]); j++ {
			result, _ := p.IsAvailablePosition(ChessBoardCell{x: i + 1, y: j + 1})
			if result {
				avalablePositions = append(avalablePositions, ChessBoardCell{x: i + 1, y: j + 1})
			}
		}
	}
	return avalablePositions
}

func main() {

	var chessBoard ChessBoard
	var knightPiece Knight

	chessBoard.Init()

	knightPiece.SetPosition(ChessBoardCell{x: 1, y: 1})
	fmt.Println("Начальная позиция коня ", knightPiece)

	availablePositions := AvailablePiecePositions(knightPiece, chessBoard)
	fmt.Println("Доступные из этого положения ходы: ", availablePositions)

	knightPiece.SetPosition(ChessBoardCell{x: 5, y: 5})
	fmt.Println("Начальная позиция коня ", knightPiece)

	availablePositions = AvailablePiecePositions(knightPiece, chessBoard)
	fmt.Println("Доступные из этого положения ходы: ", availablePositions)
}
