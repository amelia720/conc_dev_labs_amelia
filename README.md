# Concurrent Development Labs
## How to run

Open the main project folder in terminal, then run one exercise at a time using the commands below.

| Exercise | Description | Command |
| --- | --- | --- |
| Atomic | Uses an atomic variable to safely add to a shared counter. | `go run ./atomic` |
| Mutex | Uses a mutex so only one go routine changes the counter at a time. | `go run ./mutex` |
| Barrier | Makes all go routines finish Part A before starting Part B using a mutex and semaphore. | `go run ./barrier` |
| Reusable Barrier | Uses an atomic counter and two buffered channels to run a barrier over three rounds. | `go run ./barrier2` |
| Barrier Struct | Stores the barrier's channel, mutex and counters in a struct. | `go run ./barrierStruct` |
| Rendezvous | Uses two channels so all five go routines finish Part A before starting Part B. | `go run ./rendezvous` |
| Semaphore | Uses a buffered channel so only five tasks can work at a time. | `go run ./semaphore` |
| Signalling | Uses an unbuffered channel so both go routines finish Part A before either starts Part B. | `go run ./signalling` |
| Dining Philosophers | An exercise based on the Dining Philosophers problem. | `go run ./DiningPhilosophers` |
| Collatz Conjecture | An exercise based on the Collatz conjecture. | `go run ./sem-ex-collatz-conjecture` |

## Credits

Template code by Dr. Joseph Kehoe.
Modified by Amelia Hamulewicz.

I worked with Mark Lambert, Dorian Nowacki and Adam Noonan on some of the exercises. Their help is noted in the code comments.

## Licence

GNU General Public License, version 3.
See the LICENSE file for licence details.
