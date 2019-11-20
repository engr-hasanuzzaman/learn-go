package main
import(
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup){
	fmt.Println("Worker started working with id ", id)
	time.Sleep(time.Second)
	fmt.Println("Worker finished working ", id)
	wg.Done()
}

func main(){
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}