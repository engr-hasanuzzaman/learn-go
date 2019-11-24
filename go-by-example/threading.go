package main
import(
	"fmt"
	"runtime"
)

func say(s string){
	for i := 0; i < 10; i++{
		runtime.Gosched()
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
	// fmt.Println("befeor number of cpu ", runtime.NumCPU())
	// nCpu := runtime.GOMAXPROCS(4)
	go say("1")
	say("2")
}