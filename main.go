package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "game.csv", "Provide a file a csv file in the format of 'question,answer' eg '3+5,8'")
	flag.Parse()

	f, err := os.Open(file)
	if err != nil {
		exit(fmt.Sprintf("Unable to open file: %s, (%v)", file, err), 1)
	}

	defer f.Close()

	reader := csv.NewReader(f)

	lines, err := reader.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Unable to read file, %v (%v)", file, err), 1)
	}

	problems := parseFiles(lines)

	var count int
	for i, problem := range problems {
		fmt.Printf("Problem %d, %s: ", i+1, problem.question)

		var answer string
		fmt.Scanf("%s\n", (&answer))

		if strings.TrimSpace(answer) == problem.answer {
			count++
		}

	}

	fmt.Printf("You got a total of %d out of %d\n", count, len(problems))

}

func exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}

func parseFiles(lines [][]string) []problems {
	problems := make([]problems, len(lines))
	for i, line := range lines {
		problems[i] = newProblems(line[0], line[1])
	}
	return problems
}

type problems struct {
	question, answer string
}

func newProblems(q, a string) problems {
	return problems{
		question: q,
		answer:   a,
	}
}
