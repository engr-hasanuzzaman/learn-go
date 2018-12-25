package main
import . "fmt"

func main()  {
	s := make([]string, 3)
	Println("string slice before", s)
	s[0] = "sumon"
	s[1] = "sumon"
	s[2] = "sumon"
	Println("string slice after", s)

	s = append(s, "appended 1")
	Println("string slice apn", s)

	// just string
	var fstr [2]string
	fstr[0] = "first str"
	var str = [3]string{"1st", "2nd", "3rd3rd3rd3rd"}
	Println("str", str)

	ages := make([]int, 3)
	ages[2] = 30
	ages[0] = 10
	ages[1] = 20
	Println(ages)
	Println("ages[2:3]" , ages[0:1])
}