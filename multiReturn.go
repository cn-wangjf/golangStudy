package main

import (
	"fmt"
)

func main() {
	fn, mn, ln, nn := getName()
	fmt.Println(fn, mn, ln, nn)
}

func getName() (firstName, middleName, lastName, nickName string) {
	return "May", "M", "Chen", "Babe"
}
