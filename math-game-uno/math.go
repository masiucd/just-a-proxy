package mathgameuno

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// MathQuizStruct for math quiz
type MathQuizStruct struct {
	Q string
	A string
}

// Game Math game
func Game(file string) string {
	csvFile, err := os.Open(file)

	if err != nil {
		log.Fatalf("could not read csv file")
	}

	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()

	problems := printLines(lines)

	counter := 0

	for i, p := range problems {
		fmt.Printf("#%d %s = \n", i, p.Q)
		var input string
		fmt.Scanf("%s", &input)

		if input == p.A {
			counter++
			fmt.Println("correct \n")
		} else {
			fmt.Println("wrong \n")
		}
	}
 r:= 	fmt.Sprintf("you got %d out of %d possible \n", counter, len(problems))
 return r
}

func printLines(lines [][]string) []MathQuizStruct {
	res := make([]MathQuizStruct, len(lines))

	for i, line := range lines {

		res[i] = MathQuizStruct{
			Q: line[0],
			A: strings.TrimSpace(line[1]),
		}
	}
	return res
}

func runGame() {

}
