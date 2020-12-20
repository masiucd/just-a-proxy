package mathgameuno

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

// MathQuizStruct for math quiz
type MathQuizStruct struct {
	Q string
	A string
}

// Counter type
type Counter int

// Game Math game
func Game(file string) string {

	lines, err := ReadFile(file)

	if err != nil {
		log.Fatal("could not parse from ReadFile function")
	}

	problems := PrintLines(lines)

	var counter Counter = 0

	RunGame(problems, &counter)

	r := fmt.Sprintf("you got %d out of %d possible \n", counter, len(problems))
	return r

}

// PrintLines get the lines
func PrintLines(lines [][]string) []MathQuizStruct {
	res := make([]MathQuizStruct, len(lines))

	for i, line := range lines {

		res[i] = MathQuizStruct{
			Q: line[0],
			A: strings.TrimSpace(line[1]),
		}
	}
	return res
}

// RunGame start game
func RunGame(problems []MathQuizStruct, counter *Counter) {
	for i, p := range problems {
		fmt.Printf("#%d %s = \n", i, p.Q)
		var input string
		fmt.Scanf("%s", &input)

		if input == p.A {
			*counter++
			fmt.Printf("correct \n")
		} else {
			fmt.Printf("wrong \n")
		}
	}
}


// ReadFile reading input file
func ReadFile(file string) ([][]string, error) {
	csvFile, err := os.Open(file)

	if err != nil {
		log.Fatalln("could not read file")
	}
	defer csvFile.Close()

	r, err := csv.NewReader(csvFile).ReadAll()
	x := reflect.TypeOf(r)
	fmt.Println(x)

	return r, err
}
