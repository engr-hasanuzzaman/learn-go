package main
import . "fmt"

func main()  {
	// array size and type both needed for defining array
	var ages [3] int
	Println("befor assigning", ages)
	ages[0] = 10
	ages[1] = 20
	ages[2] = 30
	Println("after assign age is ", ages)
	
	// array with initialization
	var roles = [5]int{1,2,3,4}
	Println("printing rolse", roles)
}