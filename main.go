package main

import (
	"bufio"
	"fmt"
	"getting-started/agg"
	"os"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("weather_stations.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var agg agg.Aggregator = agg.NewNaiveAggregator()

	final := agg.Process(lines)

	avg := final.Sum / float64(final.Count)
	fmt.Printf("Total Sum: %.4f\n", final.Sum)
	fmt.Printf("Average: %.4f\n", avg)
	fmt.Printf("Min: %.4f\n", final.Min)
	fmt.Printf("Max: %.4f\n", final.Max)

	elapsed := time.Since(start)
	fmt.Printf("Program ran in: %s\n", elapsed)
}
