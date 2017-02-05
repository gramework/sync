// Copyright 2017 github.com/bcmills and Kirill Danshin
// See github.com/golang/go/issues/17973

package sync

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkGrameworkRWMutex(b *testing.B) {
	for ng := 1; ng <= 256; ng <<= 2 {
		b.Run(fmt.Sprint(ng), func(b *testing.B) {
			var mu RWMutex
			mu.Lock()

			var wg sync.WaitGroup
			wg.Add(ng)

			n := b.N
			quota := n / ng

			for g := ng; g > 0; g-- {
				if g == 1 {
					quota = n
				}

				go func(quota int) {
					for i := 0; i < quota; i++ {
						mu.RLock()
						mu.RUnlock()
					}
					wg.Done()
				}(quota)

				n -= quota
			}

			if n != 0 {
				b.Fatalf("Incorrect quota assignments: %v remaining", n)
			}

			b.StartTimer()
			mu.Unlock()
			wg.Wait()
			b.StopTimer()
		})
	}
}

func BenchmarkOrigRWMutex(b *testing.B) {
	for ng := 1; ng <= 256; ng <<= 2 {
		b.Run(fmt.Sprint(ng), func(b *testing.B) {
			var mu sync.RWMutex
			mu.Lock()

			var wg sync.WaitGroup
			wg.Add(ng)

			n := b.N
			quota := n / ng

			for g := ng; g > 0; g-- {
				if g == 1 {
					quota = n
				}

				go func(quota int) {
					for i := 0; i < quota; i++ {
						mu.RLock()
						mu.RUnlock()
					}
					wg.Done()
				}(quota)

				n -= quota
			}

			if n != 0 {
				b.Fatalf("Incorrect quota assignments: %v remaining", n)
			}

			b.StartTimer()
			mu.Unlock()
			wg.Wait()
			b.StopTimer()
		})
	}
}
