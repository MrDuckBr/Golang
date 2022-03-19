package main

import (
	"errors"
)

type Dicionario map[string]string

var (
	erroPalavraExistente   = errors.New("Não foi possivel adicionar a palavra, ja existente")
	erroNaoEncontrado      = errors.New("Não foi possivel encontrar a palavra")
	erroPalavraInexistente = errors.New("Palavra não está no banco")
)

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

func (d Dicionario) Busca(chave string) (string, error) {
	valor, existe := d[chave]

	if !existe {
		return "", erroNaoEncontrado
	}

	return valor, nil
}

func (d Dicionario) Adiciona(chave, valor string) error {
	_, err := d.Busca(chave)
	switch err {
	case erroNaoEncontrado:
		d[chave] = valor
	case nil:
		return erroPalavraExistente

	default:
		return err
	}
	return nil
}

func (d Dicionario) Atualiza(palavra, definicao string) error {

	_, err := d.Busca(palavra)
	switch err {
	case erroNaoEncontrado:
		return erroPalavraInexistente
	case nil:
		d[palavra] = definicao
	default:
		return err
	}
	return nil

}

func (d Dicionario) Deleta(palavra string) {
	delete(d, palavra)

}
