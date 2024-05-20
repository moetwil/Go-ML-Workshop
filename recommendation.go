package main

//import (
//	"fmt"
//	"github.com/rocketlaunchr/dataframe-go"
//	"github.com/sajari/regression"
//	"gonum.org/v1/gonum/stat"
//	"sort"
//	"strconv"
//)
//
//type Review struct {
//	UserID    int
//	BookTitle string
//	Rating    int
//}
//
//type BookData struct {
//	Name    string
//	Ratings []float64
//}
//
//func recommendation() {
//	gonumMethod()
//}
//
//func gonumMethod() {
//	reviews, _ := loadReviews()
//	bookTitlesWithIds := getUniqueBookTitles(reviews)
//	ratingsMatrix := getRatingsPerUser(reviews, bookTitlesWithIds)
//	//myMovie := getRatingsForMovie("Harry Potter and the Prisoner of Azkaban (Book 3)", bookTitlesWithIds, reviews)
//
//	fmt.Println(len(ratingsMatrix))
//	fmt.Println(ratingsMatrix[793])
//	fmt.Println(ratingsMatrix[793][bookTitlesWithIds["Harry Potter and the Prisoner of Azkaban (Book 3)"]-1])
//	fmt.Println(ratingsMatrix[3085])
//
//	//weights := getWeights(ratingsMatrix[793], ratingsMatrix[3085])
//	//fmt.Println(weights)
//	corr := stat.Correlation(ratingsMatrix[793], ratingsMatrix[3085], nil)
//
//	fmt.Printf("%.5f", corr)
//
//	bookName := "Harry Potter and the Prisoner of Azkaban (Book 3)"
//
//	var bookIndex int
//	for name, i := range bookTitlesWithIds {
//		if name == bookName {
//			bookIndex = i
//			break
//		}
//	}
//
//	// Extract the book data (ratings)
//	var bookData []float64
//	for _, row := range ratingsMatrix {
//		bookData = append(bookData, row[bookIndex])
//	}
//	fmt.Println(bookData)
//
//	// Compute correlations with other movies
//	var corrs []float64
//	for i := 0; i < len(ratingsMatrix[0]); i++ {
//		//if i != bookIndex {
//		var otherMovieData []float64
//		for _, row := range ratingsMatrix {
//			otherMovieData = append(otherMovieData, row[i])
//		}
//		corr := stat.Correlation(bookData, otherMovieData, nil)
//		corrs = append(corrs, corr)
//		//}
//	}
//
//	fmt.Println(corrs)
//
//	//movieNames := []string{"Movie1", "Movie2", "Movie3", "Movie4", "Movie5"} // Replace with actual movie names
//	//movieMatrix := [][]float64{
//	//	{5, 4, 3, 2, 1}, // Example rating data, replace with actual movie ratings
//	//	{4, 3, 2, 1, 5},
//	//	{3, 2, 1, 5, 4},
//	//	{2, 1, 5, 4, 3},
//	//	{1, 5, 4, 3, 2},
//	//}
//	//ratings := map[string]int{"Movie1": 100, "Movie2": 50, "Movie3": 80} // Replace with actual movie ratings map
//	//
//	//// Call recommend function
//	//recommendedMovies := recommendBook("Movie1", 75, movieMatrix, bookTitles, ratings)
//	//fmt.Println("Recommended Movies:", recommendedMovies)
//}
//
//func loadReviews() ([]Review, error) {
//	// Read the CSV file
//	data := readCSV("./data/recommendation/book_reviews.csv")
//	//data := readCSV("./data/recommendation/Preprocessed_data_filtered.csv")
//
//	var reviews []Review
//	for _, line := range data[1:] {
//		userId, _ := strconv.Atoi(line[0])
//		bookTitle := line[3]
//		rating, _ := strconv.Atoi(line[2])
//		reviews = append(reviews, Review{UserID: userId, BookTitle: bookTitle, Rating: rating})
//	}
//
//	//fmt.Println(reviews[0])
//
//	return reviews, nil
//}
//
//func getUniqueBookTitles(reviews []Review) map[string]int {
//	uniqueBooks := make(map[string]int)
//	i := 0
//	for _, review := range reviews {
//		// add book title to the map if it doesn't exist
//		// also give it an id
//		if _, ok := uniqueBooks[review.BookTitle]; !ok {
//			i++
//			uniqueBooks[review.BookTitle] = i
//		}
//	}
//
//	//var books []string
//	//for title := range uniqueBooks {
//	//	books = append(books, title)
//	//}
//
//	return uniqueBooks
//}
//
//func getRatingsPerUser(reviews []Review, bookTitlesWithIds map[string]int) [][]float64 {
//	var userRatings [][]float64
//	for _, review := range reviews {
//		// initialize the slice if it doesn't exist
//		ratings := make([]float64, len(bookTitlesWithIds))
//		// get the id of the book title
//		id := bookTitlesWithIds[review.BookTitle]
//		// add the rating at the correct position
//		ratings[id-1] = float64(review.Rating)
//		// append the ratings to the userRatings slice
//		userRatings = append(userRatings, ratings)
//	}
//	return userRatings
//}
//
//func getRatingsPerUserCopilot(reviews []Review, bookTitlesWithIds map[string]int) [][]float64 {
//	userRatings := make(map[int][]float64)
//	for _, review := range reviews {
//		// initialize the slice if it doesn't exist
//		if _, ok := userRatings[review.UserID]; !ok {
//			userRatings[review.UserID] = make([]float64, len(bookTitlesWithIds))
//		}
//		// get the id of the book title
//		id := bookTitlesWithIds[review.BookTitle]
//		// add the rating at the correct position
//		userRatings[review.UserID][id-1] = float64(review.Rating)
//	}
//	// convert the map to a slice of slices
//	var ratingsMatrix [][]float64
//	for _, ratings := range userRatings {
//		ratingsMatrix = append(ratingsMatrix, ratings)
//	}
//	return ratingsMatrix
//}
//
//func getRatingsForMovie(movieName string, bookTitlesWithIds map[string]int, reviews []Review) []float64 {
//	// get the id of the movie
//	id := bookTitlesWithIds[movieName]
//	// initialize the ratings slice
//	var ratings []float64
//	for _, review := range reviews {
//		// if the movie id is the same as the current review
//		// add the rating to the ratings slice
//		if bookTitlesWithIds[review.BookTitle] == id {
//			ratings = append(ratings, float64(review.Rating))
//		}
//	}
//	return ratings
//}
//
//func getWeights(ratingsA []float64, ratingsB []float64) []float64 {
//	weights := make([]float64, len(ratingsA))
//	for i, ratingA := range ratingsA {
//		if ratingA == 0 || ratingsB[i] == 0 {
//			weights[i] = 0.0
//		} else {
//			weights[i] = 1.0
//		}
//	}
//	return weights
//}
//
//func recommendBook(movieName string, minRatings int, bookMatrix [][]float64, bookNames []string, ratings map[string]int) []string {
//	// Find the index of the book
//	var bookIndex int
//	for i, name := range bookNames {
//		if name == movieName {
//			bookIndex = i
//			break
//		}
//	}
//
//	// Extract the book data (ratings)
//	var bookData []float64
//	for _, row := range bookMatrix {
//		bookData = append(bookData, row[bookIndex])
//	}
//
//	// Compute correlations with other movies
//	var corrs []float64
//	for i := 0; i < len(bookMatrix[0]); i++ {
//		if i != bookIndex {
//			var otherMovieData []float64
//			for _, row := range bookMatrix {
//				otherMovieData = append(otherMovieData, row[i])
//			}
//			corr := stat.Correlation(bookData, otherMovieData, nil)
//			corrs = append(corrs, corr)
//		}
//	}
//
//	// Create a map of movie correlations to their names
//	movieCorrelations := make(map[string]float64)
//	for i, corr := range corrs {
//		movieCorrelations[bookNames[i]] = corr
//	}
//
//	// Sort the movies by correlation in descending order
//	type movieCorrelation struct {
//		name  string
//		corr  float64
//		rates int
//	}
//	var sortedMovies []movieCorrelation
//	for name, corr := range movieCorrelations {
//		rates := ratings[name]
//		if rates >= minRatings {
//			sortedMovies = append(sortedMovies, movieCorrelation{name, corr, rates})
//		}
//	}
//	sort.Slice(sortedMovies, func(i, j int) bool {
//		return sortedMovies[i].corr > sortedMovies[j].corr
//	})
//
//	// Extract recommended movies
//	var recommendedMovies []string
//	for _, movie := range sortedMovies {
//		recommendedMovies = append(recommendedMovies, movie.name)
//	}
//
//	return recommendedMovies
//}
//
////func recommendation() {
////	// Import the dataset
////	data := readCSVAsReader("./data/recommendation/book_reviews.csv")
////
////	// Load the dataset into a DataFrame
////	ctx := context.Background()
////	opts := imports.CSVLoadOptions{
////		DictateDataType: map[string]interface{}{
////			"user_id":   float64(0),
////			"rating":    float64(0),
////			"book_type": "",
////		},
////	}
////	df, err := imports.LoadFromCSV(ctx, data, opts)
////
////	if err != nil {
////		panic(err)
////	}
////
////	// Print the first 3 and last 3 rows
////	fmt.Println(df)
////
////	// Make a list of all the unique book titles by iterating over the Series
////	var books []string
////	bookTitleColumnCopy := df.Series[3].Copy()
////	bookTitleColumnCopy.Sort(ctx)
////	booksIterator := bookTitleColumnCopy.ValuesIterator()
////	for {
////		_, value, _ := booksIterator()
////		if value == nil {
////			break
////		}
////		title := value.(string)
////		if !slices.Contains(books, title) {
////			books = append(books, title)
////		}
////	}
////
////	// Convert the book_title column to numerical values
////	bookTitleColumn := df.Series[3].(*dataframe.SeriesString)
////	convertedBookTitleColumn, _ := bookTitleColumn.ToSeriesFloat64(ctx, false, func(value interface{}) (float64, error) {
////		index := slices.Index(books, *(value.(*string))) // thank you Copilot
////		if index == -1 {
////			return 0, nil
////		}
////		return float64(index), nil
////	})
////
////	// Copy the DataFrame and replace the book_title column
////	dfCopy := df.Copy()
////	dfCopy.Series[3] = convertedBookTitleColumn
////
////	fmt.Println(convertedBookTitleColumn)
////	fmt.Println(dfCopy)
////
////	corr, _ := stats.Correlation(
////		dfCopy.Series[3].(*dataframe.SeriesFloat64).Values,
////		dfCopy.Series[2].(*dataframe.SeriesFloat64).Values,
////	)
////	fmt.Println(corr)
////
////	// Filter to one user
////	filterFn := dataframe.FilterDataFrameFn(func(vals map[interface{}]interface{}, row, nRows int) (dataframe.FilterAction, error) {
////		if vals["user_id"] == 256167.0 {
////			return dataframe.KEEP, nil
////		}
////		return dataframe.DROP, nil
////	})
////	userBooks, _ := dataframe.Filter(ctx, dfCopy, filterFn)
////
////	//fmt.Println(userBooks.(dataframe.DataFrame))
////
////	fmt.Println(userBooks)
////
////	//correlation()
////
////	// Train the model by setting the dependent and independent variables
////	//model := trainRecommendationModel(df)
////
////	//	// Make a prediction and print the result
////	//	prediction := predictRecommendation(model, 230)
////	//	fmt.Println(prediction)
////}
////
////func correlation() {
////	// get all ratings from user 1
////	s1 := []float64{1, 2, 3, 4, 5}
////	// get all ratings from user 2
////	s2 := []float64{1, 2, 3, 5, 6}
////	// calculate the correlation between the two users' ratings
////	a, _ := stats.Correlation(s1, s2)
////	fmt.Println(a)
////}
//
//func trainRecommendationModel(df *dataframe.DataFrame) {
//	// Preprocess your data (if needed)
//
//	// Implement or choose a recommendation algorithm
//	// For example, let's say we have a simple popularity-based recommender
//	// that recommends items based on their average rating
//	//averageRatings := df.GroupBy("item_id").Mean("rating")
//	//
//	//// Make recommendations for a user (or item)
//	//userID := 123
//	//userRecommendations := averageRatings.Filter(dataframe.F{"user_id", "==", userID})
//	//
//	//// Print recommendations
//	//fmt.Println("Recommendations for user", userID)
//	//fmt.Println(userRecommendations)
//
//	//// Output the trained model parameters
//	//fmt.Printf("\nRegression Formula:\n%v\n\n", regressionModel.Formula)
//	//
//	//// Return the trained model
//	//return regressionModel
//}
//
//func predictRecommendation(model regression.Regression, x int64) float64 {
//	//// Make a prediction
//	//prediction, err := model.Predict([]float64{float64(x)})
//	//
//	//// Check if there was an error
//	//if err != nil {
//	//	log.Fatal(err)
//	//
//	//}
//	//
//	//// Return the prediction
//	//return prediction
//	return 0.0
//}
