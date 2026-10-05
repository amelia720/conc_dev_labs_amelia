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
// Uses a semaphore made with a buffered channel so only five tasks can run at a time.
// --------------------------------------------
package main

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

// Acquire a slot, wait if the semaphore is full
func Acquire(sem *semaphore) { //Dorian Nowacki helped me with this
	sem.theCounter <- struct{}{}
}

// Release a slot so another task can run
func Release(sem *semaphore) {
	<-sem.theCounter
}

// Start 20 tasks, allowing up to five to run at once
func main() {
	maxGoroutines := 5
	sem := Init(maxGoroutines) // set semaphore limit to 5

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ { //loop 20 times
		wg.Add(1) // add 1 task to wait group

		//Wait for a free slot, run the task, then release the slot
		go func(i int) {
			defer wg.Done()    // mark as done at the end of the go function
			Acquire(sem)       // take a slot, wait if sem is full
			defer Release(sem) //release a slot when function ends

			// Simulate a task
			fmt.Printf("Running task %d\n", i) // print which task number is running
			time.Sleep(2 * time.Second)        // sleep for 2 seconds
		}(i)
	}
	wg.Wait() // wait for every task to finish
}
