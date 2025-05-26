package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"time"
)

// Generate normally distributed random number (Box-Muller transform)
func randNorm(mean, stddev float64) float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()
	z := math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	return z*stddev + mean
}

// Calculate median
func median(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2.0
	}
	return data[n/2]
}

// Calculate standard deviation
func stdDev(data []float64) float64 {
	n := float64(len(data))
	mean := 0.0
	for _, v := range data {
		mean += v
	}
	mean /= n

	sumSq := 0.0
	for _, v := range data {
		sumSq += (v - mean) * (v - mean)
	}
	return math.Sqrt(sumSq / (n - 1))
}

func main() {
	startTime := time.Now()
	var memStart, memEnd runtime.MemStats
	runtime.ReadMemStats(&memStart)

	rand.Seed(9999)

	n := 400
	B := 1000
	popMean := 100.0
	popSD := 10.0

	// Generate original sample
	original := make([]float64, n)
	for i := 0; i < n; i++ {
		original[i] = randNorm(popMean, popSD)
	}

	// Save sample to CSV file
	file, err := os.Create("original_data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range original {
		writer.Write([]string{fmt.Sprintf("%.5f", value)})
	}

	fmt.Println("Saved original sample to original_data.csv")

	// Perform bootstrap and estimate standard error of the median
	medians := make([]float64, B)
	for b := 0; b < B; b++ {
		sample := make([]float64, n)
		for i := 0; i < n; i++ {
			index := rand.Intn(n)
			sample[i] = original[index]
		}
		medians[b] = median(sample)
	}

	seMedian := stdDev(medians)
	fmt.Printf("Estimated SE of Median (Go): %.5f\n", seMedian)

	// Record total runtime and memory usage
	runtime.ReadMemStats(&memEnd)
	duration := time.Since(startTime)
	memUsed := memEnd.Alloc - memStart.Alloc

	fmt.Printf("Total runtime: %d ms\n", duration.Milliseconds())
	fmt.Printf("Total memory used: %d bytes\n", memUsed)
}
