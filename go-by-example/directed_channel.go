package main
import . "fmt"

func ping(c chan<- string, msg string){
	c <- msg
}

func pong(p <-chan string, pongs chan<- string){
	msg := <-p
	pongs <- msg
}

func main()  {
	pChan := make(chan string, 1)
	pongChan := make(chan string, 1)
	ping(pChan, "this is ping message")
	pong(pChan, pongChan)
	Println("The output is ", <-pongChan)
}