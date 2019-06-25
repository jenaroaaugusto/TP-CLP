package Pessoa

import "fmt"

type Pessoas struct {
	Nome     string
	Idade    int
	Endereco string
	Lulu     int
}

func Nomesdo() {
	fmt.Println("Estou aqui")
}
func (p *Pessoas) GetNome_Pessoa() string {

	// fmt.Println(p.Nome)
	return p.Nome
}
func (p *Pessoas) GetIdade() string {

	return p.Endereco
}
func GetEndereco() string {
	var p Pessoas
	// fmt.Println(p.Nome)
	return p.Nome
}
func SetNome_Pessoa(nome string) {
	var p Pessoas
	p.Nome = nome
}
func (p *Pessoas) SetIdade(idade int) {

	p.Idade = idade

}
func (p *Pessoas) SetEndereço(endereço string) {
	p.Endereco = endereço
}
