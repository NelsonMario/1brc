package agg

import (
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type OptAggregator struct {
	Workers int
}

func NewOptAggregator() Aggregator {
	return &OptAggregator{
		Workers: runtime.NumCPU(),
	}
}

func processChunk(lines []string, wg *sync.WaitGroup, statChan chan<- *Stats) {
	defer wg.Done()

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

	statChan <- stats
}

func mergeChunk(allStats []*Stats) *Stats {
	var final *Stats = &Stats{}

	for _, s := range allStats {
		final.Sum += s.Sum
		final.Count += s.Count

		if final.Count == s.Count {
			final.Min = s.Min
			final.Max = s.Max
		} else {
			if s.Min < final.Min {
				final.Min = s.Min
			}
			if s.Max > final.Max {
				final.Max = s.Max
			}
		}
	}

	return final
}

func (o *OptAggregator) Process(lines []string) *Stats {
	var wg sync.WaitGroup

	chunkSize := (len(lines) + o.Workers - 1) / o.Workers
	statsChannel := make(chan *Stats, o.Workers)

	for i := 0; i < len(lines); i += chunkSize {
		end := i + chunkSize
		if end > len(lines) {
			end = len(lines)
		}
		wg.Add(1)
		go processChunk(lines[i:end], &wg, statsChannel)
	}

	wg.Wait()
	close(statsChannel)

	var allStats []*Stats
	for stat := range statsChannel {
		allStats = append(allStats, stat)
	}

	return mergeChunk(allStats)
}
