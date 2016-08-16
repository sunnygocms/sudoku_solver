package main

import (
	"fmt"

	"github.com/sunnygocms/sudoku_solver/sudoku"
)

func main() {
	//	g := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
	//	values := sudoku.Solve(g)
	//	fmt.Println("resolution:", values)
	g := "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
	values := sudoku.Solve(g)
	fmt.Println("resolution:")
	sudoku.Display(values)
}
