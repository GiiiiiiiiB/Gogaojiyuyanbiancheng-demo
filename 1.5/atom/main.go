package main

import (
	"fmt"
	"sync"
)

type total struct {
	mu sync.Mutex
	int
}

var t = new(total)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 200; i++ {
		t.mu.Lock()
		t.int += 1
		t.mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(t.int)
}
