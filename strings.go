package main
import (
	. "fmt"
	"strings"
)

func main()  {
	str := "hello world"
	Println("contain 'or'", strings.Contains(str, "or"))
	Println("Index of or", strings.Index(str, "or"))
	Println("Count of 'or'", strings.Count(str, "or"))
	Println("replace 'or' with 'xx'", strings.Replace(str, "l", "X", len(str)))
	Println(str)
	Println("split with space", strings.Split(str, ""))
	Println("split with space", strings.Join(strings.Split(str, ""), ""))
}