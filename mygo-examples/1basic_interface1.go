package main

import "fmt"

// Define an interface to functionalities
type DoubleReturn interface {
	double() int
}
type thrippleReturn interface {
	thripple() int
}

// Insted interfaces
type MultipleReturns interface {
	DoubleReturn
	thrippleReturn
}

type basicCls struct {
	InvAmount int
}

// Define an struct to hold data
type internalCls struct {
	ReturnAmount int
	basicCls     //Embedding another structs
}

// associating a method to structs
// Should match with signature of interface method
// struct need to have associated all implementations of an interface
func (ic internalCls) double() int {
	return ic.InvAmount * 2
}

// associating a method to structs
// Should match with signature of interface method
// struct need to have associated all implementations of an interface
func (ic internalCls) thripple() int {
	return ic.InvAmount * 3
}

// a function with an interface
func goodInvestment(d DoubleReturn) int {
	return d.double()
}

// a function with an interface
func veryGoodInvestment(d MultipleReturns) int {
	return d.thripple()
}
func main() {
	//Declaring a struct
	myClass := internalCls{basicCls: basicCls{InvAmount: 10}}
	// passing a struct for an interface.
	fmt.Println(goodInvestment(myClass))
	fmt.Println(veryGoodInvestment(myClass))
}
