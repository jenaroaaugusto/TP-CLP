package Pessoa

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco string
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
func (p *Pessoa) GetEndereco() string {
	// p := Pessoa{}
	// fmt.Println(p.Nome)
	return p.Nome
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
