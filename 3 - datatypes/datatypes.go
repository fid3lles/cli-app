package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		Int types:
			int
			int8
			int16
			int32 = rune
			int64
			:=
			uint
			uint8 = byte
			uint16
			uint32
			uint64
	*/

	var number int = 100
	fmt.Println(number)

	var number2 uint16 = 100
	fmt.Println(number2)

	/*
		Float types:
			float32
			float64
	*/

	var float1 float64 = 100.60
	fmt.Println(float1)

	/*
		String types:
			string
			char = need to be declared with single quotes
	*/

	var string1 string = "test"
	fmt.Println(string1)

	char := 'a'
	//Char variables are stored as int32, that integer represents the character ASCII code
	fmt.Println(char)

	//Variables without value starts with '0' value:
	var variableWithoutValue int32
	fmt.Println(variableWithoutValue)

	//Boolean variable
	var boolean1 bool = true
	fmt.Println(boolean1)

	//This declaration infer a false value into the boolean variable
	var boolean2 bool
	fmt.Println(boolean2)

	/*
		Error variable
			error

		<nil> - standard value for error variables without value
		errors.New(<message>) - create a new error with a custom message
	*/
	var err error
	fmt.Println(err)

	var err2 error = errors.New("Error message")
	fmt.Println(err2)
}
