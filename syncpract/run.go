package syncpract

import (
	"fmt"
	"sync"
	"time"
)

/*The sync package contains the concurrency primitives
that are most useful for low level memory access
synchronization*/

func Run() {
	fmt.Println("sync package in detail")
	waitGroupExample()
	mutexExample()
}

/*waitGroup*/
/*WaitGroup is a great way to wait for a set of
concurrent operations to complete when you either
don’t care about the result of the concurrent operation,*/
func waitGroupExample() {
	var wg sync.WaitGroup //define WaitGroup variable
	wg.Add(1)             //add a goroutine in WaitGroup
	go func() {
		defer wg.Done() //notify WaitGroup that goroutine is done
		fmt.Println("Hello from Goroutine 1st")
		time.Sleep(10 * time.Millisecond)
	}()

	wg.Add(1) //add second goroutine in WaitGroup
	go func() {
		defer wg.Done()
		fmt.Println("Hello from Goroutine 2nd")
		time.Sleep(5 * time.Millisecond)
	}()

	wg.Wait() //wait for all goroutine to be completed
	fmt.Println("All gorotines are completed")

}

/*Mutex*/
/*To Lock and Unlock Crtical sections in program
Mutex is used in golang*/
func mutexExample() {

	var wg sync.WaitGroup
	var lock sync.Mutex //declaring Mutex variable
	count := 0          // count variable as critical section in this program

	IncrFunc := func() { //Increment function
		lock.Lock() //locking the count variable
		count++     //Incrementing the count
		fmt.Println("Increment count = ", count)
		defer lock.Unlock() //unlocking
	}

	DecrFunc := func() { //decrement function
		lock.Lock() //locking the count variable
		count--     //dcrementing the count
		fmt.Println("Decrement count = ", count)
		defer lock.Unlock() //unlocking
	}

	//running Increment function in goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() { //Incrementing count 10 times
			defer wg.Done()
			IncrFunc()
		}()
	}

	//running decrement function in goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() { //decrementing count 10 times
			defer wg.Done()
			DecrFunc()
		}()
	}
	wg.Wait()
	fmt.Println("Count:= ", count) //End result will be 0
}
