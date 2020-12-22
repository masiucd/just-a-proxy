package mathgamedues

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"text/template/parse"
)

func game()  {
}

func readProblems(problemsFile string)  {
	csvFile := flag.String("csv","problems.csv","csv file foramt with question.answer")	
	timeLimit:= flag.Int("limi",30,"time limit fr the quiz in seconds")

	flag.Parse()

	file,err := os.Open(*csvFile)
	if err != nil {
		log.Fatal("could not read the csv file")
	}

	r := csv.NewReader(file)
	lines,err := r.ReadAll()

	if err != nil {
		log.Fatal("could not parse the csv file")
	}

	problems := 
	
}

func parseLines(lines [][]string)  {
	
}