package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Quiz struct {
	index    int
	question string
	answer   string
}

func ReadCSV() []Quiz {
	var filename string

	flag.StringVar(&filename, "filename", "problems.csv", "input filename")
	flag.Parse()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("FATAL: Failed to open file.", filename, err)
	}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("FATAL: Failed to read csv records from file.", filename, err)
	}

	var parsedRecords []Quiz
	for index, record := range records {
		parsedRecords = append(parsedRecords, Quiz{index: index, question: record[0], answer: record[1]})
	}

	return parsedRecords
}

func PopQuiz(quizes []Quiz) []bool {
	userAnswers := make([]bool, len(quizes))
	for _, quiz := range quizes {
		var input string
		fmt.Printf("Q: %s, Answer: ", quiz.question)
		fmt.Scanln(&input)
		userAnswers[quiz.index] = input == quiz.answer
	}

	return userAnswers
}

func main() {
	quizes := ReadCSV()
	userAnswers := PopQuiz(quizes)
	var correct, incorrect int
	for _, quiz := range quizes {
		// fmt.Println(quiz.question, quiz.answer, userAnswers[quiz.index])
		if userAnswers[quiz.index] {
			correct++
		} else {
			incorrect++
		}
	}
	fmt.Printf("Summary: totalCount=%d, corrects=%d, incorrect=%d\n", len(quizes), correct, incorrect)
}
