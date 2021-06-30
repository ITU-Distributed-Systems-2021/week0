package main

import "fmt"

func main() {
	say := Echo("hello")
	fmt.Println(say)
}

func Echo(yell string) string {
	return yell
}
