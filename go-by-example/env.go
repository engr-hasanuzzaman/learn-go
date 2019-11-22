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
}
