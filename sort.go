package main

import (
	"fmt"
	"sort"
)

func main()  {
	people := []Person{
		Person{"sumon", 29},
		Person{"salma", 27},
		Person{"santo", 17},
		Person{"eti", 22},
	}

	// sort people by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})

	fmt.Println("sort by age", people)

	// sort people by name
	sort.Slice(people, func(i, j int) bool {
		return people[i].name < people[j].name
	})

	fmt.Println("sort by name", people)
}

type Person struct {
	name string
	age float32
}