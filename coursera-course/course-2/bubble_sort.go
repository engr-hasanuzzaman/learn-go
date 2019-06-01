package main
import(
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main()  {
	input := make([]int, 0, 10)
	input = TakeInput()
	BubbleSort(input)
	fmt.Println("After sorting", input)
}

func TakeInput() []int {
	fmt.Println("Enter your input array")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rLine := scanner.Text()
	return numbers(rLine)
}

// swap element
func Swap(input []int, i int)  {
	temp := input[i]
	input[i] = input[i+1]
	input[i+1] = temp
}

// sort slice inplace
func BubbleSort(input []int)  {
	for i, _ := range input {
		for j := 0; j < len(input) - (i + 1) ; j++ {
			if input[j] > input[j+1]{
				Swap(input, j)
			}
		}
	}
}

// space separated string to []int
func numbers(s string) []int {
	var n []int
	for index, f := range strings.Fields(s) {
			if index >= 10 {
				break // max 10 input
			}
			// convert string to int
			i, err := strconv.Atoi(f)
			if err == nil {
					n = append(n, i)
			}
	}
	return n
}