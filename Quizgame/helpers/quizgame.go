package helpers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

var (
	Total int
	Right int
	Wrong int
)

func Quizgame() error {
	file, err := os.Open("./helpers/problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//Initalize the reading process for the csv-file
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	row_max_length := len(record) - 1
	for i := 0; i < row_max_length; i++ {
		scanner := bufio.NewReader(os.Stdin)

		fmt.Println(record[i][0])
		userInput, err := scanner.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if (userInput == "exit") || (err != nil) {
			fmt.Println(err)
			os.Exit(1)
		}

		if userInput == record[i][1] {
			Right++
		} else if userInput != record[i][1] {
			Wrong++
		}
		Total = Wrong + Right
	}
	return nil
}
