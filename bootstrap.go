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

// 生成正态分布随机数（Box-Muller）
func randNorm(mean, stddev float64) float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()
	z := math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	return z*stddev + mean
}

// 计算中位数
func median(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2.0
	}
	return data[n/2]
}

// 计算标准差
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

	rand.Seed(9999) // 固定种子，确保生成数据可复现

	n := 400  // ✅ 样本大小改为 400
	B := 1000 // bootstrap 次数
	popMean := 100.0
	popSD := 10.0

	// 1. 生成原始样本
	original := make([]float64, n)
	for i := 0; i < n; i++ {
		original[i] = randNorm(popMean, popSD)
	}

	// 2. 保存到 CSV 文件
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

	// 3. 执行 bootstrap 并输出中位数标准误差
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

	// 4. 记录总用时和内存
	runtime.ReadMemStats(&memEnd)
	duration := time.Since(startTime)
	memUsed := memEnd.Alloc - memStart.Alloc

	fmt.Printf("Total runtime: %d ms\n", duration.Milliseconds())
	fmt.Printf("Total memory used: %d bytes\n", memUsed)
}
