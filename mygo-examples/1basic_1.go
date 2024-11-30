package main

import (
	"fmt"
)

var a int //Out side a function, only declaration

var b string    //Out side a function, only declaration
var c int = 100 // legal
// var d int :=100  illegal
// f:="MyName" illegal

func main() {
	fmt.Println("test again")

	fmt.Printf("b:'%s' done. \n", b)
	print(a)
}
