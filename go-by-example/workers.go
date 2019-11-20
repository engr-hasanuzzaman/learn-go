package main
import . "fmt"
import "time"

func worker(id int, jobs <-chan int, result chan<- int){
	for j := range jobs{
		Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		Println("worker", id, "finished job", j)
		result <- j * 2
	}
}
func main(){
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// run the worker
	for i := 1; i <= 3; i++{
		go worker(i, jobs, results)
	}

	for i := 1; i <= 5; i++{
		jobs<- i
	}

	close(jobs)
	for i := 1; i <= 5; i++{
		<-results
	}
}