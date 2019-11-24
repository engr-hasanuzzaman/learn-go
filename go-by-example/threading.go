package main
import(
	"fmt"
	"runtime"
)

func say(s string){
	for i := 0; i < 10; i++{
		fmt.Println(s)
	}
}

func sum(x int) {
  sum := 0
  for i := 0; i < x; i++ {
    sum += i
  }
  fmt.Println("sum is", sum)
}

func main()  {
	fmt.Println("befeor number of cpu ", runtime.NumCPU())
	// nCpu := runtime.GOMAXPROCS(4)
	go say("hello")
	fmt.Println("World")
	fmt.Println("World")
	fmt.Println("World")
	fmt.Println("World")
	fmt.Println("World")
	fmt.Println("World")
	fmt.Println("World")
	fmt.Println("World")
	fmt.Println("World")
	fmt.Println("World")
	// fmt.Println("Set Number of CPU is ", nCpu)
	fmt.Println("After number of cpu ", runtime.NumCPU())
	go sum(2)
	go sum(100)
	go sum(200)
	for i := 0; i < 100000; i++{
	}
}