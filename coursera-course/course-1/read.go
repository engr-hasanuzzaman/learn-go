package main
import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main()  {
	var fileName string
	type Person struct{
		fname, lname string
	}
  var pCol = make([]Person, 0)

	fmt.Println("Enter file name")
	// new scanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName = scanner.Text()
	file, _ := os.Open(fileName)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), " ")
		pCol = append(pCol, Person{temp[0], temp[1]})
	}

	for _, p := range pCol {
		fmt.Println(p.fname, p.lname)
	}
}