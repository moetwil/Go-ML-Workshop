package main

import (
	"fmt"
	"github.com/sajari/regression"
	"log"
	"strconv"
)

func linearRegression() {
	// Read the CSV file
	data := readCSV("./data/regression/advertising.csv")

	// Split the data into training and testing sets
	splitTrainTest(data, "regression")

	// Train the model by setting the dependent and independent variables
	model := trainLinearRegressionModel("Sales", "TV")

	// Make a prediction and print the result
	prediction := predictLinearRegression(model, 230)
	fmt.Println(prediction)
}

func trainLinearRegressionModel(y string, x string) regression.Regression {
	// Read the training data
	trainingData := readCSV("./data/regression/training.csv")

	// Create a new regression model and set the x and y values
	var regressionModel regression.Regression
	regressionModel.SetObserved(y)
	regressionModel.SetVar(0, x)

	// Loop over trainingData in the CSV, adding the training data to the regression value.
	for i, record := range trainingData {
		// Skip the header.
		if i == 0 {
			continue
		}

		// find y index by column name
		yIndex := -1
		for i, columnName := range trainingData[0] {
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
		for i, columnName := range trainingData[0] {
			if columnName == x {
				xIndex = i
				break
			}
		}

		// Check if the column was found in the CSV data
		if xIndex == -1 {
			log.Fatalf("Column '%s' not found in the CSV data", x)
		}

		// Get the x value
		xVal, err := strconv.ParseFloat(record[xIndex], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Add the datapoint to the model
		regressionModel.Train(regression.DataPoint(yVal, []float64{xVal}))
	}

	// Train/fit the regression model
	regressionModel.Run()

	// Output the trained model parameters
	fmt.Printf("\nRegression Formula:\n%v\n\n", regressionModel.Formula)

	// Return the trained model
	return regressionModel
}

func predictLinearRegression(model regression.Regression, x int64) float64 {
	// Make a prediction
	prediction, err := model.Predict([]float64{float64(x)})

	// Check if there was an error
	if err != nil {
		log.Fatal(err)

	}

	// Return the prediction
	return prediction
}
