package agg

type Stats struct {
	Sum   float64
	Count int
	Min   float64
	Max   float64
}

type Aggregator interface {
	Process(lines []string) *Stats
}
