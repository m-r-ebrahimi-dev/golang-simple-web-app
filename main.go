package main

import "fmt"

type User struct {
	Custom Mf
}

type Mf func(int) string

func a(b int) string {
	return fmt.Sprintf("cc: %d", b)
}

var mmmm = make(chan interface{})

func main() {
}
func b() {
	fmt.Println("aaa")
}
