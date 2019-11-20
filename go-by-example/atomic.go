package main
import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main()  {
	var ops uint64
	var wg sync.WaitGroup

	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func ()  {
			for j := 0; j < 1000; j++{
				atomic.AddUint64(&ops, 1) // ops += 1 will lead to unpredictable state
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Updated num is ", ops)
}