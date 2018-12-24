package main
import "fmt"
func main()  {
	const size int = 5
	// print start piramid using for loop
	for i := 0; i <= size; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}

	var inputSize int;
	fmt.Scanf("Inter pyramid size %d", inputSize)
	printPyramid(inputSize)
}

func printPyramid (size int){
	for i := 0; i <= size; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}