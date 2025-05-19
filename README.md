# 1BRC Optimized Aggregator with Goroutines

This package implements an optimized solution to the 1 Billion Row Challenge (1BRC)] using Go and goroutines for parallel data processing.

## Description

The core idea is to:
1. **Split** the input lines into chunks.
2. **Process** each chunk in a separate goroutine to calculate sum, count, min, and max.
3. **Merge** all partial results into a final aggregated statistic.

## Features

- Utilizes `runtime.NumCPU()` to determine the number of worker goroutines.
- Handles parsing and aggregation in parallel.
- Efficient merging of statistics from all chunks.

## Input Format

Each line in the input should follow this format:

```
CityName;FloatValue
```

Example:
```
Tokyo;35.6897
Jakarta;-6.1750
```

## Output

Returns a `*Stats` struct containing:
- `Sum`: Total of all values
- `Count`: Number of entries
- `Min`: Minimum value
- `Max`: Maximum value

## Usage

Import the package and call the `Process` method:

```go
agg := agg.NewOptAggregator()
result := agg.Process(lines) // lines is a []string from your file
```

## Performance Note

This version significantly improves speed on large datasets by leveraging Go’s concurrency model. For small datasets, the overhead may outweigh the benefit compared to a sequential version.
