package main
import(
	. "fmt"
	"strconv"
) 

func main()  {
	jobs := make(chan string)
	status := make(chan bool)

	go func(){
		for task := range jobs {
			Println("Received task  ", task)
		}
		
		Println("Processed all the tasks")
		status <- true
		return
		// Println("Recei")
	}()

	for i := 0; i < 3; i++{
		Println("task pass ", i)
		jobs <- "Jobs " + strconv.Itoa(i)
		
	}

	close(jobs)
	<- status
}