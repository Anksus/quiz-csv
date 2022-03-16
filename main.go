package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problem.csv", "A CSV file in the format of question,answer.")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit("Failed to open the csv file: %s\n")
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV.")
	}

	problems := parseLines(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, p.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

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

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
