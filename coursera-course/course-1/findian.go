package main
import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func main(){
	var myStr string
	fmt.Println("Enter a string")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	myStr = scanner.Text()

	myStr = strings.ToLower(myStr)
	sLen := len(myStr)

	if sLen < 3 {
		fmt.Println("Not Found!")
	} else if myStr[0] == 'i' && myStr[len(myStr) - 1] == 'n' && strings.IndexRune(myStr, 'a') != -1{
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}