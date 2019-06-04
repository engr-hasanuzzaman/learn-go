package main
import(
	"fmt"
	"sync"
)
var x int
var wg sync.WaitGroup

func main()  {
	wg.Add(2)
	go doStaf1()
	go doStaf2()
	wg.Wait()
	fmt.Println(x)
}

func doStaf1()  {
	x = x + 1
	wg.Done()
}

func doStaf2()  {
	x = x + 4
	wg.Done()
}