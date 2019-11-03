package main
import(
	"fmt"
	"encoding/csv"
	_ "io"
	"os"
	"log"
	"time"
) 

func main()  {
	// wait_sec := 15
	n_correct_answer := 0
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

	timer := time.NewTimer(10 * time.Second)
	questions_set := make_problem_set(lines)
	problem:
					for i, p := range questions_set {
						fmt.Printf("Q. %d: %s = ", i, p.q)
						answerCh := make(chan string)
						go func ()  {
							var ans string
							fmt.Scanf("%s", &ans)
							answerCh <- ans
						}()

						select {
						case <-timer.C:
							break problem
						case a := <-answerCh:
							if a == p.a {
								n_correct_answer += 1
							}
						}
					}

	fmt.Printf("Number of correct answer is %d out or %d \n", n_correct_answer, len(questions_set))
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