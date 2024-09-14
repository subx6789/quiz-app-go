package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

// Problem struct represents a quiz question with multiple choices and the correct answer.
type Problem struct {
	Question      string
	OptionA       string
	OptionB       string
	OptionC       string
	OptionD       string
	CorrectAnswer string
}

func main() {
	fmt.Println("Welcome to Quiz App in Golang by Subhajit Sarkar. Test your general knowledge skills!")
	// Flags for CSV file path and timer duration
	fName := flag.String("f", "quiz.csv", "path of csv file")
	timer := flag.Int("t", 30, "timer for the quiz in seconds")
	flag.Parse()
	// Load the problems from the CSV file
	problems, err := problemPuller(*fName)
	if err != nil {
		exit(fmt.Sprintf("Failed to load problems: %s\n", err.Error()))
	}
	// Initialize variables for tracking correct answers and the quiz timer
	correctAns := 0
	// Use bufio.NewScanner to handle user input with spaces
	reader := bufio.NewScanner(os.Stdin)
	ansC := make(chan string)
	// Loop through each problem and ask the user for their answer
problemLoop:
	for i, p := range problems {
		tObj := time.NewTimer(time.Duration(*timer) * time.Second)
		// Display the question and options
		fmt.Printf("%d. %s\nA) %s\nB) %s\nC) %s\nD) %s\nYour answer: ", i+1, p.Question, p.OptionA, p.OptionB, p.OptionC, p.OptionD)
		// Use a goroutine to gather the user's answer
		go func() {
			if reader.Scan() {
				ansC <- reader.Text()
			}
		}()
		// Use select to handle timeout or answer input
		select {
		case <-tObj.C:
			fmt.Println("\nTime's up!")
			break problemLoop
		case iAns := <-ansC:
			// Check if the user's answer is correct
			if iAns == p.CorrectAnswer {
				correctAns++
			}
		}
	}
	// Display the result
	fmt.Printf("\nYour score: %d out of %d\n", correctAns, len(problems))
	fmt.Println("Press enter to exit ...")
	// Wait for the user to press enter before exiting
	var exitInput string
	fmt.Scanln(&exitInput)
}

// problemPuller reads a CSV file and returns a slice of Problem structs
func problemPuller(filename string) ([]Problem, error) {
	// Open the CSV file
	fObj, err := os.Open(filename)
	if err != nil {
		// Error handling if the file cannot be opened
		return nil, fmt.Errorf("could not open the file: %s", err.Error())
	}
	defer fObj.Close() // Ensure the file is closed after reading
	// Read all the lines from the CSV
	csvR := csv.NewReader(fObj)
	cLines, err := csvR.ReadAll()
	if err != nil {
		// Error handling if CSV cannot be read
		return nil, fmt.Errorf("could not read the CSV data: %s", err.Error())
	}
	// Parse the lines into Problem structs
	return parseProblem(cLines), nil
}

// parseProblem converts CSV lines into a slice of Problem structs
func parseProblem(lines [][]string) []Problem {
	// Create a slice to hold all the problems
	r := make([]Problem, len(lines))
	// Iterate through each line and populate the Problem struct
	for i, line := range lines {
		r[i] = Problem{
			Question:      line[0],
			OptionA:       line[1],
			OptionB:       line[2],
			OptionC:       line[3],
			OptionD:       line[4],
			CorrectAnswer: line[5],
		}
	}
	return r
}

// exit gracefully exits the program with an error message
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
