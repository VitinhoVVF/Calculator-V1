package main

import (
	"fmt"
)

func soma(valor1 float32, valor2 float32) float32 {
	return valor1 + valor2

}

func sub(valor1 float32, valor2 float32) float32 {
	return valor1 - valor2

}

func div(valor1 float32, valor2 float32) float32 {

	if valor2 == 0 {
		fmt.Println("Não da para dividir por 0!")
	}

	return valor1 / valor2
}

func mult(valor1 float32, valor2 float32) float32 {
	return valor1 * valor2

}

func menu() int {
	var resposta int

	fmt.Println("Calculadora V1")

	fmt.Println("1- Soma?")
	fmt.Println("2- Subtração?")
	fmt.Println("3- Divisão?")
	fmt.Println("4- Multiplicação?")

	fmt.Println("Digite sua opção: ")
	fmt.Scanln(&resposta)

	return resposta
}

func main() {
	var valor1, valor2 float32
	var resposta int
	var resultado float32

	resposta = menu()
	fmt.Println("Digite o valor 1: ")
	fmt.Scanln(&valor1)
	fmt.Println("Digite o valor 2: ")
	fmt.Scanln(&valor2)

	switch resposta {
	case 1:
		resultado = soma(valor1, valor2)
		fmt.Println(resultado)
	case 2:
		resultado = sub(valor1, valor2)
		fmt.Println(resultado)
	case 3:
		resultado = div(valor1, valor2)
	case 4:
		resultado = mult(valor1, valor2)
		fmt.Println(resultado)
	default:
		fmt.Println("Opção invalida!")

	}

}
