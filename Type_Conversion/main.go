package main

import "fmt"

func main() {

	x := 10

	y := 10.0
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// fmt.Println(x + y) (mismatched types int and float64)

	//Type Conversion type(value)=> int(y)

	fmt.Println(x + int(y))     // float64 --> int
	fmt.Println(float64(x) + y) //int --> float64

	var z int8 = 11
	var m int16 = 11
	fmt.Println(m + int16(z))

	//a := 11
	//b := "12"
	//fmt.Println(a + int(b)) cannot convert b (variable of type string) to int

	num1 := 106 //ASCII VALUE  "j"
	str1 := string(num1)
	fmt.Printf("%v %T\n", str1, str1)
}
