package main

import (
	"fmt"

	mathgamedues "github.com/masiuciszek/Gophers/math-game-dues"
)


func main(){
	x := mathgamedues.Game("problems.csv")
	fmt.Println(x)
}