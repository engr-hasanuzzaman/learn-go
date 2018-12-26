package main
import . "fmt"

func main()  {
	Println("addition of 2, 3 is", add(2, 3))
	v1, v2 := mulReturn(10)
	Println("multiple return", v1, v2) 
}

// normalt function
func add(n, m int) int {
	return n + m
}

// function with multiple return
func mulReturn(value int) (int, int) {
	return value + 1, value - 1
}