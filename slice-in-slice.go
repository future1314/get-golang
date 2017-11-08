package main

import(
	"fmt"
	"strings"
)

func showBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

func main() {
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"}}

	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[2][1] = "O"
	game[1][1] = "X"

	showBoard(game)
}
