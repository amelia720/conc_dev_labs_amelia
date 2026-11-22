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
// Date: 01/11/2025
// Description:
// Uses mutexes to allow multiple readers or one writer at a time.
// --------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Read the shared data
func reader(id int, readerLock *sync.Mutex, writerLock *sync.Mutex,
	readerCount *int, data *int, wg *sync.WaitGroup) {

	// Add one reader
	readerLock.Lock()
	*readerCount++

	if *readerCount == 1 {
		writerLock.Lock() // first reader blocks writers
	}
	readerLock.Unlock()

	// Display what the reader is reading
	fmt.Println("Reader", id, "is reading:", *data)
	time.Sleep(time.Second)

	// Remove one reader
	readerLock.Lock()
	*readerCount--

	if *readerCount == 0 {
		writerLock.Unlock() // no readers left, allow a writer
	}
	readerLock.Unlock() // let other readers update the count

	wg.Done()
}

// Change the shared data
func writer(id int, writerLock *sync.Mutex,
	data *int, wg *sync.WaitGroup) {

	writerLock.Lock() // take the lock or wait until no readers or writers are inside

	*data++
	fmt.Println("Writer", id, "changed data to:", *data)
	time.Sleep(time.Second)

	writerLock.Unlock() // release the lock so others can enter

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var readerLock sync.Mutex
	var writerLock sync.Mutex

	readerCount := 0
	data := 0
	writerID := 0

	wg.Add(7) // five readers and two writers

	for i := range 5 {
		go reader(i, &readerLock, &writerLock,
			&readerCount, &data, &wg)

		// After reader 1 or 3, writer can start writing
		if i == 1 || i == 3 {
			go writer(writerID, &writerLock, &data, &wg)
			writerID++

			// Pause main  before launching the next readers...
			// this allows  writers to start
			time.Sleep(time.Second)
		}
	}

	wg.Wait() // wait until all goroutines finish
	fmt.Println("Final data:", data)
}
