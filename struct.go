package main
import(
	 . "fmt"
	 "encoding/json"
	)

func main()  {
	person := Person{ "hasanuzzaman", "sumon", 29, "male" }
	Println("person structure is ", person)
	Println("person full-name is", person.fullName())
	// test := Test{"sumon", 30}
	r, _ := json.Marshal(person)
	Println("marshalling structure", string(r))
}

type Person struct {
	firstName string
	lastName string
	age int
	gender string
}

type Test struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (person *Person) fullName() string{
	return person.firstName + " " + person.lastName
}
