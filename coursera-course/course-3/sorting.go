package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	_ "sync"
)

func main() {
	iElm := TakeInput()
	eLen := len(iElm)
	aNunOfElm := eLen / 4

	c := make(chan []int)
	go Sort(iElm[0:aNunOfElm], c)
	go Sort(iElm[aNunOfElm:(2*aNunOfElm)], c)
	go Sort(iElm[(2*aNunOfElm):(3*aNunOfElm)], c)
	go Sort(iElm[(3*aNunOfElm):], c)

	// take 4 sorted arry from channel
	a1 := <-c
	a2 := <-c
	a3 := <-c
	a4 := <-c

	iElm = Merge(a1, a2)
	iElm = Merge(iElm, a3)
	iElm = Merge(iElm, a4)
	fmt.Println("Final sorted array is", iElm)
}

// merge two sorted array
func Merge(a []int, b []int) []int {
	var i, j int
	var mArray []int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			mArray = append(mArray, a[i])
			i++
		} else {
			mArray = append(mArray, b[j])
			j++
		}
	}

	// add remaining element
	if i != len(a) {
		mArray = append(mArray, a[i:]...)
	}

	if j != len(b) {
		mArray = append(mArray, b[j:]...)
	}

	return mArray
}

// sort functin
func Sort(a []int, ch chan []int) {
	fmt.Println("partial array have to sort", a)
	sort.Ints(a)
	ch <- a
}

func TakeInput() []int {
	fmt.Println("Enter your input array")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rLine := scanner.Text()
	return numbers(rLine)
}

// space separated string to []int
func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		// convert string to int
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}
