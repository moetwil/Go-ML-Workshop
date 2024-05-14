package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"os"
)

func classification() {
	rawData, err := base.ParseCSVToInstances("./data/classification/heart_statlog_cleveland_hungary_final.csv", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.70)

	// print trainData and testData to see what they look like
	fmt.Println(trainData)
	fmt.Println(testData)

	// train two different models and evaluate them using the evaluateModel function
}

// use this function to evaluate the model
func evaluateModel(testData, predictions base.FixedDataGrid) {
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
}
