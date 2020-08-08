package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputPath := flag.String("input", "sample.csv", "Quiz csv file path")
	flag.Parse()

	lines, err := readCsv(*inputPath)
	check(err)

	problems := parseLines(lines)

	correctAnwers := 0

	keyboardReader := bufio.NewReader(os.Stdin)
	for i, problem := range problems {
		fmt.Printf("Question %d:\n  %s?\n", i+1, problem.question)

		answer, err := keyboardReader.ReadString('\n')
		check(err)

		if strings.TrimSpace(answer) == problem.answer {
			correctAnwers += 1
			fmt.Println("Correct!")
		} else {
			fmt.Println("Whoops!")
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
