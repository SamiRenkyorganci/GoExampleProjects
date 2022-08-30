package main

import "fmt"

// packVar2 := "Package Scope"//expected declaration, found packVar2

var packVar = "Package Scope"

func main() {

	funcVar := "Func Scope"
	if true {
		blockVar := "Block Scope"
		fmt.Println(blockVar)
	}
	// fmt.Println(blockVar) Not Scope
	fmt.Println(funcVar)
	//fmt.Println(packVar)
	myFunc()
}
func myFunc() {
	fmt.Println(packVar)
	//fmt.Println(funcVar)Not Scope
}
