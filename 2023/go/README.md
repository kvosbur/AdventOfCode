To run benchmark run: `go test *.go -bench=. -benchtime=1000x`
To run benchmark with profiling: `go test -cpuprofile cpu.prof -memprofile mem.prof *.go -bench=. -benchtime=1x`
To run benchmark with profiling: `go test -cpuprofile cpu.prof -memprofile mem.prof day12/part2_benchmark_test.go -bench=. -benchtime=1x`
To run benchmark with profiling: `go test day12/part2_benchmark_test.go -bench=. -benchtime=1x`