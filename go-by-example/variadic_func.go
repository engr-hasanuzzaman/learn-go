package main
import . "fmt"

func adder(args ...int) int {
	var sum int
	Println("passing args is ", args)
	for _, n := range args {
		sum += n
	}

	return sum
}

func main()  {
	Println("Sum of 1,2,3,4,5 is ", adder(1,2,3,4,5))
	ip := []int{1,2,34,56,67,89}
	Println("Sum of ", ip, " is ", adder(ip...))
}