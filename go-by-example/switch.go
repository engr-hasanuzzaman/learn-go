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
  fmt.Println("")

	// case without condition
	switch  {
	case age >= 20 && age < 30:
		fmt.Println("You are middle aged")
	case age >= 30:
		fmt.Printf("You are late 30")
	default:
		fmt.Printf("You are default")
	}
}