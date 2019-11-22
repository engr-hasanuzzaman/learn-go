package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, s := range os.Environ() {
		pair := strings.SplitN(s, "=", 2)
		fmt.Printf("%s=%s", pair)
	}
	fmt.Println()
	os.Setenv("foo", "this is foo")
	fmt.Println("env value of foo is ", os.Getenv("foo"))
	fmt.Println("env value of bar is ", os.Getenv("bar")) // default empty value
}
