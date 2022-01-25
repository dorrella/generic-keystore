package main

import (
	"fmt"
	"sort"
	"time"
)

// Wrapper to hold results for a series of benchmarks
type Result struct {
	duration []time.Duration
}

// one space for each time
func NewResult(size int) *Result {
	return &Result{
		duration: make([]time.Duration, size),
	}
}

// sorts times to calculate percentiles
func (r *Result) Percentiles(tests int) {
	// less function used as a parameter to sort function
	f := func(i, j int) bool {
		return r.duration[i].Nanoseconds() < r.duration[j].Nanoseconds()

	}

	sort.Slice(r.duration, f)

	//calculate each percentile
	index := 0 //index is start of each percentile
	for p := 0; p < 100; p++ {
		//this skips the 100th, but could be handled as
		//r.duration[len(r.duration) -1]

		percentile := ((tests - index) * 100) / tests
		next := index + 1
		for true {
			percentile2 := ((tests - next) * 100) / tests
			if percentile2 != percentile {
				break
			}
			next++
		}

		fastest := r.duration[index].Nanoseconds()
		slowest := r.duration[next-1].Nanoseconds()
		index = next + 1
		//trim down on output
		if (p > 5 && p < 10) || (p > 10 && p < 25) || (p > 25 && p < 50) {
			continue
		}

		if (p > 50 && p < 95) || (p > 95 && p < 99) {
			continue
		}
		fmt.Printf("p:%d start %d nanoseconds, end %d nanoseconds\n", p, fastest, slowest)
	}
}
