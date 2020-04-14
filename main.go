package main

import "fmt"

func main() {
	fmt.Println(echo("hello"))
	fmt.Println(echo("world"))
}

func echo(s string) string {
	return s
}
