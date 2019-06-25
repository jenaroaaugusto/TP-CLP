package main

import "fmt"

// import (
// 	Pes "./Pessoa"
// 	Cli "./Usuario"
// )
type Pessoas struct {
	Nome     string
	Idade    int
	Endereco string
}

type Cliente struct {
	Pessoas
	RG   string
	Data string
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

func GetRG() string {
	var p Cliente
	return p.RG
	// fmt.Println(p.Endereço)
}
func (p *Cliente) GetData() string {
	return p.Data
}

func SetDados_Pessoais(nome string, idade int, endereço string, a *[]Pessoas) {
	*a = append(*a, Pessoas{Nome: nome, Idade: idade, Endereco: endereço})

}

func (p *Cliente) SetRG(RG string) {

	p.RG = RG
}
func (p *Cliente) SetData(Data string) {

	p.Data = Data
}

func main() {

	var sum int
	sum = 0
	var pes []Pessoas

	for sum != 10 {
		fmt.Println("\n1: Gereciar Pessoas 2:Cliente 3:Produto 4:Totalizavel 5: Sair")
		var controle int
		var controle1 int

		// var pes [100]Pessoas
		fmt.Scan(&controle)
		switch controle {
		case 1:
			fmt.Printf("\n 1:Cadastrar \t 2: Visualizar \n")
			fmt.Scan(&controle1)
			switch controle1 {
			case 1:
				var endereco, nomes string
				var idade int

				fmt.Println("Digite nome:")
				fmt.Scan(&nomes)
				fmt.Print("Digite Idade:\n")
				fmt.Scan(&idade)
				fmt.Printf("Endereço:\n")
				fmt.Scan(&endereco)
				SetDados_Pessoais(nomes, idade, endereco, &pes)

			case 2:
				var i Pessoas

				// fmt.Println("Nome: %s", pes[0].Pessoas.Nome)
				fmt.Printf("Idade: %d %d \n", len(pes), cap(pes))
				fmt.Println("Endereço: %s", i.Endereco)

			}

		case 2:
			fmt.Printf("\n Sou um Cliente\n")
			// var c Cli.Cliente
			// fmt.Printf(c.RG)
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
