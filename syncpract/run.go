package syncpract

import (
	"fmt"
	"math/rand"
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
	condExample()
	poolExample()
	onceExample()
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

/*Cond*/
/*Cond helps to notify goroutines that a condition
is occured or wait for certain conditon to occur*/
func condExample() {

	con := sync.NewCond(&sync.Mutex{}) // creating new Cond vaiable with Mutex
	queue := make([]int, 0, 10)        //creating queue to store int

	popFunc := func(dt time.Duration) { //function to pop value from queue
		time.Sleep(dt) //sleep for given time duration
		con.L.Lock()   //locks to perform operation
		queue = queue[1:]
		fmt.Println("Removed from queue")
		con.L.Unlock() //unclocking
		con.Signal()   //sending signal for goroutine which are waiting

	}

	for i := 0; i < 10; i++ {
		con.L.Lock() //locking to perform operation
		for len(queue) == 2 {
			con.Wait() //wait if length of queue is 2
		}
		fmt.Println("Adding to queue")
		queue = append(queue, i)    //append queue
		go popFunc(1 * time.Second) //run popfunc in another goroutine
		con.L.Unlock()              // unlock
	}
	fmt.Println(queue) //print queue with only two values
}

/*Pool*/
/*A Pool is a set of temporary objects that may be individually
saved and retrieved.*/
func poolExample() {

	rand.Seed(time.Now().Unix()) //seed random with unix
	myPool := &sync.Pool{        //creating pool
		New: func() interface{} { //new func which gives new random object
			fmt.Println("New Random number created")
			return rand.Intn(15)
		},
	}

	num := myPool.Get()             //get the random number from pool
	println(num.(int))              //print the num
	myPool.Put(num)                 //put the num into pool
	fmt.Println(myPool.Get().(int)) //get value from pool and print value

	/*In the above the random number will be created only once*/
}

/*Once*/
/*A Pool is a set of temporary objects that may be individually saved and*/
func onceExample() {

	var num int           //num to increment
	Increfunc := func() { //increment function
		num++
	}
	var once sync.Once    //creating once object
	var wg sync.WaitGroup //waitgroup
	for i := 0; i < 100; i++ {

		wg.Add(1) //adding goroutine in waitgroup
		go func() {
			defer wg.Done()
			once.Do(Increfunc) //incrementing num using Once
		}()
	}
	wg.Wait()
	fmt.Println("Value of num in Once:= ", num)
	/*this will increment num once*/
}
