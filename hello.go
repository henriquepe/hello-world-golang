package main

// Aqui estou declarando o pacote principal da minha aplicação, acredito que seja como se fosse um server.ts
// Isso deve ser feito em todos os arquivos de um mesmo diretório

import (
	"fmt"

	"rsc.io/quote"
)

// Aqui estou importando o pacote fmt que possui funções para formatar strings, e printar strings no console
// Este pacote já vem instalado por default com o GO

// Depois de adicionar um novo módulo ao projeto que não vem por padrão,
// é necessário digitar o seguinte comando no terminal para instalar o módulo: go mod tidy

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(quote.Go())
}

// Temos o pacote main que executa a função main (obrigatório)
// Acredito que em futuras implementações esse é o ponto de partida na execução de outros serviços
