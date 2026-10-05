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
// Date modifed: 21/09/2026
// Description:
// Uses a barrier struct to make all go routines wait until everyone finishes Part A... then starts Part B.
// --------------------------------------------
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Create a barrier data type
type barrier struct {
	theChan chan bool
	theLock sync.Mutex
	total   int
	count   int
}

// creates a properly initialised barrier
// N== number of threads (go Routines)
func createBarrier(N int) barrier {
	theBarrier := barrier{
		theChan: make(chan bool),
		total:   N,
		count:   0,
	}
	return theBarrier
}

// Method belonging to barrier data type
// Blocks until everyone reaches this point then lets everyone continue
func (b *barrier) wait() {
	b.theLock.Lock()
	b.count++
	if b.count == b.total {
		b.theLock.Unlock()
		fmt.Println("here")
		for _ = range b.total - 1 {
			<-b.theChan
		}
	} else {
		fmt.Println(b.count)
		b.theLock.Unlock()
		b.theChan <- true
	}
} //wait
// Start a go routine using the shared barrier
func WorkWithRendezvous(wg *sync.WaitGroup, Num int, theBarrier *barrier) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here
	theBarrier.wait()
	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

// Start five go routines and wait for them to finish
func main() {
	var wg sync.WaitGroup
	barrier := createBarrier(5)
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, &barrier)
	}
	wg.Wait() //wait here until everyone (5 go routines) is done

}
