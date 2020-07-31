package main

type chessBoardCell struct {
	x, y int
}

type ChessPiece interface {
	MoveTo(int, int) bool
	AvailableCells() []chessBoardCell
}

type Pawn struct {
	position chessBoardCell
}

func (p *Pawn) MoveTo(x int, y int) {
	p.position.x = x
	p.position.y = y
}

func (p Pawn) AvailableCells() []chessBoardCell {

}

func (p *Pawn) SetPosition(x int, y int) {
	p.x = x
	p.y = y
}

func (p Pawn) GetPosition() {
	return p
}

func main() {

}
