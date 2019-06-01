package main
import(
	"fmt"
	"os"
	"bufio"
	"strings"
)


func main()  {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither",	"hsss"}
	var aName, action string
	var animal Animal

	for {
		aName, action = TakeInput("> ")

		switch aName {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			fmt.Println("please enter eithe cow/bird/snake as animal name")
			continue			
		}

		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("please enter eithe eat/move/speak as animal action")
			continue			
		}
	}
}

func TakeInput(msg string) (string, string) {
	scanner:= bufio.NewScanner(os.Stdin)
	fmt.Printf(msg)
	scanner.Scan()
	rLine := strings.Fields(scanner.Text())
	return rLine[0], rLine[1]
}

// declare Animal
type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat(){
	fmt.Println(a.food)
}

func (a *Animal) Move(){
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak(){
	fmt.Println(a.noise)
}