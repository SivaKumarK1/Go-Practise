package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// need to change and break down these function into smaller functions

func main() {
	csvFile := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit",30, "the time limit for the quiz in seconds")
	flag.Parse()
	file, err := os.Open(*csvFile)
	defer file.Close()
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
	score := showQues(problems,*timeLimit)
	fmt.Printf("Your Score is : %d\n", score)
}

func showQues(problems []problem, timeLimit int) int {
	score := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func()  {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh<-answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {	
				score++
			}
			if answer != p.a {
				fmt.Printf("Incorrect Answer!! \nDeducting From Score\n")
				score--
			}
		
		}
	}
	return score
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
