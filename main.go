package main

import (
	"fmt"

	mathgameuno "github.com/masiuciszek/Gophers/math-game-uno"
)


func main(){
	game:=mathgameuno.Game("problems.csv")
	fmt.Println(game)
}