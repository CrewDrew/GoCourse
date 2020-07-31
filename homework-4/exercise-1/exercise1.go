package main

// ChessBoardCell : coordinates of one cell of chess board
type ChessBoardCell struct {
	x, y int
}

// ChessPiece interface to chess figures
type ChessPiece interface {
	SetPosition(int, int)
	AvailablePositions() []ChessBoardCell
}

//Pawn : chess fugure
type Pawn struct {
	position ChessBoardCell
}

// AvailablePositions return available cell for figure
func (p Pawn) AvailablePositions() []ChessBoardCell {
	var availablePositions []ChessBoardCell

	if p.position.y+1 < 8 {
		availablePositions = append(availablePositions, ChessBoardCell{x: p.position.x, y: p.position.y + 1})
	}

	return availablePositions
}

// SetPosition set the position to Pawn
func (p *Pawn) SetPosition(x int, y int) {
	p.position.x = x
	p.position.y = y
}

// GetPosition return current chess position
func (p Pawn) GetPosition() ChessBoardCell {
	return p.position
}

func main() {

}
