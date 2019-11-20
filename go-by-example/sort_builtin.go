package main
import(
	"fmt"
	"sort"
)

func main(){
	intInputs := []int{10,2,-4,45, 23, 5,89}
	fmt.Println("Before sorting the array is ", intInputs)
	sort.Ints(intInputs)
	fmt.Println("After sorting the array is ", intInputs)

	strInputs := []string {"This", "is","sample","string", "and", "nothing", "else", "-0", "22", "1"}
	fmt.Println("Before sorting the array is ", strInputs)
	sort.Strings(strInputs)
	fmt.Println("After sorting the array is ", strInputs)
}