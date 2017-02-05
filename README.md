# sync
Drop-in replacement for `sync`'s RWLock

# Benchmarks

```
~/gopath/src/github.com/gramework/sync % go test -bench . -benchmem
testing: warning: no tests to run
BenchmarkGrameworkRWMutex/1-4 			100000000	        22.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGrameworkRWMutex/4-4 			20000000	        64.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGrameworkRWMutex/16-4         	20000000	        64.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGrameworkRWMutex/64-4         	20000000	        64.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGrameworkRWMutex/256-4        	20000000	        65.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrigRWMutex/1-4               	100000000	        23.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrigRWMutex/4-4               	20000000	        70.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrigRWMutex/16-4              	20000000	        71.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrigRWMutex/64-4              	20000000	        73.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrigRWMutex/256-4             	20000000	        73.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/gramework/sync	16.126s
```

(note: first two lines was modified: I've put two tabs after the test name.)