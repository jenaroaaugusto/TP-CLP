package main

import "fmt"

//import "errors"

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
	ID   int
	RG   string
	Data string
	Pessoas
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

	fmt.Print("\n")

	*cli = append(*cli, Cliente{RG: rg, Data: data, Pessoas: Pessoas{nome, idade, endereço}})

}

func main() {

	var sum int
	sum = 0
	var pes []Pessoas
	var cli []Cliente

	for sum != 10 {
		fmt.Println("\n|\t|Sistema De Gestão|\t|\n")
		fmt.Println("\n1: Gereciar Pessoas 2:Cliente 3:Produto 4:Totalizavel 5: Sair")
		var controle int
		var controle1 int
		var controle3 int
		var controle4 int
		var controlepalavras string
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

			fmt.Printf("\n 1:Cadastrar Cliente \t 2: Visualizar 3:Alterar  \n")
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
				fmt.Print(cli[0].Pessoas.Nome)
				break

			case 2:
				//Visuzalizar os dados
				for i, V := range cli {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID %d Nome: %s \t Idade: %d \nEndereço:%s \n", i, V.Pessoas.Nome, V.Pessoas.Idade, V.Pessoas.Endereco)
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}
				break
			case 3:
				//Realizar alteração
				fmt.Printf("\n|\t| Alteração dos Clientes |\t|\n")
				fmt.Printf("Escolha pelo ID")
				for i, V := range cli {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID %d Nome: %s \t Idade: %d \nEndereço:%s \n RG:%s \t Data:", i, V.Pessoas.Nome, V.Pessoas.Idade, V.Pessoas.Endereco, V.RG)
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}
				fmt.Printf("\nDigite ID\n")
				fmt.Scan(&controle3)
				// s := len(cli) Tratar os erros
				for interacao !=10
					fmt.Printf("\n Alterar selecione-:> 1:Nome 2:Idade 3:Endereço 4:RG  5:Concluido\n")
					fmt.Scan(&controle4) 
					if controle4 == 1{
						fmt.Printf("Digite o novo nome:")
						fmt.Scan(&controlepalavras)
						cli[controle3].Pessoas.Nome=controlepalavras

					}else if controle4== 2{
						fmt.Printf("Digite o nova Idade:")
						fmt.Scan(&controle)
						cli[controle3].Pessoas.Idade=controle
					}else if controle4 == 3{

					}else if controle4 == 4{

					}else if controle4 == 5{
						interacao= 10
						break
					}

					
					

			}
		case 3:
			fmt.Printf("\n ||| Produto ||| \n")
			fmt.Printf("\n 1:Cadastrar Produto \t 2: Lista de Cliente \t 3:Comprar 4:  \n")
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
