package main

/*
Autor: Ivan Herrera Garcia
Fecha: 17/01/2023
*/

import (
	"fmt"
	"math"
)

//Imports para la impresion y utilizar la funcion pow de math

func main() {
	var n1, n2, control int //Variables

	fmt.Printf("Selecciona una opcion\n1. Suma\n2. Resta\n3. Division\n4. Multiplicacion\n5. Elevacion al cuadrado\n")
	fmt.Scanf("%d\n", &control) //Variable de control para el switch

	switch control {
	case 1:
		fmt.Println("Ingresa los numeros a sumar:")
		fmt.Println("Ingresa el valor del numero 1")
		fmt.Scanf("%d\n", &n1)
		fmt.Println("Ingresa el valor del numero 2")
		fmt.Scanf("%d\n", &n2)
		fmt.Println("El resultado es:", (n1 + n2)) //Suma
	case 2:
		fmt.Println("Ingresa los numeros a restar:")
		fmt.Println("Ingresa el valor del numero 1")
		fmt.Scanf("%d\n", &n1)
		fmt.Println("Ingresa el valor del numero 2")
		fmt.Scanf("%d\n", &n2)
		fmt.Println("El resultado es:", (n1 - n2)) //Resta
	case 3:
		fmt.Println("Ingresa los numeros a dividir:")
		fmt.Println("Ingresa el valor del numero 1")
		fmt.Scanf("%d\n", &n1)
		fmt.Println("Ingresa el valor del numero 2")
		fmt.Scanf("%d\n", &n2)
		fmt.Printf("El resultado es: %f", (float64(n1) / float64(n2))) //Division
	case 4:
		fmt.Println("Ingresa los numeros a multiplicar:")
		fmt.Println("Ingresa el valor del numero 1")
		fmt.Scanf("%d\n", &n1)
		fmt.Println("Ingresa el valor del numero 2")
		fmt.Scanf("%d\n", &n2)
		fmt.Println("El resultado es:", (n1 * n2)) //Multiplicacion
	case 5:
		fmt.Println("Ingrese el numero a elevar:")
		fmt.Println("Ingresa el valor del numero")
		fmt.Scanf("%d\n", &n1)
		fmt.Printf("El numero elevado es %.2f", math.Pow(float64(n1), 2))
	default:
		fmt.Printf("No selecciono un valor valido")
	}
}
