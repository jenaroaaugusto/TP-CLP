package main

import "fmt"
import Pessoas "./Pessoa"

// import us "/usuario/usuario"
// import "time"

// import (
// 	L "./Pessoa"
// 	C "./usuario"
// )

type Pessoa struct {
	// contains filtered or unexported fields
	Nome     string
	Idade    int
	Endereço string
}

func (p *Pessoa) GetNome_Pessoa() string {
	// p := Pessoa{}
	// fmt.Println(p.Nome)
	return p.Nome
}
func (p *Pessoa) GetIdade() string {
	// p := Pessoa{}
	// fmt.Println(p.Endereço)
	return p.Endereço
}
func (p *Pessoa) GetEndereco() {
	// p := Pessoa{}
	fmt.Println(p.Nome)
}
func (p *Pessoa) SetNome_Pessoa(nome string) {
	// p := Pessoa{}
	p.Nome = nome
}
func (p *Pessoa) SetIdade(idade int) {
	// p := Pessoa{}
	p.Idade = idade

}
func (p *Pessoa) SetEndereço(endereço string) {
	p.Endereço = endereço
}

type Cliente struct {
	Pessoa
	RG   string
	Data string
}

func (p *Pessoa) GetIdade() int {
	// p := Pessoa{}
	return p.Endereço
	// fmt.Println(p.Endereço)
}

func (p *Cliente) SetRG(RG string) {
	// p.RG := Pessoa{}
	p.RG = RG
}
func (p *Cliente) SetData(Data string) {
	// p.RG := Pessoa{}
	p.Data = Data
}

func main() {
	// var nome, cargo string
	// var idade int
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
