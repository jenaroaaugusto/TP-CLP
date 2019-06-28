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
	Numero    int
	Data      string
	cliente   Cliente
	ItemVenda []ItemVenda
	idlista   int
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
func removeVenda(slice []Venda, s int) []Venda {

	slice = append(slice[:s], slice[s+1:]...)
	return slice
}
func Somas(valor float32, Numero int) float32 {
	rar := float32(float32(Numero) * valor)
	return rar
}

// &ven, &cli, id, idp1, idp2
func Total(NumerodaVenda int, ven *[]Venda, id int, idp1 int, idp2 int, datavenda string, nome string, idade int, rg string, data string, endereço string, iten *[]ItemVenda, nam string, cod int, valor float32) {

	*ven = append(*ven, Venda{Numero: NumerodaVenda, Data: datavenda, cliente: Cliente{RG: rg, Data: data, Pessoas: Pessoas{nome, idade, endereço}}})
}
func lista(NumerodaVenda int, vcs *[]ItemVenda, nam string, cod int, valor float32, idp2 int) {
	var resultado float32
	resultado = Somas(valor, idp2)
	*vcs = append(*vcs, ItemVenda{Produto: Produto{Codigo: cod, Nome: nam}, Quantidade: NumerodaVenda, Valor: resultado})

}


func main() {

	var sum int
	sum = 0
	var pes []Pessoas
	var cli []Cliente
	var pro []Produto
	var ven []Venda

	var iten []ItemVenda

	var id, idp1, idp2 int //idp3, idp4, idp5 int
	var datavenda string

	for sum != 10 {
		var controle1 int
		var controle3 int
		var controle4 int
		var controle int
		var controlepalavras string
		var controleFloat float32
		var NumerodaVenda int

		fmt.Println("\n|\t|Sistema De Gestão|\t|\n")
		fmt.Println("\n 2:Cliente 3:Produto 4:Venda 5: Sair")
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
			fmt.Printf("\n 1:Cadastrar Produto \t 2: Lista de Produtos \t 3:Remover um produto \t 4:Alterar Produtos \n")
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
				fmt.Printf("\n|\t| Alteração dos Clientes |\t|\n")
				fmt.Printf("Escolha pelo ID")
				for i, V := range pro {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID: %d \t Codigo %d \t Nome: %s \t Valor: %.2f  \n", i, V.Codigo, V.Nome, V.Valor)
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}
				fmt.Printf("\nDigite ID: \n")
				fmt.Scan(&controle3)
				// s := len(cli) Tratar os erros
				var interacao int
				for interacao != 10 {
					fmt.Printf("\n Alterar selecione-:> 1:Codigo 2:Nome 3:Valor 4:Concluido\n")
					fmt.Scan(&controle4)
					if controle4 == 1 {
						fmt.Printf("Digite o novo código:")
						fmt.Scan(&controle)
						pro[controle3].Codigo = controle
					} else if controle4 == 2 {
						fmt.Printf("Digite o novo nome: ")
						fmt.Scan(&controlepalavras)
						pro[controle3].Nome = controlepalavras
					} else if controle4 == 3 {
						fmt.Printf("Digite o novo Valor:")
						fmt.Scan(&controleFloat)
						pro[controle3].Valor = controleFloat
					} else if controle4 == 4 {
						interacao = 10
						break
					}
				}
			}

		//Venda
		case 4:
			fmt.Printf("\nVenda\n")
			fmt.Printf("\n 1:Fazer Venda \t 2: Visualizar 3:Remover  \n")
			fmt.Scan(&controle1)
			switch controle1 {
			// Fazer a venda

			case 1:

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
				var nam string
				var cod int
				cod = pro[idp1].Codigo
				nam = pro[idp1].Nome
				valor := pro[idp1].Valor

				fmt.Printf("Data:\n")
				fmt.Scan(&datavenda)
				lista(NumerodaVenda, &iten, nam, cod, valor, idp2)
				Total(NumerodaVenda, &ven, id, idp1, idp2, datavenda, nome2, Idade2, rg2, data2, end2, &iten, nam, cod, valor)

				copy(ven[NumerodaVenda].ItemVenda, iten)
				fmt.Println(ven)
				NumerodaVenda = NumerodaVenda + 1

			case 2:

				for i, V := range ven {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID: %d \t Numero da Venda:%d Data: %s Cliente: Nome:%s RG %s \n  ", i, V.Numero, V.Data, V.cliente.Nome, V.cliente.RG)
					for j := 0; j < len(V.ItemVenda); j++ {
						fmt.Printf("ID %d Produto: Codigo %d \t Nome:%s \t Valor: %.2f  \n", j, V.ItemVenda[j].Produto.Codigo, V.ItemVenda[j].Produto.Nome, V.ItemVenda[j].Valor)
					}
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}
			case 3:
				fmt.Printf("\n|\t| Remover Venda|\t|\n")
				fmt.Printf("Escolha pelo ID")
				for i, V := range ven {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID: %d \t Numero da Venda:%d Data: %s Cliente: Nome:%s RG %s \n  ", i, V.Numero, V.Data, V.cliente.Nome, V.cliente.RG)
					for j := 0; j < len(V.ItemVenda); j++ {
						fmt.Printf("ID %d Produto: Codigo %d \t Nome:%s \t Valor: %.2f  \n", j, V.ItemVenda[j].Produto.Codigo, V.ItemVenda[j].Produto.Nome, V.ItemVenda[j].Valor)
					}
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}
				fmt.Printf("\nDigite ID\n")
				fmt.Scan(&controle3)

				ven := removeVenda(ven, controle3)

				for i, V := range ven {
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
					fmt.Printf("ID: %d \t Numero da Venda:%d Data: %s Cliente: Nome:%s RG %s \n  ", i, V.Numero, V.Data, V.cliente.Nome, V.cliente.RG)
					for j := 0; j < len(V.ItemVenda); j++ {
						fmt.Printf("ID %d Produto: Codigo %d \t Nome:%s \t Valor: %.2f  \n", j, V.ItemVenda[j].Produto.Codigo, V.ItemVenda[j].Produto.Nome, V.ItemVenda[j].Valor)
					}
					fmt.Printf("\n-----------------------------------------------------------------------------\n")
				}

			}

		//Sair
		case 5:
			sum = 10
		}
		
	}
	defer func() {
        if err := recover(); err != nil {
            fmt.Println("Ocorreu um erro: ", err)
        }
    }()
}
