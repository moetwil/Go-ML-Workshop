package main

import (
	"encoding/csv"
	"fmt"
	"gonum.org/v1/gonum/stat"
	"log"
	"os"
	"sort"
	"strconv"
)

type Review struct {
	UserID    string
	BookTitle string
	Rating    float64
}

type UserRatings map[string]float64
type BookRatings map[string]UserRatings

func recommendation() {
	// Adjust the filename to point to your actual CSV file
	//filename := "data/recommendation/book_reviews.csv"
	filename := "data/recommendation/Preprocessed_data_filtered.csv"

	reviews, err := readCsvToReview(filename)
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	bookRatings := organizeData(reviews)
	correlations := computeCorrelations(bookRatings)

	// Example usage
	//bookTitle := "Some Book Title"
	//bookTitle := "Harry Potter and the Prisoner of Azkaban (Book 3)"
	//bookTitle := "A Game of Thrones (A Song of Ice and Fire, Book 1)"
	bookTitle := "Wild Animus"
	minCorrelations := 5

	recommendations := recommend(bookTitle, bookRatings, correlations, minCorrelations)
	fmt.Printf("Recommendations for book '%s':\n", bookTitle)
	for _, rec := range recommendations {
		fmt.Println(rec)
	}
}

func readCsvToReview(filename string) ([]Review, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var reviews []Review
	for _, record := range records[1:] {
		rating, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, Review{
			UserID:    record[0],
			BookTitle: record[3],
			Rating:    rating,
		})
	}
	return reviews, nil
}

func organizeData(reviews []Review) BookRatings {
	bookRatings := make(BookRatings)
	for _, review := range reviews {
		if _, exists := bookRatings[review.BookTitle]; !exists {
			bookRatings[review.BookTitle] = make(UserRatings)
		}
		bookRatings[review.BookTitle][review.UserID] = review.Rating
	}
	return bookRatings
}

func mean(data []float64) float64 {
	return stat.Mean(data, nil)
}

func stddev(data []float64) float64 {
	return stat.StdDev(data, nil)
}

func correlation(a, b []float64) float64 {
	if len(a) != len(b) {
		return 0.0
	}
	//meanA := mean(a)
	//meanB := mean(b)
	stdDevA := stddev(a)
	stdDevB := stddev(b)

	if stdDevA == 0 || stdDevB == 0 {
		return 0.0
	}

	covariance := stat.Covariance(a, b, nil)
	return covariance / (stdDevA * stdDevB)
}

func computeCorrelations(bookRatings BookRatings) map[string]map[string]float64 {
	correlations := make(map[string]map[string]float64)
	for bookA, ratingsA := range bookRatings {
		correlations[bookA] = make(map[string]float64)
		for bookB, ratingsB := range bookRatings {
			if bookA != bookB {
				var commonUsers []string
				for user := range ratingsA {
					if _, exists := ratingsB[user]; exists {
						commonUsers = append(commonUsers, user)
					}
				}
				if len(commonUsers) > 1 {
					var ratingsListA, ratingsListB []float64
					for _, user := range commonUsers {
						ratingsListA = append(ratingsListA, ratingsA[user])
						ratingsListB = append(ratingsListB, ratingsB[user])
					}
					correlations[bookA][bookB] = correlation(ratingsListA, ratingsListB)
				}
			}
		}
	}
	return correlations
}

func recommend(bookTitle string, bookRatings BookRatings, correlations map[string]map[string]float64, minCorrelations int) []string {
	recommendations := make(map[string]float64)
	for book, corr := range correlations[bookTitle] {
		if corr > 0 {
			recommendations[book] = corr
		}
	}

	type recommendation struct {
		BookTitle   string
		Correlation float64
	}

	var recommendationList []recommendation
	for book, corr := range recommendations {
		recommendationList = append(recommendationList, recommendation{BookTitle: book, Correlation: corr})
	}

	sort.Slice(recommendationList, func(i, j int) bool {
		return recommendationList[i].Correlation > recommendationList[j].Correlation
	})

	var topRecommendations []string
	for i := 0; i < minCorrelations && i < len(recommendationList); i++ {
		topRecommendations = append(topRecommendations, recommendationList[i].BookTitle)
	}

	return topRecommendations
}
