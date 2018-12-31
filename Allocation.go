package main
import (
	"fmt"
	"reflect"
)

func main()  {
	num := new(int) // new create pointer
	fmt.Println("Pointer is", num, "value is", *num)
	*num = 1234
	fmt.Println("after assign value, Pointer is", num, "value is", *num)
	fmt.Printf("type of num is %T \n", *num)
	// type checking 
	fmt.Println("type of *num 1st way", firstWay(*num)) 
	fmt.Println("type of *num 2nd way", secondWay(*num)) 
	fmt.Println("type of *num 2nd way", thirdWay(*num)) 
}

// three ways to check type
// N.B: empty interface means any value
func firstWay(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// using built-in reflect package
func secondWay(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// using type switch
func thirdWay(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
