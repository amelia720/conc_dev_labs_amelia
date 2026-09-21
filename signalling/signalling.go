//--------------------------------------------
// Author: Amelia Hamulewicz (C00296605@setu.ie)
// Created on 21/09/2025
// Modified by: Amelia Hamulewicz
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

//Global variables shared between functions --A BAD IDEA

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool) // create an unbuffered channel

	doStuffOne := func() bool {
		fmt.Println("StuffOne - Part A") // 1st print
		//wait here
		barrier <- true                 // Send a signal and wait until StuffTwo receives it
		fmt.Println("StuffOne - PartB") // 3rd print
		wg.Done()                       // mark 1 go routine as done in the wait group
		return true
	}
	doStuffTwo := func() bool {
		time.Sleep(time.Second * 5)
		fmt.Println("StuffTwo - Part A") // 2nd print
		//wait here

		<-barrier                       // receive signal from stuffOne (StuffTwo continues)
		fmt.Println("StuffTwo - PartB") // 4th print
		wg.Done()                       // mark 2 go routine as done in the wait group
		return true
	}
	wg.Add(2)       // waitgroup expects 2 go routines
	go doStuffOne() // perform stuffOne first
	go doStuffTwo() // perform stuffTwo next (after stuffOne sends its signal)
	wg.Wait()       //wait here until everyone (10 go routines) is done

}
