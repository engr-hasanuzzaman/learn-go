package main
import(
	"fmt"
	_ "runtime"
)

func say(s string){
	for i := 0; i < 10; i++{
		fmt.Println(s)
	}
}

func main()  {
	say("hellow")
	fmt.Println("World")
}