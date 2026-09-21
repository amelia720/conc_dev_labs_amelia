//--------------------------------------------
// Author: Amelia Hamulewicz (C00296605@setu.ie)
// Created on 21/09/2025
// Modified by: Amelia Hamulewicz
// Group I worked with: Mark Lambert, Dorian Nowacki, Adam Noonan
// Issues:
// The barrier is implemented!
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

var count int // count completed go routines
var countLock sync.Mutex
var barrier = semaphore.NewWeighted(1) //
var ctx2 = context.TODO()              // Create the context used by the barrier.

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup) bool {
	time.Sleep(time.Second)      // wait 1 second as part A does its work
	fmt.Println("Part A", goNum) // print that part A finished
	//we wait here until everyone has completed part A
	countLock.Lock() // Lock before reading or changing count (Mark Lambert helped me with this)
	count++          //increment count
	if count == 10 { // Check if this is the final goroutine.
		countLock.Unlock() // unlock before continuing
		barrier.Release(1) // create the first permit
	} else {
		countLock.Unlock()       // unlock before waiting
		barrier.Acquire(ctx2, 1) // wait for the permit
		barrier.Release(1)       // pass the permit to the next goroutine
	}
	fmt.Println("PartB", goNum) // This can only run after all goroutines complete Part A.
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup // create a wait group
	wg.Add(totalRoutines) // add 10 go routines to the wait group
	//we will need some of these
	ctx := context.TODO() //create context and store it in ctx
	// Mark Lambert explained that the context is basically
	// something that gets the compiler not to throw an error
	var theLock sync.Mutex                             // create the mutex
	sem := semaphore.NewWeighted(int64(totalRoutines)) // create a weighted semaphore
	theLock.Lock()                                     // lock before reading or changing go routines
	sem.Acquire(ctx, 1)
	barrier.Acquire(ctx2, 1)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg) // i is the goroutine number and &wg gives it the WaitGroup.
	}
	sem.Release(1)   //return the permit
	theLock.Unlock() // unlock after creating the go routines

	wg.Wait() //wait for everyone to finish before exiting
}
