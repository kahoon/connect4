package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	valid "github.com/asaskevich/govalidator"
)

const (
	CONNECT = 4
	X       = 6
	Y       = 7
	B       = 2
	W       = 1
)

var (
	TEXT    = [3]string{"Undefined", "WHITE", "BLACK"}
	SYMBOL  = [3]string{"O ", "W ", "B "}
	game    = [X][Y]int{{}}
	current = W
	command = bufio.NewScanner(os.Stdin)
)

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func display() {
	clear()
	for x := X - 1; x >= 0; x-- {
		for y := 0; y < Y; y++ {
			fmt.Print(SYMBOL[game[x][y]])
		}
		fmt.Print("\n")
	}
	for y := 0; y < Y; y++ {
		fmt.Printf("%d ", y+1)
	}
	fmt.Print("\n\n")
}

func read() string {
	command.Scan()
	return command.Text()
}

func turn() {
	if current == W {
		current = B
	} else {
		current = W
	}
}

func input() int {
	for {
		fmt.Printf("Enter column number for %s: ", TEXT[current])
		column := read()
		if valid.IsInt(column) {
			if value, err := strconv.Atoi(column); err == nil {
				if value > 0 && value <= Y {
					return value
				}
			}
		}
	}
}

func update(column int) bool {
	for x := 0; x < X; x++ {
		if game[x][column-1] == 0 {
			game[x][column-1] = current
			return true
		}
	}
	return false
}

func win() {
	display()
	fmt.Printf("%s Wins!!!\n\n", TEXT[current])
	os.Exit(0)
}

func analyze() {
	var match int
	var color int
	// right
	for x := 0; x < X; x++ {
		for y := 0; y <= Y-CONNECT; y++ {
			match = 0
			color = game[x][y]
			if color == 0 {
				continue
			}
			for z := y; z < Y; z++ {
				if game[x][z] == color {
					match += 1
				} else {
					break
				}
			}
			if match >= CONNECT {
				win()
			}
		}
	}
	// up
	for y := 0; y < Y; y++ {
		for x := 0; x <= X-CONNECT; x++ {
			match = 0
			color = game[x][y]
			if color == 0 {
				continue
			}
			for z := x; z < X; z++ {
				if game[z][y] == color {
					match += 1
				} else {
					break
				}
			}
			if match >= CONNECT {
				win()
			}
		}
	}
	// diagonal right
	for x := 0; x <= X-CONNECT; x++ {
		for y := 0; y <= Y-CONNECT; y++ {
			match = 0
			color = game[x][y]
			if color == 0 {
				continue
			}
			for i, j := x, y; i < X && j < Y; i, j = i+1, j+1 {
				if game[i][j] == color {
					match += 1
				} else {
					break
				}
			}
			if match >= CONNECT {
				win()
			}
		}
	}
	// diagonal left
	for x := 0; x <= X-CONNECT; x++ {
		for y := CONNECT - 1; y >= 0; y-- {
			match = 0
			color = game[x][y]
			if color == 0 {
				continue
			}
			for i, j := x, y; i < X && j >= 0; i, j = i+1, j-1 {
				if game[i][j] == color {
					match += 1
				} else {
					break
				}
			}
			if match >= CONNECT {
				win()
			}
		}
	}
}

func main() {
	for {
		display()
		column := input()
		if update(column) {
			analyze()
			turn()
		}
	}
}
