package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

//Var funciona num escopo global
var ErroSaldoInsuficiente = errors.New("Operacao recusada: não é possivel retirar")

func (c *Carteira) Retirar(valor Bitcoin) error {
	if valor > c.saldo {
		return ErroSaldoInsuficiente
	}
	c.saldo -= valor
	return nil
}
