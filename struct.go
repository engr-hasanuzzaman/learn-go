package main
import(
	 . "fmt"
	 "encoding/json"
	)

func main()  {
	person := Person{ "hasanuzzaman", "sumon", "male", 29 }
	Println("person structure is ", person)
	Println("person full-name is", person.fullName())
	test := Test{"sumon", 30}
	r, _ := json.Marshal(test)
	Println("marshalling structure", string(r))
}

type Person struct {
	firstName, lastName, gender string
	age int
}

type Test struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (person *Person) fullName() string{
	return person.firstName + " " + person.lastName
}
