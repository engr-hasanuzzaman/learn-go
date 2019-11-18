package main
import(
	"fmt"
	"time"
)

func main(){
	ticker := time.NewTicker(500*time.Millisecond)
	done := make(chan bool)
	go func ()  {
		for {
			select{
			case tick := <-ticker.C:
				fmt.Println("tick ", tick)
			case <- done:
				fmt.Println("stop")
				return
			}
		}
	}()

	time.Sleep(1600*time.Millisecond)
	ticker.Stop()
	done<-true
	fmt.Println("Tricker stopped")
}