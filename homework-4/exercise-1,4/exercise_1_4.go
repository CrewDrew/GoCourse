package main

import "errors"

// ChessBoardCell : coordinates of one cell of chess board
type ChessBoardCell struct {
	x, y int
}

// ChessPiece interface to chess figures
type ChessPiece interface {
	SetPosition(ChessBoardCell)
	AvailablePositions() ([]ChessBoardCell, error)
}

//Pawn : chess fugure
type Pawn struct {
	position ChessBoardCell
}

// AvailablePositions return available cell for figure
func (p Pawn) AvailablePositions() ([]ChessBoardCell, error) {
	var availablePositions []ChessBoardCell

	if (p.position == ChessBoardCell{x: 0, y: 0}) {
		errors.New("Невозможно получить доступные ячейки для фигуры вне доски")
	}

	if p.position.y+1 <= 8 {
		availablePositions = append(availablePositions, ChessBoardCell{x: p.position.x, y: p.position.y + 1})
	}

	return availablePositions, nil
}

// SetPosition set the position to Pawn
func (p *Pawn) SetPosition(newPosition ChessBoardCell) error {
	if newPosition.x > 8 {
		return errors.New("Координата X выходит за пределы доски")
	}
	if newPosition.y > 8 {
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

// AvailablePositions return available cell for figure
func (p Knight) AvailablePositions() ([]ChessBoardCell, error) {
	var availablePositions []ChessBoardCell

	if (p.position == ChessBoardCell{x: 0, y: 0}) {
		errors.New("Невозможно получить доступные ячейки для фигуры вне доски")
	}

	if p.position.y+1 <= 8 {
		availablePositions = append(availablePositions, ChessBoardCell{x: p.position.x, y: p.position.y + 1})
	}

	return availablePositions, nil
}

// SetPosition set the position to Pawn
func (p *Knight) setPosition(newPosition ChessBoardCell) {
	p.position = newPosition
}

// GetPosition return current chess position
func (p Knight) GetPosition() ChessBoardCell {
	return p.position
}

// SetPiecePosition set the position to Pawn
func SetPiecePosition(p ChessPiece, newPosition ChessBoardCell) error {
	availablePositions, err := p.AvailablePositions()

	if err != nil {
		return errors.New("Не удалось определить доступные ячейки")
	}

	if availablePositions != nil {
		for _, position := range availablePositions {
			if position == newPosition {
				p.SetPosition(newPosition)
				return nil
			}
		}
		return errors.New("Нельзя поместить фигуру в данную ячейку")
	}
	return errors.New("Нет доступных переходов")
}

func main() {

}
