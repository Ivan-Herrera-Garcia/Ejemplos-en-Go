package main

/*
Autor: Ivan Herrera Garcia
Fecha: 18/01/2023
*/
import (
	"fmt"
	"strings"
)

/*
Funciones que se pueden utilizar en strings
strings.ToUpper(): Devuelve una nueva cadena, en mayúsculas
strings.ToLower(): Devuelve una nueva cadena, en minúsculas
strings.HasSuffix(): Comprueba si una cadena termina con una subcadena
strings.HasPrefix(): Comprueba si una cadena comienza con una subcadena
strings.Contains():Comprueba si una cadena contiene una subcadena
strings.Count(): cuenta cuántas veces aparece una subcadena en una cadena
strings.Join(): Se utiliza para unir varias cadenas y crear una nueva
strings.Split(): Se usa para crear una matriz de cadenas a partir de una cadena, dividiendo la original en un carácter específico, como una coma o un espacio
strings.ReplaceAll(): Se usa para reemplazar una porción en una cadena y reemplazarla por una nueva.
*/

func main() {
	cadena := "Hola mundo"
	fmt.Print("\nCadena original: '", cadena, "', aplicando ToUpper: '", strings.ToUpper(cadena), "'\n")
	fmt.Print("\nCadena original: '", cadena, "', aplicando ToLower: '", strings.ToLower(cadena), "'\n")
	fmt.Print("\nContiene una subcadena 'Hola mundo' por espacio ' ' : ", strings.Contains(cadena, " "), "\n")
	fmt.Print("\nConteo de caracteres en una cadena ('o'): ", strings.Count(cadena, "o"), "\n")
	fmt.Print("\nImpresion de partes de la cadena por posiciones:")
	fmt.Print("\nDe 0:4: ", cadena[0:4], "\n")
	fmt.Print("\nDe 4: ", cadena[4:], "\n")
	fmt.Print("\nUso de HasPrefix ('Hola mundo'): ", strings.HasPrefix(cadena, "Hola mundo"), "\n")
	fmt.Print("\nUso de HasPrefix ('Hola Mundo'): ", strings.HasPrefix(cadena, "Hola Mundo"), "\n")
	fmt.Print("\nUso de HasSuffix ('mundo'): ", strings.HasSuffix(cadena, "mundo"), "\n")
	fmt.Print("\nUso de HasSuffix ('Mundo'): ", strings.HasSuffix(cadena, "Mundo"), "\n")

}
