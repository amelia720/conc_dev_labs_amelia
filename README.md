# Concurrent Development Labs
## How to run

Open the main project folder in terminal, then run one exercise at a time using the commands below.

| Exercise | Description | Command |
| --- | --- | --- |
| 1. Atomic | Uses an atomic variable to safely add to a shared counter. | `go run ./atomic` |
| 2. Mutex | Uses a mutex so only one go routine changes the counter at a time. | `go run ./mutex` |
| 3. Barrier | Makes all go routines finish Part A before starting Part B using a mutex and semaphore. | `go run ./barrier` |
| 4. Reusable Barrier | Uses an atomic counter and two buffered channels to run a barrier over three rounds. | `go run ./barrier2` |
| 5. Barrier Struct | Stores the barrier's channel, mutex and counters in a struct. | `go run ./barrierStruct` |
| 6. Rendezvous | Uses two channels so all five go routines finish Part A before starting Part B. | `go run ./rendezvous` |
| 7. Semaphore | Uses a buffered channel so only five tasks can work at a time. | `go run ./semaphore` |
| 8. Signalling | Uses an unbuffered channel so both go routines finish Part A before either starts Part B. | `go run ./signalling` |
| 9. Dining Philosophers | An exercise based on the Dining Philosophers problem. | `go run ./DiningPhilosophers` |
| 10. Collatz Conjecture | An exercise based on the Collatz conjecture. | `go run ./sem-ex-collatz-conjecture` |
| 11. Producer-Consumer | Uses a buffered channel to pass events from a producer to a consumer. | `go run ./Producer-Consumer` |
| 12. Readers-Writers | Uses mutexes to allow multiple readers or one writer at a time. | `go run ./Reader-Writer` |

## Credits

Template code by Dr. Joseph Kehoe.
Modified by Amelia Hamulewicz.

I worked with Mark Lambert, Dorian Nowacki and Adam Noonan on some of the exercises. Their help is noted in the code comments.

## Licence

GNU General Public License, version 3.
See the LICENSE file for licence details.
