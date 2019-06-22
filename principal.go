package main

import "fmt"
import "os"
import Pessoas "./Pessoa"

func main() {
	// var nome, cargo string
	// var idade int
	// Outra pessoa
	pes := Pessoas.Pessoa1{Nome: "Luana"}
	fmt.Print(pes)
	os.Exit(1)
	var sum int
	sum = 0
	var controle int
	// Cliente := Pessoa{}
	// fmt.Println("Sou um programa em Go, eu existo!")
	// fmt.Println("Qual e o seu nome e idade:")
	// fmt.Scanln(&Cliente.Nome, &Cliente.Idade)
	// fmt.Println("Certo, agora informe seu endereço: ")
	// fmt.Scan(&Cliente.Endereco)
	// fmt.Printf("Nome %s Idade %d seu endereço %s", Cliente.Nome, Cliente.Idade, Cliente.Endereco)
	// // fmt.Println()

	for sum != 10 {
		fmt.Println("\n1:Pessoa 2:Cliente 3:Produto 4:Totalizavel 5: Sair")
		fmt.Scan(&controle)
		switch controle {
		case 1:

			fmt.Printf("\n Sou uma Pessoas\n")
			Pessoas.GetNome()
		case 2:
			fmt.Printf("\n Sou um Cliente\n")
		case 3:
			fmt.Printf("\nProduto\n ")
		case 4:
			fmt.Printf("\nTotalizavel\n")
		case 5:
			sum = 10
		}
	}

	// checkNumber(getValue())
}
