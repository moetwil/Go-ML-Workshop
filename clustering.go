package main

import (
	"fmt"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Point represents a point in 2D space
type Point struct {
	X, Y float64
}

// KMeans represents the K-means algorithm
type KMeans struct {
	K         int
	Points    []Point
	Centroids []Point
}

// NewKMeans creates a new KMeans instance
func NewKMeans(k int, points []Point) *KMeans {
	return &KMeans{
		K:         k,
		Points:    points,
		Centroids: make([]Point, k),
	}
}

// InitializeCentroids randomly initializes the centroids
func (km *KMeans) InitializeCentroids() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < km.K; i++ {
		km.Centroids[i] = km.Points[rand.Intn(len(km.Points))]
	}
}

// AssignClusters assigns each point to the nearest centroid
func (km *KMeans) AssignClusters() [][]Point {
	clusters := make([][]Point, km.K)
	for _, p := range km.Points {
		minDist := math.Inf(1)
		clusterIdx := 0
		for i, c := range km.Centroids {
			dist := distance(p, c)
			if dist < minDist {
				minDist = dist
				clusterIdx = i
			}
		}
		clusters[clusterIdx] = append(clusters[clusterIdx], p)
	}
	return clusters
}

// UpdateCentroids updates the centroids to the mean of the clusters
func (km *KMeans) UpdateCentroids(clusters [][]Point) {
	for i, cluster := range clusters {
		if len(cluster) == 0 {
			continue
		}
		var sumX, sumY float64
		for _, p := range cluster {
			sumX += p.X
			sumY += p.Y
		}
		km.Centroids[i] = Point{X: sumX / float64(len(cluster)), Y: sumY / float64(len(cluster))}
	}
}

// Fit runs the K-means algorithm until convergence
func (km *KMeans) Fit() [][]Point {
	km.InitializeCentroids()
	var clusters [][]Point
	for {
		clusters = km.AssignClusters()
		oldCentroids := km.Centroids
		km.UpdateCentroids(clusters)
		if equalCentroids(oldCentroids, km.Centroids) {
			break
		}
	}
	return clusters
}

// distance calculates the Euclidean distance between two points
func distance(p1, p2 Point) float64 {
	return math.Sqrt((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))
}

// equalCentroids checks if two sets of centroids are equal
func equalCentroids(c1, c2 []Point) bool {
	for i, p := range c1 {
		if p != c2[i] {
			return false
		}
	}
	return true
}

// plotClusters visualizes the clusters
func plotClusters(clusters [][]Point, centroids []Point) {
	p := plot.New()

	p.Title.Text = "K-Means Clustering"
	p.X.Label.Text = "Finishing"
	p.Y.Label.Text = "Heading Accuracy"

	colors := []color.RGBA{
		{R: 255, G: 0, B: 0, A: 255},     // red
		{R: 0, G: 0, B: 255, A: 255},     // blue
		{R: 0, G: 255, B: 0, A: 255},     // green
		{R: 255, G: 165, B: 0, A: 255},   // orange
		{R: 128, G: 0, B: 128, A: 255},   // purple
		{R: 165, G: 42, B: 42, A: 255},   // brown
		{R: 255, G: 192, B: 203, A: 255}, // pink
		{R: 128, G: 128, B: 128, A: 255}, // gray
		{R: 128, G: 128, B: 0, A: 255},   // olive
		{R: 0, G: 255, B: 255, A: 255},   // cyan
	}

	for i, cluster := range clusters {
		pts := make(plotter.XYs, len(cluster))
		for j, pt := range cluster {
			pts[j].X = pt.X
			pts[j].Y = pt.Y
		}
		scatter, err := plotter.NewScatter(pts)
		if err != nil {
			panic(err)
		}
		scatter.GlyphStyle.Color = colors[i%len(colors)]
		p.Add(scatter)
	}

	centroidPts := make(plotter.XYs, len(centroids))
	for i, c := range centroids {
		centroidPts[i].X = c.X
		centroidPts[i].Y = c.Y
	}
	centroidScatter, err := plotter.NewScatter(centroidPts)
	if err != nil {
		panic(err)
	}
	centroidScatter.GlyphStyle.Shape = draw.CircleGlyph{}
	centroidScatter.GlyphStyle.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255} // black
	centroidScatter.GlyphStyle.Radius = vg.Points(4)
	p.Add(centroidScatter)

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "clusters.png"); err != nil {
		panic(err)
	}
}

// Extract relevant columns from CSV data
func extractPoints(records [][]string) ([]Point, error) {
	var points []Point
	for _, record := range records[1:] { // Skip header row
		finishing, err := strconv.ParseFloat(record[0], 64) // Adjust index as per column order
		if err != nil {
			return nil, err
		}
		headingAccuracy, err := strconv.ParseFloat(record[1], 64) // Adjust index as per column order
		if err != nil {
			return nil, err
		}
		points = append(points, Point{X: finishing, Y: headingAccuracy})
	}
	return points, nil
}

func clustering() {
	records := readCSV("./data/clustering/Soccer2019C.csv")
	points, err := extractPoints(records)
	if err != nil {
		log.Fatalf("failed to extract points: %v", err)
	}

	kmeans := NewKMeans(3, points) // Adjust K as needed
	clusters := kmeans.Fit()

	fmt.Println("Centroids:")
	for _, c := range kmeans.Centroids {
		fmt.Printf("(%.2f, %.2f)\n", c.X, c.Y)
	}

	plotClusters(clusters, kmeans.Centroids)
}

// HDBSCAN?
