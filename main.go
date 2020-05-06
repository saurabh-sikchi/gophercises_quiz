package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffleLines := flag.Bool("shuffle", false, "display problems in a random order")
	flag.Parse()

	_ = shuffleLines
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s", *csvFilename))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the csv file")
	}

	problems := parseLines(lines)
	if *shuffleLines {
		shuffle(problems)
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, prob := range problems {

		answerCh := make(chan string)
		go func() {
			fmt.Printf("Problem #%d: %s = ", i+1, prob.q)
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou ran out of time and scored %d out of %d.\n", correct, len(problems))
			return
		case ans := <-answerCh:
			if ans == prob.a {
				correct += 1
			}
		}

	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func shuffle(vals []problem) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
