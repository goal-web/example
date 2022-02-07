package tests

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

func TestBenchKafka(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(int2 int) {
			http.Get(fmt.Sprintf("http://localhost:8008/queue?info=bench-%d", int2))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func BenchmarkKafka(b *testing.B) {
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(int2 int) {
			http.Get(fmt.Sprintf("http://localhost:8008/queue?info=bench-%d", int2))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
