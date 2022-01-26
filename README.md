# Generic KeyStore

Threadsafe Key/Value store written in go using generics.

## Note

Requires generics with at least version 1.18 Beta 1. [See here](https://go.dev/blog/go1.18beta1)

Go 1.18 beta can be installed or can be run in the Go playground dev branch mode.

## Install
Can be installed to /usr/env/go from the downloads page or


Beta can be installed with `go install golang.org/dl/go1.18beta1@latest`

and use `go1.8beta1` as an alias for go.

[See this](https://go.dev/doc/tutorial/generics#installing_beta)

# KeyStore Library

## API

### Create KeyStore

Create a reference to a keystore with uint32 keys and 
string values.

**Note:** keys must be comparable to be valid [See here](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#comparable-types-in-constraints)

```
ks := keystore.NewkeyStore[uint32, string]
```

### Put

Put a key/value pair in keystore.

```
ks := keystore.NewkeyStore[uint32, string]
ks.Put(1, "this string has key 1")
```

### PutExpires

Put a key/value pair that has an expiration value. Adds key/value, then
waits `d Duration` seconds and deletes them.

```
ks := keystore.NewkeyStore[uint32, string]
timer, err := time.ParseDuration("300s")
if err != nil {
    panic(err)
}

ks.PutExpires(1, "this string has key 1", timer)
```

### Get

Retrieve a key/value pair.

```
v, ok := ks.Get(1)

if !ok {
   panic("key not found")
}

fmt.Println(v)
```

### Delete

Delete a key/value pair.

```
ks := keystore.NewkeyStore[uint32, string]
ks.add(1, "this string has key 1")

///...
ks.delete(1)
```

## Performance

### Performance Targets

* Retrieve a key with a 95th percentile time of less than 1 millisecond
* Retrieve a key with a 99th percentile time of less than 5 milliseconds
* Handle up to 10 million key/value pairs

### Benchmarks

```
cd benchmarks
go run .

go test -v -bench=.
```

##### Results

```
benchmarking 10000000 reads with ints
total time 6290 ms
p:0 start 81 nanoseconds, end 81 nanoseconds
p:1 start 82 nanoseconds, end 151 nanoseconds
p:2 start 151 nanoseconds, end 155 nanoseconds
p:3 start 155 nanoseconds, end 158 nanoseconds
p:4 start 158 nanoseconds, end 161 nanoseconds
p:5 start 161 nanoseconds, end 163 nanoseconds
p:6 start 163 nanoseconds, end 166 nanoseconds
p:7 start 166 nanoseconds, end 169 nanoseconds
p:8 start 169 nanoseconds, end 171 nanoseconds
p:9 start 171 nanoseconds, end 174 nanoseconds
p:10 start 174 nanoseconds, end 177 nanoseconds
p:25 start 220 nanoseconds, end 221 nanoseconds
p:50 start 252 nanoseconds, end 253 nanoseconds
p:75 start 294 nanoseconds, end 297 nanoseconds
p:90 start 337 nanoseconds, end 342 nanoseconds
p:95 start 368 nanoseconds, end 378 nanoseconds
p:99 start 448 nanoseconds, end 523 nanoseconds

benchmarking 10000000 reads with strings
total time 6288 ms
p:0 start 79 nanoseconds, end 79 nanoseconds
p:1 start 82 nanoseconds, end 156 nanoseconds
p:2 start 156 nanoseconds, end 162 nanoseconds
p:3 start 162 nanoseconds, end 167 nanoseconds
p:4 start 167 nanoseconds, end 171 nanoseconds
p:5 start 171 nanoseconds, end 175 nanoseconds
p:6 start 175 nanoseconds, end 178 nanoseconds
p:7 start 178 nanoseconds, end 182 nanoseconds
p:8 start 182 nanoseconds, end 185 nanoseconds
p:9 start 185 nanoseconds, end 189 nanoseconds
p:10 start 189 nanoseconds, end 193 nanoseconds
p:25 start 234 nanoseconds, end 235 nanoseconds
p:50 start 270 nanoseconds, end 272 nanoseconds
p:75 start 315 nanoseconds, end 317 nanoseconds
p:90 start 359 nanoseconds, end 364 nanoseconds
p:95 start 394 nanoseconds, end 405 nanoseconds
p:99 start 468 nanoseconds, end 524 nanoseconds

benchmarking 10000000 writes with ints
total time 10490 ms
p:0 start 107 nanoseconds, end 107 nanoseconds
p:1 start 107 nanoseconds, end 217 nanoseconds
p:2 start 217 nanoseconds, end 241 nanoseconds
p:3 start 241 nanoseconds, end 253 nanoseconds
p:4 start 253 nanoseconds, end 264 nanoseconds
p:5 start 264 nanoseconds, end 276 nanoseconds
p:6 start 276 nanoseconds, end 290 nanoseconds
p:7 start 291 nanoseconds, end 308 nanoseconds
p:8 start 308 nanoseconds, end 323 nanoseconds
p:9 start 323 nanoseconds, end 336 nanoseconds
p:10 start 336 nanoseconds, end 347 nanoseconds
p:25 start 424 nanoseconds, end 429 nanoseconds
p:50 start 739 nanoseconds, end 769 nanoseconds
p:75 start 2688 nanoseconds, end 2909 nanoseconds
p:90 start 173215727 nanoseconds, end 239557534 nanoseconds
p:95 start 525304254 nanoseconds, end 646196462 nanoseconds
p:99 start 947351314 nanoseconds, end 1037640330 nanoseconds

benchmarking 10000000 writes with strings
total time 12402 ms
p:0 start 112 nanoseconds, end 112 nanoseconds
p:1 start 113 nanoseconds, end 214 nanoseconds
p:2 start 214 nanoseconds, end 236 nanoseconds
p:3 start 236 nanoseconds, end 248 nanoseconds
p:4 start 248 nanoseconds, end 257 nanoseconds
p:5 start 257 nanoseconds, end 264 nanoseconds
p:6 start 264 nanoseconds, end 271 nanoseconds
p:7 start 271 nanoseconds, end 278 nanoseconds
p:8 start 278 nanoseconds, end 287 nanoseconds
p:9 start 287 nanoseconds, end 296 nanoseconds
p:10 start 296 nanoseconds, end 308 nanoseconds
p:25 start 421 nanoseconds, end 427 nanoseconds
p:50 start 768 nanoseconds, end 805 nanoseconds
p:75 start 3798 nanoseconds, end 4967 nanoseconds
p:90 start 546270766 nanoseconds, end 665991939 nanoseconds
p:95 start 902864032 nanoseconds, end 1016694596 nanoseconds
p:99 start 1344537788 nanoseconds, end 1654484084 nanoseconds

benchmarking 10000000 writes with strings in batches of 10000
total time 13295 ms
p:0 start 107 nanoseconds, end 107 nanoseconds
p:1 start 107 nanoseconds, end 213 nanoseconds
p:2 start 213 nanoseconds, end 236 nanoseconds
p:3 start 236 nanoseconds, end 247 nanoseconds
p:4 start 247 nanoseconds, end 255 nanoseconds
p:5 start 255 nanoseconds, end 263 nanoseconds
p:6 start 263 nanoseconds, end 270 nanoseconds
p:7 start 270 nanoseconds, end 278 nanoseconds
p:8 start 278 nanoseconds, end 287 nanoseconds
p:9 start 287 nanoseconds, end 298 nanoseconds
p:10 start 298 nanoseconds, end 310 nanoseconds
p:25 start 403 nanoseconds, end 408 nanoseconds
p:50 start 552 nanoseconds, end 564 nanoseconds
p:75 start 1222 nanoseconds, end 1271 nanoseconds
p:90 start 3123 nanoseconds, end 3543 nanoseconds
p:95 start 850673 nanoseconds, end 1743444 nanoseconds
p:99 start 87962455 nanoseconds, end 189293598 nanoseconds

benchmarking 10000000 writes with strings in batches of 500
total time 15941 ms
p:0 start 109 nanoseconds, end 109 nanoseconds
p:1 start 111 nanoseconds, end 221 nanoseconds
p:2 start 221 nanoseconds, end 242 nanoseconds
p:3 start 242 nanoseconds, end 254 nanoseconds
p:4 start 254 nanoseconds, end 263 nanoseconds
p:5 start 263 nanoseconds, end 271 nanoseconds
p:6 start 271 nanoseconds, end 278 nanoseconds
p:7 start 278 nanoseconds, end 286 nanoseconds
p:8 start 286 nanoseconds, end 294 nanoseconds
p:9 start 294 nanoseconds, end 302 nanoseconds
p:10 start 302 nanoseconds, end 311 nanoseconds
p:25 start 419 nanoseconds, end 423 nanoseconds
p:50 start 532 nanoseconds, end 540 nanoseconds
p:75 start 1009 nanoseconds, end 1052 nanoseconds
p:90 start 1998 nanoseconds, end 2183 nanoseconds
p:95 start 5216 nanoseconds, end 15385 nanoseconds
p:99 start 4081450 nanoseconds, end 14893089 nanoseconds

goos: linux
goarch: amd64
pkg: github.com/dorrella/generic-keystore/bechmarks
cpu: AMD Phenom(tm) II X4 955 Processor
BenchmarkIntReads
BenchmarkIntReads-3       	       1	9331071360 ns/op
BenchmarkIntWrites
BenchmarkIntWrites-3      	       1	10246216549 ns/op
BenchmarkStringReads
BenchmarkStringReads-3    	       1	10106898313 ns/op
BenchmarkStringWrites
BenchmarkStringWrites-3   	       1	11995732892 ns/op
PASS
ok  	github.com/dorrella/generic-keystore/bechmarks	42.095s
```

## Design

The design uses a the `sync.RWMutex` to allow faster read access.
This bottle necks with a write happens during read, but assuming
reads are more likely than writes, this should be ok.

Data is stored in a basic map[K]v object.

Expirations are handled by creating a thread that sleeps a number
of seconds and then calls `Delete(key)`.

### Other Considerations

#### Sync.Map

The sync library also provides an `interface{}` based map that is
designed for caches. It is unclear if using that would be faster,
but the `sync` library documentation is clear that it is heavily
optimised for cache use.

#### Memoization

It may be possible to create a concurrent non-blocking cache similar to the
one described in "The Go Programming Language" by Donovan and Kernighan to
create a smaller non-blocking version of the `Get` calls. Whether that is
a net gain for read times is unclear.

#### Sharding

Sharding the map would allow the locks to be split across some number
of shards based on the key. This would mean each shard having a lock
instead of one for the whole keystore.

#### Batching Writes/Deletes

Benchmarks show that while read times are low, write times quickly bottleneck
in large numbers. Adding support for bulk writes/deletes should allow less
locking/unlocking in the datapath.

A messaging queue of some sort might also be useful.


#### RateLimiting Writes

Benchmarking shows that limiting writes greatly normalizes the wait times on writes.

### Profile

todo run go profile for bottlenecks

## Memory Usage


### Map[K]V

The golang hashmap is pretty good when it comes to performace, so it is expected to grow
fairly linearly in relation to the key/value sizes with access of `O(1 + n/k)` where `n/k`
is the load factor.

### Expirations

Starting a thread of each expiration has overhead, but golang was designed to run large number
of threads. It may be more effiecient to store the times as either `map[K]time.Time` or

```
type struct Expiration{
    key K
	Time time.Time
}

type KeyStore{
    a []*Expiration

```

would use one thread to iterate over for loop a or sort/wait/delete algorithm.
