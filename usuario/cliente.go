package Usuario

import "fmt"

// import
import (
	Pessoas "../Pessoa"
)

type Cliente struct {
	Pessoas.Pessoas
	RG   string
	Data string
}

func GetRG() string {
	var p Cliente
	return p.RG
	// fmt.Println(p.Endereço)
}
func (p *Cliente) GetData() string {
	return p.Data
}

func (p *Cliente) SetRG(RG string) {
	// p.RG := Pessoa{}
	p.RG = RG
}
func (p *Cliente) SetData(Data string) {
	// p.RG := Pessoa{}
	p.Data = Data
}

func (p *Cliente) Estou_vivo() {
	fmt.Print("\nEu vim aqui\n")
	fmt.Println(p.Pessoas.Nome)

}
