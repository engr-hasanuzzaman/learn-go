package main
import(
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)

func main()  {
	// declare scanner
	scanner := bufio.NewScanner(os.Stdin)

	// declare map
	dic := make(map[string] string)

	// read name
	fmt.Println("Pleae enter your name")
	scanner.Scan()
	name := scanner.Text()

	// read address
	fmt.Println("Please enter your address")
	scanner.Scan()
	address := scanner.Text()

	dic["name"] = name
	dic["address"] = address
	jVal, _ := json.Marshal(dic)
	fmt.Printf("%s\n", jVal)
}
