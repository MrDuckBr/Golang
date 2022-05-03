package main

import (
	"fmt"
	"time"
)

type VerificadorWebsite func(string) bool
type resultado struct {
	string
	bool
}

//Canal são estruturas de dados nativas do GO, dar uma olhada depois

func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)
	canalResultado := make(chan resultado)
	for _, url := range urls {
		go func(u string) {
			canalResultado <- resultado{u, vw(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		resultado := <-canalResultado
		fmt.Println()
		resultados[resultado.string] = resultado.bool
	}

	time.Sleep(2 * time.Second)
	return resultados
}
