package main

import (
	"fmt"
	"testing"
)

func TestCarteira(t *testing.T) {

	verificaSaldo := func(t *testing.T, c Carteira, esperado Bitcoin) {
		t.Helper()
		resultado := c.Saldo()

		if resultado != esperado {
			t.Errorf("Resuldado %s , esperado %s", resultado, esperado)
		}
	}

	confirmaErro := func(t *testing.T, resultado error, esperado error) {
		if resultado == nil {
			fmt.Print("Esperava um erro mas ele não aconteceu")
		}

		if resultado != esperado {
			t.Errorf("Resuldado %s , esperado %s", resultado, esperado)
		}

	}

	t.Run("Depositar na carteira", func(t *testing.T) {

		carteira := Carteira{}

		carteira.Depositar(Bitcoin(10))
		esperado := Bitcoin(10)
		verificaSaldo(t, carteira, esperado)

	})

	t.Run("Retirar da Carteira", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}

		carteira.Retirar(Bitcoin(10))
		esperado := Bitcoin(10)
		verificaSaldo(t, carteira, esperado)

	})

	t.Run("Retirar mais do que se tem na carteira", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}

		erro := carteira.Retirar(Bitcoin(100))

		verificaSaldo(t, carteira, Bitcoin(20))
		confirmaErro(t, erro, ErroSaldoInsuficiente)

	})

}
