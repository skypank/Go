// print odd even
package main

import (
	"fmt"
	"sync"
)

func printEvenNum(wg *sync.WaitGroup, msgEven, msgOdd chan bool) {
	fmt.Println("printEven")
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		// receive communication on even channel
		val, status := <-msgEven
		// if communication is received, print 
		if val {
			fmt.Println("even:", i)
		}
		// if channel is open, communicate back to odd channel
		if status {
			msgOdd <- true
		} else {
			fmt.Println("closing odd")
			// if not, close the odd channel also
			close(msgOdd)
		}
	}
}
func printOddNum(wg *sync.WaitGroup, msgOdd, msgEven chan bool) {
	fmt.Println("printOdd")
	defer wg.Done()
	for i := 1; i < 10; i += 2 {
		// wait for communication on this channel
		if <-msgOdd {
			fmt.Println("odd:", i)
			if i == 9 {
				fmt.Println("closing even")
				// when printing is done, close even channel
				close(msgEven)
			} else {
				// communicate to even channel
				msgEven <- true
			}
		}
	}
}
func main() {
	msgOdd := make(chan bool)
	msgEven := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go printEvenNum(&wg, msgEven, msgOdd)
	go printOddNum(&wg, msgOdd, msgEven)
	// trigger the odd channel, that will start communication
	msgOdd <- true
	wg.Wait()
}
