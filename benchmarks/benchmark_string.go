package main

import (
	"fmt"
	"sync"
	"time"

	keystore "github.com/dorrella/generic-keystore"
)

// string for benchmarking
// could do something fuzzy, but that might throw off the timing
var big_string = "abcdefghijklmnopkqrstuvwxyznowiknowmyabcsnexttimewontyousingwithme"

// benchmarks writing 10 million strings in parallel
// using uint64 keys and string values
func benchmarkStringWrites(verbose bool) {
	ks := keystore.NewKeyStore[uint64, string]()
	tests := 10000000 //ten million
	wg := sync.WaitGroup{}
	res := NewResult(tests)

	if verbose {
		fmt.Printf("benchmarking %d writes with strings\n", tests)
	}
	time_start := time.Now()

	for i := uint64(0); i < uint64(tests); i++ {
		wg.Add(1)
		go func(i uint64) {
			//must pass i for race conditions
			t := time.Now()
			ks.Put(i, big_string)
			res.duration[i] = time.Since(t)
			wg.Done()
		}(uint64(i))
	}

	wg.Wait()

	time_total := time.Since(time_start)
	if verbose {
		fmt.Printf("total time %d ms\n", time_total.Milliseconds())
		res.Percentiles(tests)
	}

}

// benchmarks writes, but limits parallel writes to
// bufflen using a channel based semaphor
func benchmarkStringBufferedWrites(verbose bool, buff_len int) {
	ks := keystore.NewKeyStore[uint64, string]()
	tests := 10000000 //ten million
	wg := sync.WaitGroup{}
	res := NewResult(tests)

	//semaphore to reduce threads/parallel writes
	var sem = make(chan struct{}, buff_len)

	if verbose {
		fmt.Printf("benchmarking %d writes with strings in batches of %d\n", tests, buff_len)
	}
	time_start := time.Now()

	for i := uint64(0); i < uint64(tests); i++ {
		wg.Add(1)
		go func(i uint64) {
			//get semaphor first
			sem <- struct{}{}

			t := time.Now()
			ks.Put(i, big_string)
			res.duration[i] = time.Since(t)
			wg.Done()

			//release semaphor token
			<-sem
		}(uint64(i))
	}

	wg.Wait()

	time_total := time.Since(time_start)
	if verbose {
		fmt.Printf("total time %d ms\n", time_total.Milliseconds())
		res.Percentiles(tests)
	}

}

// Benchmarks 10 million parallel reads
// using uint64 keys and string
func benchmarkStringReads(verbose bool) {
	ks := keystore.NewKeyStore[uint64, string]()
	tests := 10000000 //ten million
	wg := sync.WaitGroup{}
	res := NewResult(tests)

	if verbose {
		fmt.Printf("benchmarking %d reads with strings\n", tests)
	}

	//populate
	for i := uint64(0); i < uint64(tests); i++ {
		ks.Put(i, big_string)
	}

	time_start := time.Now()
	for i := 0; i < tests; i++ {
		wg.Add(1)
		go func(i uint64) {
			//must pass i for race conditions
			t := time.Now()
			_, _ = ks.Get(i)
			//check k == -v?
			res.duration[i] = time.Since(t)
			wg.Done()
		}(uint64(i))
	}

	wg.Wait()

	time_total := time.Since(time_start)
	if verbose {
		fmt.Printf("total time %d ms\n", time_total.Milliseconds())
		res.Percentiles(tests)
	}
}
