package main

import "fmt"

func main() {
	say := echo("hello")
	fmt.Println(say)
}

func echo(yell string) string {
	return yell
}
