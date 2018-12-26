package main
import(
	. "fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		Fprintf(w, "Welcome to go web development")
	})

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request){
		Fprintf(w, "Welcome to go index")
	})

	http.ListenAndServe(":8080", nil)
}