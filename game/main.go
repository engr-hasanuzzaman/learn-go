package main
import(
	"fmt"
	"encoding/csv"
	"io"
	"os"
	"log"
) 

func main()  {
	csv_file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Could not open file ", err)
	}

	r := csv.NewReader(csv_file)

	for {
		record, err := r.Read()
		
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("CSV read error")
		}

		fmt.Printf("Question is %s and answer is %s \n", record[0], record[1])
	}
	fmt.Println("This is main file")
}