package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFile := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit",30, "the time limit for the quiz in seconds")
	flag.Parse()
	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFile))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the privided CSV file")
	}
	problems := parseLines(lines)
	
	score := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Printf("Correct Answer! \n Incrementing Score\n")
			score++
		}
		if answer != p.a {
			fmt.Printf("Incorrect Answer!! \n Deducting From Score\n")
			score--
		}
	}
	fmt.Printf("Your Score is : %d\n", score)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
