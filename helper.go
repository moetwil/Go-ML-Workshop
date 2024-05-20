package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

/*
Description: This function can be used to read data from a CSV file.
Params: name - the name of the file to read
*/
func readCSV(name string) [][]string {
	// Open the file
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	// Close the file when the function returns
	defer file.Close()

	// Read the CSV data
	data := csv.NewReader(file)

	// Parse the CSV data
	records, err := data.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Return the records
	return records
}

func readCSVAsReader(name string) *strings.Reader {
	// Open the file
	file, err := os.ReadFile(name) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	data := string(file)
	return strings.NewReader(data)
}

/*
Description: This function can be used to split the data in a CSV file into a training and testing set.
Params:
- records - the data to split
- problemType - the type of problem (e.g. regression, classification)
*/
func splitTrainTest(records [][]string, problemType string) {
	// Define the directory path based on the problem type
	dataDir := fmt.Sprintf("./data/%s", problemType)

	// Create the directory if it does not exist
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dataDir, 0755); err != nil {
			log.Fatalf("Failed to create directory: %s", err)
		}
	}

	// Save the header
	header := records[0]

	// Shuffle the records to randomize the data
	shuffled := make([][]string, len(records)-1)
	perm := rand.Perm(len(records) - 1)
	for i, v := range perm {
		shuffled[v] = records[i+1]
	}

	// Split the training set
	trainingIdx := (len(shuffled)) * 4 / 5
	trainingSet := shuffled[1 : trainingIdx+1]

	// Split the testing set
	testingSet := shuffled[trainingIdx+1:]

	// Define the file paths for training and testing data
	trainingPath := fmt.Sprintf("%s/training.csv", dataDir)
	testingPath := fmt.Sprintf("%s/testing.csv", dataDir)

	// Save the split sets to separate CSV files
	sets := map[string][][]string{
		trainingPath: trainingSet,
		testingPath:  testingSet,
	}

	for fn, dataset := range sets {
		f, err := os.Create(fn)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		out := csv.NewWriter(f)
		defer out.Flush()

		// Write the header to the CSV file
		if err := out.Write(header); err != nil {
			log.Fatal(err)
		}

		// Write the dataset to the CSV file
		if err := out.WriteAll(dataset); err != nil {
			log.Fatal(err)
		}
	}
}
