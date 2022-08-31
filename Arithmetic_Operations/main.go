// addition + => 15 + 10 => 15, 10 operand, + operator
// substruction -
// product *
// division /
// remainder %
package main

import "fmt"

func main() {
	x, y := 15, 10
	//fmt.Printf("%T , %v\n", x, x)
	//fmt.Printf("%T , %v\n", y, y)
	fmt.Printf("%T , %v\n", (x + y), (x + y))
	fmt.Printf("%T , %v\n", (x - y), (x - y))
	fmt.Printf("%T , %v\n", (x * y), (x * y))
	fmt.Printf("%T , %v\n", (x / y), (x / y))
	fmt.Printf("%T , %v\n", (x % y), (x % y))
	z := 5.0 / 2
	fmt.Printf("%T , %v\n", z, z)

	i := 10
	fmt.Println("i", i)
	i = i + 1
	fmt.Println("i", i)
	i++ //GO ONLY POSTFIX
	fmt.Println("i", i)
	i-- //GO ONLY POSTFIX
	fmt.Println("i", i)
	//fmt.Println(i++) ERROR: unexpected ++  https://stackoverflow.com/a/25800388
}
