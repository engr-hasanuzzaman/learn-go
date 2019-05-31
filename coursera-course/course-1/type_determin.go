package main
import(
	"fmt"
	"reflect"
)

func main()  {
	array := [...]int{1,2,3} // this declare an array
	slice := []int{1,2,3} // this declare an slice

	fmt.Printf("type of array is %s and slice is %s\n", reflect.TypeOf(array).Kind(), reflect.TypeOf(slice).Kind())
}