package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	// uncomment these rows if u want new train and test data
	records := readData("./data/advertising.csv")
	splitTrainTest(records)

	// train the model
	model := trainModel("Sales", "TV")

	// make a prediction and print the result
	fmt.Println(predict(model, 230))

}

func readData(name string) [][]string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data := csv.NewReader(file)

	records, err := data.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func splitTrainTest(records [][]string) {
	// save the header
	header := records[0]

	// shuffle the records
	shuffled := make([][]string, len(records)-1)
	perm := rand.Perm(len(records) - 1)
	for i, v := range perm {
		shuffled[v] = records[i+1]
	}
	// split the training set
	trainingIdx := (len(shuffled)) * 4 / 5
	trainingSet := shuffled[1 : trainingIdx+1]
	// split the testing set
	testingSet := shuffled[trainingIdx+1:]
	// we write the splitted sets in separate files
	sets := map[string][][]string{
		"./data/training.csv": trainingSet,
		"./data/testing.csv":  testingSet,
	}
	for fn, dataset := range sets {
		f, err := os.Create(fn)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		out := csv.NewWriter(f)
		if err := out.Write(header); err != nil {
			log.Fatal(err)
		}
		if err := out.WriteAll(dataset); err != nil {
			log.Fatal(err)
		}
		out.Flush()
	}
}

func trainModel(y string, x string) regression.Regression {
	// open the file
	f, err := os.Open("./data/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := csv.NewReader(f)
	records, err := data.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// create a new regression model and set the x and y values
	var r regression.Regression
	r.SetObserved(y)
	r.SetVar(0, x)

	// Loop over records in the CSV, adding the training data to the regression value.
	for i, record := range records {
		// Skip the header.
		if i == 0 {
			continue
		}

		// find y index by column name
		yIndex := -1
		for i, columnName := range records[0] {
			if columnName == y {
				yIndex = i
				break
			}
		}

		// get the y value
		yVal, err := strconv.ParseFloat(record[yIndex], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Find the index of the column with the name 'x'
		xIndex := -1
		for i, columnName := range records[0] {
			if columnName == x {
				xIndex = i
				break
			}
		}

		if xIndex == -1 {
			log.Fatalf("Column '%s' not found in the CSV data", x)
		}

		// get the x value
		xVal, err := strconv.ParseFloat(record[xIndex], 64)
		if err != nil {
			log.Fatal(err)
		}

		// add the datapoint to the model
		r.Train(regression.DataPoint(yVal, []float64{xVal}))
	}

	// Train/fit the regression model
	r.Run()

	// Output the trained model parameters
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)

	return r
}

func predict(model regression.Regression, x int64) float64 {
	prediction, err := model.Predict([]float64{float64(x)})
	if err != nil {
		log.Fatal(err)

	}
	return prediction
}
