package main
import "fmt"

func main()  {
	if age := 29; age > 50{
		fmt.Println("your are an old man")
	} else if age > 40{
		fmt.Println("you are a middle age man")
	} else if age > 20 {
		fmt.Println("you are a young man")
	}	else{
		fmt.Println("you are a child")
	}
}