//Copyright (C) 2026 Amelia Hamulewicz

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
// Author: Amelia Hamulewicz (C00296605@setu.ie)
// Date: 11/10/2025
// Description:
// Uses a mutex so multiple go routines can safely add to the same counter.
// --------------------------------------------

package main

import (
	"fmt"
	"sync"
)

// Create events and add them to the buffer
func producer(theBuffer chan int, wg *sync.WaitGroup) {
	for i := range 10 {
		theBuffer <- i // wait here if buffer is full
	}
	close(theBuffer) // no more events will be added
	wg.Done()
}

// Take events from the buffer and consume them
func consumer(theBuffer chan int, wg *sync.WaitGroup) {
	for event := range theBuffer {
		fmt.Println("Consumed event", event)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	theBuffer := make(chan int, 20) // buffer holds up to 20 events

	wg.Add(2)
	go producer(theBuffer, &wg)
	go consumer(theBuffer, &wg)

	wg.Wait() // wait until both go routines finish
}
