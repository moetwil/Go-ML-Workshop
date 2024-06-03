package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
	"os"
)

func classification() {
	rawData, err := base.ParseCSVToInstances("./data/classification/heart_statlog_cleveland_hungary_final.csv", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.70)

	fmt.Println("KNN Classifier:")
	knnClassifier(trainData, testData)

	fmt.Println("\nRandom Forest Classifier:")
	randomForestClassifier(trainData, testData)
}

func knnClassifier(trainData, testData base.FixedDataGrid) {
	knnClassifier := knn.NewKnnClassifier("euclidean", "linear", 5)

	err := knnClassifier.Fit(trainData)
	if err != nil {
		return
	}

	predictions, err := knnClassifier.Predict(testData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	evaluateModel(testData, predictions)
}

func randomForestClassifier(trainData, testData base.FixedDataGrid) {
	rf := ensemble.NewRandomForest(500, 5)

	err := rf.Fit(trainData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	predictions, err := rf.Predict(testData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	evaluateModel(testData, predictions)
}

func evaluateModel(testData, predictions base.FixedDataGrid) {
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
}
