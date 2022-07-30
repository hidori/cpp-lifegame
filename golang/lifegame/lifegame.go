package lifegame

type Lifegame struct {
	Width  int
	Height int
	Cells  []bool
}

func New(width int, height int, cells []bool) *Lifegame {
	return &Lifegame{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}

func (g *Lifegame) Yield() {
	cells := make([]bool, g.Width*g.Height)

	for i := 0; i < g.Width*g.Height; i++ {
		neighbors := g.CountNeighbors(i)

		if g.Cells[i] {
			if 2 == neighbors || neighbors == 3 {
				cells[i] = true
			}
		} else {
			if neighbors == 3 {
				cells[i] = true
			}
		}
	}

	g.Cells = cells
}

func (g *Lifegame) CountNeighbors(i int) int {
	neighbors := 0
	x := i % g.Width
	y := i / g.Height

	for ny := y - 1; ny <= y+1; ny++ {
		if ny < 0 || g.Height <= ny {
			continue
		}

		for nx := x - 1; nx <= x+1; nx++ {
			if (nx < 0 || g.Width <= nx) || (nx == x && ny == y) {
				continue
			}

			ni := ny*g.Width + nx

			if g.Cells[ni] {
				neighbors++
			}
		}
	}

	return neighbors
}
