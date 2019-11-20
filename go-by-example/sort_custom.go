package main
import(
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int{
	return len(s)
}

func (s byLength) Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool{
	return len(s[i]) > len(s[j])
}

func main(){
	strInputs := []string {"This", "is","sample","string", "and", "nothing", "else", "-0", "22", "1"}
	fmt.Println("Before sorting ", strInputs)
	sort.Sort(byLength(strInputs))
	fmt.Println("After sorting ", strInputs)
}