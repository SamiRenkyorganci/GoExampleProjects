// Comparison operators
// == Equal != Not equal
// < Less than > Greater than
// <= Less than or equal >= Greater than or equal

// Logical Operators
// &&    conditional AND    p && q  is  "if p then q else false"
// ||    conditional OR     p || q  is  "if p then true else q"
// !     NOT                !p      is  "not p"
package main

import "fmt"

func main() {
	x, y := 3, 7
	//x, y := "a", "b"
	fmt.Printf("%T , %v\n", x == y, x == y) //result  --> bool
	fmt.Printf("%T, %v\n", x != y, x != y)
	fmt.Printf("%T, %v\n", x < y, x < y)
	fmt.Printf("%T, %v\n", x > y, x > y)
	fmt.Printf("%T, %v\n", x >= y, x >= y)
	fmt.Printf("%T, %v\n", x <= y, x <= y)
	fmt.Println("------------------------------")
	set1 := (x == y) // false
	set2 := (x > y)  // false
	set3 := false
	fmt.Printf("%T, %v\n", set1, set1)
	fmt.Printf("%T, %v\n", set2, set2)
	fmt.Printf("%v\n", (set1 && set2)) // AND --> Called Logical AND operator. If both the operands are non-zero, then condition becomes true  --> true
	fmt.Printf("%v\n", (set3 || set2)) // OR --> Called Logical OR Operator. If any of the two operands is non-zero, then condition becomes true.
	fmt.Printf("%v\n", (!set3))        // NOT Called Logical NOT Operator. Use to reverses the logical state of its operand. If a condition is true, then Logical NOT operator will make it false.
}
