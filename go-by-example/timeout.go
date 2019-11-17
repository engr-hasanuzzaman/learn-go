package main
import(
	"fmt"
	"time"
)

func main()  {
	c := make(chan string, 1)

	go func ()  {
		time.Sleep(2 * time.Second)
		c <- "result 1"
	}()

	select {
	case res := <- c:
		fmt.Println(res)
	case res := <-time.After(1 * time.Second):
		fmt.Println("Timeout 1 ", res)
	}
	c1 := make(chan string, 1)
	go func ()  {
		time.Sleep(2*time.Second)
		c1 <- "result 2"
	}()
	
	
	select{
	case res := <-c1:
		fmt.Println(res)
	case res := <-time.After(3*time.Second):
		fmt.Println(res)
	}
}