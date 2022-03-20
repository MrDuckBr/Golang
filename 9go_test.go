package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestContagem(t *testing.T) {
	t.Run("Imprime até 3 e vai", func(t *testing.T) {

		buffer := &bytes.Buffer{}

		Contagem(buffer, &SpyContagemOperacoes{})

		resultado := buffer.String()
		esperado := `3
2
1
Go!`

		if resultado != esperado {
			t.Errorf("Resultado: %s  , esperado: %s", resultado, esperado)
		}

	})

	t.Run("Pausa antes de cada impressao", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperacoes{}

		Contagem(spyImpressoraSleep, spyImpressoraSleep)
		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas) {
			t.Errorf("esperado %v chamadas, resultado %v", esperado, spyImpressoraSleep.Chamadas)
		}
	})

}

func TestSleeperConfiguravel(t *testing.T) {
	tempoPausa := 5 * time.Second

	tempoSpy := &TempoSpy{}

	sleeper := SleeperConfiguravel{tempoPausa, tempoSpy.Pausa}
	sleeper.Pausa()

	if tempoSpy.duracaoPausa != tempoPausa {
		t.Errorf("deveria ter pausado por %v, mas pausou por %v", tempoPausa, tempoSpy.duracaoPausa)
	}
}
