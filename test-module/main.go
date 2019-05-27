package main
import "test/adder"
import . "fmt"

func main(){
	Println("sum of 2 and 3 is", adder.Add(2, 3))
}