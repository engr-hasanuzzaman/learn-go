package main
import(
	"fmt"
	"encoding/csv"
	_ "io"
	"os"
	"log"
) 

func main()  {
	// wait_sec := 15
	n_correct_answer := 0
	q_counter := 0
	// user_answer := ""
	csv_file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Could not open file ", err)
	}

	r := csv.NewReader(csv_file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal("error ", err)
	}

	questions_set := make_problem_set(lines)
	fmt.Println(questions_set[0].a)
	// for {
	// 	record, err := r.Read()
		
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	if err != nil {
	// 		log.Fatal("CSV read error")
	// 	}
	// 	fmt.Printf("What is %s = \n", record[0])
	// 	q_counter += 1
	// 	fmt.Scanf("%s", &user_answer)
	// 	if user_answer == record[1] {
	// 		n_correct_answer += 1
	// 	}
	// }

	fmt.Printf("Number of correct answer is %d out or %d \n", n_correct_answer, q_counter)
}

/* stuct for keeping problem */ 
type problem struct {
	q string
	a string
}

/* read csv and make array of problems*/
func make_problem_set(lines [][]string) []problem {
	res := make([]problem, len(lines))
	for i, line := range lines {
		res[i] = problem{line[0], line[1]}
	}
	
	return res
} 