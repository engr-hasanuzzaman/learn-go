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

	// int array slice using make
	ages := make([]int, 3)
	ages[2] = 30
	ages[0] = 10
	ages[1] = 20
	Println(ages)
	Println("ages[2:3]" , ages[0:1])

	// array slice
	s1 := [5]int {1, 2, 3, 4, 5}
	Println("s1 slice", s1[1:3], s1[3:])

	// loop over array
	for i, value := range s1 {
		Println("index = ", i, "value is", value) 
	}

	// make with slice size
	list1 := make([]int, 5, 10)
	Println("size of list1 is", len(list1))
	list1 = append(list1, 0, 6)
	list1 = append(list1, 1, 7)
	list1 = append(list1, 2, 8)
	list1 = append(list1, 3, 9)
	Println(list1)

	// array copy
	a1 := []int {1,2,3,4,5}
	a2 := make([]int, 10)
	copy(a2, a1)
	Println(a2)
}