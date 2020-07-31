package main

type chessBoardCell{
	X, Y int
}

type ChessPiece interface {
	MoveTo(int, int) bool
	AvailableCells() []chessBoardCell 
}

type Pawn{
	x, y int
}

func (p *Pawn) MoveTo(x int, y int){
	p.x = x
	p.y = y
}

func (p Pawn) AvailableCells() []chessBoardCell {
	
}

func main() {

}