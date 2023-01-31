package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "csv file in the format 'problem,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		fmt.Printf("Failed to open file: %s \n", *csvFileName)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		fmt.Println("Failed to parse csv file")
	}

	fmt.Println(fmt.Sprintf("%v", lines))

	problems := parseCsvLines(lines)

	fmt.Printf("%v \n", problems)

	var count int
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)
		var answer string
		fmt.Scanf("%s \n", &answer)

		if answer == problem.a {
			count++
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
