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

	// Uncomment these rows if u want new train and test data
	records := readData("./data/housing.csv")
	splitTrainTest(records)
	model := trainModel("price", "bedrooms")

	fmt.Println(predict(model, 2))

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
	// Open the CSV file from the disk
	f, err := os.Open("./data/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader
	salesData := csv.NewReader(f)

	// Read all the records
	records, err := salesData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// print size of the records
	fmt.Println(len(records))

	// In this case, we are going to try and model median house value (y) by the total rooms feature.
	var r regression.Regression
	r.SetObserved(y)
	r.SetVar(0, x)

	// Loop over records in the CSV, adding the training data to the regression value.
	for i, record := range records {
		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the house price (y)
		price, err := strconv.ParseFloat(record[0], 64)
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

		// Parse the total rooms value
		xVal, err := strconv.ParseFloat(record[xIndex], 64)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(xVal)
		r.Train(regression.DataPoint(price, []float64{xVal}))
	}

	// Train/fit the regression model
	r.Run()

	// Output the trained model parameters
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)

	return r

}

func predict(model regression.Regression, numRooms int64) float64 {
	//	make prediction
	prediction, err := model.Predict([]float64{float64(numRooms)})
	if err != nil {
		log.Fatal(err)

	}
	return prediction
}
