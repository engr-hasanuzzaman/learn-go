package main
import (
	"fmt"
)

func main()  {
	var age int
	fmt.Printf("Enter your age ")
	fmt.Scanf("%d", &age)
	fmt.Printf("You are ")
	switch age {
	case 20:
		fmt.Printf("20")
	case 30:
		fmt.Printf("30")
	default:
		fmt.Printf("Other")
	}
}