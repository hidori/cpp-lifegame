package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/hidori/lifegame/golang/lifegame"
)

func main() {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	const (
		width  = 8
		height = 8
	)

	cells := generateCells(width, height)
	game := lifegame.New(width, height, cells)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printCells(game)
		game.Yield()
		scanner.Scan()
	}
}

func generateCells(width int, height int) []bool {
	cells := make([]bool, width*height)

	for i := 0; i < width*height; i++ {
		cells[i] = rand.Intn(2) == 1
	}

	return cells
}

func printCells(g *lifegame.Lifegame) {
	for i := 0; i < g.Width*g.Height; i++ {
		if g.Cells[i] {
			fmt.Print("@")
		} else {
			fmt.Print(".")
		}

		if i%g.Width == (g.Width - 1) {
			fmt.Println()
		}
	}

	fmt.Println()
}
