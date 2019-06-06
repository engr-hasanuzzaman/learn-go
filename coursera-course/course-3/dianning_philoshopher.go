package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch chan int
var mutex sync.Mutex

func main() {
	num := 5
	ch = make(chan int, 2)
	ch <- 1
	ch <- 1
	mutex = sync.Mutex{}
	// Create forks
	forks := make([]*fork, num)
	for i := 0; i < num; i++ {
		forks[i] = new(fork)
	}

	philosophers := make([]*philosopher, num)
	
	for i := 0; i < num; i++ {
		philosophers[i] = &philosopher{
			id: i + 1, nOfEat: 0, leftFork: forks[i], rightFork: forks[(i+1)%num]}
	}

	for i := 0; i < 16; i++ {
		wg.Add(1)
		if eatPermissionFromHost() {
			go philosophers[i%num].eat()
		}
	}

	// wait for goroutine complete
	wg.Wait()
}

type fork struct{ sync.Mutex }

type philosopher struct {
	id                  int
	nOfEat 							int
	leftFork, rightFork *fork
}

func (p philosopher) eat() {
	mutex.Lock()
	if p.nOfEat >= 3 {
		mutex.Unlock()
		return
	}
	mutex.Unlock()

	p.leftFork.Lock()
	p.rightFork.Lock()

	fmt.Println("starting to eat ", p.id)
	time.Sleep(time.Second)
	p.nOfEat = p.nOfEat + 1
	ch <- 1
	p.rightFork.Unlock()
	p.leftFork.Unlock()

	fmt.Println("finishing eating", p.id)
	time.Sleep(time.Second)
	wg.Done()
}

func eatPermissionFromHost() bool{
	return 1 == <- ch
}