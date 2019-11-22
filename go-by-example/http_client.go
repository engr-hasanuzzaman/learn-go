package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://kichole.com")

	if err != nil {
		panic(err)
	}

	fmt.Println(res.Body)
	s := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	for i := 0; s.Scan() && i < 5; i++ {
		fmt.Println(s.Text())
	}
}
