// Package main provides the logic of the quiz game, no other packages are required.
// The goal of this exercise is to read a provided .csv file,
// Request the user for input
// Check whether the answers provided by the user are correct
// Report to the user the correct amount of questions
// Part 2 of this exercise contains a timer, meaning the user has a set amount of time 'x'
// to provide the correct answers, after the time expires the amount of correct answers is
// reported to the user.

// Given this exercise I hope to learn the following:
// - Use flag package to set a file name and time limit
// - Use os, io and embed (csv) package to read from files
// - Use context and time package to create a timer
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var filename string
	var timeLimit int

	flag.StringVar(&filename, "filename", "problems.csv", "Questionaire filename, supports '.csv'")
	flag.IntVar(&timeLimit, "'time limit'", 30, "Quiz time limit")
	flag.Parse()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	questionaire, err := csvToQuestionaire(file)
	if err != nil {
		return err
	}

	correctAnswerCount, err := questionUser(questionaire, timeLimit)
	if err != nil {
		return err
	}

	fmt.Printf("correct number of answers given: %v", correctAnswerCount)
	return nil
}

func questionUser(questionaire map[string]int, timeLimit int) (int, error) {
	var correct int
	answerChan := make(chan int)

	// Start timer go routine sending a message on channel timer.C when timer expires
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for question, answer := range questionaire {
		// Use another go routine to ask questions in parallel to using the timer.
		go askQuestion(answerChan, question)

		select {
		case <-timer.C: // as soon as the timer expires return to main method execution
			fmt.Println("quiz time expired")
			return correct, nil
		case userAnswer := <-answerChan:
			if userAnswer == answer {
				correct++
			}
		}
	}
	return correct, nil
}

func askQuestion(answerChan chan int, question string) {
	var answer string
	for {
		fmt.Printf("What is %v?\n", question)
		fmt.Scanf("%v\n", &answer)

		answer = strings.TrimSpace(answer)

		if answer == "" {
			fmt.Println("Passed on question")
			answerChan <- 0
			break
		}

		answerInt, err := strconv.Atoi(answer)
		if err == nil {
			fmt.Printf("Provided answer is: %v\n", answerInt)
			answerChan <- answerInt
			break
		}
		fmt.Println("Please provide number format")
	}
}

func csvToQuestionaire(file io.Reader) (map[string]int, error) {
	csvReader := csv.NewReader(file)
	csvReader.Comma = ','
	csvReader.TrimLeadingSpace = true

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	questionaire := map[string]int{}
	for _, record := range records {
		answerInt, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		questionaire[record[0]] = answerInt
	}

	return questionaire, err
}
