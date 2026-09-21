//--------------------------------------------
// Author: Amelia Hamulewicz (C00296605@setu.ie)
// Created on 21/09/2025
// Modified by: Amelia Hamulewicz
// Group I worked with: Mark Lambert, Dorian Nowacki, Adam Noonan
//--------------------------------------------

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

func main() {
	var wg sync.WaitGroup
	// barrier := make(chan bool) //go routines wait on this channel
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N)
	}

	// all a's wait for b (Mark Lambert helped here)
	for range threadCount {
		<-aArrived
	}
	// all b's wait for a
	for range threadCount {
		bArrived <- true
	}

	wg.Wait() //wait here until everyone (10 go routines) is done

}
