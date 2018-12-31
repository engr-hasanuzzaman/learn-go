package main
import . "fmt"

func main()  {
	num := new(int) // new create pointer
	Println("Pointer is", num, "value is", *num)
	*num = 1234
	Println("after assign value, Pointer is", num, "value is", *num)
	Printf("type of num is %T \n", *num)
} 