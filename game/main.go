package main
import(
	"fmt"
	"encoding/csv"
	"io"
	"os"
	"log"
) 

func main()  {
	// wait_sec := 15
	n_correct_answer := 0
	q_counter := 0
	user_answer := ""
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
		fmt.Printf("What is %s = \n", record[0])
		q_counter += 1
		fmt.Scanf("%s", &user_answer)
		if user_answer == record[1] {
			n_correct_answer += 1
		}
	}

	fmt.Printf("Number of correct answer is %d out or %d \n", n_correct_answer, q_counter)
}