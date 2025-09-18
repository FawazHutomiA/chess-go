package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	g := NewGame()
	reader := bufio.NewReader(os.Stdin)
	current := White
	for {
		ClearScreen()
		fmt.Printf("Current player: %s\n", colorName(current))
		g.Board.Print()
		if g.GameOver {
			fmt.Printf("Game over! Winner: %s\n", colorName(g.Winner))
			return
		}
		fmt.Print("Enter move (e.g. b2 b3  or 1,2 2,2): ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "quit" || line == "exit" {
			fmt.Println("Exiting...")
			return
		}
		from, to, err := ParseMove(line)
		if err != nil {
			fmt.Println("Invalid input:", err)
			Pause()
			continue
		}
		p := g.Board.Get(from)
		if p == nil {
			fmt.Println("No piece at source")
			Pause()
			continue
		}
		if p.Color != current {
			fmt.Println("Not your piece")
			Pause()
			continue
		}
		ok, reason := g.Board.CanMove(from, to)
		if !ok {
			fmt.Println("Illegal move:", reason)
			Pause()
			continue
		}
		captured := g.Board.Move(from, to)
		if captured != nil && captured.Type == King {
			g.GameOver = true
			g.Winner = current
		}
		current = opposite(current)
	}
}

func colorName(c Color) string {
	if c == White {
		return "White"
	}
	return "Black"
}

func ClearScreen() {
	// crude cross-platform clear
	fmt.Print("\033[H\033[2J")
}

func Pause() {
	fmt.Print("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
