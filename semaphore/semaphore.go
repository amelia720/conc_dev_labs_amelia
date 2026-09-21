package main

//--------------------------------------------
// Author: Amelia Hamulewicz (C00296605@setu.ie)
// Created on 21/09/2025
// Modified by: Amelia Hamulewicz
// Group I worked with: Mark Lambert, Dorian Nowacki, Adam Noonan
//--------------------------------------------

import (
	"fmt"
	"sync"
	"time"
)

// make struct containing channel
// add init, acquire and release
type semaphore struct {
	theCounter chan struct{}
}

// Init creates the semaphore and sets its maximum size.
func Init(max int) *semaphore {
	return &semaphore{
		theCounter: make(chan struct{}, max),
	}
}

func Acquire(sem *semaphore) { //Dorian Nowacki helped me with this
	sem.theCounter <- struct{}{} // send a signal into the channel
}
func Release(sem *semaphore) {
	<-sem.theCounter // receive the signal from the channel
}

func main() {
	maxGoroutines := 5
	sem := Init(maxGoroutines) // create a buffered channel struct that acts
	// as a semaphore that allows up to 5 goroutines

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ { //loop 20 times
		wg.Add(1) // add 1 task to wait group
		go func(i int) {
			defer wg.Done() // mark as done at the end of the go function
			Acquire(sem)    // send a signal into the channel
			defer Release(sem)

			// Simulate a task
			fmt.Printf("Running task %d\n", i) // print which task number is running
			time.Sleep(2 * time.Second)        // sleep for 2 seconds
		}(i)
	}
	wg.Wait() // wait for every task to finish
}
