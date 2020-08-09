package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	inputPath := flag.String("input", "sample.csv", "Quiz csv file path")
	timeLimit := flag.Int("limit", 30, "time limit for answering all questions, in seconds")
	flag.Parse()

	lines, err := readCsv(*inputPath)
	check(err)

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correctAnwers := 0

	for i, problem := range problems {
		fmt.Printf("Question %d:\n  %s?\n", i+1, problem.question)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou got %d correct answers out of %d\n", correctAnwers, len(lines))
			return
		case answer := <-answerCh:
			if strings.TrimSpace(answer) == problem.answer {
				correctAnwers += 1
				fmt.Println("Correct!")
			} else {
				fmt.Println("Whoops!")
			}
		}
	}
	fmt.Printf("\nYou got %d correct answers out of %d\n", correctAnwers, len(lines))
}

type problem struct {
	question string
	answer   string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)

	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, err
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems
}
