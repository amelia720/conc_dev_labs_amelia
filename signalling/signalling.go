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
//Uses an unbuffered channel so both go routines finish Part A before either starts Part B.
// --------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Start two go routines and use a channel to make both finish Part A before starting Part B
func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool) // create an unbuffered channel

	doStuffOne := func() bool {
		fmt.Println("StuffOne - Part A")
		//wait here
		barrier <- true // Send a signal and wait until StuffTwo receives it
		fmt.Println("StuffOne - PartB")
		wg.Done() // mark 1 go routine as done in the wait group
		return true
	}
	doStuffTwo := func() bool {
		time.Sleep(time.Second * 5)
		fmt.Println("StuffTwo - Part A")
		//wait here

		<-barrier // receive signal from stuffOne (StuffTwo continues)
		fmt.Println("StuffTwo - PartB")
		wg.Done() // mark 2 go routine as done in the wait group
		return true
	}
	wg.Add(2)       // waitgroup expects 2 go routines
	go doStuffOne() // perform stuffOne first
	go doStuffTwo() // perform stuffTwo next (after stuffOne sends its signal)
	wg.Wait()       //wait here until everyone (10 go routines) is done

}
