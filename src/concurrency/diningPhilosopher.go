//Dining philosopher
/* Imagine you have five philosophers sitting around a roundtable.
The philosophers do only two kinds of activities. One: they contemplate, and two: they eat.

However, they have only five forks between themselves to eat their food with.
Each philosopher requires both the fork to his left and the fork to his right to eat his food.

Design a solution where each philosopher gets a chance to eat his food without causing a deadlock.*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// Philosophers class
type philosopher struct {
	name                    string
	leftFork, rightFork     *fork
	leftForkId, rightForkId string
}

type fork struct{ sync.Mutex }

var wg sync.WaitGroup

func main() {
	// Decide number of philosopher and forks
	numberOfPhilopsopherAndForks := 5

	// Create philospoher, assign them 2 forks and send them to the dining table
	philosophers := make([]*philosopher, numberOfPhilopsopherAndForks)

	// create forks
	forks := make([]*fork, numberOfPhilopsopherAndForks)
	for i := 0; i < numberOfPhilopsopherAndForks; i++ {
		forks[i] = new(fork)
	}

	for i := 0; i < numberOfPhilopsopherAndForks; i++ {
		philosophers[i] = &philosopher{
			name:        fmt.Sprintf("phil %d", i),
			leftFork:    forks[i],
			rightFork:   forks[(i+1)%numberOfPhilopsopherAndForks],
			leftForkId:  fmt.Sprintf("%d", i),
			rightForkId: fmt.Sprintf("%d", (i+1)%numberOfPhilopsopherAndForks),
		}
		wg.Add(1)
		go philosophers[i].eat()

	}
	wg.Wait()

}

func (ph *philosopher) eat() {
	fmt.Println(ph.name, "Started :")
	defer wg.Done()
	for j := 0; j < 5; j++ {
		fmt.Println("\t", ph.name, "waiting for left fork", ph.leftForkId)
		ph.leftFork.Lock()
		fmt.Println("\t", ph.name, "waiting for right fork", ph.rightForkId)
		ph.rightFork.Lock()
		fmt.Println("\t\t", ph.name, "is eating")
		time.Sleep(time.Second)
		fmt.Println("\t\t", ph.name, " :thinking")
		ph.rightFork.Unlock()
		ph.leftFork.Unlock()

	}
	fmt.Println(ph.name, "Finished :")
}
// Problem simulation
