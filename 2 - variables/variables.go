package main

import "fmt"

func main() {
	//Inferring data type of variable with declaration
	var variable1 string = "Var1"
	fmt.Println(variable1)

	//Inferring data type of variable without declaration
	variable2 := "Var2"
	fmt.Println(variable2)

	//Declaring a lot of variables
	var (
		variable3 string = "test1"
		variable4 string = "test2"
	)
	fmt.Println(variable3, variable4)

	variable5, variable6 := "test3", "test4"
	fmt.Println(variable5, variable6)

	//Declaring contant variables
	const constant1 string = "Const1"
	fmt.Println(constant1)

	//In GoLang we can exchange data between variables without using an auxiliary variable
	variable5, variable6 = variable6, variable5
	fmt.Println(variable5, variable6)
}
