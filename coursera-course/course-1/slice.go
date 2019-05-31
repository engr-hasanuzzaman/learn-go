package main
import (
	"fmt"
	"strconv"
)

func main(){
	slice := make([]int, 0, 3)
	for {
		var input string
		fmt.Scan(&input)
		if input == "X" { break }
		lastVal, _ := strconv.Atoi(input)

		for i, v := range slice {
			if v > lastVal {
				slice[i] = lastVal
				lastVal = v
			}
		}

		slice = append(slice, lastVal)
		fmt.Println(slice)
	}
}