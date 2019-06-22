package Pessoa

import "fmt"

type Pessoa1 struct {
	Nome     string
	Endereco string
}

func Nome(nome string, idade string) (string, string) {
	var nomes string
	var endereco string
	p1 := Pessoa1{"0", "j"}
	p := &p1
	p.Nome = nome

	p.Endereco = idade
	nomes = p.Nome
	endereco = p.Endereco
	// return nome
	// fmt.Printf("Nome %s seu cargo %s", nomes, endereco)
	return nomes, endereco
}
func GetNome() {
	// var aq Pessoa
	var non, loq string
	non = "Jenaro"
	loq = "Belizario"
	// var aq Pessoa
	// aq Pessoa Pessoa

	fmt.Printf("Nome %s seu cargo %s", non, loq)
}
