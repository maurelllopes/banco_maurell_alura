package main

import (
	"fmt"

	"github.com/maurelllopes/banco_maurell_alura/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaMaurell := contas.ContaPoupanca{}
	contaMaurell.Depositar(100)
	PagarBoleto(&contaMaurell, 60)
	fmt.Println(contaMaurell.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 400)

	fmt.Println(contaDaLuisa.ObterSaldo())

}
