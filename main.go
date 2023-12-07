package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Board struct {
	Fields [3][3]string
}

const emptyField string = "\u00A0\u00A0"

func NewBoard() *Board{
	return &Board{
		Fields: [3][3]string {
			{emptyField, emptyField, emptyField},
			{emptyField, emptyField, emptyField},
			{emptyField, emptyField, emptyField},
		},
	}
}

func getInput(x *int) {
	for {
		fmt.Print("Select a number(1-3):")
		fmt.Scan(x)
		if *x >= 1 && *x <= 3 {
			break
		}

		fmt.Println("Invalid input try again.\n")
	}
	
}

func (board *Board) tryToMakeMove(row, collum int, character string) bool{
	row --
	collum --
	if board.Fields[row][collum] == emptyField {
		board.Fields[row][collum] = character 
		return true
	}

	return false
}

func clearCMD() {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "cls") 
    } else {
        cmd = exec.Command("clear")
    }

    cmd.Stdout = os.Stdout
    cmd.Run()
}

func main() {
	for{
		fmt.Println("------------MENU------------")
		fmt.Println("Press any key to quit, or p to play:")
		str := ""
		fmt.Scan(&str)
		if str != "p" {
			break
		}

		runGame()
	}
	
}

func runGame() {
	clearCMD()
	b := NewBoard()
	player := 0

	character := "❌"

	b.PrintBoard()

	for !b.Won() && !b.Full() {
		
		if player == 1 {
			character = "⭕️"
		} else {
			character = "❌"
		}

		var row, collum int

		fmt.Println("\nWaiting for " + character + "'s move...")

		fmt.Println("\nRow")
		getInput(&row)

		fmt.Println("\nCollum")
		getInput(&collum)

		if b.tryToMakeMove(row, collum, character) {

			if player == 1 {
				player = 0
			} else {
				player = 1
			}
			clearCMD()
			b.PrintBoard()

			continue
		}

		fmt.Println("\nField already full try again")
	}

	if b.Won() {
		fmt.Println("\n🎉" + character + " has won the game 🎉\n")
	} else {
		fmt.Println("\nThe game was a tie \n")
	}
}

func (board *Board) Won() bool {
	for i := 0; i < 3; i++ {
		rowWin := board.Fields[i][0] != emptyField && board.Fields[i][0] == board.Fields[i][1] && board.Fields[i][1] == board.Fields[i][2]
		collumWin := board.Fields[0][i] != emptyField && board.Fields[0][i] == board.Fields[1][i] && board.Fields[1][i] == board.Fields[2][i]

		if rowWin || collumWin {
			return true
		}
	}

	diagonalWin1 := board.Fields[0][0] == board.Fields[1][1] && board.Fields[1][1] == board.Fields[2][2]
	diagonalWin2 := board.Fields[0][2] == board.Fields[1][1] && board.Fields[1][1] == board.Fields[2][0]

	if board.Fields[1][1] != emptyField && (diagonalWin1 || diagonalWin2) {
		return true
	}

	return false
}

func (board *Board) Full() bool {
	for _, row := range board.Fields {
		for _, element := range row {
			if element == emptyField {
				return false
			}
		}
	}

	return true
}

func (board *Board) PrintBoard() {
	rowSeperator := "\n"//"\n--+---+--\n"
	collumSeperator := "" //" | "
	
	res := [3]string{} 

	for i, row := range board.Fields {
		res[i] = strings.Join(row[:], collumSeperator)
	}

	fmt.Println(strings.Join(res[:], rowSeperator))
}
