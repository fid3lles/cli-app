package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo primeiro módulo!")
	auxiliar.Writeline()

	erro := checkmail.ValidateFormat("1")
	fmt.Println(erro)
}
