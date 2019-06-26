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
	ID       int
	Nome     string
	Idade    int
	Endereco string
	RG       string
	Data     string
}
type Produto struct {
	ID     int
	Codigo int
	Nome   string
	Valor  float32
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

func SetDados_Clientes(nome string, idade int, endereço string, rg string, data string, cli *[]Cliente) {
	// var c Cliente
	// fmt.Print(c.Pessoas.Idade)
	*cli = append(*cli, Cliente{Nome: nome, Idade: idade, Endereco: endereço, RG: rg, Data: data})
}

func main() {

	var sum int
	sum = 0
	var pes []Pessoas
	var cli []Cliente

	for sum != 10 {
		fmt.Println("\n1: Gereciar Pessoas 2:Cliente 3:Produto 4:Totalizavel 5: Sair")
		var controle int
		var controle1 int

		// var pes [100]Pessoas
		fmt.Scan(&controle)
		switch controle {
		case 1:
			fmt.Printf("\n ||| Pessoas ||| \n")
			fmt.Printf("\n 1:Cadastrar \t 2: Visualizar \n")
			fmt.Scan(&controle1)
			switch controle1 {
			case 1:
				var endereco, nomes string
				var idade int

				fmt.Printf("Digite nome:\n")
				fmt.Scan(&nomes)
				fmt.Printf("Digite Idade:\n")
				fmt.Scan(&idade)
				fmt.Printf("Endereço:\n")
				fmt.Scan(&endereco)
				SetDados_Pessoais(nomes, idade, endereco, &pes)
				break

			case 2:
				for i, V := range pes {
					fmt.Printf("ID %d Nome: %s \t Idade: %d \nEndereço:%s \n", i, V.Nome, V.Idade, V.Endereco)
				}
				break
			}

		case 2:
			fmt.Printf("\n ||| Cliente ||| \n")

			fmt.Printf("\n 1:Cadastrar \t 2: Visualizar 3:Associar \n")
			fmt.Scan(&controle1)
			switch controle1 {
			case 1:
				var endereco, nomes, rg, data string
				var idade int

				fmt.Printf("Digite nome:\n")
				fmt.Scan(&nomes)
				fmt.Printf("Digite Idade:\n")
				fmt.Scan(&idade)
				fmt.Printf("Endereço:\n")
				fmt.Scan(&endereco)
				fmt.Printf("RG:")
				fmt.Scan(&rg)
				fmt.Printf("Data:")
				fmt.Scan(&data)
				SetDados_Clientes(nomes, idade, endereco, rg, data, &cli)
				break

			case 2:
				for i, V := range cli {
					fmt.Printf("ID %d Nome: %s \t Idade: %d \nEndereço:%s \n", i, V.Nome)
				}
				break
			}
		case 3:
			fmt.Printf("\n ||| Produto ||| \n")
			fmt.Printf("\n 1:Cadastrar Produto \t 2: Visualizar 3:Comprar 4: Lista de Cliente \n")
			fmt.Scan(&controle1)
			switch controle1 {
			case 1:
				var endereco, nomes, rg, data string
				var idade int

				fmt.Printf("Digite nome:\n")
				fmt.Scan(&nomes)
				fmt.Printf("Digite Preço:\n")
				fmt.Scan(&idade)
				fmt.Printf("Endereço:\n")
				fmt.Scan(&endereco)
				fmt.Printf("RG:")
				fmt.Scan(&rg)
				fmt.Printf("Data:")
				fmt.Scan(&data)
				SetDados_Clientes(nomes, idade, endereco, rg, data, &cli)
				break

			case 2:
				for i, V := range cli {
					fmt.Printf("ID %d Nome: %s \t Idade: %d \nEndereço:%s \n", i, V.Nome)
				}
				break
			}

		case 4:
			fmt.Printf("\nTotalizavel\n")
		case 5:
			sum = 10
		}
	}

	// checkNumber(getValue())
}
