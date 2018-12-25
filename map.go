package main
import . "fmt"

func main()  {
	m1 := make(map[string] string)
	m1["name"] = "hasan"
	m1["age"] = "29"
	m1["marriage status"] = "married"
	Println("size of map is", len(m1))
	Println(m1)
	delete(m1, "age")
	Println(m1)

	// loop over map with range that return key, value
	for k, v := range m1 {
		Println("k", k, "value", v)
	}
}