package mathgamedues

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Q string
	A string
}

// Game run the game
func Game(fileName string) string {

	score := 0
	problems, timer := readProblems(fileName)

	gameRun := runQuestions(problems, score, timer)
	return gameRun
}

func readProblems(problemsFile string) ([]Problem, *time.Timer) {

	csvFile := flag.String("csv", problemsFile, "csv file foramt with question.answer")

	timeLimit := flag.Int("limit", 30, "the limit for the game timer")

	flag.Parse()

	file, err := os.Open(*csvFile)

	if err != nil {
		log.Fatal("failed to open file %s\n", *csvFile)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal("failed to parse provided csv file")
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	problems := parseLines(lines)

	return problems, timer

}

func parseLines(lines [][]string) []Problem {
	x := make([]Problem, len(lines))

	for i, v := range lines {
		x[i] = Problem{
			Q: v[0],
			A: strings.TrimSpace(v[1]),
		}
	}

	return x
}

func runQuestions(problems []Problem, score int, timer *time.Timer) string {
problemLabel: // label
	for i, v := range problems {
		fmt.Printf("Problem #%d %s =", i, v.Q)

		answerCn := make(chan string)
		go func() {
			var input string
			fmt.Scanf("%s\n", &input)
			// closure here
			answerCn <- input
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemLabel
		case input := <-answerCn:
			if input == v.A {
				score++
				fmt.Println("correct")
			}
		}
	}

	x := fmt.Sprintf("you got %d out of %d \n", score, len(problems))
	return x

}
