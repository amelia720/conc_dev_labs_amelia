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
// Uses two channels so all go routines finish Part A before starting Part B.
// --------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Global variables shared between functions --A BAD IDEA
var aArrived = make(chan bool) // signals that a goroutine finished Part A
var bArrived = make(chan bool) // allows a goroutine to start Part B

// print Part A, signal, then print Part B
func WorkWithRendezvous(wg *sync.WaitGroup, Num int) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5)) // choose a random wait time from 0 to 4 seconds
	time.Sleep(X * time.Second)     //wait random time amount
	fmt.Println("Part A", Num)

	//Rendezvous here
	aArrived <- true // tell main that this goroutine finished Part A
	<-bArrived       // wait until all goroutines finish Part A

	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

// Start five go routines and let them continue once all finish Part A
func main() {
	var wg sync.WaitGroup
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount { //starts go routines... doesnt wait for them to finish
		go WorkWithRendezvous(&wg, N)
	}

	// all a's wait for b (Mark Lambert helped here)
	for range threadCount {
		<-aArrived //wait for one go routine to signal that it finsihed part A
	}
	// all b's wait for a
	for range threadCount {
		bArrived <- true //send signal (allow one go routine start part B
	}

	wg.Wait() //wait here until everyone (10 go routines) is done

}
