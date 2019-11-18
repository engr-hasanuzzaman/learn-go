package main
import(
	"fmt"
	"time"
)

func main()  {
	timer1 := time.NewTimer(2*time.Second)
	<-timer1.C
	fmt.Println("Timer 1 is expired")

	timer2 := time.NewTimer(2*time.Second)
	go func(){
		<- timer2.C
		fmt.Println("Timer 2 expired")
	}()

	if ok := timer2.Stop(); ok{
		fmt.Println("Stopped timer 2")
	}
	
}