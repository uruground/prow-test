package main

import "fmt"

func main() {
	fmt.Println(echo("hello"))
}

func echo(s string) string {
	return s
}
