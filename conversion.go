package main
import (
	. "fmt"
	"strconv"
)

func main()  {
	rInt := 10
	rFloat := 23.0
	fStr := "123.5"
	Println("convert int to float", float64(rInt))
	Println("convert float to int", int(rFloat))
	nInt, _ := strconv.Atoi("10")
	Println("convert rand str to int", nInt)

	nFloat, _ := strconv.ParseFloat(fStr, 64)
	Println("parse float from str", nFloat)
	Println("convert int to string", strconv.Itoa(1234))
}