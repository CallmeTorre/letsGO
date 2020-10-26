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

type result struct {
	correct int
	wrong   int
}

type problem struct {
	question string
	answer   string
}

var csvFilename *string
var timeLimit *int

func init() {
	csvFilename = flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer'")
	timeLimit = flag.Int("limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()
}

func main() {

	file := openFile(csvFilename)

	lines := readFile(file)

	problems := parseLines(lines)

	fmt.Println(problems)

	shuffleProblems(&problems)

	fmt.Println(problems)

	/*timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	finalResult := printProblems(problems, timer)

	printResults(*finalResult, len(problems))*/
}

func openFile(csvFilename *string) *os.File {
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}
	return file
}

func readFile(file *os.File) [][]string {
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse de provided CSV file")
	}
	return lines
}

func parseLines(lines [][]string) []problem {
	result := make([]problem, len(lines))
	for i, line := range lines {
		result[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return result
}

func shuffleProblems(problems *[]problem) {
	var source rand.Source = rand.NewSource(time.Now().UnixNano())
	var r *rand.Rand = rand.New(source)

	for i := range *problems {
		newIndex := r.Intn(len(*problems) - 1)
		(*problems)[i], (*problems)[newIndex] = (*problems)[newIndex], (*problems)[i]
	}

}

func printProblems(problems []problem, timer *time.Timer) *result {
	finalResult := &result{correct: 0, wrong: 0}
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		answerChannel := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			return finalResult
		case answer := <-answerChannel:
			if answer == problem.answer {
				finalResult.correct++
			} else {
				finalResult.wrong++
			}
		}
	}
	return finalResult
}

func printResults(finalResult result, problemsLen int) {
	fmt.Printf("Correct answers: %d\n", finalResult.correct)
	fmt.Printf("Wrong answers: %d\n", finalResult.wrong)
	fmt.Printf("Total questions: %d\n", problemsLen)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
