package main

/**
concurrency is making progress on more than one task simultaneously at a time.
Go has rich support for concurrency using "go-routines" and "channels"
GoRoutine is a function that is capable of running concurrently with other functions

GoRoutines are extremely cheap when compared to threads (resource-wise speaking), GoRoutines are only a few kb in stack size
and the stack can grow and shrink according to needs of the application whereas in the case of threads the stack size has to
be specified and fixed.

GoRoutines are multiplexed into fewer OS threads, i.e.There might only one thread in a program of thousands of GoRoutines
GoRoutines communicate using channels
channels are like pipes that link the main routine with the subroutine
*
*/
func main() {

}
