package main

import (
	"fmt"
	"sync"
)

// boundeBuffer

type boundedBuffer struct {
	buffer []int
	cap    int
	m      sync.Mutex
	c      sync.Cond
}

//producer
func producer(bb *boundedBuffer, item int, wg *sync.WaitGroup) {
	defer wg.Done()
	// write into buffer
	bb.c.L.Lock()
	defer bb.c.L.Unlock()
	for len(bb.buffer) == bb.cap {
		fmt.Println("Producer waiting")
		bb.c.Wait()
	}
	bb.buffer = append(bb.buffer, item)
	fmt.Println("producing to buffer", item)
	bb.c.Signal()
}

// consumer
func consumer(bb *boundedBuffer, wg *sync.WaitGroup) {
	defer wg.Done()
	bb.c.L.Lock()
	defer bb.c.L.Unlock()
	// take off from buffer
	for len(bb.buffer) == 0 {
		fmt.Println("consumer waiting")
		bb.c.Wait()
	}
	item := bb.buffer[0]
	bb.buffer = bb.buffer[1:]
	fmt.Println("consuming from buffer", item)
	bb.c.Signal()
}

func initBuffer(size int) *boundedBuffer {
	bb := new(boundedBuffer)
	bb.cap = size
	bb.c = sync.Cond{L: &bb.m}
	return bb
}

func main() {
	bb := initBuffer(10)
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go producer(bb, i, &wg)
	}
	for i := 1; i < 3; i++ {
		wg.Add(1)
		go consumer(bb, &wg)
	}
	wg.Wait()
}
