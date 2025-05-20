package main

import (
	"strconv"
	"strings"
)

type NaiveAggregator struct{}

func NewNaiveAggregator() Aggregator {
	return &NaiveAggregator{}
}

func (n *NaiveAggregator) Process(lines []string) *Stats {
	var stats *Stats = &Stats{}

	for _, line := range lines {
		parts := strings.Split(line, ";")
		if len(parts) != 2 {
			continue
		}

		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			panic(err)
		}

		stats.Sum += value
		stats.Count += 1
		if stats.Count == 1 || stats.Max < value {
			stats.Max = value
		}
		if stats.Count == 1 || stats.Min > value {
			stats.Min = value
		}
	}

	return stats
}
