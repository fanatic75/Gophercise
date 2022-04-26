package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var csvFilePath string
var limit int
var score = 0

func main() {

	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	fd, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	fileR := csv.NewReader(fd)
	questions, erro := fileR.ReadAll()
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Press enter to start Program")
	scanner.Scan()

	quit := time.NewTimer(time.Duration(limit) * time.Second)
	ansChan := make(chan string)
	go func() {
		scanner.Scan()
		ansChan <- scanner.Text()
	}()
	for i, questionAns := range questions {
		question := questionAns[0]

		fmt.Printf("\nProblem #%d: %s = ", (i + 1), question)
		select {
		case <-quit.C:
			fmt.Printf("\nYou scored %d out of %d", score, len(questions))
			return

		case ansInput := <-ansChan:

			if answer := questionAns[1]; answer == ansInput {
				score++
			}

		}

	}
	fmt.Printf("\nYou scored %d out of %d", score, len(questions))

	defer fd.Close()

}

func init() {
	getCSVPath()
	getLimit()
}

func getCSVPath() {
	csvUsage := `a csv file in the format of 'question,answer'`
	flag.StringVar(&csvFilePath, "csv", "problems.csv", csvUsage)

}

func getLimit() {
	limitUsage := `the time limit for the quiz in seconds`
	flag.IntVar(&limit, "limit", 30, limitUsage)
}
