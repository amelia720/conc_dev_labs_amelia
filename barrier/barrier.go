//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// --------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Amelia Hamulewicz (C00296605@setu.ie)
// Group I worked with: Mark Lambert, Dorian Nowacki, Adam Noonan
// Date modifed: 21/09/2026
// Description:
// Use a barrier so that all go routines print Part A before they print Part B
// --------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Complete Part A, wait at barrier, and then complete Part B
func doStuff(goNum int, wg *sync.WaitGroup, count *int, countLock *sync.Mutex, barrier *semaphore.Weighted, ctx2 context.Context) bool {
	time.Sleep(time.Second)      // wait 1 second as part A does its work
	fmt.Println("Part A", goNum) // print that part A finished
	// Only one go routine can increment count at a time....
	countLock.Lock() // Lock before reading or changing count (Mark Lambert helped me with this)

	*count++
	if *count == 10 { // All 10 go routines have completed part A.... so, open the barrier and continue to Part B
		countLock.Unlock()
		barrier.Release(1)
	} else {
		countLock.Unlock()
		barrier.Acquire(ctx2, 1) // wait until last go routine opens the barrier
		barrier.Release(1)       // pass the permit to the next goroutine
	}
	fmt.Println("PartB", goNum) // This can only run after all goroutines complete Part A.
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines) // Add 10 go routines to the wait group

	ctx := context.TODO() // Create context that will be used by sem
	// Mark Lambert explained to me that the context is basically something that gets the compiler not to throw an error

	var count int // number of go routines that have completed Part A
	var countLock sync.Mutex
	var barrier = semaphore.NewWeighted(1)
	var ctx2 = context.TODO() // Create the context used by the barrier.

	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))

	theLock.Lock() // lock before reading or changing go routines
	sem.Acquire(ctx, 1)
	barrier.Acquire(ctx2, 1)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, &count, &countLock, barrier, ctx2)
	}
	sem.Release(1)   //return the permit
	theLock.Unlock() // unlock after creating the go routines

	wg.Wait() //wait for everyone to finish before exiting
}
