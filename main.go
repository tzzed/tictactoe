package main

import (
	"fmt"
	"tictac/game"
)

func playerInput(s [][]int, input int) {
	i := 1
	for row := range s {
		for col := range s[row] {
			if i == input {
				s[row][col] = 1
			}
			i++
		}
	}
}

/*
Initialize Matrix for each player
return a slice 3x3
*/
func CreateTable() [][]int {
	return [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
}

func main() {

	fmt.Println("  Tictactoe Game  ")

	table := CreateTable()
	p := game.NewPlayer(1, CreateTable(), false, 0, 'X')
	p2 := game.NewPlayer(2, CreateTable(), false, 0, 'O')
	var input int
	p.PrintGrid(p2)
	for !p.Isfull && !p2.Isfull {
		input = p.KeyboardInput()
		for !p.SetCase(input, p, p2) {
			input = p2.KeyboardInput()
		}
		playerInput(table, input)
		if p.Turn(input) {
			fmt.Printf("Player n°%d Wins\n", p.ID)
			return
		}
		p.PrintGrid(p2)
		input = p2.KeyboardInput()
		for !p2.SetCase(input, p, p2) {
			input = p2.KeyboardInput()
		}
		playerInput(table, input)
		p.PrintGrid(p2)
		if p2.Turn(input) {
			fmt.Printf("Player n°%d Wins\n", p2.ID)
			return
		}

	}

}
