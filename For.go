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
	fmt.Println("Inter pyramid size")
	fmt.Scan(&inputSize)
	printPyramid(inputSize)
}

func printPyramid (size int){
	i := 1
	for i <= size {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
		i++
	}
}