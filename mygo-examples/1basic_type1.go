package main

import "fmt"

type subjects []string

func (subjects subjects) print() {
	for _, s := range subjects {
		fmt.Println(s)
	}
}

func main() {
	mysubjects := subjects{"Mathematics", "Economics"}
	mysubjects.print()
}
