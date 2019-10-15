package main

import (
	divi "Practica1/calculadora/pkg/ops/div"
	multi "Practica1/calculadora/pkg/ops/mul"
	resta "Practica1/calculadora/pkg/ops/res"
	suma "Practica1/calculadora/pkg/ops/sum"
	"fmt"
)

func main() {
	fmt.Println(`
	1.- Suma	
	2.- Resta	
	3.- Multiplicación	
	4.- Division`)

	fmt.Println("Que Operación desea realizar?")
	var ope int
	fmt.Scanln(&ope)
	var num, num1 float32

	if ope != 1 && ope != 2 && ope != 3 && ope != 4 {
		fmt.Println("No hay una operacion asociada a ese numero seleccionado")
	} else {
		fmt.Println("Ingrese Su Primer Numero Para La Operacion:")
		fmt.Scanln(&num)
		fmt.Println("Ingrese Su Segundo Numero Para La Operacion:")
		fmt.Scanln(&num1)

		switch ope {
		case 1:
			result := suma.Suma(num, num1)
			fmt.Println("El resultado es: ", result)
		case 2:
			result := resta.Resta(num, num1)
			fmt.Println("El resultado es: ", result)
		case 3:
			result := multi.Multiplicacion(num, num1)
			fmt.Println("El resultado es: ", result)
		case 4:
			result := divi.Division(num, num1)
			fmt.Println("El resultado es: ", result)
		}

	}
}
