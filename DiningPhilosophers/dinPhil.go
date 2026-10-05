//Barrier.go Template Code
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
// Date modifed: 05/10/2024
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

// Philosopher is thinking for a random amount of time
func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5)) // set random number from 1-5
	time.Sleep(X * time.Second)     //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

// Philosopher is eating for a random amount of time
func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5)) // set random number from 1-5
	time.Sleep(X * time.Second)     //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

// Acquires 2 forks needed by a philosopher
func getForks(index int, forks map[int]chan bool) {
	if index == 0 { //if you/re the first philosopher
		forks[(index+1)%5] <- true //philosopher 0 takes fork 1
		forks[index] <- true       //philosopher 0 takes fork 0 (his own fork)
	} else {
		forks[index] <- true       // take your fork
		forks[(index+1)%5] <- true // take the second-closest fork to you
	}
}

// release both forks so other philosophers can use them
func putForks(index int, forks map[int]chan bool) {
	<-forks[index]
	<-forks[(index+1)%5]
}

// This function repeatedly makes a philosopher think, get forks, eat, and put forks back onto the table.
func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	for {
		think(index)
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1)
	} //set up forks
	for N := range philCount {
		go doPhilStuff(N, &wg, forks)
	} //start philosophers
	wg.Wait() //wait here until everyone (10 go routines) is done

} //main
