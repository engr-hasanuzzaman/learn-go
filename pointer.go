package main
import . "fmt"

func main()  {
	x := 10
	Println("initial value of x", x)
	value(x)
	Println("after calling value func", x)
	ref(&x)
	Println("after calling ref func", x)

	// create pointer
	y := new(int)
	ref(y)
	Println("value of y is", *y)
}

func value(x int){
	x = 100
}

func ref(x *int){
	*x = 200
}

