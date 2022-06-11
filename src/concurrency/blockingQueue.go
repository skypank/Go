package main

import (
	"fmt"
	"sync"
)

type BlockingQueue struct {
	queue []interface{}
	m     sync.Mutex
	cond  sync.Cond
	cap   int
}

func createBQ(capacity int) *BlockingQueue {
	q := new(BlockingQueue)
	q.cond = sync.Cond{L: &q.m}
	q.cap = capacity
	return q
}

func (bq *BlockingQueue) push(item interface{}, wg *sync.WaitGroup) {
	bq.cond.L.Lock()
	defer bq.cond.L.Unlock()
	for bq.isFull() {
		fmt.Println("Blocked for pushing", item)
		bq.cond.Wait()

	}
	bq.queue = append(bq.queue, item)
	fmt.Println("Pushed", item)
	bq.cond.Signal()
	wg.Done()
}

func (bq *BlockingQueue) pop(wg *sync.WaitGroup) interface{} {
	bq.cond.L.Lock()
	defer bq.cond.L.Unlock()
	for bq.IsEmpty() {
		fmt.Println("Blocked for pop")
		bq.cond.Wait()
	}
	topElement := bq.queue[0]
	bq.queue = bq.queue[1:]
	bq.cond.Signal()
	fmt.Println("Poped", topElement)
	wg.Done()
	return topElement

}

func (bq *BlockingQueue) isFull() bool {
	if len(bq.queue) == bq.cap {
		return true
	}
	return false
}
func (bq *BlockingQueue) IsEmpty() bool {
	if len(bq.queue) == 0 {
		return true
	}
	return false
}

func main() {
	// create a blocking queue of capacity 5
	queue := createBQ(5)
	wg := sync.WaitGroup{}
	for i := 0; i < 9; i++ {
		wg.Add(1)
		go queue.push(i, &wg)
	}
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go queue.pop(&wg)
	}
	wg.Wait()

}
