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

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by:
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *atomic.Int32, max int, wg *sync.WaitGroup, theChan chan bool, secondChan chan bool) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A
	if arrived.Add(1) == int32(max) { // increment our counter and check if its the last go routine (10)
		<-secondChan    //remove signal from outer gate
		theChan <- true //open 2nd door
		<-theChan       // remove signal from 2nd door
		theChan <- true //put signal back (aka close door)
	} else { //not all here yet we wait until signal
		<-theChan       //stop here until the last arrival opens the gate
		theChan <- true //once we get through send signal to next routine to continue
	} //end of if-else

	if arrived.Add(-1) == 0 { // everyone has now passed the first gate
		<-theChan //remove signal (close it for the next round)
		secondChan <- true
		<-secondChan
		secondChan <- true
	} else { //not all here yet we wait until signal
		<-secondChan
		secondChan <- true //once we get through send signal to next routine to continue
	} //end of if-else
	fmt.Println("PartB", goNum)
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	var arrived atomic.Int32 //shared counter... starting value is 0
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	theChan := make(chan bool, 1)    //use buffered channel with 1 storage slot (inner)
	secondChan := make(chan bool, 1) //use buffered channel with 1 storage slot (outer)

	secondChan <- true             //hand key to outer channel
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, totalRoutines, &wg, theChan, secondChan)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
