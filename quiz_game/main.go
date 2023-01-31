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
	csvFileName := flag.String("csv", "problems.csv", "csv file in the format 'problem,answer'")
	timeLimit := flag.Int("limit", 5, "Quiz duration")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		fmt.Println("Failed to parse csv file")
	}

	fmt.Println(fmt.Sprintf("%v", lines))

	problems := parseCsvLines(lines)
	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))
	var count int

loop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.q)

		answerChan := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s \n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			break loop
		case answer := <-answerChan:
			if answer == problem.a {
				count++
			}
		}
	}

	fmt.Printf("You scored %d out of %d \n", count, len(problems))
}

type problem struct {
	q string
	a string
}

func parseCsvLines(input [][]string) []problem {
	result := make([]problem, len(input))
	for i, innerSlice := range input {
		result[i] = problem{
			q: innerSlice[0],
			a: strings.TrimSpace(innerSlice[1]),
		}
	}
	return result
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)

}
