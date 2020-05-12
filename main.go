package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var file string
	var duration int
	flag.StringVar(&file, "file", "game.csv", "Provide a file a csv file in the format of 'question,answer' eg '3+5,8'")
	flag.IntVar(&duration, "duration", 30, "specify the duration for the code to run")
	flag.Parse()

	// open file for reading
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

	timer := time.NewTimer(time.Duration(duration) * time.Second)
	correct := 0
problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

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
