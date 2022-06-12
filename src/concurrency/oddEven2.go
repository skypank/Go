// print odd even
package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func recoverPanic(wg *sync.WaitGroup) {
	wg.Done()
	if err := recover(); err != nil {
		fmt.Println("send on closed channel occured, but recoverd")
	}
}
func printNumbers(start int, inChan <-chan bool, outChan chan<- bool, wg *sync.WaitGroup) {
	// since we are printing this list forever, to end printing, we have to close the channel
	// which will cause a panic, we can catch the panic and close program gracefully
	// recover call has to be deffered always
	defer recoverPanic(wg)
	for range inChan {
		fmt.Println(start)
		start += 2
		time.Sleep(time.Second)
		outChan <- true
	}
}

func main() {
	msgOdd := make(chan bool)
	msgEven := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)
	// trigger a routine, where thread will wait on even channel,
	// once it recieves a msg, it will print even number and communicate to odd channel
	go printNumbers(0, msgEven, msgOdd, &wg)
	// this thread will wait on odd channel, once it recieves a msg on it, it will print odd number and communicate back to even channel
	go printNumbers(1, msgOdd, msgEven, &wg)
	// trigger the channel, that will start the first msg on even channel
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press enter to cancel")
	fmt.Println("---------------------")
	msgEven <- true
	reader.ReadString('\n')
	fmt.Println("finished")
	close(msgEven)
	close(msgOdd)
	wg.Wait()
}
