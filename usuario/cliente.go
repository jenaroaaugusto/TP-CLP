package Usuario

import "fmt"

// import
import (
	P "../Pessoa"
)

func Estou_vivo() {
	fmt.Print("\nEu vim aqui\n")
	// fmt.Println(*P.GetNome)
	P.GetNome()
}
