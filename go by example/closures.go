package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(intSeq()())
	fmt.Println(nextInt())
	newInts := intSeq()

	fmt.Println("lalala")
	fmt.Println("hello world")
	fmt.Println(newInts())
}
