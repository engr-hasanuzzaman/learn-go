package main
import . "fmt"

func main()  {
	Println("addition of 2, 3 is", add(2, 3))
	v1, v2 := mulReturn(10)
	Println("multiple return", v1, v2)

	// closure
	num3 := 0
	sub := func () int {
		num3 += 10
		return num3
	}

	Println(sub())
	Println(sub())

	// defer prevent executing func. After execution end of caller method 
	//  defer method will be called
	defer printOne() 
	printTwo()
	Println("end of main")
	foo()
}

// normalt function
func add(n, m int) int {
	return n + m
}

// function with multiple return
func mulReturn(value int) (int, int) {
	return value + 1, value - 1
}

// method defering
func printOne(){
	Println(1)
}

func printTwo(){
	Println(2)
}

// defer panic and recover
func foo(){
	defer func ()  {
		err := recover()
		Println("panic has been occured", err)
		Println("recovering")
	}()

	panic("I am paniccing")
}