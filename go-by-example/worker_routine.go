package main
import(
	. "fmt"
	"strconv"
) 

func main()  {
	jobs := make(chan string)
	status := make(chan bool)

	go func(){
		for {
			if task, more := <-jobs; more {
				Println("Received task  ", task)
			}else{
				// Println("Received task ", task)
				Println("Processed all the tasks")
				status <- true
				return
			}
		}
		// Println("Recei")
	}()

	for i := 0; i < 3; i++{
		Println("task pass ", i)
		jobs <- "Jobs " + strconv.Itoa(i)
		
	}

	close(jobs)
	<- status
}