package main

import (
	"fmt"
	"sync"
)

type single struct {
	//some elements
}

var singleInstance *single

var mutex = &sync.Mutex{}

func getInstance() *single {
	if singleInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if singleInstance == nil {
			fmt.Println("creating instance")
			singleInstance = &single{}
		} else {
			fmt.Println("already created")
		}

	} else {
		fmt.Println("singleton already created")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	fmt.Scanln() // to enter post last thread execution, else we have to add wg's
}

// There is a nil-check at the start for making sure singleInstance is empty first time around.
// This is to prevent expensive lock operations every time the getinstance method is called.
// If this check fails, then it means that the singleInstance field is already populated.

// The singleInstance struct is created within the lock.

// There is another nil-check after the lock is acquired.
// This is to ensure that if more than one goroutine bypasses the first check, only one goroutine can create the singleton instance.
// Otherwise, all goroutines will create their own instances of the singleton struct.
