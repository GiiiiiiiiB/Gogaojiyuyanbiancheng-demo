package main

import (
	"fmt"
	"sync"
)

type total struct {
	mu sync.Mutex
	int
}

func worker(wg *sync.WaitGroup, total *total) {
	defer wg.Done()
	for i := 0; i < 200; i++ {
		total.mu.Lock()
		total.int += 1
		total.mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	t := new(total)
	wg.Add(2)
	go worker(&wg, t)
	go worker(&wg, t)
	wg.Wait()
	fmt.Println(t.int)
}
