package main

import (
	"testing"
)

func TestPerimetro(t *testing.T) {
	resultado := Perimetro(Retangulo{10.0, 10.0})
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("Esperado: %.2f , Resultado: %.2f", esperado, resultado)
	}
}

func TestArea(t *testing.T) {

	//Struct anonima
	formaArea := []struct {
		forma    Forma
		esperado float64
	}{
		//Pode nomear arbritariamente para saber oque está sendo passado
		{forma: Retangulo{Largura: 10.0, Altura: 10.0}, esperado: 100.0},
		{forma: Circulo{Raio: 10}, esperado: 314.1592653589793},
		{forma: Triangulo{Base: 12, Altura: 6}, esperado: 36.0},
	}

	for _, valor := range formaArea {
		resultado := valor.forma.Area()

		if resultado != valor.esperado {
			t.Errorf("%#v Esperado: %.2f , Resultado: %.2f", valor.forma, valor.esperado, resultado)

		}
	}

}

//https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/estruturas-metodos-e-interfaces#refatoracao-adicional
