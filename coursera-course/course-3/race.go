package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go doStaf1()
	go doStaf2()
	wg.Wait()
	fmt.Println(x)
}

func doStaf1() {
	x = x + 1
	wg.Done()
}

func doStaf2() {
	x = x + 4
	wg.Done()
}

/*
A race condition occurs when two or more goroutines can access shared data and they try to change it at the same time.
Because the goroutine scheduling algorithm can swap between routines at any time, you don't know the order in which the routines will attempt to access the shared data.
Most importantly, inteleaving does not happen on go code level rather than it is on machine instruction level that can perform interleaved for single go statement.

Therefore, the result of the change in data is dependent on the scheduling algorithm,
i.e. both routines are "racing" to access/change the data.

For example, on the above program the value of x is not predictable as any of doStaf1 or doStaf2 can complte their task in unpredictable order.
The possible value after executing the code may be 5 or 4 or 1

When x is 5:
	It could be happen that doStaf1 will complete the task then doStaf2. So, when doStaf1 compte execution the velue of x is 1.
	So, when doStaf2 read the x the value is 1 that will be 5 after adding 4 with it

When x is 1:
	It could be happen that reading the value of x for both function are interleaved. So, in both function the value of x is 0.
	Now, if doStaf2 complte the code first(at this point x == 4) then doStaf1 will override the value 4 with 1 as the initial value of x is 0 to this routine.

When x is 4:
	It could be happen that reading the value of x for both function are interleaved. So, in both function the value of x is 0.
	Now, if doStaf1 complte the code first(at this point x == 1) then doStaf2 will override the value 1 with 4 as the initial value of x is 0 to this routine.
*/
