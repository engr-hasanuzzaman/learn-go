package main
import . "fmt"

func printSerias(n int){
	if n == 0{
		return
	}

	printSerias(n-1)
	Println(n)
}

func main()  {
	printSerias(10)
}