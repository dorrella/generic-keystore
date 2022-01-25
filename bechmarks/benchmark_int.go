package main

import (
	"fmt"
	"sync"
	"time"

	keystore "github.com/dorrella/generic-keystore"
)

// Benchmark 10 million writes using key int
// and key values
func benchmarkIntWrites(verbose bool) {
	ks := keystore.NewKeyStore[int, int]()
	tests := 10000000 //ten million
	wg := sync.WaitGroup{}
	res := NewResult(tests)

	if verbose {
		fmt.Printf("benchmarking %d writes with ints\n", tests)
	}
	time_start := time.Now()

	for i := 0; i < tests; i++ {
		wg.Add(1)
		go func(i int) {
			//must pass i for race conditions
			t := time.Now()
			ks.Put(i, -i)
			res.duration[i] = time.Since(t)
			wg.Done()
		}(i)
	}

	wg.Wait()

	time_total := time.Since(time_start)
	if verbose {
		fmt.Printf("total time %d ms\n", time_total.Milliseconds())
		res.Percentiles(tests)
	}

}

// benchmark 10 million reads using int key
// and int values
func benchmarkIntReads(verbose bool) {
	//ten million
	ks := keystore.NewKeyStore[int, int]()
	tests := 10000000
	wg := sync.WaitGroup{}
	res := NewResult(tests)

	if verbose {
		fmt.Printf("benchmarking %d reads with ints\n", tests)
	}

	//populate
	for i := 0; i < tests; i++ {
		ks.Put(i, -i)
	}

	time_start := time.Now()
	for i := 0; i < tests; i++ {
		wg.Add(1)
		go func(i int) {
			//must pass i for race conditions
			t := time.Now()
			_, _ = ks.Get(i)
			//check k == -v?
			res.duration[i] = time.Since(t)
			wg.Done()
		}(i)
	}

	wg.Wait()

	time_total := time.Since(time_start)
	if verbose {
		fmt.Printf("total time %d ms\n", time_total.Milliseconds())
		res.Percentiles(tests)
	}
}
