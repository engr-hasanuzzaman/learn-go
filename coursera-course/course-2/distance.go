package main
import(
	"fmt"
)

func main()  {
	var a, v0, s0, t float64
	fmt.Println("Enter acceleration")
	fmt.Scan(&a)
	fmt.Println("Enter initial velocity")
	fmt.Scan(&v0)
	fmt.Println("Enter initial displacement")
	fmt.Scan(&s0)

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("Enter time")
	fmt.Scan(&t)
	fmt.Println(fn(t))
}

// 
func GenDisplaceFn(a, v0, s0 float64)(func(float64) float64)  {
	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v0 * t) + s0
	}
}