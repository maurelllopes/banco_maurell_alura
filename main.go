package main

import (
	"fmt"

	"github.com/maurelllopes/banco_maurell_alura/contas"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())

}
