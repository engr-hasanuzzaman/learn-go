
package main
import(
	_ "fmt"
	"os"
)

func main(){
	panic("This is obvious panic") // this will happen alway. Run with commenting this line also

	currDir, _ := os.Getwd()

	if _, err := os.Create( currDir + "/test/test.text"); err != nil{
		panic("File creation failed")
	}
}