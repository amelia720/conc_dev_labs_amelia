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
// Date modifed: 21/09/2025
// Description:
// Uses a mutex so multiple go routines can safely add to the same counter.
// --------------------------------------------

package main

import (
	"fmt"
	"sync"
)

// use a mutex to safely add to the shared counter
func adds(n int, theLock *sync.Mutex, total *int64, wg *sync.WaitGroup) bool {
	for i := 0; i < n; i++ {
		theLock.Lock()
		*total++
		theLock.Unlock()
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {

	//theLock will be passed by reference between go routines
	var theLock sync.Mutex
	var wg sync.WaitGroup
	var total int64

	// init it to number of go routines
	wg.Add(10)

	//for loop using range option
	for i := range 10 {
		//starting
		fmt.Println(i)
		go adds(1000, &theLock, &total, &wg)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total)
}
