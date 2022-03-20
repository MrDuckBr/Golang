package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	escrita        = "escrita"
	pausa          = "pausa"
	inicioContagem = 3
	ultimaPalavra  = "Go!"
)

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

//Alteracao da assinatura do metodo time.alguma coisa
type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

type SleeperPadrao struct {
}

type Sleeper interface {
	Pausa()
}

func (d *SleeperPadrao) Pausa() {
	time.Sleep(1 * time.Second)
}

func Contagem(saida io.Writer, sleep Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleep.Pausa()
		fmt.Fprintln(saida, i)
	}

	sleep.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}
