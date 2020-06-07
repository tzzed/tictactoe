package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Player struct with field of a Player
type Player struct {
	ID     int
	Mat    [][]int
	Isfull bool
	Wins   int
	Op     byte
}

// NewPlayer is the constructor
func NewPlayer(id int, mat [][]int, isfull bool, wins int, op byte) *Player {
	return &Player{ID: id, Mat: mat, Isfull: isfull, Wins: wins, Op: op}
}

// FindByCol check columun if there are 3 entrie for 1 player
func (p *Player) FindByCol() {
	check := 0
	for col := 0; col < 3; col++ {
		for row := 0; row < 3; row++ {
			if p.Mat[row][col] == 1 {
				if check++; check == 3 {
					p.Isfull = true
					p.Wins++
				}
			} else {
				check = 0
			}
		}
	}
}

// FindInDiag check diags if there are 3 entrie for 1 player
func (p *Player) FindInDiag() {
	if p.Mat[1][1] != 1 {
		return
	}
	check := 0
	for row, col := 0, 0; row < 3 && col < 3; row, col = col+1, row+1 {
		if p.Mat[row][col] == 1 {
			if check++; check == 3 {
				p.Isfull = true
				p.Wins++
			}
		}
	}

	//reset the check
	check = 0
	for row, col := 2, 2; row > -1 && col > -1; row, col = col-1, row-1 {
		if p.Mat[row][col] == 1 {
			if check++; check == 3 {
				p.Isfull = true
				p.Wins++
			}
		}

	}
}

// FindByRow check row by row if there are 3 entrie for 1 player
func (p *Player) FindByRow() {
	check := 0
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if p.Mat[row][col] == 1 {
				if check++; check == 3 {
					p.Isfull = true
					p.Wins++
				}
			} else {
				check = 0
			}
		}

	}
}

// SetCase the case which is selected by the player
func (p *Player) SetCase(c int, p1 *Player, p2 *Player) bool {
	sec := 0
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if sec++; sec == c {
				if p1.Mat[row][col] == 0 && p2.Mat[row][col] == 0 {
					p.Mat[row][col] = 1
					return true
				} else {
					fmt.Printf("Please select another Case n°%d is full\n", c)
					return false
				}

			}
		}
	}
	return false
}

// PrintGrid show the player the main map
func (p *Player) PrintGrid(p2 *Player) {
	i := 1
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if p.Mat[row][col] == 1 {
				fmt.Printf("  %c |", p.Op)
			}
			if p2.Mat[row][col] == 1 {
				fmt.Printf("  %c |", p2.Op)

			}
			if p2.Mat[row][col] == 0 && p.Mat[row][col] == 0 {
				fmt.Printf("  %d |", i)
			}
			i++
		}
		fmt.Println("\n------------------")
	}
}

// Turn select the turn players
func (p *Player) Turn(input int) bool {
	p.FindByRow()
	p.FindByCol()
	p.FindInDiag()
	return p.Isfull

}

// KeyboardInput the selected case by the playser
func (p *Player) KeyboardInput() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Player n°%d Enter your case: ", p.ID)
	number, _ := reader.ReadString('\n')
	input := strings.TrimSuffix(number, "\n")
	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("please enter a number")
	}
	return i
}

func init() {

}
