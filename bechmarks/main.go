// benchmarks for keyvalue store
package main

import "fmt"

func main() {
	//reads
	benchmarkIntReads(true)
	fmt.Println()
	benchmarkStringReads(true)
	fmt.Println()

	//writes
	benchmarkIntWrites(true)
	fmt.Println()
	benchmarkStringWrites(true)
	fmt.Println()

	//rate limited writes
	benchmarkStringBufferedWrites(true, 10000)
	fmt.Println()
	benchmarkStringBufferedWrites(true, 500)
}
