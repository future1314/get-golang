package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	counts 	map[string]int
	mux 	sync.Mutex
}

func (s *SafeCounter) Inc(key string) {
	s.mux.Lock()
	s.counts[key]++
	s.mux.Unlock()
}

func (s *SafeCounter) Get(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.counts[key]
}

func main() {

	s := SafeCounter{counts: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go s.Inc("zhangsan")
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println(s.Get("zhangsan"))
}
