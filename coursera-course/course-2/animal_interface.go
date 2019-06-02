package main
import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"errors"
)


func main()  {
	// contains animal collection
	animals := make(map[string] Animal)
	// var command, aName, action string

	for {
		// action contain animal type for createanimal command
		command, aName, action, err := TakeInput("> ")
		
		// error handling for user input
		if err != nil {
			fmt.Println("Please enter test with three string value")
			continue
		}

		switch command {
		case "newanimal":
			if _, ok := animals[aName]; ok {
				fmt.Println("Animal exist with name ", aName)
				continue
			}
			if tAnimal, err := CreateAnima(action); err != nil {
				continue
			} else {
				animals[aName] = tAnimal
			}
		case "query":
			if animal, ok := animals[aName]; ok {
				AnimalAction(animal, action)
				continue
			}
			fmt.Println("Animal does not exist with name ", aName)
			continue
		default:
			fmt.Println("please enter eithe newanimal/query command")
			continue			
		}
	}
}

func CreateAnima(aType string) (Animal, error) {
	switch aType {
	case "cow":
		fmt.Println("Created it!")
		return Cow{"grass", "walk", "moo"}, nil
	case "bird":
		fmt.Println("Created it!")
		return Bird{"worms", "fly", "peep"}, nil
	case "snake":
		fmt.Println("Created it!")
		return Snake{"mice", "slither",	"hsss"}, nil
	default:
		fmt.Println("please enter eithe cow/bird/snake as animal type")
		return Cow{},  errors.New("animal: bad animal type")		
	}
}

func AnimalAction(animal Animal, action string)  {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("please enter eithe eat/move/speak as animal action")
	}
}

func TakeInput(msg string) (string, string, string, error) {
	scanner:= bufio.NewScanner(os.Stdin)
	fmt.Printf(msg)
	scanner.Scan()
	rLine := strings.Fields(scanner.Text())
	if len(rLine) == 3 {
		return rLine[0], rLine[1], rLine[2], nil
	} else {
		return "", "", "",  errors.New("command: bad command")
	}
}

// declare Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// declare Cow
type Cow struct {
	food, locomotion, noise string
}

func (a Cow) Eat(){
	fmt.Println(a.food)
}

func (a Cow) Move(){
	fmt.Println(a.locomotion)
}

func (a Cow) Speak(){
	fmt.Println(a.noise)
}

// declare Bird
type Bird struct {
	food, locomotion, noise string
}

func (a Bird) Eat(){
	fmt.Println(a.food)
}

func (a Bird) Move(){
	fmt.Println(a.locomotion)
}

func (a Bird) Speak(){
	fmt.Println(a.noise)
}

// declare Snake
type Snake struct {
	food, locomotion, noise string
}

func (a Snake) Eat(){
	fmt.Println(a.food)
}

func (a Snake) Move(){
	fmt.Println(a.locomotion)
}

func (a Snake) Speak(){
	fmt.Println(a.noise)
}