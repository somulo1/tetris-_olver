package main

import (
	"fmt"
	"os"
	"time"

	"tetris-optimizer/helperfunctions"
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("error: please provide the path to the tetromino file")
		return
	}
	fileName := os.Args[1]
	tet, err := helperfunctions.InputFileReader(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	Trimtetros, err := helperfunctions.TrimTetrominos(tet)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	solvedTetrominos, err := helperfunctions.SolveTetris(Trimtetros)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, t := range solvedTetrominos.Tet {
		for _, s := range t {
			fmt.Printf("%s", s)
		}
		fmt.Println()
	}
	elapsed := time.Since(start)
	fmt.Println()
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
