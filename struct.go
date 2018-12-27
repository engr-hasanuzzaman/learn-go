package main
import . "fmt"

func main()  {
	person := Person{ "hasanuzzaman", "sumon", 29, "male" }
	Println("person structure is ", person)
	Println("person full-name is", person.fullName())
}

type Person struct {
	firstName string
	lastName string
	age int
	gender string
}

func (person *Person) fullName() string{
	return person.firstName + " " + person.lastName
}
