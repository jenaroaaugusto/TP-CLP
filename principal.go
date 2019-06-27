package main

import "fmt"

//import "math/rand"

//import "errors"

// import (
// 	Pes "./Pessoa"
// 	Cli "./Usuario"
// )
type ItemVenda struct {
	Produto    Produto
	Quantidade int
	Valor      float32
}
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

// type Totalizavel interface{
// 	// func Total()
// 	// type Venda struct{}
// }
// Venda(subclasse  de  Totalizavel):  número  (int),  data  (Date),  cliente
// (Cliente), itens (lista ou array de ItemVenda).O método total deve calcular a soma dos totais de cada item;
type Venda struct {
	Numero  int
	Data    string
	cliente Cliente
	// ItemVenda []ItemVenda
}

func SetDados_Pessoais(nome string, idade int, endereço string, a *[]Pessoas) {
	*a = append(*a, Pessoas{Nome: nome, Idade: idade, Endereco: endereço})

}
func SetDados_Clientes(nome string, idade int, endereço string, rg string, data string, cli *[]Cliente) {

	fmt.Print("\n")

	*cli = append(*cli, Cliente{RG: rg, Data: data, Pessoas: Pessoas{nome, idade, endereço}})

}
func SetDados_Produto(codigo int, nome string, valor float32, pro *[]Produto) {
	*pro = append(*pro, Produto{Codigo: codigo, Nome: nome, Valor: valor})

}
func SetVenda() {}
func removeCliente(slice []Cliente, s int) []Cliente {

	slice = append(slice[:s], slice[s+1:]...)
	return slice
}
func removeProduto(slice []Produto, s int) []Produto {

	slice = append(slice[:s], slice[s+1:]...)
	return slice
}

// &ven, &cli, id, idp1, idp2
func Total(NumerodaVenda int, ven *[]Venda, id int, idp1 int, idp2 int, datavenda string, nome string, idade int, rg string, data string, endereço string) {

	*ven = append(*ven, Venda{Numero: NumerodaVenda, Data: datavenda, cliente: Cliente{RG: rg, Data: data, Pessoas: Pessoas{nome, idade, endereço}}})
}

func main() {

	var sum int
	sum = 0
	var pes []Pessoas
	var cli []Cliente
	var pro []Produto
	// var proaux []Produto
	var ven []Venda
	// var iten
	var id, idp1, idp2 int //idp3, idp4, idp5 int
	var datavenda string

	for sum != 10 {
		var controle1 int
		var controle3 int
		var controle4 int
		var controle int
		var controlepalavras string
		var NumerodaVenda int

		fmt.Println("\n|\t|Sistema De Gestão|\t|\n")
		fmt.Println("\n 2:Cliente 3:Produto 4:Totalizavel 5: Sair")
		fmt.Scan(&controle)

		switch controle {
		//Gerencias Pessoas
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

		//Cliente
		case 2:
			fmt.Printf("\n ||| Cliente ||| \n")

			fmt.Printf("\n 1:Cadastrar Cliente \t 2: Visualizar 3:Alterar 4:Remover  \n")
			fmt.Scan(&controle1)
			switch controle1 {
			//Cadastrar
			case 1:
				var endereco, nomes, rg, data string
				var idade int

				fmt.Printf("Digite nome:\n")
				fmt.Scan(&nomes)
				fmt.Printf("Digite Idade:\n")
				fmt.Scan(&idade)
				fmt.Printf("Endereço:\n")
				fmt.Scan(&endereco)
				fmt.Printf("RG:\n")
				fmt.Scan(&rg)
				fmt.Printf("Data:\n")
				fmt.Scan(&data)
				SetDados_Clientes(nomes, idade, endereco, rg, data, &cli)
				fmt.Print(cli[0].Pessoas.Nome)

			//Visuzalizar os dados
			case 2:
				for i, V := range cli {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID %d Nome: %s \t Idade: %d \nEndereço:%s \n", i, V.Pessoas.Nome, V.Pessoas.Idade, V.Pessoas.Endereco)
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}

			//Alteração
			case 3:
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
				var interacao int
				for interacao != 10 {
					fmt.Printf("\n Alterar selecione-:> 1:Nome 2:Idade 3:Endereço 4:RG  5:Concluido\n")
					fmt.Scan(&controle4)
					if controle4 == 1 {
						fmt.Printf("Digite o novo nome:")
						fmt.Scan(&controlepalavras)
						cli[controle3].Pessoas.Nome = controlepalavras

					} else if controle4 == 2 {
						fmt.Printf("Digite o nova Idade:")
						fmt.Scan(&controle)
						cli[controle3].Pessoas.Idade = controle
					} else if controle4 == 3 {
						fmt.Printf("Digite o novo Endereço:")
						fmt.Scanf(controlepalavras)
						cli[controle3].Pessoas.Endereco = controlepalavras

					} else if controle4 == 4 {
						fmt.Printf("Alterar RG:")
						fmt.Scanf(controlepalavras)
						cli[controle3].RG = controlepalavras
					} else if controle4 == 5 {
						interacao = 10
						break
					}

				}
				// Remover
			case 4:
				fmt.Printf("\n|\t| Remover Cliente |\t|\n")
				fmt.Printf("Escolha pelo ID")
				for i, V := range cli {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID %d Nome: %s \t Idade: %d \nEndereço:%s \n RG:%s \t Data:", i, V.Pessoas.Nome, V.Pessoas.Idade, V.Pessoas.Endereco, V.RG)
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}
				fmt.Printf("\nDigite ID\n")
				fmt.Scan(&controle3)
				// s := len(cli) Tratar os erros
				// var interacao int
				cli := removeCliente(cli, controle3)

				for i, V := range cli {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID %d Nome: %s \t Idade: %d \nEndereço:%s \n RG:%s \t Data:", i, V.Pessoas.Nome, V.Pessoas.Idade, V.Pessoas.Endereco, V.RG)
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}
			}

		//Produto
		case 3:
			fmt.Printf("\n ||| Produto ||| \n")
			fmt.Printf("\n 1:Cadastrar Produto \t 2: Lista de Produtos \t 3:Remover um produto 4:Alterar Produtos \n")
			fmt.Scan(&controle1)
			switch controle1 {
			//Cadastrar
			case 1:
				var Codigo int
				var Nome string
				var Valor float32
				fmt.Printf("Digite Codigo do Produto: \t Nome do Produto: \n")
				fmt.Scan(&Codigo)
				fmt.Printf("Digite Nome do Produto: \n")
				fmt.Scan(&Nome)
				fmt.Printf("Digite Preço:\n")
				fmt.Scan(&Valor)

				SetDados_Produto(Codigo, Nome, Valor, &pro)
				break

			//Lista de Produtos
			case 2:
				fmt.Printf("\n |--| Lista de Produtos |--| \n")

				for i, V := range pro {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID: %d \t Codigo %d \t Nome: %s \t Valor: %.2f  \n", i, V.Codigo, V.Nome, V.Valor)
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}
				break
			case 3:
				fmt.Printf("\n|\t| Remover Produto |\t|\n")
				fmt.Printf("Escolha pelo ID")
				for i, V := range pro {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID: %d \t Codigo %d \t Nome: %s \t Valor: %.2f  \n", i, V.Codigo, V.Nome, V.Valor)
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}
				fmt.Printf("\nDigite ID\n")
				fmt.Scan(&controle3)

				pro := removeProduto(pro, controle3)

				for i, V := range pro {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID: %d \t Codigo %d \t Nome: %s \t Valor: %.2f  \n", i, V.Codigo, V.Nome, V.Valor)
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}
			// ALterar Produto
			case 4:

			}

		//Venda
		case 4:
			fmt.Printf("\nVenda\n")
			fmt.Printf("\n 1:Fazer Venda \t 2: Visualizar 3:Alterar 4:Remover  \n")
			fmt.Scan(&controle1)
			switch controle1 {
			// Fazer a venda

			case 1:
				NumerodaVenda = NumerodaVenda + 1
				for i, V := range cli {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID %d Nome: %s \t Idade: %d \nRG:%s \n", i, V.Pessoas.Nome, V.Pessoas.Idade, V.RG)
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}
				fmt.Printf("Selecionar Cliente ID :\n")
				fmt.Scan(&id)
				nome2 := cli[id].Pessoas.Nome
				Idade2 := cli[id].Pessoas.Idade
				rg2 := cli[id].RG
				data2 := cli[id].Data
				end2 := cli[id].Pessoas.Endereco

				// SetVenda(&Venda,&cli[],id)
				// var id, idp1,idp2,idp3,idp4,idp5 int
				fmt.Printf("\n |--| Lista de Produtos |--| \n")

				for i, V := range pro {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID: %d \t Codigo %d \t Nome: %s \t Valor: %.2f  \n", i, V.Codigo, V.Nome, V.Valor)
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}

				fmt.Printf("Produto ID - Numero de produtos:\n")
				fmt.Scan(&idp1, &idp2)
				fmt.Printf("Data:\n")
				fmt.Scan(&datavenda)
				Total(NumerodaVenda, &ven, id, idp1, idp2, datavenda, nome2, Idade2, rg2, data2, end2)
			}

		//Sair
		case 5:
			sum = 10
		}
		// checkNumber(getValue())
	}
}
