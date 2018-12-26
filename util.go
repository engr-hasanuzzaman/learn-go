package main
import(
	. "fmt"
	"os"
	"log"
	"io/ioutil"
)

func main()  {
	Println("create new file using os")
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("this is file testing")
	// log.Fatal("test fatal log")
	file.Close()

	// read file
	stream, err := ioutil.ReadFile("test.txt")
	if err != nil{
		log.Fatal("error reading file", err)
	}
  Println("print stream", stream)
	readString := string(stream)
	Println("stream to string", readString)
}