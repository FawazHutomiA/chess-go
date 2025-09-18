package main

import "testing"

func TestBoardInit(t *testing.T) {
	g := NewGame()
	if g.Board.Get(Pos{R: 1, C: 0}).Type != Pawn {
		t.Fatal("expected black pawn at 1,0")
	}
	if g.Board.Get(Pos{R: 6, C: 0}).Type != Pawn {
		t.Fatal("expected white pawn at 6,0")
	}
	if g.Board.Get(Pos{R: 0, C: 4}).Type != King {
		t.Fatal("expected black king")
	}
	if g.Board.Get(Pos{R: 7, C: 4}).Type != King {
		t.Fatal("expected white king")
	}
}

func TestPawnMove(t *testing.T) {
	g := NewGame()
	// langkah 1 kotak ke depan
	ok, _ := g.Board.CanMove(Pos{R: 6, C: 0}, Pos{R: 5, C: 0})
	if !ok {
		t.Fatal("pawn should move forward 1 square")
	}
	// langkah 2 kotak dari posisi awal
	ok, _ = g.Board.CanMove(Pos{R: 6, C: 0}, Pos{R: 4, C: 0})
	if !ok {
		t.Fatal("pawn should move 2 squares from start row")
	}
}

func TestKnightMove(t *testing.T) {
	g := NewGame()
	ok, _ := g.Board.CanMove(Pos{R: 0, C: 1}, Pos{R: 2, C: 2})
	if !ok {
		t.Fatal("knight should move")
	}
}

func TestIllegalMoves(t *testing.T) {
	g := NewGame()
	ok, _ := g.Board.CanMove(Pos{R: 7, C: 0}, Pos{R: 7, C: 1})
	if ok {
		t.Fatal("should not capture own piece")
	}
}

func TestKingCaptureEndsGame(t *testing.T) {
	g := NewGame()
	g.Board.Set(Pos{R: 4, C: 4}, &Piece{Type: King, Color: White})
	g.Board.Set(Pos{R: 4, C: 0}, &Piece{Type: Rook, Color: Black})
	ok, _ := g.Board.CanMove(Pos{R: 4, C: 0}, Pos{R: 4, C: 4})
	if !ok {
		t.Fatal("rook should be able to capture king")
	}
	capt := g.Board.Move(Pos{R: 4, C: 0}, Pos{R: 4, C: 4})
	if capt == nil || capt.Type != King {
		t.Fatal("king must be captured")
	}
}
