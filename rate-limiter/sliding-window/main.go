package main

import (
	"fmt"
	"time"
)

func main() {
	limit := 2
	window := time.Minute
	var timestamps []time.Time

	requestTime := []time.Duration{
		0,                 // 0s
		20 * time.Second,  // 20s
		40 * time.Second,  // 40s
		60 * time.Second,  // 60s (1m)
		80 * time.Second,  // 80s (1m20s)
		100 * time.Second, //100s (1m40s)
	}

	start := time.Now()

	for i, delay := range requestTime {
		now := start.Add(delay)

		cutoff := now.Add(-window)
		fmt.Printf("\n[%d] CUTOFF %v\n", i+1, formatTimestamp(cutoff))
		valid := make([]time.Time, 0)
		for _, ts := range timestamps {
			if ts.After(cutoff) {
				valid = append(valid, ts)
			}
		}
		timestamps = valid

		fmt.Printf("\n[%d] Request at %v\n", i+1, delay)
		fmt.Printf("Timestamps before add: %v\n", formatTimestamps(timestamps))

		if len(timestamps) >= limit {
			fmt.Println("REJECTED (limit reached)")
			continue
		}

		timestamps = append(timestamps, now)
		fmt.Println("ALLOWED")
		fmt.Printf("Timestamps after add : %v\n", formatTimestamps(timestamps))
	}
}
func formatTimestamp(t time.Time) string {
	formatted := fmt.Sprintf("%ds", int(time.Since(t).Seconds()*-1))

	return formatted
}

func formatTimestamps(t []time.Time) []string {
	formatted := make([]string, len(t))
	for i, ts := range t {
		formatted[i] = fmt.Sprintf("%ds", int(time.Since(ts).Seconds()*-1))
	}

	return formatted
}
