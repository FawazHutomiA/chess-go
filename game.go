package main

import (
	"fmt"
	"unicode"
)

// Color
type Color int

const (
	White Color = iota
	Black
)

// Piece types
const (
	Pawn   = 'P'
	Rook   = 'R'
	Knight = 'N'
	Bishop = 'B'
	Queen  = 'Q'
	King   = 'K'
)

// Piece
type Piece struct {
	Type  rune
	Color Color
}

func (p *Piece) String() string {
	if p == nil {
		return "."
	}
	if p.Color == White {
		return string(p.Type)
	}
	return string(unicode.ToLower(p.Type))
}

// Board
type Board struct {
	Cells [8][8]*Piece
}

// Position
type Pos struct{ R, C int }

// Game
type Game struct {
	Board    *Board
	GameOver bool
	Winner   Color
}

// Init game
func NewGame() *Game {
	b := &Board{}
	setupBackRank := []rune{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}
	for c, t := range setupBackRank {
		b.Cells[0][c] = &Piece{Type: t, Color: Black}
		b.Cells[1][c] = &Piece{Type: Pawn, Color: Black}
		b.Cells[6][c] = &Piece{Type: Pawn, Color: White}
		b.Cells[7][c] = &Piece{Type: t, Color: White}
	}
	return &Game{Board: b}
}

// Print board
func (b *Board) Print() {
	fmt.Println("    a  b  c  d  e  f  g  h")
	for r := 0; r < 8; r++ {
		fmt.Printf("%d ", 8-r)
		for c := 0; c < 8; c++ {
			p := b.Cells[r][c]

			// background selang-seling
			if (r+c)%2 == 0 {
				fmt.Print("\033[47m") // putih
			} else {
				fmt.Print("\033[40m") // hitam
			}

			symbol := " . "
			if p != nil {
				if p.Color == White {
					symbol = fmt.Sprintf(" \033[1;34m%c\033[0m ", p.Type) // biru terang
				} else {
					symbol = fmt.Sprintf(" \033[1;31m%c\033[0m ", p.Type) // merah terang
				}
			}

			fmt.Print(symbol)
			fmt.Print("\033[0m") // reset background
		}
		fmt.Printf(" %d\n", 8-r)
	}
	fmt.Println("    a  b  c  d  e  f  g  h")
}

func (b *Board) Get(pos Pos) *Piece {
	if !inBounds(pos) {
		return nil
	}
	return b.Cells[pos.R][pos.C]
}

func (b *Board) Set(pos Pos, p *Piece) {
	if inBounds(pos) {
		b.Cells[pos.R][pos.C] = p
	}
}

func inBounds(p Pos) bool { return p.R >= 0 && p.R < 8 && p.C >= 0 && p.C < 8 }

func (b *Board) Move(from, to Pos) *Piece {
	if !inBounds(from) || !inBounds(to) {
		return nil
	}
	capt := b.Get(to)
	b.Set(to, b.Get(from))
	b.Set(from, nil)
	return capt
}

// --- Move Rules ---
func (b *Board) CanMove(from, to Pos) (bool, string) {
	if !inBounds(from) || !inBounds(to) {
		return false, "out of bounds"
	}
	p := b.Get(from)
	if p == nil {
		return false, "no piece at source"
	}
	if from.R == to.R && from.C == to.C {
		return false, "same square"
	}
	target := b.Get(to)
	if target != nil && target.Color == p.Color {
		return false, "cannot capture own piece"
	}

	switch p.Type {
	case Pawn:
		return pawnCanMove(b, from, to, p.Color)
	case Rook:
		return straightCanMove(b, from, to)
	case Bishop:
		return diagCanMove(b, from, to)
	case Queen:
		if ok, _ := straightCanMove(b, from, to); ok {
			return true, ""
		}
		return diagCanMove(b, from, to)
	case Knight:
		return knightCanMove(from, to)
	case King:
		return kingCanMove(from, to)
	default:
		return false, "unknown piece"
	}
}

// Pawn
func pawnCanMove(b *Board, from, to Pos, col Color) (bool, string) {
	dir := -1
	startRow := 6
	if col == Black {
		dir = 1
		startRow = 1
	}
	if to.C == from.C {
		if to.R == from.R+dir && b.Get(to) == nil {
			return true, ""
		}
		if from.R == startRow && to.R == from.R+2*dir {
			mid := Pos{R: from.R + dir, C: from.C}
			if b.Get(mid) == nil && b.Get(to) == nil {
				return true, ""
			}
		}
		return false, "blocked or invalid pawn forward move"
	}
	if (to.C == from.C+1 || to.C == from.C-1) && to.R == from.R+dir {
		t := b.Get(to)
		if t != nil && t.Color != col {
			return true, ""
		}
		return false, "no piece to capture diagonally"
	}
	return false, "invalid pawn move"
}

// Rook & Queen straight
func straightCanMove(b *Board, from, to Pos) (bool, string) {
	if from.R != to.R && from.C != to.C {
		return false, "not straight"
	}
	rs := sign(to.R - from.R)
	cs := sign(to.C - from.C)
	cur := Pos{R: from.R + rs, C: from.C + cs}
	for !(cur.R == to.R && cur.C == to.C) {
		if b.Get(cur) != nil {
			return false, "path blocked"
		}
		cur.R += rs
		cur.C += cs
	}
	return true, ""
}

// Bishop & Queen diagonal
func diagCanMove(b *Board, from, to Pos) (bool, string) {
	dr := to.R - from.R
	dc := to.C - from.C
	if abs(dr) != abs(dc) {
		return false, "not diagonal"
	}
	rs := sign(dr)
	cs := sign(dc)
	cur := Pos{R: from.R + rs, C: from.C + cs}
	for !(cur.R == to.R && cur.C == to.C) {
		if b.Get(cur) != nil {
			return false, "path blocked"
		}
		cur.R += rs
		cur.C += cs
	}
	return true, ""
}

// Knight
func knightCanMove(from, to Pos) (bool, string) {
	dr := abs(to.R - from.R)
	dc := abs(to.C - from.C)
	if (dr == 2 && dc == 1) || (dr == 1 && dc == 2) {
		return true, ""
	}
	return false, "invalid knight move"
}

// King
func kingCanMove(from, to Pos) (bool, string) {
	dr := abs(to.R - from.R)
	dc := abs(to.C - from.C)
	if dr <= 1 && dc <= 1 {
		return true, ""
	}
	return false, "king moves one square"
}

// Helpers
func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
