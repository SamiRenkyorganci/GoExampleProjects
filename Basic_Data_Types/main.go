/*
--- BASIC TYPES (BUILT-IN TYPES)

# NUMERIC TYPES

uint8       the set of all unsigned  8-bit integers (0 to 255) = 256
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
int8        the set of all signed  8-bit integers (-128 to 127) = 256
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers
complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts
byte        alias for uint8
rune        alias for int32

# STRING TYPES

A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes.
The number of bytes is called the length of the string and is never negative.

# BOOLEAN TYPES

A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false.

--- COMPOSITE TYPES

Slice - Struct - Pointer - Function - Interface - Map - Channel - Array
*/
package main

import "fmt"

func main() {

	// name := "Sami"
	// age := 25
	// job := true
	// job="true" Wrong string
	var month uint8 = 12
	var days uint16 = 256
	var weight = 72.4
	// var (
	// 	name string = "Sami"
	// 	age  int    = 25

	// 	job  bool   = true
	// )

	// var name ,age,job ="Sami",12,true

	//Zero Values string --> "" ,numeric --> 0
	name, age, job := "Sami", 12, true
	fmt.Println(name, age, job, month, days, weight)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", job)
	fmt.Printf("%T\n", weight)

}
