package main
import "fmt"

func main(){
	var fNum float32
	fmt.Printf("Enter your floating number \n")
	fmt.Scan(&fNum)
	fmt.Printf("You truncated number is %d\n", int32(fNum))
}