# Testing performance aspects

## Benchmarking

Benchmark tests run the code and collect data about how long it takes to run. The `go test -bench=.` command reports in nanoseconds. Usually, we run benchmarks with different values or configurations that will affect code performance. We should run benchmarks several times to confirm results are consistent (small variance). Ensure the number of executions for each run is enough to be statistically meaningful.

Benchmark tests are similar to unit tests. Filenames end in `_test` and benchmarks and unit tests can be in the same file.

Benchmark tests begin with `Benchmark` and accept a `*testing.B` (`(b *testing.B)` by convention). Benchmark tests look something like:

```golang
   // i could be any type, could be set outside the function, could be passed as a parameter
   func BenchmarkMyFunc(b *testing.B) {
      i := 100
      for n := 0; n < b.N; n++ {
         MyFunc(i)
      }
   }
```

The `-bench=.` parameter runs all benchmarks. Pass a value other than `.` to match a regex of test names (example: `go test -bench=My` to run benchmarks that include `My` in their name).

When running benchmark tests, `go test` runs the test, increasing the number of loops to run (`b.N`) until the benchmark runs for at least one second. You can set the minimum time with the `-benchtime` parameter (example, `go test -bench=. -benchtime=10s`).

You can use timer methods (`b.StartTimer`, `b.StopTimer`, `b.ResetTimer`) to control the benchmarking timer and exclude setup or other code that's out of scope of your benchmarking goals.

See the docs for more details on other options.

Reference links:

* [Benchmarking Golang to Improve Function Performance](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)
* [How to Write Benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

So, let's write a benchmark of encrypting with AES and RC4. See `014-PerfTesting/encrypt.go` and `.../encrypt_test.go`.

I confirmed the code runs with `go run .` Then I ran `go test -bench=.`

```
goos: linux
goarch: amd64
pkg: perfTest
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkEncryptAES-4             464145              2542 ns/op
BenchmarkEncryptRC4-4             325524              3386 ns/op
PASS
ok      perfTest        3.196s
```

Let's run each test ten times to see how variable the results are. `go test -bench=. -count=10` (I inserted the line break between the two types.)

```
goos: linux
goarch: amd64
pkg: perfTest
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkEncryptAES-4             436233              2446 ns/op
BenchmarkEncryptAES-4             465284              2474 ns/op
BenchmarkEncryptAES-4             401414              2595 ns/op
BenchmarkEncryptAES-4             395916              2569 ns/op
BenchmarkEncryptAES-4             421144              2513 ns/op
BenchmarkEncryptAES-4             434274              2590 ns/op
BenchmarkEncryptAES-4             463372              2540 ns/op
BenchmarkEncryptAES-4             482533              2489 ns/op
BenchmarkEncryptAES-4             470952              2538 ns/op
BenchmarkEncryptAES-4             484734              2520 ns/op

BenchmarkEncryptRC4-4             356094              3203 ns/op
BenchmarkEncryptRC4-4             337200              3346 ns/op
BenchmarkEncryptRC4-4             363393              3346 ns/op
BenchmarkEncryptRC4-4             365127              3239 ns/op
BenchmarkEncryptRC4-4             368600              3373 ns/op
BenchmarkEncryptRC4-4             322765              3226 ns/op
BenchmarkEncryptRC4-4             331986              3289 ns/op
BenchmarkEncryptRC4-4             308529              3399 ns/op
BenchmarkEncryptRC4-4             350782              3367 ns/op
BenchmarkEncryptRC4-4             346321              3384 ns/op
PASS
ok      perfTest        35.088s
```

Over 300,000 iterations per test sounds like a statistically valid sample size. Doing some math on those numbers shows the results are tightly clustered.

|         | AES    | RC4    |
| ------- | ------ | ------ |
| Average | 2537.4 | 3317.2 |
| StdDev  | 41.446 | 63.428 |
| SD/Avg  | 1.63%  | 1.91%  |
| Max     | 2595   | 3203   |
| Min     | 2446   | 3339   |
| Max-Min |  149   |  196   |

**COMMIT:** TEST: write basic benchmark tests and see how they run

## Application profiling

We can profile several things about the application.

* `-benchmem` -- memory allocation statistics
* `-trace traceFileName` -- execution traces
* `-<type>profile outFileName`
  * `-blockprofile` -- goroutine blocking
  * `-cpuprofile` -- cpu time
  * `-memprofile` -- more detailed that `-benchmem`
  * `-coverprofile` -- test coverage (text file; see `013-MoreTesting.md` notes on test coverage)
  * `-mutexprofile` -- mutex contention

The `-trace` and `<type>profile` options write binary files. Use `go tool pprof <filename>` which has a lot of options to analyze the file. Use `help` to see options`

Run `go test -bench=. -benchmem` and get the following output (environment info snipped), which adds bytes allocated per operation and number of memory allocations per operation (how many variables, basically). Understanding memory allocation patterns can be significant when tuning high performance applications because allocs take time.

```
BenchmarkEncryptAES-4             445838              2746 ns/op             816 B/op          8 allocs/op
BenchmarkEncryptRC4-4             258068              3887 ns/op            1392 B/op          2 allocs/op
PASS
ok      perfTest        3.307s
```

`go test -bench=. -memprofile mem.prof` then `go tool pprof mem.prof`.

In pprof, run `svg` to get a graph. Requires Graphviz, so `sudo apt-get install graphviz`. The graph shows memory allocations by function, where size of the box is relative size vs. others for allocs within the function. In the example, several small-box functions allocate little or no memory. `rc4.NewCipher` is the biggest single allocator (~690MB). While `encryptRC4()` calls it, so is responsible for ~750MB of allocation, it only allocs ~128MB itself, so is smaller. The graph also shows percentages. ([Intepreting the Callgraph](https://github.com/google/pprof/blob/main/doc/README.md#interpreting-the-callgraph))

Run `text` to get a table like below. The `flat` and `flat%` columns are local allocation, `cum` and `cum%` is total including called methods (example, `perfTest.encryptRC4` `cum` includes `crypto/rc4.NewCipher`) and `sum` is total percentage (so we see the two methods in `crypto` account for ~71% of all memory allocated).raw

```
Showing nodes accounting for 1269.81MB, 100% of 1269.81MB total
Showing top 10 nodes out of 11
      flat  flat%   sum%        cum   cum%
  630.69MB 49.67% 49.67%   630.69MB 49.67%  crypto/rc4.NewCipher
  281.54MB 22.17% 71.84%   281.54MB 22.17%  crypto/aes.newCipher
  161.54MB 12.72% 84.56%   510.58MB 40.21%  perfTest.encryptAES
  128.53MB 10.12% 94.68%   759.22MB 59.79%  perfTest.encryptRC4
   67.50MB  5.32%   100%    67.50MB  5.32%  crypto/cipher.newCFB
         0     0%   100%   281.54MB 22.17%  crypto/aes.NewCipher
         0     0%   100%    67.50MB  5.32%  crypto/cipher.NewCFBEncrypter (inline)
         0     0%   100%   510.58MB 40.21%  perfTest.BenchmarkEncryptAES
         0     0%   100%   759.22MB 59.79%  perfTest.BenchmarkEncryptRC4
         0     0%   100%  1269.81MB   100%  testing.(*B).launch
```

`go test -bench=. -cpuprofile cpu.prof` gives similar data for cpu time. Looking at this profile (text) we see `runtime.mallocgc` is the second largest total consumer of time (460ms), which points to memory allocation cost in time terms.

**COMMIT:** TEST: collect and analyze performance testing statistics
